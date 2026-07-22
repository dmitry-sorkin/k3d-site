// Layout-independent "/" hotkey for focusing the Material search bar.
//
// The built-in Material keyboard binding listens for `event.key === "/"`.
// On a non-Latin layout (e.g. Russian "ЙЦУКЕН") that key produces "ш" or
// "э" instead of "/", so the binding never fires. We re-implement the
// shortcut against the *physical* key (event.code === "Slash"), which is
// independent of the active keyboard layout.

(function () {
  var isMac = /Mac|iPod|iPhone|iPad/.test(navigator.platform);

  function focusSearch() {
    var input =
      document.querySelector(".md-search__input") ||
      document.querySelector('input[type="search"]');
    if (!input) return false;
    // Material toggles its own state machine; calling .focus() alone works,
    // but clicking the search toggle first makes the overlay behave the same
    // as the built-in "/" shortcut.
    var toggle = document.querySelector(".md-search__form") ||
      document.querySelector("[data-md-component='search']");
    if (toggle && !document.querySelector(".md-search--active")) {
      toggle.dispatchEvent(new MouseEvent("click", { bubbles: true }));
    }
    input.focus();
    input.select && input.select();
    return true;
  }

  document.addEventListener("keydown", function (event) {
    // Only the bare Slash key, no modifiers. Shift would change the produced
    // character ("?" on US layout), and we want a clean "/" press.
    if (event.key !== "Slash" && event.code !== "Slash") return;
    if (event.ctrlKey || event.altKey || event.metaKey || event.shiftKey) return;

    var target = event.target;
    if (target && (target.isContentEditable ||
                   target.tagName === "INPUT" ||
                   target.tagName === "TEXTAREA" ||
                   target.tagName === "SELECT")) {
      return;
    }

    if (focusSearch()) {
      event.preventDefault();
      event.stopPropagation();
    }
  }, true);
})();
