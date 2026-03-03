const calibrator_version = 'v2.1';
window.calibrator_version = calibrator_version;
var savedSegmentsInfo = null;

function download(filename, text) {
	var element = document.createElement('a');
	element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
	element.setAttribute('download', filename);

	element.style.display = 'none';
	document.body.appendChild(element);

	element.click();

	document.body.removeChild(element);
}

function saveTextAsFile(filename, text) {
	var textFileAsBlob = new Blob([text], { type: 'text/plain' });

	var downloadLink = document.createElement("a");
	downloadLink.download = filename;
	if (window.webkitURL != null) {
		// Chrome allows the link to be clicked without actually adding it to the DOM.
		downloadLink.href = window.webkitURL.createObjectURL(textFileAsBlob);
	} else {
		// Firefox requires the link to be added to the DOM before it can be clicked.
		downloadLink.href = window.URL.createObjectURL(textFileAsBlob);
		downloadLink.onclick = destroyClickedElement;
		downloadLink.style.display = "none";
		document.body.appendChild(downloadLink);
	}

	downloadLink.click();
}

var currentStream = null;
var currentWriter = null;
const encoder = new TextEncoder();

function beginSaveFile(filename) {
	if (currentStream != null) {
		currentStream.abort();
	}
	currentStream = streamSaver.createWriteStream(filename);
	currentWriter = currentStream.getWriter();
}

function writeToFile(str) {
	if (currentWriter != null) {
		currentWriter.write(encoder.encode(str));
	}
}

function finishFile() {
	if (currentWriter != null) {
		currentWriter.ready.then(() => {
			currentWriter.close();
			currentWriter = null;
			currentStream = null;
		})
			.catch((err) => {
				showError(err);
			});
	}
}

function showError(value) {
	var container = document.getElementById("resultContainer");
	var output = document.createElement("textarea");
	output.id = "gCode";
	// output.name = "gCode";
	output.cols = "80";
	output.rows = "10";
	output.value = value;
	// output.className = "css-class-name"; // set the CSS class
	container.appendChild(output); //appendChild
}

function destroyClickedElement(event) {
	// remove the link from the DOM
	document.body.removeChild(event.target);
}

var formFields = [
	"k3d_la_bedX",
	"k3d_la_bedY",
	"k3d_la_delta",
	"k3d_la_nozzleDiameter",
	"k3d_la_zOffset",
	"k3d_la_hotendTemperature",
	"k3d_la_bedTemperature",
	"k3d_la_cooling",
	"k3d_la_flow",
	"k3d_la_minPrintSpeed",
	"k3d_la_maxPrintSpeed",
	"k3d_la_initKFactor",
	"k3d_la_endKFactor",
	"k3d_la_smoothTime",
	"k3d_la_numSegments",
	"k3d_la_initAdditionalParameter",
	"k3d_la_endAdditionalParameter",
	"k3d_la_delta",
	"k3d_la_marlin",
	"k3d_la_klipper",
	"k3d_la_rrf",
	"k3d_la_targetNone",
	"k3d_la_targetSmoothTime",
	"k3d_la_targetAcceleration",
	"k3d_la_targetFlowrate",
	"k3d_la_acceleration",
	"k3d_la_startGcode",
	"k3d_la_endGcode",
	"k3d_la_numPerimeters",
];
var segmentFields = [
	"k3d_la_initKFactor",
	"k3d_la_endKFactor",
	"k3d_la_numSegments"
];
var segmentKeys = [
	"la_values",
	"num_segments"
];

var checkboxes = [
	"k3d_la_delta",
	"k3d_la_marlin",
	"k3d_la_klipper",
	"k3d_la_rrf",
	"k3d_la_targetNone",
	"k3d_la_targetSmoothTime",
	"k3d_la_targetAcceleration",
	"k3d_la_targetFlowrate"
];

var firmwareCheckboxes = [
	"k3d_la_marlin",
	"k3d_la_klipper",
	"k3d_la_rrf",
];

var targetCheckboxes = [
	"k3d_la_targetNone",
	"k3d_la_targetSmoothTime",
	"k3d_la_targetAcceleration",
	"k3d_la_targetFlowrate"
];


var saveForm = function () {
	for (var elementId of formFields) {
		var element = document.getElementById(elementId);
		if (element) {
			var saveValue = element.value;
			if (checkboxes.includes(elementId)) {
				saveValue = element.checked;
			}
			localStorage.setItem(elementId, saveValue);
		}
	}
}

function loadForm() {
	// Flag that indicates whether at least one checkbox was loaded from localStorage
	let anyCheckboxLoaded = false;

	for (var elementId of formFields) {
		let loadValue = localStorage.getItem(elementId);
		if (loadValue === undefined) {
			continue;
		}

		var element = document.getElementById(elementId);
		if (element) {
			if (checkboxes.includes(elementId)) {
				// We found a checkbox that was loaded from storage
				anyCheckboxLoaded = true;
				if (loadValue == 'true') {
					element.checked = true;
				} else {
					element.checked = false;
				}

			} else {
				if (loadValue != null) {
					element.value = loadValue;
				}
			}
		}
	}
}

