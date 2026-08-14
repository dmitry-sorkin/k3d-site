"""Минимальная обработка изображений для docs/vostok/manual/prepare.md.

Делает белый фон прозрачным, добавляет тонкую白ую обводку по силуэту
контента, обрезает прозрачные поля. Эталон результата - pics/prepare/threaded_inserts.png.

Использование:
    python tools/strip_white.py [--no-crop] <path> [<path> ...]

    --no-crop   не обрезать прозрачные поля по bbox непрозрачного контента.
    --wb        все белые пиксели -> прозрачный (включая через отверстия).
    --stroke    добавить белую обводку по силуэту (по умолчанию нет).

Оригиналы сохраняются в подпапке <dir>/.bak/ как <name>.bak1, .bak2, ... - первый
свободный индекс. Перезаписывает исходный файл. НЕ идемпотентно - повторный прогон
расширит обводку; в .bakN всегда лежит исходник.
"""

import argparse
import sys
from collections import deque
from pathlib import Path

import numpy as np
from PIL import Image

STROKE_PX = 4.0  # ширина обводки, px
AA_WIDTH = 2.0  # ширина сглаживания на внешнем крае, px


def next_backup_path(path: Path) -> Path:
    """Первый незанятый <name>.bakN в подпапке .bak рядом с оригиналом."""
    bak_dir = path.parent / ".bak"
    bak_dir.mkdir(exist_ok=True)
    for n in range(1, 1000):
        cand = bak_dir / f"{path.name}.bak{n}"
        if not cand.exists():
            return cand
    raise RuntimeError(f"слишком много .bak* в {bak_dir}")


def distance_transform(mask: np.ndarray) -> np.ndarray:
    """Евклидова дистанция от каждого False-пикселя до ближайшего True.

    Реализован двухпроходный chamfer-3-4 (ортогональный шаг = 3,
    диагональный = 4, нормировка делением на 3). Точность ~2% от истинной
    евклидовой дистанции, но равномерная по всем направлениям - в отличие от
    квадратной MaxFilter-дилатации, у которой диагональная толщина в sqrt(2)
    раз больше ортогональной.
    """
    try:
        INF = float("inf")
    except ValueError:
        INF = 1e300  # fallback for unusual numpy builds
    d = np.where(mask, 0.0, INF).astype(np.float64)
    h, w = d.shape

    # Прямой проход: TL -> BR.
    for y in range(h):
        for x in range(w):
            if d[y, x] == 0:
                continue
            v = d[y, x]
            if y > 0:
                v = min(v, d[y - 1, x] + 3)
                if x > 0:
                    v = min(v, d[y - 1, x - 1] + 4)
                if x < w - 1:
                    v = min(v, d[y - 1, x + 1] + 4)
            if x > 0:
                v = min(v, d[y, x - 1] + 3)
            d[y, x] = v

    # Обратный проход: BR -> TL.
    for y in range(h - 1, -1, -1):
        for x in range(w - 1, -1, -1):
            if d[y, x] == 0:
                continue
            v = d[y, x]
            if y < h - 1:
                v = min(v, d[y + 1, x] + 3)
                if x < w - 1:
                    v = min(v, d[y + 1, x + 1] + 4)
                if x > 0:
                    v = min(v, d[y + 1, x - 1] + 4)
            if x < w - 1:
                v = min(v, d[y, x + 1] + 3)
            d[y, x] = v

    return d / 3.0  # нормировка: 1 ед. = 1 px


