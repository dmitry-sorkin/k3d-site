#!/usr/bin/env python3
"""Извлекает посты из темы "Галерея" Telegram-экспорта и раскладывает их по
`docs/vostok/gallery/posts/post-NNN/`.

Использование:
    python tools/gallery/extract.py --json .tmp/vostok_chat_export/result.json
                                  --topic-id 361053
                                  --out docs/vostok/gallery/posts

Структура выходной папки поста:
    caption.txt          # метаданные + текст сообщения
    photo-01.jpg
    video-01.mp4
    ...

После извлечения пересоберите галерею:
    python tools/gallery/build.py
"""

from __future__ import annotations

import argparse
import json
import re
import shutil
import sys
from datetime import datetime
from pathlib import Path
from typing import Any

DEFAULT_CHAT_JSON = ".tmp/vostok_chat_export/result.json"
DEFAULT_TOPIC_ID = 361053
DEFAULT_OUT = "docs/vostok/gallery/posts"


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Extract gallery posts from Telegram chat export")
    parser.add_argument("--json", default=DEFAULT_CHAT_JSON, help="Path to result.json")
    parser.add_argument("--topic-id", type=int, default=DEFAULT_TOPIC_ID, help="First message ID of the gallery topic")
    parser.add_argument("--out", default=DEFAULT_OUT, help="Output directory for posts")
    return parser.parse_args()


def load_json(path: Path) -> dict:
    with open(path, "r", encoding="utf-8") as f:
        return json.load(f)


def text_entities_to_markdown(entities: list[dict[str, Any]] | str) -> str:
    """Превращает text_entities Telegram-экспорта в plain Markdown."""
    if isinstance(entities, str):
        return entities
    parts = []
    for e in entities:
        if isinstance(e, str):
            parts.append(e)
            continue
        t = e.get("text", "")
        tp = e.get("type", "plain")
        if tp == "text_link":
            parts.append(f"[{t}]({e.get('href', '')})")
        elif tp == "link":
            parts.append(f"<{t}>")
        elif tp == "bold":
            parts.append(f"**{t}**")
        elif tp == "italic":
            parts.append(f"*{t}*")
        elif tp == "code":
            parts.append(f"`{t}`")
        elif tp == "pre":
            parts.append(f"```\n{t}\n```")
        else:
            parts.append(t)
    return "".join(parts)


def get_message_text(m: dict) -> str:
    text = m.get("text", "")
    if isinstance(text, list):
        return text_entities_to_markdown(text)
    if isinstance(text, str) and not text and "text_entities" in m:
        return text_entities_to_markdown(m["text_entities"])
    return text


def media_files(m: dict, base_dir: Path) -> list[tuple[Path, str]]:
    """Возвращает список (абсолютный путь к медиа, имя для сохранения)."""
    result = []
    # фото
    photo = m.get("photo")
    if photo:
        src = base_dir / photo
        if src.exists():
            result.append((src, f"photo-{len(result)+1:02d}{src.suffix.lower()}"))
    # видео/файлы
    media_type = m.get("media_type", "")
    file_path = m.get("file")
    if file_path and media_type in {"video_file", "video_message", "animation"}:
        src = base_dir / file_path
        if src.exists():
            ext = src.suffix.lower() or ".mp4"
            result.append((src, f"video-{len(result)+1:02d}{ext}"))
    # альбомы — не встречаются в пересланных постах, но на всякий случай
    if "photo" not in m and "media_type" not in m and "file" in m:
        # стикеры и прочее игнорируем
        pass
    return result


def extract(args: argparse.Namespace) -> int:
    json_path = Path(args.json)
    if not json_path.exists():
        print(f"Файл не найден: {json_path}", file=sys.stderr)
        return 1

    base_dir = json_path.parent
    out_dir = Path(args.out)
    out_dir.mkdir(parents=True, exist_ok=True)

    data = load_json(json_path)
    messages = data.get("messages", [])
    topic_id = args.topic_id

    # Собираем сообщения темы: либо reply_to_top совпадает с topic_id,
    # либо это само первое сообщение темы.
    topic_messages = [
        m for m in messages
        if m.get("id") == topic_id
        or m.get("reply_to_top") == topic_id
        or m.get("reply_to_message_id") == topic_id
    ]

    # Также, если в сообщении есть forward из канала, оно может быть вложено
    # в другую тему? Пока оставляем reply_to_top как основной критерий.
    if not topic_messages:
        print(f"Сообщения темы {topic_id} не найдены в {json_path}", file=sys.stderr)
        return 1

    # Сортируем по дате
    topic_messages.sort(key=lambda m: m.get("date", ""))

    # Удаляем старые посты? Нет — пусть пользователь сам решает.
    # Нумеруем заново, начиная с 001.
    existing = sorted(d for d in out_dir.iterdir() if d.is_dir() and d.name.startswith("post-"))
    start_num = len(existing)

    for i, m in enumerate(topic_messages, start=1):
        text = get_message_text(m).strip()
        if not text and not media_files(m, base_dir):
            continue  # пустые сообщения пропускаем

        num = start_num + i
        post_dir = out_dir / f"post-{num:03d}"
        post_dir.mkdir(exist_ok=True)

        # метаданные
        author = m.get("from", "Автор не указан")
        date = m.get("date", "")[:10]
        msg_id = str(m.get("id", ""))

        caption_lines = [
            f"<!-- version: - -->",
            f"<!-- size: - -->",
            f"<!-- hotends: - -->",
            f"<!-- author: {author} -->",
            f"<!-- date: {date} -->",
            f"<!-- first_id: {msg_id} -->",
            "",
        ]
        caption_lines.append(text)
        (post_dir / "caption.txt").write_text("\n".join(caption_lines), encoding="utf-8")

        # медиа
        for src, dst_name in media_files(m, base_dir):
            shutil.copy2(src, post_dir / dst_name)

    print(f"Извлечено {len(topic_messages)} сообщений темы {topic_id} в {out_dir}")
    return 0


if __name__ == "__main__":
    raise SystemExit(extract(parse_args()))