function initForm() {
	loadForm();

	// Переменные для отслеживания: был ли хотя бы один checked в каждом списке
	let firmwareChecked = false;
	let targetChecked = false;

	// Обработка formFields (основные поля формы)
	for (var elementId of formFields) {
		var element = document.getElementById(elementId);
		element.addEventListener('change', function (e) {
			saveForm();

			var el = e.target;
			var id = el.id;

			if (segmentFields.indexOf(id) != -1) {
				checkSegments();
			} else {
				checkGo();
			}
		});
	}

	// Обработка segmentFields — слушатели фокуса
	for (var elementId of segmentFields) {
		var element = document.getElementById(elementId);
		element.addEventListener('focusin', function (e) {
			checkSegments();
		});
		element.addEventListener('focusout', function (e) {
			if (e.relatedTarget == undefined || segmentFields.indexOf(e.relatedTarget.id) === -1) {
				checkGo();
			}
		});
	}

	// Обработка firmwareCheckboxes и targetCheckboxes с подсчётом
	for (var id of firmwareCheckboxes) {
		var el = document.getElementById(id);
		if (el && el.checked) {
			firmwareChecked = true;
			break;  // уже нашли — можно не проверять остальные
		}
	}

	for (var id of targetCheckboxes) {
		var el = document.getElementById(id);
		if (el && el.checked) {
			targetChecked = true;
			break;
		}
	}

	// Установка значений по умолчанию, если ни один чекбокс не выбран
	if (!firmwareChecked) {
		var k3d_la_klipper = document.getElementById('k3d_la_klipper');
		if (k3d_la_klipper) k3d_la_klipper.checked = true;
	}

	if (!targetChecked) {
		var k3d_la_targetSmoothTime = document.getElementById('k3d_la_targetSmoothTime');
		if (k3d_la_targetSmoothTime) k3d_la_targetSmoothTime.checked = true;
	}

	// Навешиваем слушатели изменения на все чекбоксы из checkboxes
	for (var id of checkboxes) {
		var el = document.getElementById(id);
		if (el) {
			el.addEventListener('change', function (e) {
				updateInputDisability();
			});
		}
	}

	// Вызываем один раз сразу после загрузки страницы, чтобы установить корректное состояние
	updateInputDisability();
}

