# AGENTS

Russian-language documentation site for the K3D 3D-printing community, built with a custom-named MkDocs Material toolchain.

## Build

- Config: `properdocs.yml` (NOT `mkdocs.yml`). Build tool is the `properdocs` CLI, a renamed mkdocs.
- Source: `docs/`. Output: `site/` (gitignored).
- Python deps live in `.github/workflows/requirements.txt` (mkdocs 1.6.1, properdocs 1.6.7, mkdocs-materialx, glightbox, redirects, charts, table-reader, document-dates, pillow).
- Local build:

  ```
  python -m venv .venv && .venv/Scripts/activate
  pip install -r .github/workflows/requirements.txt
  properdocs build --clean
  ```

  Dev server: `properdocs serve` (port 8000 by default). No tests, no linter — this is a static site.
- Deploy: `.github/workflows/main.yaml` builds on push/PR to `main`, `build-stage-python`, `test-lolo`; deploys via rsync-over-SSH only on push to `main`. Secrets: `SSH_PRIVATE_KEY`, `SSH_USERNAME`, `SSH_SERVER`, `DEPLOY_PATH`.

## Layout

```
docs/
  index.md            # landing page, no nav
  authors.yml         # author registry (sorkin, makarov, alexander)
  part-navi/          # 3D recommender (printers, extruder, electronics, mechanics, filament, dryers, consumables)
  vostok/             # K3D VOSTOK printer: index, faq, releases, manual/, gallery/, v8/ (archive)
  calibrations/       # PID, IS, flow, accuracy/, VFA, retractions/, la/ (PA), max_flowrate/
  projects/           # ps_translation, custom_pulleys, printheads (ebp/ehp/esp), feeders (bemege/minifeeder/feeder965)
  printers/           # per-printer mods (kobra3, k1, neptune3pro/4, kp3s, qidi_q1_pro, qidi_q2)
  articles/           # long-form posts (print_strong, custom_gcode, nozzle_test_2024, ultra_high_flow_hotends, ad5m_ad5x_hotend_test, cooling_problems)
  redirects/          # stub .md files whose real target is in properdocs.yml `redirect_maps`
  pics/, assets/, models/   # images and 3D models referenced from MD
  stylesheets/        # extra9.css referenced from theme
  javascripts/        # mathjax loader
overrides/            # Material theme partial overrides (partials/integrations/, .icons/)
tools/gallery/           # tracked build helpers (gallery/extract.py, gallery/build.py, gallery/README.md)
tools/affiliate.py       # epn.bz: превращает прямые ссылки на AliExpress в аффилиатные (одна страница за запуск)
tools/strip_white.py     # image prep for the VOSTOK part viewplaceholders
site/                 # build output, gitignored
.github/workflows/    # CI
```

Nav structure is defined explicitly in `properdocs.yml` under `nav:` — folder names mirror nav sections but you must add new pages to both.

## Conventions

- All content is Russian. Keep filenames ASCII; titles and headings in Russian.
- Every page starts with YAML frontmatter. Minimum expected keys (see `docs/index.md`):

  ```
  ---
  authors:
    - sorkin
  title: ...
  description: ...
  ---
  ```

  Optional: `hide: [navigation, toc]`, `show_created`, `show_updated`, `show_author`. Authors must exist in `docs/authors.yml`.
- Use Material admonitions and these custom ones defined in config: `fire`, `speed`, `time`, `clone` (icons: fire, speedometer, clock-outline, mirror).
- Mermaid, Vega-Lite, MathJax, tabs, keys, superfences, critic/caret/mark/tilde are enabled — use them instead of screenshots or hand-drawn diagrams.
- Inline images: place in the same section's folder or in `docs/pics/`, reference with relative paths.

### Mermaid diagrams

- **Always set `labelColor` and `lineColor` to the theme CSS variables** in the diagram's `themeVariables` frontmatter — otherwise colors don't follow light/dark mode and are unreadable on dark backgrounds. The safe template for any mermaid block:

  ````
  ```mermaid
  ---
  config:
      themeVariables:
          labelColor: 'var(--md-mermaid-label-fg-color)'
          lineColor: 'var(--md-mermaid-label-fg-color)'
  ---
  ...diagram...
  ```
  ````

  For `treeView` diagrams the variables go under `themeVariables.treeView` (they are not inherited from the top level):

  ````
  ```mermaid
  ---
  config:
      themeVariables:
          treeView:
              labelColor: 'var(--md-mermaid-label-fg-color)'
              lineColor: 'var(--md-mermaid-label-fg-color)'
  ---
  treeView-beta
  ...
  ```
  ````