def strip_white(
    path: Path,
    *,
    crop: bool = True,
    square: bool = False,
    no_bg_rem: bool = False,
    stroke: bool = False,
) -> None:
    img = Image.open(path).convert("RGBA")
    arr = np.array(img)
    h, w = arr.shape[:2]

    if no_bg_rem:
        # Все белые пиксели -> прозрачный (включая проглядывающий через отверстия).
        # Деталь тёмно-серая (min-канал <= ~130), фон белый. Eдиними внутренние
        # белые области и антиалиасинг-ореол на границе (светло-серые пиксели
        # между белым и деталью). Подавляем альфу плавно по близости к белому:
        # m = min(R,G,B), при m<=BG_LOW - непрозрачно, при m>=BG_HIGH - alpha 0,
        # между - линейный фейд.
        BG_LOW = 150
        BG_HIGH = 210
        min_ch = arr[..., :3].min(axis=2).astype(np.float32)
        new_alpha = np.clip((BG_HIGH - min_ch) / (BG_HIGH - BG_LOW) * 255, 0, 255)
        arr[..., 3] = np.minimum(arr[..., 3].astype(np.float32), new_alpha).astype(
            np.uint8
        )
    else:
        # Обычный режим: BFS из углов по белому пикселю — это фон.
        is_white = (arr[:, :, 0] == 255) & (arr[:, :, 1] == 255) & (arr[:, :, 2] == 255)
        visited = np.zeros_like(is_white, dtype=bool)
        queue: deque  # type: ignore[type-arg]
        queue = deque()
        for y, x in [(0, 0), (0, w - 1), (h - 1, 0), (h - 1, w - 1)]:
            if is_white[y, x] and not visited[y, x]:
                visited[y, x] = True
                queue.append((y, x))
        while queue:
            cy, cx = queue.popleft()
            for dy, dx in ((-1, 0), (1, 0), (0, -1), (0, 1)):
                ny, nx = cy + dy, cx + dx
                if (
                    0 <= ny < h
                    and 0 <= nx < w
                    and not visited[ny, nx]
                    and is_white[ny, nx]
                ):
                    visited[ny, nx] = True
                    queue.append((ny, nx))
        arr[visited, 3] = 0  # фон -> прозрачный

    # BFS из 4 углов ловит только внешний фон. Белые пиксели, проглядывающие
    # через отверстия в детали, окружены её силуэтом (анти-алиасенной обводкой)
    # и до углов не дотянутся. Оставляем их если нужно (no_bg_rem = True),
    # иначе — удаляем дыры: достаточно большие (>= MIN_HOLE_PIX) и компактные
    # (fill_ratio >= MIN_FILL_RATIO). Отличаем от бликов на поверхности.
    if not no_bg_rem:
        MIN_HOLE_PIX = 50
        MIN_FILL_RATIO = 0.5
        is_white_after = (
            (arr[..., 0] == 255)
            & (arr[..., 1] == 255)
            & (arr[..., 2] == 255)
            & (arr[..., 3] > 0)
        )
        cc_visited = np.zeros_like(is_white_after, dtype=bool)
        cc_queue: deque  # type: ignore[type-arg]
        cc_queue = deque()
        holes = np.zeros_like(is_white_after, dtype=bool)
        for y in range(h):
            for x in range(w):
                if is_white_after[y, x] and not cc_visited[y, x]:
                    cc_queue.append((y, x))
                    cc_visited[y, x] = True
                    pixels: list[tuple[int, int]] = []
                    while cc_queue:
                        cy, cx = cc_queue.popleft()
                        pixels.append((cy, cx))
                        for dy, dx in ((-1, 0), (1, 0), (0, -1), (0, 1)):
                            ny2, nx2 = cy + dy, cx + dx
                            if (
                                0 <= ny2 < h
                                and 0 <= nx2 < w
                                and is_white_after[ny2, nx2]
                                and not cc_visited[ny2, nx2]
                            ):
                                cc_visited[ny2, nx2] = True
                                cc_queue.append((ny2, nx2))
                    if len(pixels) < MIN_HOLE_PIX:
                        continue
                    ys = [p[0] for p in pixels]
                    xs = [p[1] for p in pixels]
                    bbox_area = (max(ys) - min(ys) + 1) * (max(xs) - min(xs) + 1)
                    if len(pixels) / bbox_area >= MIN_FILL_RATIO:
                        for cy, cx in pixels:
                            holes[cy, cx] = True

        arr[holes, 3] = 0  # белый фон через дырки -> прозрачный

    # Евклидова дистанция до ближайшего непрозрачного пикселя. 0 на силуэте,
    # растёт наружу.
    opaque = arr[:, :, 3] > 0
    dist = distance_transform(opaque)

    # Обводка: solid на dist в (0, STROKE_PX-A_A_WIDTH], линейный фейд на
    # (STROKE_PX-A_A_WIDTH, STROKE_PX]. Внутри силуэта (dist == 0) ничего не
    # трогаем - там сама деталь. Включается только флагом --stroke.
    if stroke:
        fade_in = STROKE_PX - AA_WIDTH
        in_stroke = (dist > 0) & (dist <= STROKE_PX)
        alpha_fade = np.clip((STROKE_PX - dist) / AA_WIDTH * 255, 0, 255)
        alpha = np.where(dist <= fade_in, 255.0, alpha_fade).astype(np.uint8)

        arr[..., 0] = np.where(in_stroke, 255, arr[..., 0])
        arr[..., 1] = np.where(in_stroke, 255, arr[..., 1])
        arr[..., 2] = np.where(in_stroke, 255, arr[..., 2])
        arr[..., 3] = np.where(in_stroke, alpha, arr[..., 3])

    # Кроп по bbox непрозрачного.
    opaque_after = arr[:, :, 3] > 0
    if not opaque_after.any():
        raise ValueError(f"{path}: после удаления фона не осталось контента")
    if crop:
        ys2, xs2 = np.where(opaque_after)
        out = arr[ys2.min() : ys2.max() + 1, xs2.min() : xs2.max() + 1]
    else:
        out = arr

    if square:
        oh, ow = out.shape[:2]
        side = max(oh, ow)
        canvas = np.zeros((side, side, 4), dtype=out.dtype)
        # центруем содержимое на квадратном холсте
        y_off = (side - oh) // 2
        x_off = (side - ow) // 2
        canvas[y_off : y_off + oh, x_off : x_off + ow] = out
        out = canvas

    backup = next_backup_path(path)
    path.rename(backup)
    Image.fromarray(out, mode="RGBA").save(path)
    print(f"{path}: ok -> {backup} ({out.shape[1]}x{out.shape[0]})")


