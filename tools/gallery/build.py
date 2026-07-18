#!/usr/bin/env python3
"""Генерирует docs/vostok/gallery/index.md из папок posts/post-*/caption.txt и медиа.

Использование:
    python scripts/build_gallery.py

Структура поста:
    docs/vostok/gallery/posts/post-<NNN>/
        caption.txt     # метаданные в HTML-комментариях + текст сообщения
        photo-01.jpg
        video-01.mp4
        ...

caption.txt должен содержать в начале строки вида:
    <!-- author: Имя Автора -->
    <!-- date: 2025-08-16 -->
    <!-- first_id: 361174 -->
    <!-- version: 9.x -->
    <!-- size: M (400x250) -->
    <!-- hotends: CHC Pro -->

Затем идёт текст сообщения (может занимать несколько строк).
"""

from __future__ import annotations

import re
import sys
from pathlib import Path

GALLERY_DIR = Path("docs/vostok/gallery")
POSTS_DIR = GALLERY_DIR / "posts"
INDEX_PATH = GALLERY_DIR / "index.md"

FRONTMATTER = """---
authors:
  - sorkin
icon: material/image-multiple
title: Галерея
description: Принтеры VOSTOK, собранные участниками сообщества
show_created: false
show_updated: false
show_author: false
---

# 

<div class="gallery" markdown>

"""

FOOTER = """</div>

<script>
// Переключение постов по якорю в URL. Показывается либо пост, заголовок
// которого указан в hash, либо первый пост при открытии страницы без якоря.
(function () {
  function update() {
    var posts = document.querySelectorAll('.gallery-post');
    if (!posts.length) return;
    var id = window.location.hash.replace('#', '');
    var target = id ? document.getElementById(id) : null;
    var active = target ? target.closest('.gallery-post') : null;
    for (var i = 0; i < posts.length; i++) {
      posts[i].classList.toggle('active', posts[i] === active);
    }
    if (!active && posts[0]) posts[0].classList.add('active');
  }
  document.addEventListener('DOMContentLoaded', update);
  window.addEventListener('hashchange', update);
})();
</script>
"""


def parse_caption(path: Path) -> dict:
    text = path.read_text(encoding="utf-8")
    meta: dict[str, str] = {}
    comment_re = re.compile(r"^\s*<!--\s*(\w+)\s*:\s*(.*?)\s*-->\s*$")
    lines = text.splitlines()
    body_start = 0
    for i, line in enumerate(lines):
        m = comment_re.match(line)
        if m:
            key, value = m.group(1).lower(), m.group(2).strip()
            meta[key] = value
            body_start = i + 1
        else:
            break
    body = "\n".join(lines[body_start:]).strip()
    return {
        "author": meta.get("author", "Автор не указан"),
        "date": meta.get("date", ""),
        "first_id": meta.get("first_id", ""),
        "version": meta.get("version", "-"),
        "size": meta.get("size", "-"),
        "hotends": meta.get("hotends", "-"),
        "body": body,
    }


def media_markup(post_dir: Path) -> str:
    """Возвращает markdown-разметку медиа файлов поста в виде grid cards.

    Каждый файл оформляется как отдельный пункт списка, чтобы grid cards
    рендерился в <ul><li>...</li></ul> и CSS-сетка работала корректно.
    Возвращает строку с уже обёрнутым <div class="grid ...">...первая
    ячейка подставляется отдельно, поэтому оставляем только медиа.
    """
    files = sorted(
        p for p in post_dir.iterdir()
        if p.is_file() and p.name.lower() not in ("caption.txt", "index.md", "readme.md")
    )

    if not files:
        return ""

    cols = "cols-3"

    lines = [f'<div class="grid cards no-gap {cols}" markdown>', ""]
    for f in files:
        rel = f"posts/{post_dir.name}/{f.name}"
        suffix = f.suffix.lower()
        if suffix in {".mp4", ".mov", ".webm"}:
            lines.append(f'- <video src="{rel}" controls preload="none"></video>')
        elif suffix in {".jpg", ".jpeg", ".png", ".webp", ".gif"}:
            alt = post_dir.name
            lines.append(f'- ![{alt}]({rel}){{ loading="lazy" }}')
        else:
            lines.append(f'- [{f.name}]({rel})')
    lines.append("")
    lines.append("</div>")
    return "\n".join(lines)