function initLang(key) {
	var values = window.lang.values;
	switch (key) {
		case 'en':
			values['header.title'] = 'K3D Pressure Advance Calibrator';

			values['table.header.parameter'] = 'Parameter';
			values['table.header.value'] = 'Value';
			values['table.header.description'] = 'Description';

			values['table.bed_size.title'] = 'Print bed<br>size';
			values['table.bed_size.description'] = '[mm] For Cartesian printers - the size of the print area along the X/Y axes<br>For delta printers - the diameter of the bed';
			values['table.firmware.title'] = 'Firmware';
			values['table.firmware.description'] = 'The firmware installed on your printer<br>If you don\'t know, it\'s most likely Klipper<br>For Bambu Lab printers select Marlin';
			values['table.delta.title'] = 'Origin at<br>center of bed';
			values['table.delta.description'] = 'For Cartesian printers this should be disabled<br>For deltas and AD5M enabled';
			values['table.nozzle_diameter.title'] = 'Nozzle diameter';
			values['table.nozzle_diameter.description'] = '[mm] Nozzle diameter installed on your printer. Affects:<br>- First layer line width = D * 1.5<br>- Model line width = D * 1.05<br>- Layer height = D * 0.5<br>When using additional calibration target "flow rate", model line width and layer thickness will be increased to 1.5*D and 0.75*D respectively<br>If you want thinner lines or layers, you can specify a smaller nozzle diameter';
			values['table.temperatures.title'] = 'Temperatures';
			values['table.temperatures.description'] = '[°C] Hotend and bed temperatures during printing of the calibration model<br>If the printer has an active thermal chamber, its temperature will be set equal to 0°C';
			values['table.fan_speed.title'] = 'Fan speed';
			values['table.fan_speed.description'] = '[%] Fan speed in percent<br>To prevent errors related to hotend cooling, the specified fan speed will only be reached by the 4th layer';
			values['table.flow.title'] = 'Flow';
			values['table.flow.description'] = '[%] Flow in percent. Needed to compensate for overall over- or under-extrusion';
			values['table.speed.title'] = 'Printing head<br>movement speed';
			values['table.speed.description'] = '[mm/s] Printing speed of slow and fast sections when checking PA value<br>Also, travel speed will be equal to fast section print speed, and first layer print speed will match slow section speed<br>If "flow" is selected as target, then fast section speeds will be redefined to achieve the specified flow rate';
			values['table.acceleration.title'] = 'Acceleration';
			values['table.acceleration.description'] = '[mm/s^2] Acceleration with which the test model will be printed<br>In general, it should be equal to the maximum acceleration obtained from input shaping calibration or based on absence of echo after corners<br>For adaptive PA calibration, it is recommended to perform 3 calibrations: at 1000 acceleration, half of maximum acceleration, and at maximum acceleration<br>It is not recommended to use values below 1000 since the methodology may work incorrectly';
			values['table.num_perimeters.title'] = 'Number of<br>perimeters';
			values['table.num_perimeters.description'] = 'How many perimeters the main part of the model will have<br>In general, 2 works well. To determine thickness variation through light inspection, set to 1. For elastomers, try 3–4 if the model behaves poorly at 2.';
			values['table.la_values.title'] = 'LA/PA coefficients';
			values['table.la_values.description'] = 'Initial and final LA/PA coefficient values for testing<br>On direct drive extruders, for solid material PA calibration, start with 0.0 - 0.1. If that range isn’t sufficient, check 0.0 - 0.2<br>For elastomer PA tuning, increase end value up to 0.5–1.0<br>On Bowden extruders, begin calibration with a range like 0.0–1.0';
			values['table.z_offset.title'] = 'Z-offset';
			values['table.z_offset.description'] = '[mm] Vertical offset applied to entire model to compensate incorrect distance between nozzle and bed on the first layer. Positive moves model upward, negative downward.<br>If your start G-code already sets Z-offset, leave this field at 0.';
			values['table.num_segments.title'] = 'Number of<br>segments';
			values['table.num_segments.description'] = 'Number of segments in tower structure. During each segment, LA/PA factor remains constant. Segments are visually separated for easier analysis of the model<br>Generally recommended to use 11 for quick checks or 21 for wider value ranges.';
			values['table.additional_calibration_target.title'] = 'Additional<br>calibration target';
			values['table.additional_calibration_target.description'] = 'With all add-on targets, PA factor changes vertically. The add-on parameter varies from minimum at rear wall to maximum at left wall (clockwise print direction)<br>No – no extra parameters change; walls remain identical<br>Smoothing time – smoothing time between walls changes. Works only on Klipper firmware<br>Acceleration – print acceleration changes between walls<br>Flow rate – model generated with increased layer thickness and line width. Fast segment speeds adjusted so flow matches given parameter value';
			values['table.additional_parameter.title'] = 'Values of<br>additional<br>parameter';
			values['table.additional_parameter.description'] = 'Minimum and maximum values of the additional calibration target parameter<br>If additional target is off, these do nothing';
			values['table.start_gcode.title'] = 'Start G-code';
			values['table.gcodes.description'] = 'To properly initialize and complete the print job, correct G-codes must be entered for your printer. Easiest way is copying them from your slicer. If unsupported elements exist in G-codes, they need manual removal — replacing unsupported placeholders with actual values, calculating math expressions, expanding conditions etc.<hr><u>Supported placeholders:</u><br>- <code>$HOTTEMP</code>, <code>{temperature}</code>, <code>{first_layer_temperature}</code>, <code>{nozzle_temperature}</code>, and similar will be replaced with hotend temp;<br>- <code>$BEDTEMP</code>, <code>{bed_temperature}</code>, <code>{first_layer_bed_temperature}</code>, <code>{hot_plate_temp}</code>, <code>{textured_plate_temp_initial_layer}</code>, and similar will be replaced with bed temp;<br>- <code>{chamber_temperature}</code>, <code>{chamber_minimal_temperature}</code> will be replaced with 0;<br>- <code>$FLOW</code> will be replaced with flow value;<br>- Various variations of above placeholders such as square brackets instead of curly braces, array index notation inside braces, arithmetic operations within braces, etc.<hr><u>Unsupported features:</u><br>- Other placeholders;<br>- Conditional statements (<code>if</code>);<br>- Ternary operators (<code>a ? b : c</code>);<br>- Mathematical functions and expressions';
			values['table.end_gcode.title'] = 'End G-code';
			values['table.smooth_time.title'] = 'PA Smoothing Time';
			values['table.smooth_time.description'] = '[s] Pressure advance smoothing time<br>In Klipper firmware reduces load on feeder mechanism but excessive smoothing may cause defects like dents before and after seams. So this value should be calibrated via setting additional target “smoothing time”<br>In Marlin and RRF there\'s no control over PA smoothing time, hence this parameter does nothing there<br>When selecting “smoothing time” as additional calibration target, this parameter’s value will be ignored';

			values['generator.generate_and_download'] = 'Generate & Download';
			values['generator.generate_button_loading'] = 'Generator loading...';
			values['generator.segment'] = '; Segment %d: K-Factor: %s\n';
			values['generator.additional_target.smooth_time'] = '; Smoothing Time:\n';
			values['generator.additional_target.acceleration'] = '; Acceleration:\n';
			values['generator.additional_target.flowrate'] = '; Volumetric Flow Rate:\n';
			values['generator.additional_target.side'] = '; Back: %s\n; Right: %s\n; Front: %s\n; Left: %s\n';
			values['generator.reset_to_default'] = 'Reset Settings';

			values['error.bed_size_x.format'] = 'Format error: Print bed size along X-axis';
			values['error.bed_size_x.value'] = 'Value error: Print bed size along X-axis (less than 100 or more than 1000 mm)';
			values['error.bed_size_y.format'] = 'Format error: Print bed size along Y-axis';
			values['error.bed_size_y.value'] = 'Value error: Print bed size along Y-axis (less than 100 or more than 1000 mm)';
			values['error.nozzle_diameter.format'] = 'Format error: Nozzle diameter';
			values['error.nozzle_diameter.value'] = 'Value error: Nozzle diameter (less than 0.1 or more than 2.0 mm)';
			values['error.hotend_temp.format'] = 'Format error: Hotend temperature';
			values['error.hotend_temp.too_low'] = 'Value error: Hotend temperature (below 100°C)';
			values['error.hotend_temp.too_high'] = 'Value error: Hotend temperature (above 500°C)';
			values['error.bed_temp.format'] = 'Format error: Bed temperature';
			values['error.bed_temp.too_high'] = 'Value error: Bed temperature (above 150°C)';
			values['error.fan_speed.format'] = 'Format error: Fan speed';
			values['error.num_segments.format'] = 'Format error: Number of segments';
			values['error.num_segments.value'] = 'Value error: Number of segments (less than 2 or more than 100)';
			values['error.z_offset.format'] = 'Format error: Z-offset';
			values['error.z_offset.value'] = 'Value error: Z-offset (less than -0.5 or more than 0.5 mm)';
			values['error.flow.format'] = 'Format error: Flow';
			values['error.flow.value'] = 'Value error: Flow (less than 50 or more than 150%)';
			values['error.min_print_speed.format'] = 'Format error: Slow section movement speed';
			values['error.min_print_speed.value'] = 'Value error: Slow section movement speed (less than 5 or more than 1000 mm/s)';
			values['error.max_print_speed.format'] = 'Format error: Fast section movement speed';
			values['error.max_print_speed.value'] = 'Value error: Fast section movement speed (less than 5 or more than 1000 mm/s)';
			values['error.firmware.not_set'] = 'Error: Firmware not selected';
			values['error.acceleration.format'] = 'Format error: Acceleration';
			values['error.acceleration.value'] = 'Value error: Acceleration (less than 500 or more than 50000 mm/s²)';
			values['error.num_perimeters.format'] = 'Format error: Number of perimeters';
			values['error.num_perimeters.value'] = 'Value error: Number of perimeters (less than 1 or more than 5)';
			values['error.init_la.format'] = 'Format error: Initial LA/PA coefficient value';
			values['error.init_la.value'] = 'Value error: Initial LA coefficient value (less than 0.0 or more than 2.0)';
			values['error.end_la.format'] = 'Format error: Final LA/PA coefficient value';
			values['error.end_la.value'] = 'Value error: Final LA/PA coefficient value (less than 0.0 or more than 2.0)';
			values['error.smooth_time.format'] = 'Format error: LA/PA smoothing time';
			values['error.smooth_time.value'] = 'Value error: LA/PA smoothing time (less than 0.005 or more than 0.2)';
			values['error.additional_target.not_set'] = 'Error: Additional calibration target not selected';
			values['error.init_additional_parameter.format'] = 'Format error: Initial additional parameter value';
			values['error.init_additional_parameter.value'] = 'Value error: Initial additional parameter value. Valid ranges: smoothing time 0.001–1.0, acceleration 500–50000, flow rate 1–200';
			values['error.end_additional_parameter.format'] = 'Format error: Final additional parameter value';
			values['error.end_additional_parameter.value'] = 'Value error: Final additional parameter value. Valid ranges: smoothing time 0.001–1.0, acceleration 500–50000, flow rate 1–200';
			values['error.unsupported_placeholders'] = 'Error: Unsupported placeholders detected in Start or End G-code:';
			values['error.unsupported_conditions'] = 'Error: Conditional statements (if) detected in Start or End G-code';
			values['error.unsupported_ternary_operators'] = 'Error: Ternary operators (a ? b : c) detected in Start or End G-code';
			values['error.unsupported_math'] = 'Error: Math operations (min, max, round, etc.) detected in Start or End G-code';
			values['error.model_size_limited'] = 'Model size must be increased to allow the print head to accelerate to the speed required for testing the specified volumetric flow rate. However, the resulting part size turned out to be larger than the print bed. Increase acceleration or decrease the volumetric flow rate values for testing';

			values['info.delta_print_speed_too_small'] = 'Speed difference across one wall too small. Consider increasing initial volumetric flow rate for better results';
			values['info.smooth_time_not_supported_by_firmware'] = 'Pressure advance smoothing time not supported by selected firmware';
			values['info.layer_change_calibration_info'] = ';Layer:%s PA:%s';
			values['info.model_size_changed'] = 'Model size has been increased to allow the print head to accelerate to the speed required for testing the specified volumetric flow rate. If the model size does not suit you, increase acceleration or decrease the maximum value of the volumetric flow rate being tested. Model size: ';
			break;
		case 'ru':
			values['header.title'] = 'K3D калибровщик Pressure Advance';

			values['table.header.parameter'] = 'Параметр';
			values['table.header.value'] = 'Значение';
			values['table.header.description'] = 'Описание';

			values['table.bed_size.title'] = 'Размер области<br> печати';
			values['table.bed_size.description'] = '[мм] Для декартовых принтеров - размер области печати по осям X/Y<br>Для дельта-принтеров - диаметр стола';
			values['table.firmware.title'] = 'Прошивка';
			values['table.firmware.description'] = 'Прошивка, установленная на вашем принтере<br>Если не знаете, то, скорее всего, Klipper<br>Для принтеров Bambu Lab выберите Marlin';
			values['table.delta.title'] = 'Начало координат<br>в центре стола';
			values['table.delta.description'] = 'Для декартовых принтеров должно быть выключено<br>Для дельт и AD5M включено';
			values['table.nozzle_diameter.title'] = 'Диаметр сопла';
			values['table.nozzle_diameter.description'] = '[мм] Диаметр сопла, установленного на вашем принтере. Влияет на:<br>- Ширина линии первого слоя = D * 1.5<br>- Ширина линии модели = D * 1.05<br>- Толщина слоя = D * 0.5<br>При использовании доп. цели калибровки "расход", ширина линии модели и толщина слоя будут увеличены до 1.5*D и 0.75*D соответственно<br>Если хотите использовать линии уже или слой тоньше, то можете указать заниженный диаметр сопла';
			values['table.temperatures.title'] = 'Температуры';
			values['table.temperatures.description'] = '[°C] Температуры хотэнда и стола при печати калибровочной модели<br>Если в принтере есть активная термокамера, то её температура будет установлена равной 0°C';
			values['table.fan_speed.title'] = 'Скорость<br>вентилятора';
			values['table.fan_speed.description'] = '[%] Обороты вентилятора в процентах<br>Для предотвращения ошибок, связанных с охлаждением хотэнда, указанная скорость вентилятора будет достигнута только к 4 слою';
			values['table.flow.title'] = 'Поток';
			values['table.flow.description'] = '[%] Поток в процентах. Нужен для компенсации общей пере- или недоэкструзии';
			values['table.speed.title'] = 'Скорость движения<br>печатающей головы';
			values['table.speed.description'] = '[мм/с] Скорость печати медленных (slow) и быстрых (fast) участков при проверке значения PA<br>Также, скорость холостых перемещений будет равна скорости печати быстрых участков, а скорость печати первого слоя - скорости печати медленных участков<br>Если выбрана цель "расход", то скорость быстрых участков будет переопределена, чтобы достичь указанного расхода';
			values['table.acceleration.title'] = 'Ускорение';
			values['table.acceleration.description'] = '[мм/с^2] Ускорение, с которым будет напечатана тестовая модель<br>В общем случае стоит поставить равной максимальному ускорению вашего принтера, полученному при калибровке input shaping\'а или по критерию отстуствия эхо после углов<br>При калибровке адаптивного PA рекомендуется провести 3 калибровки: при ускорении 1000, половине от максимального ускорения и при максимальном ускорении<br>Не рекомендуется ставить значения ниже 1000 т.к. тогда методика может отработать некорректно';
			values['table.num_perimeters.title'] = 'Количество<br>периметров';
			values['table.num_perimeters.description'] = 'Сколько периметров будет у основно части модели<br>В общем случае хорошо работает 2. Для определения перепада толщины на просвет поставьте 1. Для эластомеров можно попробовать 3-4, если модель ведёт при 2';
			values['table.la_values.title'] = 'Коэффициенты<br>LA/PA';
			values['table.la_values.description'] = 'Начальное и конечное значение коэффициентов LA/PA для проверки<br>На принтерах с директ экструдерами, при калибровке PA для твёрдого материала, стоит установить 0.0 - 0.1. Если этого диапазона не хватит, то проверить 0.0 - 0.2<br>При подборе PA для эластомеров, стоит увеличить конечное значение до 0.5-1.0<br>На принтерах с боуден экструдерами, начинать калибровку лучше с диапазона 0.0-1.0';
			values['table.z_offset.title'] = 'Z-offset';
			values['table.z_offset.description'] = '[мм] Смещение всей модели по вертикали для компенсации неправильного расстояния от сопла до стола на первом слое. Положительное значение смещает модель вверх, отрицательное - вниз.<br>Если в вашем стартовом G-коде уже есть установка Z-offset\'а, то оставьте 0';
			values['table.num_segments.title'] = 'Количество<br>сегментов';
			values['table.num_segments.description'] = 'Количество сегментов башенки. В течение сегмента коэффициент LA/PA остаётся неизменным. Сегменты визуально разделены для упрощения анализа модели<br>В общем случае рекомендуется установить 11 для быстрой проверки, или 21 для проверки с более широким диапазоном значений';
			values['table.additional_calibration_target.title'] = 'Дополнительная<br>цель калибровки';
			values['table.additional_calibration_target.description'] = 'При всех доп. целях, коэффициент PA меняется по высоте. А значение доп. параметра меняется от минимального на задней   стенке модели, до максимального на левой стенке (направление печати по часовой стрелке)<br>Нет - никакого доп. параметра меняться не будет, все стенки одинаковые<br>Время сглаживания - между стенками будет меняться время сглаживания PA. Работает только для прошивки Klipper<br>Ускорение - между стенками будет меняться ускорение печати<br>Расход - модель будет сгенерирована с увеличенными толщиной слоя и шириной линии. Скорость быстрых участков будет меняться так, чтобы расход на них соответствовал значению доп. параметра';
			values['table.additional_parameter.title'] = 'Значения<br>дополнительного<br>параметра';
			values['table.additional_parameter.description'] = 'Минимальное и максимальное значение параметра дополнительной цели калибровки<br>Если дополнительная цель выключена, то не делают ничего';
			values['table.start_gcode.title'] = 'Начальный G-код';
			values['table.gcodes.description'] = 'Для того, чтобы печать была правильно инициализирована и завершена, необходимо указать корректные для вашего принтера G-коды. Проще всего скопировать их из вашего слайсера. Если G-коды в слайсере не содержат неподдерживаемых элементов, то их можно копировать без изменений. Если неподдерживаемые элементы есть, то придётся от них избавиться - вручную заменить неподдерживаемые плейсхолдеры на значения, посчитать результаты математических функций, раскрыть условия и т.д.<hr><u>Поддерживаемые плейсхолдеры:</u><br>- <code>$HOTTEMP</code>, <code>{temperature}</code>, <code>{first_layer_temperature}</code>, <code>{nozzle_temperature}</code> и подобные будут заменены на температуру хотэнда;<br>- <code>$BEDTEMP</code>, <code>{bed_temperature}</code>, <code>{first_layer_bed_temperature}</code>, <code>{hot_plate_temp}</code>, <code>{textured_plate_temp_initial_layer}</code> и подобные будут заменены на температуру стола;<br>- <code>{chamber_temperature}</code>, <code>{chamber_minimal_temperature}</code> будут заменены на 0;<br>- <code>$FLOW</code> будет заменён на значение потока;<br>- Разные вариации указанных выше плейсхолдеров. Например, в квадратных скобках вместо фигурных, с указанием номера элемента массива, с арифметикой внутри фигурных скобок и т.д.<hr><u>Не поддерживаются:</u><br>- Другие плейсхолдеры;<br>- Условные операторы (<code>if</code>);<br>- Тернарные операторы (<code>a ? b : c</code>);<br>- Математические функции и выражения';
			values['table.end_gcode.title'] = 'Конечный G-код';
			values['table.smooth_time.title'] = 'Время сглаживания<br>PA';
			values['table.smooth_time.description'] = '[с] Время сглаживания PA<br>В прошивке Klipper уменьшает нагрузку на подающий механизм, но завышенное значение сглаживания может вызвать дефекты в виде ямок перед и после шва. Поэтому это значение стоит откалибровать с помощью установки дополнительной цели "время сглаживания"<br>В Marlin и RRF нет управления временем сглаживания PA, поэтому там этот параметр ничего не делает<br>При выборе доп. цели калибровки "время сглаживания", значение этого параметра будет проигнорировано';

			values['generator.generate_and_download'] = 'Генерировать и скачать';
			values['generator.generate_button_loading'] = 'Генератор загружается...';
			values['generator.segment'] = '; Сегмент %d: K-Factor: %s\n';
			values['generator.additional_target.smooth_time'] = '; Время сглаживания:\n';
			values['generator.additional_target.acceleration'] = '; Ускорение:\n';
			values['generator.additional_target.flowrate'] = '; Объёмный расход:\n';
			values['generator.additional_target.side'] = '; Сзади: %s\n; Справа: %s\n; Спереди: %s\n; Слева: %s\n';
			values['generator.reset_to_default'] = 'Сбросить настройки';

			values['error.bed_size_x.format'] = 'Ошибка формата: Размер области печати по оси Х';
			values['error.bed_size_x.value'] = 'Ошибка значения: Размер области печати по оси X (меньше 100 или больше 1000 мм)';
			values['error.bed_size_y.format'] = 'Ошибка формата: Размер области печати по оси Y';
			values['error.bed_size_y.value'] = 'Ошибка значения: Размер области печати по оси Y (меньше 100 или больше 1000 мм)';
			values['error.nozzle_diameter.format'] = 'Ошибка формата: Диаметр сопла';
			values['error.nozzle_diameter.value'] = 'Ошибка значения: Диаметр сопла (меньше 0.1 или больше 2.0 мм)';
			values['error.hotend_temp.format'] = 'Ошибка формата: Температура хотэнда';
			values['error.hotend_temp.too_low'] = 'Ошибка значения: Температура хотэнда (меньше 100°C)';
			values['error.hotend_temp.too_high'] = 'Ошибка значения: Температура хотэнда (выше 500°C)';
			values['error.bed_temp.format'] = 'Ошибка формата: Температура стола';
			values['error.bed_temp.too_high'] = 'Ошибка значения: Температура стола (выше 150°C)';
			values['error.fan_speed.format'] = 'Ошибка формата: Скорость вентилятора';
			values['error.num_segments.format'] = 'Ошибка формата: Количество сегментов';
			values['error.num_segments.value'] = 'Ошибка значения: Количество сегментов (меньше 2 или больше 100)';
			values['error.z_offset.format'] = 'Ошибка формата: Z-offset';
			values['error.z_offset.value'] = 'Ошибка значения: Z-offset (меньше -0.5 или больше 0.5 мм)';
			values['error.flow.format'] = 'Ошибка формата: Поток';
			values['error.flow.value'] = 'Ошибка значения: Поток (меньше 50 или больше 150%)';
			values['error.min_print_speed.format'] = 'Ошибка формата: Скорость движения печатающей головы на медленных участках';
			values['error.min_print_speed.value'] = 'Ошибка значения: Скорость движения печатающей головы на медленных участках (менее 5 или более 1000 мм/с)';
			values['error.max_print_speed.format'] = 'Ошибка формата: Скорость движения печатающей головы на быстрых участках';
			values['error.max_print_speed.value'] = 'Ошибка значения: Скорость движения печатающей головы на быстрых участках (менее 5 или более 1000 мм/с)';
			values['error.firmware.not_set'] = 'Ошибка: прошивка не выбрана';
			values['error.acceleration.format'] = 'Ошибка формата: ускорение';
			values['error.acceleration.value'] = 'Ошибка значения: ускорение (меньше 500 или больше 50000 мм/с^2)';
			values['error.num_perimeters.format'] = 'Ошибка формата: количество периметров';
			values['error.num_perimeters.value'] = 'Ошибка значения: количество периметров (меньше 1 или больше 5)';
			values['error.init_la.format'] = 'Ошибка формата: Начальное значение коэффициента LA/PA';
			values['error.init_la.value'] = 'Ошибка значения: Начальное значение коэффициента LA (меньше 0.0 или больше 2.0)';
			values['error.end_la.format'] = 'Ошибка формата: Конечное значение коэффициента LA/PA';
			values['error.end_la.value'] = 'Ошибка значения: Конечное значение коэффициента LA/PA (меньше 0.0 или больше 2.0)';
			values['error.smooth_time.format'] = 'Ошибка формата: Время сглаживания LA/PA';
			values['error.smooth_time.value'] = 'Ошибка значения: Время сглаживания LA/PA (меньше 0.005 или больше 0.2)';
			values['error.additional_target.not_set'] = 'Ошибка: Не выбрана дополнительная цель калибровки';
			values['error.init_additional_parameter.format'] = 'Ошибка формата: Начальное значение дополнительного параметра';
			values['error.init_additional_parameter.value'] = 'Ошибка значения: Начальное значение дополнительного параметра. Для времени сглаживания допустимы значения от 0.001 до 1.0, для ускорения от 500 до 50000, для расхода от 1 до 200';
			values['error.end_additional_parameter.format'] = 'Ошибка формата: Конечное значение дополнительного параметра';
			values['error.end_additional_parameter.value'] = 'Ошибка значения: Конечное значение дополнительного параметра. Для времени сглаживания допустимы значения от 0.001 до 1.0, для ускорения от 500 до 50000, для расхода от 1 до 200';
			values['error.unsupported_placeholders'] = 'Ошибка: В начальном или конечном G-коде замечены неподдерживаемые плейсхолдеры:';
			values['error.unsupported_conditions'] = 'Ошибка: В начальном или конечном G-коде замечены условные операторы (if)';
			values['error.unsupported_ternary_operators'] = 'Ошибка: В начальном или конечном G-коде замечены тернарные операторы (a ? b : c)';
			values['error.unsupported_math'] = 'Ошибка: В начальном или конечном G-коде замечены математические операции (min, max, round и подобные)';
			values['error.model_size_limited'] = 'Размер модели должен быть увеличен для того, чтобы печатающая голова успевала разогнаться до скорости, необходимой для проверки заданного объёмного расхода. Но получаемый размер детали оказался больше стола. Увеличьте ускорение или уменьшите значения объёмного расхода для проверки';

			values['info.delta_print_speed_too_small'] = 'Перепад скоростей на одной из стенок слишком мал. Рекомендуется увеличить начальное значение объёмного расхода для проверки';
			values['info.smooth_time_not_supported_by_firmware'] = 'Время сглаживания pressure advance не поддерживается выбранной прошивкой';
			values['info.layer_change_calibration_info'] = ';Слой:%s PA:%s';
			values['info.model_size_changed'] = 'Размер модели был увеличен для того, чтобы печатающая голова успевала разогнаться до скорости, необходимой для проверки заданного объёмного расхода. Если вас не устраивает размер модели, то увеличьте ускорения или уменьшите максимальное значение проверяемого объёмного расхода. Размер модели: ';
			break;
	}

	document.title = window.lang.getString('header.title');
	var el = document.getElementsByClassName('lang');
	for (var i = 0; i < el.length; i++) {
		var item = el[i];
		item.innerHTML = window.lang.getString(item.id);
	}
	document.getElementById('generateButton').innerHTML = window.lang.getString('generator.generate_and_download');
	document.getElementById('resetButton').innerHTML = window.lang.getString('generator.reset_to_default');
	document.getElementById('generateButtonLoading').innerHTML = window.lang.getString('generator.generate_button_loading');
}