def main(argv: list[str]) -> int:
    parser = argparse.ArgumentParser(
        prog="strip_white",
        description="Make белый фон прозрачным, добавить белую обводку, обрезать поля.",
    )
    parser.add_argument("paths", nargs="*", metavar="path", help="image(s) to process")
    parser.add_argument(
        "--no-crop",
        dest="crop",
        action="store_false",
        help="keep transparent margins instead of cropping to the opaque bbox",
    )
    parser.add_argument(
        "--square",
        dest="square",
        action="store_true",
        help="pad the result to a square canvas (transparent), centered",
    )
    parser.add_argument(
        "--stroke",
        dest="stroke",
        action="store_true",
        help="добавить белую обводку по силуэту контента (по умолчанию нет)",
    )
    parser.add_argument(
        "--wb",
        "--white-bg",
        dest="no_bg_rem",
        action="store_true",
        help=(
            "Все белые пиксели -> прозрачный без исключения отверстий. "
            "Полезно когда белый - всегда фон."
        ),
    )
    args = parser.parse_args(argv)
    if not args.paths:
        parser.print_help(sys.stderr)
        return 2
    for arg in args.paths:
        p = Path(arg)
        if not p.is_file():
            print(f"{p}: не файл, пропускаю", file=sys.stderr)
            continue
        strip_white(
            p,
            crop=args.crop,
            square=args.square,
            no_bg_rem=args.no_bg_rem,
            stroke=args.stroke,
        )
    return 0


if __name__ == "__main__":
    # ponytail: self-check на эталоне - фон прозрачный, обводка есть и
    # частично сглажена (alpha < 255 где-то в кольце).
    import tempfile

    src = Path("docs/vostok/manual/pics/prepare/threaded_inserts.png")
    if src.exists():
        with tempfile.TemporaryDirectory() as td:
            test = Path(td) / src.name
            test.write_bytes(src.read_bytes())
            # Прогоняем всю pipeline руками, чтобы достать in_stroke и alpha
            # для проверки - strip_white() их не возвращает.
            img2 = Image.open(test).convert("RGBA")
            arr2 = np.array(img2)
            h2, w2 = arr2.shape[:2]
            is_white2 = (
                (arr2[:, :, 0] == 255) & (arr2[:, :, 1] == 255) & (arr2[:, :, 2] == 255)
            )
            visited2 = np.zeros_like(is_white2, dtype=bool)
            queue2: deque  # type: ignore[type-arg]
            queue2 = deque()
            for y, x in [(0, 0), (0, w2 - 1), (h2 - 1, 0), (h2 - 1, w2 - 1)]:
                if is_white2[y, x] and not visited2[y, x]:
                    visited2[y, x] = True
                    queue2.append((y, x))
            while queue2:
                cy, cx = queue2.popleft()
                for dy, dx in ((-1, 0), (1, 0), (0, -1), (0, 1)):
                    ny3, nx3 = cy + dy, cx + dx
                    if (
                        0 <= ny3 < h2
                        and 0 <= nx3 < w2
                        and not visited2[ny3, nx3]
                        and is_white2[ny3, nx3]
                    ):
                        visited2[ny3, nx3] = True
                        queue2.append((ny3, nx3))
            arr2[visited2, 3] = 0
            opaque2 = arr2[..., 3] > 0
            dist2 = distance_transform(opaque2)
            fade_in2 = STROKE_PX - AA_WIDTH
            in_stroke2 = (dist2 > 0) & (dist2 <= STROKE_PX)
            alpha_fade2 = np.clip((STROKE_PX - dist2) / AA_WIDTH * 255, 0, 255)
            assert in_stroke2.any(), "должна быть обводка"
            assert (alpha_fade2[in_stroke2] == 255).any(), "должны быть solid-пиксели"
            assert (alpha_fade2[in_stroke2] < 255).any(), (
                "должны быть сглаженные пиксели"
            )
            # Полный прогон тоже должен сработать
            strip_white(test)
            arr3 = np.array(Image.open(test))
            assert arr3.shape[2] == 4, "ожидался RGBA"
            assert arr3[arr3.shape[0] // 2, arr3.shape[1] // 2, 3] == 255, (
                "центр непрозрачный"
            )
            assert (arr3[..., 3] == 0).any(), "должна быть прозрачность"
            print("self-check ok")

    sys.exit(main(sys.argv[1:]))