def escape_md(text: str) -> str:
    # Не экранируем специально, чтобы сохранить ссылки и форматирование из текста.
    return text


def render_post(post_dir: Path) -> str:
    caption = parse_caption(post_dir / "caption.txt")
    anchor = post_dir.name
    author = caption["author"]
    date = caption["date"]
    first_id = caption["first_id"]
    version = caption["version"]
    size = caption["size"]
    hotends = caption["hotends"]
    body = caption["body"]

    telegram_link = f"https://t.me/k3d_vostok/{first_id}" if first_id else ""

    # Считаем количество медиа. Сетка всегда 3 ячейки по горизонтали:
    # при нехватке медиа оставшиеся ячейки в последнем ряду остаются пустыми.
    files = sorted(
        f for f in post_dir.iterdir()
        if f.is_file() and f.name.lower() not in ("caption.txt", "index.md", "readme.md")
    )
    cols = "cols-3"

    lines = [
        '<article class="gallery-post" markdown>',
        f"## {author} {{ #{anchor} }}",
        "",
        "---",
        "",
        f'<div class="grid cards no-gap {cols}" markdown>',
        "",
        "-  ",
        "",
        '    !!! info ""',
        "",
        f"        - Автор: {author}",
        f"        - Дата: {date or '-'}",
    ]
    if telegram_link:
        lines.append(f"        - [Оригинал в Telegram]({telegram_link}){{ target=\"_blank\" }}")
    else:
        lines.append("        - Оригинал в Telegram: -")
    lines.append(f"        - Версия: {version}")
    lines.append(f"        - Типоразмер: {size}")
    lines.append(f"        - Хотэнды: {hotends}")
    if body:
        body_br = "<br>".join(body.splitlines())
        lines.append("")
        lines.append("        - Комментарий автора: (1)")
        lines.append("          { .annotate }")
        lines.append("")
        lines.append(f"              1. {body_br}")
    lines.append("")
    for f in files:
        rel = f"posts/{post_dir.name}/{f.name}"
        suffix = f.suffix.lower()
        if suffix in {".mp4", ".mov", ".webm"}:
            lines.append(f'- <video src="{rel}" controls preload="none"></video>')
        elif suffix in {".jpg", ".jpeg", ".png", ".webp", ".gif"}:
            alt = post_dir.name
            lines.append(f'- ![{alt}]({rel}){{ loading="lazy" }}')
        else:
            lines.append(f'- [{f.name}]({rel})')
    lines.append("")
    lines.append("</div>")
    lines.append("")
    lines.append("</article>")
    lines.append("")
    return "\n".join(lines)


def main() -> int:
    if not POSTS_DIR.is_dir():
        print(f"Папка с постами не найдена: {POSTS_DIR}", file=sys.stderr)
        return 1

    post_dirs = sorted(
        d for d in POSTS_DIR.iterdir()
        if d.is_dir() and d.name.startswith("post-")
    )
    # Сортируем по числовой части, если она есть, иначе по имени.
    def sort_key(d: Path):
        m = re.search(r"(\d+)", d.name)
        return (int(m.group(1)) if m else 0, d.name)
    post_dirs.sort(key=sort_key, reverse=True)

    if not post_dirs:
        print("Посты не найдены.", file=sys.stderr)
        return 1

    parts = [FRONTMATTER]
    for post_dir in post_dirs:
        parts.append(render_post(post_dir))
    parts.append(FOOTER)

    INDEX_PATH.write_text("\n".join(parts), encoding="utf-8")
    print(f"Сгенерирован {INDEX_PATH} ({len(post_dirs)} постов)")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
