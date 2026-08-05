/* Регистрация icon-пака mdi (Material Design Icons) для mermaid
 * и отключение санитайзера Strict-режима (он вырезает <use> иконок).
 *
 * Тема material грузит mermaid лениво с unpkg и не вызывает registerIconPacks,
 * поэтому перехватываем присваивание window.mermaid: setter срабатывает
 * в момент загрузки UMD-скрипта — до initialize/render, без гонок. */
(() => {
	const iconsUrl = "https://unpkg.com/@iconify-json/mdi@1/icons.json";

	const register = () => {
		const m = window.mermaid;
		if (!m || typeof m.registerIconPacks !== "function" || m.__iconsRegistered)
			return;
		m.__iconsRegistered = true;

		// Strict-сантайзер mermaid вырезает <use> из SVG — иконки не видны.
		// Принудительно включаем loose до инициализации темы.
		if (typeof m.initialize === "function" && !m.__initializePatched) {
			const origInit = m.initialize.bind(m);
			m.__initializePatched = true;
			m.initialize = (config) => {
				origInit({ ...config, securityLevel: "loose" });
			};
		}

		m.registerIconPacks([
			{
				name: "mdi",
				loader: () => fetch(iconsUrl).then((res) => res.json()),
			},
		]);
	};

	let mermaid;
	Object.defineProperty(window, "mermaid", {
		configurable: true,
		get: () => mermaid,
		set: (v) => {
			mermaid = v;
			register();
		},
	});
	register();
})();
