---
authors:
  - sorkin
hide:
  - navigation
  - toc
title: K3D - ВСЁ О 3D ПЕЧАТИ
description: Главная страница
show_created: false
show_updated: false
show_author: false
glightbox: false
---

<div class="lp-hero">
	<video class="lp-hero__video" autoplay muted loop playsinline>
		<source src="/vostok/pics/index_main.mp4" type="video/mp4">
	</video>
	<div class="lp-hero__overlay"></div>
	<div class="lp-hero__content">
		<img class="lp-hero__logo" src="/assets/k3d_logo.png" alt="K3D">
		<span class="lp-hero__title">ВСЁ О 3D ПЕЧАТИ</span>
	</div>
</div>

<script>
(() => {
  const v = document.querySelector('.lp-hero__video');
  if (!v) return;
  if (v.readyState >= 2) {
    v.classList.add('is-loaded');
  } else {
    v.addEventListener('loadeddata', () => v.classList.add('is-loaded'), { once: true });
  }
})();
</script>

<div class="lp-search-note">
	Нажмите <kbd>/</kbd> и начните печатать — поиск работает по всем статьям и гайдам сайта
</div>

<div class="lp-cards">
	<a class="lp-card" href="/part-navi/printers/">
		<div class="lp-card-body">
			<div class="lp-num">01</div>
			<h3>Подобрать 3D принтер</h3>
			<p>Какой принтер купить под ваш бюджет и задачи: актуальные модели, честные плюсы и минусы</p>
		</div>
	</a>
	<a class="lp-card" href="/part-navi/filament/">
		<div class="lp-card-body">
			<div class="lp-num">02</div>
			<h3>Подобрать филамент</h3>
			<p>Проверенные пластики по типам и брендам: что брать, а от чего лучше держаться подальше</p>
		</div>
	</a>
	<a class="lp-card" href="/vostok/">
		<div class="lp-card-body">
			<div class="lp-num">03</div>
			<h3>K3D VOSTOK</h3>
			<p>Проект быстрого крупногабаритного IDEX 3D-принтера с открытыми исходниками</p>
		</div>
	</a>
	<a class="lp-card" href="/calibrations/">
		<div class="lp-card-body">
			<div class="lp-num">04</div>
			<h3>Калибровки</h3>
			<p>От PID до Pressure Advance: пошаговые гайды и калибровщики для настройки принтера</p>
		</div>
	</a>
</div>

<div class="grid cards cols-6 lp-secondary" markdown>

-   :material-printer-3d:{ .lg .middle } __[3D Рекомендатор](./part-navi/printers.md)__

    ---

    [:material-printer-3d: 3D Принтеры](./part-navi/printers.md)    [:material-printer-3d-nozzle: Экструдеры](./part-navi/extruder.md)    [:fontawesome-solid-microchip: Электроника](./part-navi/electronics.md)    [:fontawesome-solid-wrench: Механика](./part-navi/mechanics.md)    [:fontawesome-solid-drum-steelpan: Филаменты](./part-navi/filament.md)    [:material-heat-wave: Сушка филамента](./part-navi/filament_dryers.md)    [:material-water: Расходники](./part-navi/consumables.md)

-   :material-printer-3d:{ .lg .middle } __[VOSTOK](./vostok/index.md)__

    ---

    [:material-printer-3d: О VOSTOK](./vostok/index.md)    [:lucide-help-circle: Частые вопросы](./vostok/faq.md)    [:material-image-multiple: Галерея](./vostok/gallery/index.md)    [:lucide-download: Релизы](./vostok/releases/index.md)    [:lucide-list-todo: Начало](./vostok/manual/start.md)    [:lucide-package-open: Подготовка](./vostok/manual/prepare.md)    [:lucide-plug: Схемы](./vostok/manual/wiring.md)    [:lucide-code-2: Прошивка](./vostok/manual/firmware.md)    [:custom-orcaslicer: Слайсер](./vostok/manual/slicer_configuration.md)    [:material-swap-horizontal: СБСИ](./vostok/manual/fast_tool_swaps.md)    [:lucide-command: Макросы](./vostok/manual/macros.md)

-   :material-notification-clear-all:{ .lg .middle } __[Калибровки](./calibrations/index.md)__

    ---

    [:material-notification-clear-all: Порядок](./calibrations/index.md)    [:fontawesome-solid-temperature-low: PID](./calibrations/pid.md)    [:fontawesome-solid-wave-square: Input Shaping](./calibrations/manual_is_calibration.md)    [:material-water-check: Поток](./calibrations/flow.md)    [:fontawesome-solid-maximize: Точность](./calibrations/accuracy/index.md)    [:material-sine-wave: Рябь](./calibrations/vfa.md)    [:material-arrow-left-right-bold: Смещение голов](./calibrations/printhead_offset.md)    [:octicons-package-16: Откаты](./calibrations/retractions/calibrator.md)    [:octicons-package-16: Pressure Advance](./calibrations/la/calibrator.md)    [:octicons-package-16: Объёмный расход](./calibrations/max_flowrate/calibrator.md)

-   :material-translate:{ .lg .middle } __[Проекты](./projects/ps_translation/index.md)__

    ---

    [:material-translate: Перевод PrusaSlicer](./projects/ps_translation/index.md)    [:octicons-gear-24: Шкивы](./projects/custom_pulleys/index.md)    [:material-printer-3d-nozzle: Печатающие головы](./projects/printheads.md)    [:material-subdirectory-arrow-right: EBP](./projects/ebp/index.md)    [:material-subdirectory-arrow-right: EH2P](./projects/ehp/index.md)    [:material-subdirectory-arrow-right: ESP](./projects/esp/index.md)    [:fontawesome-solid-gears: Фидеры](./projects/feeders.md)    [:material-subdirectory-arrow-right: BeMeGe](./projects/bemege/index.md)    [:material-subdirectory-arrow-right: Minifeeder](./projects/minifeeder/index.md)    [:material-subdirectory-arrow-right: Feeder965](./projects/feeder965/index.md)

-   :material-weight-lifter:{ .lg .middle } __[Статьи](./articles/print_strong.md)__

    ---

    [:material-weight-lifter: Прочность](./articles/print_strong.md)    [:material-code-block-tags: G-коды](./articles/custom_gcode.md)    [:material-printer-3d-nozzle: Сопла Volcano](./articles/nozzle_test_2024.md)    [:material-printer-3d-nozzle: UHF хотэнды](./articles/ultra_high_flow_hotends.md)    [:material-printer-3d-nozzle: Хотэнды AD5M/AD5X](./articles/ad5m_ad5x_hotend_test.md)    [:material-fan: Дефекты охлаждения](./articles/cooling_problems.md)

-   :material-printer-3d:{ .lg .middle } __[Доработки](./printers/qidi_q2/index.md)__

    ---

    [:material-printer-3d: QIDI Q2](./printers/qidi_q2/index.md)    [:material-printer-3d: QIDI Q1 Pro](./printers/qidi_q1_pro.md)    [:material-printer-3d: Creality K1](./printers/k1.md)    [:material-printer-3d: Anycubic Kobra 3](./printers/kobra3.md)    [:material-printer-3d: Elegoo Neptune 4](./printers/neptune4.md)    [:material-printer-3d: Elegoo Neptune 3 Pro](./printers/neptune3pro.md)    [:material-printer-3d: Kingroon KP3S](./printers/kp3s.md)

</div>