function setSegmentsPreview(segments) {
	savedSegmentsInfo = segments;
	setSegmentsPreviewVisible(true);
}

function setSegmentsPreviewVisible(visible) {
	if (savedSegmentsInfo == null || savedSegmentsInfo == undefined) {
		visible = false;
	}
	if (visible) {
		document.getElementById('table.' + segmentKeys[0] + '.description').rowSpan = segmentKeys.length;
		document.getElementById('table.' + segmentKeys[0] + '.description').innerHTML = '<span>' + savedSegmentsInfo.replaceAll('\n', '<br>') + '</span>';

		for (var i = 1; i < segmentKeys.length; i++) {
			document.getElementById('table.' + segmentKeys[i] + '.description').style.display = 'none';
		}
	} else {
		document.getElementById('table.' + segmentKeys[0] + '.description').rowSpan = 1;
		for (var i = 0; i < segmentKeys.length; i++) {
			var id = 'table.' + segmentKeys[i] + '.description';
			document.getElementById(id).style.display = '';
			document.getElementById(id).innerHTML = window.lang.getString(id);
		}
	}
}

function reset() {
	for (var elementId of formFields) {
		localStorage.removeItem(elementId);
	}

	window.location.reload(false);
}

function init() {
	initForm();

	const urlParams = new URLSearchParams(window.location.search);
	var lang = urlParams.get('lang');
	if (lang == undefined) {
		lang = 'ru';
	}
	console.log("lang=" + lang);

	window.lang = {
		values: {},
		getString: function (key) {
			var ret = window.lang.values[key];
			if (key == 'header.title') {
				return ret + ' ' + calibrator_version;
			}
			return ret;
		}
	};

	initLang(lang);

	setTimeout(function () {
		if (checkGo != undefined && window.lang != undefined) {
			checkGo();
		} else {
			setTimeout(this, 100);
		}
	}, 100);
}

function updateInputDisability() {
	var klipperRadio = document.getElementById('k3d_la_klipper');
	var targetSmoothRadio = document.getElementById('k3d_la_targetSmoothTime');
	var targetNone = document.getElementById('k3d_la_targetNone');
	var smoothTimeInput = document.getElementById('k3d_la_smoothTime');
	var initAdditionalParameterInput = document.getElementById('k3d_la_initAdditionalParameter');
	var endAdditionalParameterInput = document.getElementById('k3d_la_endAdditionalParameter');

	if (klipperRadio && klipperRadio.checked &&	targetSmoothRadio && !targetSmoothRadio.checked) {
		smoothTimeInput.disabled = false;
	} else {
		smoothTimeInput.disabled = true;
	}

	if (targetNone && targetNone.checked && initAdditionalParameterInput && endAdditionalParameterInput) {
		initAdditionalParameterInput.disabled = true;
		endAdditionalParameterInput.disabled = true;
	} else {
		initAdditionalParameterInput.disabled = false;
		endAdditionalParameterInput.disabled = false;
	}
}