- **TreeView icons (Material, `mdi` pack)**: supported via `docs/javascripts/mermaid-icons.js` (registered in `extra_javascript`). It registers the `mdi` icon pack with `mermaid.registerIconPacks()` and forces `securityLevel: "loose"` (mermaid's strict sanitizer strips the `<use>` elements icons are rendered with — without this patch icons show as empty space). To use an icon on a node, append `icon(mdi:name)` after the label, e.g. `"vostok" icon(mdi:folder)`. Common icons: `mdi:folder`, `mdi:file`, `mdi:table`, `mdi:cube`, `mdi:file-document`, `mdi:file-image`. Icons are fetched on demand from unpkg (`@iconify-json/mdi`), so the site must be online.
- Mermaid trims `\s`-class whitespace (including `&nbsp;` U+00A0) from the start/end of labels. To pad a label to equal width, use the `&nbsp;` entity — it survives trimming as ASCII and decodes to a real space at render time.

## Redirects

If you move a page, add a mapping in `properdocs.yml` under `plugins.redirects.redirect_maps` AND keep the old stub in `docs/redirects/` so links from external sites don't 404. External links (forum, Telegram, bugtracker) live in `redirects/` with a single external `target=` link in frontmatter.

## Theme overrides

- Custom dir is `overrides/`. Material will pick up `partials/<name>.html` automatically.
- Extra CSS goes in `docs/stylesheets/` and is registered in `properdocs.yml` under `extra_css`.
- Extra JS goes in `docs/javascripts/` (registered under `extra_javascript`) — CDN scripts are listed inline in the config.

## Don't

- Don't rename `properdocs.yml` to `mkdocs.yml` — the `properdocs` package entry-point expects this name.
- Don't commit `site/`, `.cache/`, `.venv/`, `.obsidian/workspace`, `*.ods#`, lock files, or `.vs/`, `.idea/`, `.continue/`.
- Don't add plugins/extensions without updating `.github/workflows/requirements.txt` and the `Build` section above.
- Don't create top-level dirs in `docs/`; new sections belong under one of the existing nav folders, and must be added to `nav:` in `properdocs.yml`.

## VOSTOK build gallery (`docs/vostok/gallery/`)

The gallery is a single Material page that shows one post at a time, with the right-hand TOC acting as the post picker.

Structure:

```
docs/vostok/gallery/
  index.md                        # regenerated by tools/gallery/build.py
  posts/
    post-NNN/                     # NNN = 001, 002, ..., ordered by post date
      caption.txt                 # raw caption + author/date/first_id metadata
      photo-NN.jpg, video-NN.mp4  # media
tools/gallery/
  extract.py                      # parses .tmp/vostok_chat_export/... → posts/
  build.py                        # assembles index.md from posts/ folders
  README.md                       # how to add posts manually or from Telegram export
```

To rebuild from a fresh Telegram export (re-running `extract.py` is safe, it only adds new numbered folders):

```bash
python tools/gallery/extract.py
python tools/gallery/build.py
python -m properdocs build --clean
```

To add a single post by hand when there is no fresh export: create `posts/post-NNN/`, drop the media in, write `caption.txt` with the template in `tools/gallery/README.md`, then re-run `build.py`. The `<!-- author: ... -->`, `<!-- date: ... -->`, `<!-- first_id: ... -->` lines in `caption.txt` are stripped at build time and feed the per-post header card + Telegram link.

The gallery uses CSS for one-post-at-a-time display: each post is `<article class="gallery-post">` with an `h2[id^="post-"]`, and the rules in `docs/stylesheets/extra9.css` (block «Галерея сборок VOSTOK») hide every article except the one whose H2 matches the URL fragment (via `:has()` + `:target`). A tiny inline script in `index.md` provides the same behavior as a fallback for browsers without `:has()` support.

## Image prep (`tools/strip_white.py`)

One-shot script for the part-viewplaceholders in `docs/vostok/manual/prepare.md`. Takes an opaque-PNG (white background, the part on it), returns an RGBA where the background is transparent, a thin anti-aliased white stroke hugs the object's silhouette, and empty margins are cropped. Reference output: `docs/vostok/manual/pics/prepare/threaded_inserts.png`.

```
python tools/strip_white.py <path> [<path> ...]
```

Pipeline:

1. 4-connected BFS from the 4 corners over pure-white pixels marks the outer background → alpha 0. White highlights inside the part are not reachable from the corners, so they stay opaque.
2. Hole removal: connected components of remaining opaque white pixels with size ≥ `MIN_HOLE_PIX` (50) and fill-ratio ≥ `MIN_FILL_RATIO` (0.5) are removed. These are the white pixels visible through holes in the part — the part's own outline around the hole is an anti-aliased gray ring that the corner BFS can't reach, so the inner white stays opaque. The size + fill-ratio heuristic distinguishes them from white highlights on the part's surface (which are small and irregular).
3. Two-pass chamfer-3-4 distance transform (forward TL→BR, backward BR→TL, orthogonal cost 3, diagonal cost 4, normalized by 3) gives an approximate Euclidean distance field from every transparent pixel to the nearest opaque one. Chamfer 3-4 is uniform across directions — unlike the square `MaxFilter` dilation it replaces, which makes diagonal strokes √2 thicker than orthogonal ones.
4. Stroke region: `0 < dist ≤ STROKE_PX`. Solid alpha 255 for `dist ≤ STROKE_PX - AA_WIDTH`, linear fade to 0 across the last `AA_WIDTH` pixels. Knobs at the top of the file: `STROKE_PX = 4.0`, `AA_WIDTH = 2.0`.
5. Crop to the bounding box of the new opaque region.

Originals are renamed `.bak1`, `.bak2`, … (first free index) inside a `.bak/` subfolder next to the source before the new file is written. The script is **not** idempotent — a second run treats the previous stroke as part of the foreground and grows the silhouette outward by `STROKE_PX + AA_WIDTH`. Original always survives in `.bak/<name>.bak1`. Pass `--no-crop` to skip the margin crop in step 5, and `--square` to pad the result to a square canvas (transparent margins, content centered) for use in `grid cards` layouts where the aspect ratio would otherwise misalign tiles.

## Affiliate links (`tools/affiliate.py`)

Turns bare marketplace links pasted into a page into affiliate links via the epn.bz API (`https://epn.bz/api/ru/`, API v2). Run this when the user says «сделай ссылки аффилиатными», «обнови аффилиатные ссылки», or pastes raw AliExpress links into a page.

```bash
python tools/affiliate.py docs/part-navi/filament_dryers.md
```

**Hard rules — do not work around them:**

- **Exactly one page per run.** The script takes one path, must live under `docs/`, and refuses zero, multiple, or directory arguments. Never loop over files, never pass `docs/`, never try to convert the whole site. If the user asks for several pages, run the script once per page, one at a time, and check the report of each run before starting the next.
- **Only AliExpress is converted.** Only links whose host is `aliexpress.ru` / `aliexpress.com` / `aliexpress.us` are touched (`/aff/redirect` paths are skipped). Every other marketplace link — Ozon, Яндекс.Маркет, DNS, ChipDip, krepcom, vseinstrumenti — is deliberately left exactly as it is, even if the user expects otherwise. Why: EPN's product offer for Яндекс.Маркет is `disabled`, there is no product offer for Ozon at all, and `/offers/by-link` matches by domain and returns a wrong offer for both (`offer 2423` «Яндекс Браузер», `offer 2327` «Ozon Банк Дебетовая карта») — creative creation fails on them. Don't extend the marketplace list without checking the offer first.
- Never commit `.cache/` (tokens) — it is gitignored.

How it works per link: `GET /offers/by-link` → must resolve to offer `1` (AliExpress), else the link is reported and skipped → `POST /creative/create` with `type: deeplink` → the returned `attributes.code` is already a ready affiliate link with the ОРД `erid` marker baked in (`https://ali.click/xxxxxxx?erid=2SDnjc…`), so no separate shortening call is needed. The file is rewritten in place, CRLF preserved.

### Credentials (one-time setup, done by the user)

`EPN_CLIENT_ID` and `EPN_CLIENT_SECRET` come from ЛК вебмастера → Инструменты → Разработчикам → API → «Получить API 2.0 ключи». The script reads `os.environ` first and falls back to `HKCU\Environment` via PowerShell — a `setx` done after the current terminal started still works, no restart needed. Never ask the user to paste keys into the chat and never write them to a tracked file.

Tokens are cached in `.cache/epn_tokens.json` and refreshed automatically on `401` (`POST /token/refresh`). **`GET /ssid` is limited to 1 request per day per IP** (a second one needs a captcha), so do not delete `.cache/epn_tokens.json` casually — if it is gone and the daily ssid was already spent, the API is unavailable until the next day. `access_token` lives 24 h, `refresh_token` 15 days.

### Constants at the top of the script

| Constant | Value | Meaning |
| --- | --- | --- |
| `ALIEXPRESS_OFFER_ID` | `1` | AliExpress in EPN |
| `ALIEXPRESS_PLACEMENT_ID` | `1455321` | the confirmed `k3d.tech` placement for that offer |

A placement belongs to one offer: passing the placement of another offer returns `422 Offer does not match placement!`. The user has a second `k3d.tech` placement (`1510797`) — it belongs to Яндекс.Маркет, do not use it here. To re-resolve: `GET /placements?statuses=allow&fields=id,url,status` (note `fields` is required) plus `GET /placements/offers` for the offer ids.

### Reading the report

Exit code is `0` when everything that was found got converted (or nothing matched), `1` when at least one link was left unprocessed.

```text
Файл: docs/part-navi/filament_dryers.md
Ссылок на AliExpress: 1 (уникальных: 1)

  ✔ строка 48  https://aliexpress.ru/item/1005008712869586.html
      → https://ali.click/fm26o1y?erid=2SDnjcqFuat

Заменено: 1. Файл docs/part-navi/filament_dryers.md перезаписан.
```

- `✔` — replaced. `✖` — left as is, with the reason; the rest of the file is still written.
- `прямых ссылок на AliExpress нет — файл не изменён` — nothing to do; this also means the page was already processed (re-runs are safe: affiliate links sit on `ali.click`/`alli.pub` hosts and are not matched by the link regex).
- Errors worth relaying to the user instead of retrying blindly: `422 Offer does not match placement!` → `ALIEXPRESS_PLACEMENT_ID` is stale or belongs to another offer; `422 Offer not found!` / offer id ≠ 1 → the link is not really an AliExpress product offer; `429` → rate limit, the script already retries with backoff.

### After running

1. Confirm in the report that every AliExpress link on the page was either replaced or skipped, and that the script wrote the file.
2. `python -m properdocs build --clean` to make sure the page still builds. Commit the page; `.cache/` stays untracked.

**Do not audit the rest of the file.** The user edits these pages himself before publishing, so the working tree routinely carries his own uncommitted changes — removed rows, reworded comments, dropped external links, added ad-marking lines. That is normal and none of it is the script's business. Never attribute those changes to the script, never ask whether they were intentional, never "restore" them. Report only what the script itself did: the URLs from its report.

Also not worth reporting:

- **The same link standing in several rows.** One AliExpress lot can contain several variants (different thread pitch, size, colour), so a single item legitimately serves e.g. both the `Гайка Tr8x8` and the `Гайка Tr8x2` row. Normal, not an error.
- **Whether the linked item matches the row.** The user checks his own product links; verifying them means re-litigating his edits. If a link looks wrong for its row, at most mention it once as a question — do not frame it as a bug in the script or in the page.

Notes and limits:

- `POST /creative/create` is limited to 1 request per second; the script sleeps `CREATIVE_PAUSE` (1.2 s) between calls. Don't remove that sleep. **`GET /creatives` is limited the same way** (`429005 Request is limited`) — when paging through the account's creatives (≈2000 of them, 100 per page), sleep ~3 s between requests and expect an empty page plus `429` if you don't.
- The creative `description` is mandatory and ≤100 chars; it is generated as `«<page title>» [<path>:<line>]` so the entry is identifiable in ЛК epn.bz.
- `shortcut:` no cross-check against creatives that already exist in the account — a re-run after a partial failure creates a duplicate creative for the same URL. Duplicates point at the same product and are harmless; add a `GET /creatives` lookup if the ЛК clutter starts to matter.
- Query parameters on the source URL are passed to EPN as-is and are not echoed in the returned `code`.
