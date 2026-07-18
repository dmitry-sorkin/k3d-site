# Работа с галереей VOSTOK

Галерея — это одна страница `docs/vostok/gallery/index.md`, на которой показывается один пост за раз. Переключение между постами происходит через правую панель оглавления (TOC), как будто это обычные заголовки статьи.

## Структура

```
docs/vostok/gallery/
  index.md          # собирается автоматически
  posts/
    post-001/
      caption.txt   # текст и метаданные поста
      photo-01.jpg
      video-01.mp4
      ...
    post-002/
      ...
```

## caption.txt

Каждый пост начинается с HTML-комментариев — метаданных:

```
<!-- version: 9.x -->
<!-- size: M (400x250) -->
<!-- hotends: CHC Pro -->
<!-- author: Иван Иванов -->
<!-- date: 2025-08-16 -->
<!-- first_id: 361174 -->

Текст сообщения. Может занимать несколько строк и содержать Markdown-ссылки.
```

- `version` — версия VOSTOK (`-`, если неизвестна).
- `size` — типоразмер (`-`, если неизвестен).
- `hotends` — хотэнды (`-`, если неизвестны).
- `author` — автор поста, попадает в заголовок и TOC.
- `date` — дата в формате `YYYY-MM-DD`.
- `first_id` — ID сообщения в Telegram. Используется для ссылки `https://t.me/k3d_vostok/<first_id>`.

После последнего комментария идёт пустая строка, затем текст сообщения.

## Добавление нового поста вручную

1. Создайте папку `docs/vostok/gallery/posts/post-NNN/` (следующий свободный номер).
2. Положите в неё фото и/или видео:
   - фото: `photo-01.jpg`, `photo-02.jpg`, ...
   - видео: `video-01.mp4`, `video-02.mp4`, ...
3. Напишите `caption.txt` с метаданными и текстом.
4. Пересоберите `index.md`:
   ```bash
   python tools/gallery/build.py
   ```
5. Проверьте сборку сайта:
   ```bash
   properdocs build --clean
   ```

## Массовое извлечение из Telegram-экспорта

Если есть свежий JSON-экспорт чата (`result.json`):

```bash
python tools/gallery/extract.py --json .tmp/vostok_chat_export/result.json --topic-id 361053
python tools/gallery/build.py
properdocs build --clean
```

Скрипт `extract.py` найдёт все сообщения темы `361053` и раскладывает их по папкам `post-NNN/`.

## Как это работает

- Каждый пост — это `<article class="gallery-post">` с заголовком `<h2 id="post-NNN">`.
- CSS в `docs/stylesheets/gallery.css` скрывает все посты, кроме выбранного (через `:target` и `:has()`).
- Минимальный inline-скрипт в `index.md` дублирует логику переключения для старых браузеров без `:has()`.
- Правая панель TOC Material автоматически строится из заголовков `h2` и ссылается на `#post-NNN`.
