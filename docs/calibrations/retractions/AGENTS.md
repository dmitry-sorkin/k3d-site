# AGENTS — docs/calibrations/retractions/

Retraction tower calibrator: an HTML form with embedded WASM logic, embedded in
the K3D site as an MkDocs Material page. Generates G-code for retraction tower
calibration.

## ⚠ Source lost

The Go source for this calibrator (`main.go`) is **not present in this
repository** and was not lost in any commit reachable from any branch,
worktree, dangling commit, or unreachable blob (verified with `git fsck`
and `git log --diff-filter=AMD`). Only the precompiled
`assets/wasm/rct_lib.wasm` artifact is here.

The calibrator was originally developed at <https://github.com/dmitry-sorkin/k3d_rct>
and was meant to be moved into this repository when the WASM-based interface
rework landed (commit `6af3512 rct_interface_rework`, January 2024), but the
`.go` source never made it into a commit here. Treat the WASM as the only
authoritative artifact and **do not change the calibrator logic without the
source**: any behavior change to G-code generation requires re-deriving the
source from `rct_lib.wasm` (manual decompile) or finding an external copy.

**Consequence:** `calibrator_version` in `assets/js/lib.js` is the version
label carried by the JS UI. The WASM has its own embedded version. They are
not kept in sync from this repo — both are immutable artifacts. Do not bump
the JS version expecting a meaningful change in G-code output.

## Layout

```
docs/calibrations/retractions/
  index.md             Russian article: survey of retraction methods
  calibrator.md        MkDocs page hosting the live UI (the actual app)
  rct.html             Stub redirect → k3d.tech/.../calibrator/
  rct_old.html         Archived pre-WASM form, kept for reference only
  assets/
    css/rct_style.css
    js/
      wasm_exec.js     Vendored Go runtime glue (do not edit)
      gwaloader.js     Loads rct_lib.wasm, toggles the Generate button
      lib.js           Form persistence, language strings, segment preview
      streamsavermin.js  Vendored streaming file-save polyfill (do not edit)
    wasm/
      .htaccess        application/wasm MIME for Apache hosts
      rct_lib.wasm     Compiled artifact (source lost — see above)
  pics/                Images referenced from index.md / calibrator.md
```

## Data flow

`calibrator.md` (HTML form, IDs without the `k3d_la_` prefix used by the PA
calibrator) ⇄ `lib.js` (localStorage, validation, lang strings)
⇄ `rct_lib.wasm` via `syscall/js` ⇄ `streamsavermin.js` (browser downloads
`.gcode`).

The WASM side reads values via `js.Global().Get("document").Call("getElementById", id)`,
the same pattern as the PA calibrator but with plain IDs (no `k3d_la_`
prefix). Language strings cross the same boundary through
`getString(key)` → `window.lang.getString(key)`.

## Invariants

1. **Form element IDs are the public API across the JS layer and the WASM
   artifact.** An ID appearing in `calibrator.md`, `formFields` in
   `lib.js`, and `getElementById` calls inside `rct_lib.wasm` must match.
   Renaming requires either editing the WASM (impossible without source)
   or coordinating the JS rename with a future WASM rebuild.
2. **`assets/wasm/rct_lib.wasm` is the only artifact** — there is no
   `build.sh`/`build.bat` in this repo because the source is lost.
3. **Generated G-code uses `$LA`, `$BEDTEMP`, `$HOTTEMP`, `$G29`, `$FLOW`
   placeholders** in start G-code; the WASM expands them to the appropriate
   M-codes for the selected firmware.
4. **The 5-cell preview row** (`init_retract_length`, `end_retract_length`,
   `init_retract_speed`, `end_retract_speed`, `num_segments` descriptions)
   is updated live by `checkSegments()` via `setSegmentsPreview()`.

## Conventions

- Russian comments in JS, Russian UI strings, Russian prose in `index.md` /
  `calibrator.md`. English only for identifiers, paths, tool names.
- The HTML form uses **plain IDs without a prefix** (e.g. `bedX`, not
  `k3d_la_bedX`). This differs from the PA calibrator; do not "unify" the
  naming without rebuilding the WASM, or the runtime will silently fail
  to read form values.
- Three languages: `de`, `en`, `ru`. Default in `init()` is `ru`. Each
  switch case in `initLang` has its own `break`.

## Don't

- Don't edit `wasm_exec.js` or `streamsavermin.js` — vendored.
- Don't add a new language without also teaching the WASM's
  `getLangString` callers the new keys (impossible without source).
- Don't ship a `lib.js` change that depends on a WASM behavior change —
  the WASM cannot be rebuilt from this repo.
- Don't rename HTML form IDs without rebuilding the WASM, or you'll break
  the G-code generator silently.
