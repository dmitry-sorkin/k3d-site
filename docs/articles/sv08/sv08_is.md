---

authors:
  - makarov
title: Графики Input Shaping Sovol SV08
icon: material/printer-3d
---

<style>
  .graphs-controls { display: flex; flex-wrap: wrap; gap: 0.6em 1.5em;
    padding: 0.9em 1.1em; background: transparent;
    border: 1px solid var(--md-default-fg-color--light, #cccccc);
    border-radius: 6px; margin-bottom: 0.8em; align-items: flex-start;
    color: var(--md-default-fg-color, #222); }
  .graphs-controls label { font-size: 0.9rem; }
  .graphs-controls select { margin-left: 0.3em; padding: 0.2em 0.4em; }
  .graphs-ticks { display: flex; flex-wrap: wrap; gap: 0.3em 0.6em;
    padding: 0.7em 1em; background: transparent;
    border: 1px solid var(--md-default-fg-color--light, #cccccc);
    border-radius: 6px; margin-bottom: 0.8em;
    color: var(--md-default-fg-color, #222); }
  .graphs-ticks .curve-group { width: 100%; display: flex; flex-wrap: wrap;
    gap: 0.3em 0.6em; align-items: center; }
  .graphs-ticks .curve-group-hdr { font-weight: 700; font-size: 0.88rem;
    width: 100%; margin-top: 0.3em; }
  .graphs-ticks .curve-row { display: flex; align-items: center; gap: 0.4em;
    font-size: 0.85rem; }
  .graphs-ticks .curve-row input[type="checkbox"] { width: 15px; height: 15px; cursor: pointer; }
  .graphs-ticks .curve-row .curve-swatch { flex: 0 0 auto; display: block; }
  .graphs-ticks .curve-row span { font-family: monospace; }
  .vega-embed .vega-actions { display: none !important; }
</style>

# Графики Input Shaping Sovol SV08

График получен путём обработки кривых зависимости PSD_XYZ от частоты, снятых в 27 точках области печати: X={25,175,350}, Y={25,175,350}, Z={20,175,340}. Кривые распределяются на три группы по девять штук в соответствии с выбором в выпадающем списке "изменение вдоль оси". Например, при выборе оси Z, в одну группу попадут кривые, снятые в 9 точках с одинаковой Z координатой. Для кривых внутри каждой группы строится одна огибающая, которая и отображается на графике.

<div id="gr_loader" style="color:#555;font-style:italic;margin:1em 0;">Загрузка визуализации…</div>

<div id="gr_controls" class="graphs-controls" style="display:none;">
  <div><label>Изменение вдоль оси:
    <select id="gr_group">
      <option value="x">по X</option>
      <option value="y">по Y</option>
      <option value="z" selected>по Z</option>
    </select>
  </label></div>
  <div><label>Шейпер:
    <select id="gr_excite">
      <option value="both" selected>X и Y</option>
      <option value="x">только X</option>
      <option value="y">только Y</option>
    </select>
  </label></div>
</div>

<div id="gr_ticks" class="graphs-ticks" style="display:none;"></div>
<div id="gr_chart" style="width:100%;"></div>

<script>
(async function() {
  const DATA_URL = '../data/envelopes.json';

  // Three fixed coordinates per axis.
  const XS = [25, 175, 325];
  const YS = [25, 175, 325];
  const ZS = [20, 175, 320];
  const fixedVals = { x: XS, y: YS, z: ZS };

  // Bright, distinct colors per fixed-coordinate value.
  //   low  (25 / 20)   -> green
  //   mid  (175)       -> orange
  //   high (325 / 320) -> purple
  const COLORS = ['#2CA02C', '#FF7F0E', '#9467BD'];

  // Vega-Lite dash patterns. Each pattern is a string of comma-separated
  // lengths; pairs (length, gap) repeat.
  //   X-возбуждение -> пунктир
  //   Y-возбуждение -> сплошная
  const DASH_BY_EXCITE = { x_axis: '4,3', y_axis: '1,0' };
  const DASH_BY_EXCITE_LIST = [DASH_BY_EXCITE.x_axis, DASH_BY_EXCITE.y_axis];
  const EXCITE_KEYS = ['x_axis', 'y_axis'];
  const EXCITE_LABELS = { x_axis: 'Шейпер оси X', y_axis: 'Шейпер оси Y' };

  const raw = await fetch(DATA_URL).then(r => r.json());

  // ── Vega loader ───────────────────────────────────────────────────────
  async function loadVega() {
    if (window.vegaEmbed) return;
    const load = (src) => new Promise((ok, fail) => {
      const s = document.createElement('script'); s.src = src;
      s.onload = ok; s.onerror = fail; document.head.appendChild(s);
    });
    document.getElementById('gr_loader').textContent = 'Загрузка библиотеки визуализации…';
    await load('https://cdn.jsdelivr.net/npm/vega@5/build/vega.min.js');
    await load('https://cdn.jsdelivr.net/npm/vega-lite@5/build/vega-lite.min.js');
    await load('https://cdn.jsdelivr.net/npm/vega-embed@6/build/vega-embed.min.js');
  }

  // ── Build the long-form dataset for the current view ──────────────────
  // One row per (freq, fixed-coord, excite). _label is the combined
  // legend entry, _val is the fixed coord (e.g. 20/175/320), _excite
  // is 'x_axis'/'y_axis'. _label is unique per curve, so it can drive
  // the single combined color encoding.
  function buildData(groupDim, excite) {
    const vals = fixedVals[groupDim];
    const freqs = raw.freq;
    const showX = excite === 'both' || excite === 'x';
    const showY = excite === 'both' || excite === 'y';
    const rows = [];
    vals.forEach((v) => {
      const vStr = String(v);
      if (showX) {
        const psd = raw.x_axis[groupDim][vStr];
        for (let k = 0; k < freqs.length; k++) {
          rows.push({
            freq: freqs[k], psd: psd[k],
            _val: vStr, _excite: 'x_axis',
            _label: `${groupDim.toUpperCase()}=${v} (${EXCITE_LABELS.x_axis})`,
          });
        }
      }
      if (showY) {
        const psd = raw.y_axis[groupDim][vStr];
        for (let k = 0; k < freqs.length; k++) {
          rows.push({
            freq: freqs[k], psd: psd[k],
            _val: vStr, _excite: 'y_axis',
            _label: `${groupDim.toUpperCase()}=${v} (${EXCITE_LABELS.y_axis})`,
          });
        }
      }
    });
    return rows;
  }

  // ── Read current checkbox state ──────────────────────────────────────
  // Returns { x_axis: { '20': true, '175': true, '320': true }, y_axis: {...} }
  function getTickStates() {
    const s = { x_axis: {}, y_axis: {} };
    document.querySelectorAll('#gr_ticks input[type="checkbox"]').forEach(cb => {
      const ex = cb.dataset.excite;     // 'x_axis' | 'y_axis'
      const v  = cb.dataset.val;        // string like '20'
      s[ex][v] = cb.checked;
    });
    return s;
  }

  // ── Build curves panel: groups by shaper, each row = checkbox + swatch + label ──
  function buildCurvesPanel(groupDim, excite) {
    const wrap = document.getElementById('gr_ticks');
    wrap.innerHTML = '';
    wrap.style.display = 'flex';

    const vals = fixedVals[groupDim];
    const dimLabel = groupDim.toUpperCase();
    const showX = excite === 'both' || excite === 'x';
    const showY = excite === 'both' || excite === 'y';
    const svgNS = 'http://www.w3.org/2000/svg';

    const groups = [
      ['x_axis', 'Шейпер оси X', showX, '4,3'],
      ['y_axis', 'Шейпер оси Y', showY, '1,0'],
    ].filter(([_, __, show]) => show);

    groups.forEach(([exKey, exLabel, _, dashStr]) => {
      const hdr = document.createElement('div');
      hdr.className = 'curve-group-hdr';
      hdr.textContent = exLabel;
      wrap.appendChild(hdr);

      const group = document.createElement('div');
      group.className = 'curve-group';
      vals.forEach((v, vi) => {
        const row = document.createElement('div');
        row.className = 'curve-row';

        const cb = document.createElement('input');
        cb.type = 'checkbox';
        cb.id = `gr_tick_${exKey}_${v}`;
        cb.checked = true;
        cb.dataset.excite = exKey;
        cb.dataset.val = String(v);
        cb.addEventListener('change', render);

        const svg = document.createElementNS(svgNS, 'svg');
        svg.setAttribute('width', '36');
        svg.setAttribute('height', '10');
        svg.classList.add('curve-swatch');
        const path = document.createElementNS(svgNS, 'path');
        path.setAttribute('d', 'M0,5 L36,5');
        path.setAttribute('stroke', COLORS[vi]);
        path.setAttribute('stroke-width', '3');
        path.setAttribute('fill', 'none');
        path.setAttribute('stroke-dasharray', dashStr);
        svg.appendChild(path);

        const sp = document.createElement('span');
        sp.textContent = `${dimLabel}=${v}`;

        row.appendChild(cb);
        row.appendChild(svg);
        row.appendChild(sp);
        group.appendChild(row);
      });
      wrap.appendChild(group);
    });
  }

  // ── Build Vega-Lite spec ─────────────────────────────────────────────
  // Single line layer. color encodes _label (combined curve name),
  // strokeDash encodes _excite. Visibility via pre-filtered data.
  function makeSpec(groupDim, excite, tickStates) {
    const vals = fixedVals[groupDim];
    const baseData = buildData(groupDim, excite);
    const dimLabel = groupDim.toUpperCase();
    const showX = excite === 'both' || excite === 'x';
    const showY = excite === 'both' || excite === 'y';

    // Add _visible per row, then pre-filter
    const data = baseData
      .map(d => ({
        ...d,
        _visible: tickStates[d._excite][String(d._val)] !== false ? 1 : 0,
      }))
      .filter(d => d._visible > 0);

    // yMax from visible rows
    let yMax = 0;
    data.forEach(d => { if (d.psd > yMax) yMax = d.psd; });
    yMax = Math.max(yMax * 1.08, 1);

    // Build legend domain/range: one entry per visible curve
    const legendDomain = [];
    const legendRange = [];
    vals.forEach((v) => {
      const vi = vals.indexOf(v);
      if (showX) {
        const lbl = `${dimLabel}=${v} (${EXCITE_LABELS.x_axis})`;
        legendDomain.push(lbl); legendRange.push(COLORS[vi]);
      }
      if (showY) {
        const lbl = `${dimLabel}=${v} (${EXCITE_LABELS.y_axis})`;
        legendDomain.push(lbl); legendRange.push(COLORS[vi]);
      }
    });

    const layers = [{
      data: { values: data },
      mark: { type: 'line', strokeWidth: 2.4, clip: true, tooltip: true },
      encoding: {
        x: { field: 'freq', type: 'quantitative',
             axis: { title: 'Частота (Гц)', titleFontSize: 13, labelFontSize: 11 },
             scale: { domain: [0, 200] } },
        y: { field: 'psd', type: 'quantitative',
             axis: { title: 'Power spectral density',
                     titleFontSize: 13, labelFontSize: 11, format: '~e' },
             scale: { domain: [0, yMax] } },
        color: {
          field: '_label', type: 'nominal',
          scale: { domain: legendDomain, range: legendRange },
          legend: null,
        },
        strokeDash: {
          field: '_excite', type: 'nominal',
          scale: { domain: EXCITE_KEYS, range: DASH_BY_EXCITE_LIST },
          legend: null,
        },
        detail: { field: '_label', type: 'nominal' },
      },
    }];

    return {
      $schema: 'https://vega.github.io/schema/vega-lite/v5.json',
      width: 'container',
      height: 480,
      autosize: { type: 'fit', contains: 'padding' },
      config: {
        axis: { labelFontSize: 11, titleFontSize: 13 },
        legend: { labelFontSize: 12, titleFontSize: 13 },
        view: { stroke: 'transparent' },
      },
      title: null,
      layer: layers,
    };
  }

  // ── State + render ───────────────────────────────────────────────────
  let curGroup = 'z', curExcite = 'both';

  async function render() {
    document.getElementById('gr_loader').style.display = 'none';
    document.getElementById('gr_controls').style.display = 'flex';
    const tickStates = getTickStates();
    const spec = makeSpec(curGroup, curExcite, tickStates);
    await vegaEmbed('#gr_chart', spec, {
      actions: { export: true, source: false, compiled: false, editor: false },
      renderer: 'svg',
    });
  }

  function initControls() {
    const sGroup  = document.getElementById('gr_group');
    const sExcite = document.getElementById('gr_excite');
    const onChange = () => {
      curGroup = sGroup.value;
      curExcite = sExcite.value;
      buildCurvesPanel(curGroup, curExcite);
      render();
    };
    sGroup.addEventListener('change', onChange);
    sExcite.addEventListener('change', onChange);
  }

  await loadVega();
  initControls();
  buildCurvesPanel(curGroup, curExcite);
  render();
})();
</script>