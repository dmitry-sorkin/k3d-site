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
  vostok/             # K3D VOSTOK printer: index, faq, releases, manual/, v8/ (archive)
  calibrations/       # PID, IS, flow, accuracy/, VFA, retractions/, la/ (PA), max_flowrate/
  projects/           # ps_translation, custom_pulleys, printheads (ebp/ehp/esp), feeders (bemege/minifeeder/feeder965)
  printers/           # per-printer mods (kobra3, k1, neptune3pro/4, kp3s, qidi_q1_pro, qidi_q2)
  articles/           # long-form posts (print_strong, custom_gcode, nozzle_test_2024, ultra_high_flow_hotends, ad5m_ad5x_hotend_test, cooling_problems)
  redirects/          # stub .md files whose real target is in properdocs.yml `redirect_maps`
  pics/, assets/, models/   # images and 3D models referenced from MD
  stylesheets/        # extra7.css referenced from theme
  javascripts/        # mathjax loader
overrides/            # Material theme partial overrides (partials/integrations/, .icons/)
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