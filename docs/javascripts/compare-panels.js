// Кнопки «Сравнить с …» на странице рекомендатора принтеров.
//
// В Markdown эти блоки пишутся обычными адмонициями, подряд, после описания
// принтера:
//
//     ??? compare "Сравнить с Bambu Lab A1"
//
//         Текст сравнения.
//
//     ??? compare "Сравнить с QIDI Q2"
//
//         Текст сравнения.
//
// Сама по себе каждая адмониция — отдельный <details>, и панель с текстом
// лежит внутри него. Поэтому CSS не может ни собрать кнопки в одну строку с
// общей подписью, ни показать панель под этой строкой, ни закрыть соседей.
// Этот скрипт перестраивает каждую группу подряд идущих .compare в:
//
//     <div class="compare-group">
//       <div class="compare-row">
//         <span class="compare-label">Сравнить с:</span>
//         <button class="compare-chip">…</button>
//       </div>
//       <div class="compare-panels"><div class="compare-panel" hidden>…</div></div>
//     </div>
//
// Открыта всегда не более одной панели. Без JS блоки остаются рабочими
// адмонициями — деградация мягкая.

(function () {
  var PREFIX = /^сравнить\s+с\s+/i;

  // Сквозной счётчик, а не локальный для группы: иначе панели разных
  // групп получат одинаковые id и aria-controls будет указывать не туда.
  var panelSeq = 0;

  function init() {
    // Группируем подряд идущие .compare: у соседних по DOM они идут
    // непосредственно друг за другом (nextElementSibling), поэтому
    // посторонние элементы и текст между блоками разрывают группу.
    var groups = [];
    var current = null;

    Array.prototype.forEach.call(
      document.querySelectorAll("details.compare"),
      function (details) {
        if (current && current[current.length - 1].nextElementSibling === details) {
          current.push(details);
        } else {
          current = [details];
          groups.push(current);
        }
      }
    );

    groups.forEach(buildGroup);
  }

  function buildGroup(run) {
    var group = document.createElement("div");
    group.className = "compare-group";

    var row = document.createElement("div");
    row.className = "compare-row";

    var label = document.createElement("span");
    label.className = "compare-label";
    label.textContent = "Сравнить с:";
    row.appendChild(label);

    var panels = document.createElement("div");
    panels.className = "compare-panels";

    var chips = [];

    run.forEach(function (details) {
      var summary = details.querySelector(":scope > summary");
      var title = (summary ? summary.textContent : "").trim().replace(PREFIX, "");
      var id = "compare-panel-" + (++panelSeq);

      // Панель забирает тело адмониции, кроме заголовка.
      var panel = document.createElement("div");
      panel.className = "compare-panel";
      panel.id = id;
      panel.hidden = true;

      Array.prototype.forEach.call(details.children, function (child) {
        if (child !== summary) panel.appendChild(child.cloneNode(true));
      });

      var chip = document.createElement("button");
      chip.type = "button";
      chip.className = "compare-chip";
      chip.textContent = title;
      chip.setAttribute("aria-expanded", "false");
      chip.setAttribute("aria-controls", id);

      // Клик мышью не переводит фокус на кнопку: в некоторых браузерах
      // фокус утаскивает к себе прокрутку, и экран прыгает. Клавиатурный
      // фокус (Tab / Space) не затрагивается, состояние видно по
      // aria-expanded и стилю [aria-expanded="true"].
      chip.addEventListener("mousedown", function (event) {
        event.preventDefault();
      });

      chip.addEventListener("click", function () {
        var wasOpen = chip.getAttribute("aria-expanded") === "true";

        closeAll();
        if (!wasOpen) {
          chip.setAttribute("aria-expanded", "true");
          panel.hidden = false;
        }
      });

      row.appendChild(chip);
      panels.appendChild(panel);
      chips.push(chip);
    });

    function closeAll() {
      chips.forEach(function (chip) {
        chip.setAttribute("aria-expanded", "false");
      });
      Array.prototype.forEach.call(panels.children, function (panel) {
        panel.hidden = true;
      });
    }

    group.appendChild(row);
    group.appendChild(panels);

    run[0].parentNode.insertBefore(group, run[0]);
    run.forEach(function (details) {
      details.parentNode.removeChild(details);
    });
  }

  if (window.document$ && typeof window.document$.subscribe === "function") {
    window.document$.subscribe(init);
  } else if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init);
  } else {
    init();
  }
})();