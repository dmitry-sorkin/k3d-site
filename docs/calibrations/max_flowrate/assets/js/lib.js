const calibrator_version = 'v0.3-beta';
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
    output.cols = "80";
    output.value = value;
    
    // Считаем строки аналитически
    var lines = value.split('\n');
    var totalLines = 0;
    
    for (var i = 0; i < lines.length; i++) {
        // Для каждой строки считаем сколько ей нужно строк при ширине 80 символов
        totalLines += Math.ceil(lines[i].length / 80);
    }
    
    // Устанавливаем количество строк
    output.rows = Math.min(Math.max(3, totalLines), 50) + 1;
    
    container.appendChild(output);
}

function destroyClickedElement(event) {
	// remove the link from the DOM
	document.body.removeChild(event.target);
}

var formFields = [
	"k3d_mfr_bedX",
	"k3d_mfr_bedY",
	"k3d_mfr_delta",
	"k3d_mfr_nozzleDiameter",
	"k3d_mfr_zOffset",
	"k3d_mfr_hotendTemperature",
	"k3d_mfr_bedTemperature",
	"k3d_mfr_initFlowrate",
	"k3d_mfr_endFlowrate",
	"k3d_mfr_flowrateDelta",
	"k3d_mfr_extrusionLength",
	"k3d_mfr_marlin",
	"k3d_mfr_klipper",
	"k3d_mfr_rrf",
	"k3d_mfr_startGcode",
	"k3d_mfr_endGcode",
];

var checkboxes = [
	"k3d_mfr_delta",
	"k3d_mfr_marlin",
	"k3d_mfr_klipper",
	"k3d_mfr_rrf",
];

var firmwareCheckboxes = [
	"k3d_mfr_marlin",
	"k3d_mfr_klipper",
	"k3d_mfr_rrf",
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
	// Flag that indicates whether at least one advancedParametersCheckbox was loaded from localStorage
	let anyCheckboxLoaded = false;

	for (var elementId of formFields) {
		let loadValue = localStorage.getItem(elementId);
		if (loadValue === undefined) {
			continue;
		}

		var element = document.getElementById(elementId);
		if (element) {
			if (checkboxes.includes(elementId)) {
				// We found a advancedParametersCheckbox that was loaded from storage
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

			
				checkGo();
		});
	}

	// Обработка firmwareCheckboxes с подсчётом
	for (var id of firmwareCheckboxes) {
		var el = document.getElementById(id);
		if (el && el.checked) {
			firmwareChecked = true;
			break;  // уже нашли — можно не проверять остальные
		}
	}

	// Установка значений по умолчанию, если ни один чекбокс не выбран
	if (!firmwareChecked) {
		var k3d_mfr_klipper = document.getElementById('k3d_mfr_klipper');
		if (k3d_mfr_klipper) k3d_mfr_klipper.checked = true;
	}
}

function initLang(key) {
	var values = window.lang.values;
	switch (key) {
		case 'en':
			values['header.title'] = 'K3D Flow Rate Calibrator';
			values['profile.select.label'] = 'Profile: ';
			values['lang.label'] = 'Language: ';

			values['table.bed_size.title'] = 'Print Area Size';
			values['table.bed_size.description'] = '[mm] For Cartesian printers - print area size along X/Y axes<br>For Delta printers - bed diameter';
			values['table.firmware.title'] = 'Firmware';
			values['table.firmware.description'] = 'Firmware installed on your printer<br>If you don\'t know, it\'s most likely Klipper<br>For Bambu Lab printers, select Marlin';
			values['table.delta.title'] = 'Coordinate Origin<br>at Bed Center';
			values['table.delta.description'] = 'For Cartesian printers should be OFF<br>For Deltas and AD5M should be ON';
			values['table.nozzle_diameter.title'] = 'Nozzle Diameter';
			values['table.nozzle_diameter.description'] = '[mm] Diameter of the nozzle installed on your printer. Affects line width and layer height when printing samples and helper elements';
			values['table.temperatures.title'] = 'Temperatures';
			values['table.temperatures.description'] = '[°C] Hotend and bed temperatures when printing the calibration model<br>If the printer has an active chamber, its temperature will be set to 0°C';
			values['table.flowrate.title'] = 'Volumetric Flow Rate';
			values['table.flowrate.description'] = '[mm^3/s] Initial and final volumetric flow rate value for testing<br>- For V6-like hotends, recommended from 4 to 20<br>- For Volcano-like, recommended from 4 to 30<br>- For Goliath-like, from 4 to 55<br>If a CHT nozzle is installed, consider increasing the maximum test value';
			values['table.z_offset.title'] = 'Z-offset';
			values['table.z_offset.description'] = '[mm] Vertical offset of the entire model to compensate for incorrect nozzle-to-bed distance on the first layer. Positive value moves the model up, negative - down.<br>If your start G-code already sets a Z-offset, leave it as 0';
			values['table.speed.title'] = 'Print Head Movement Speed';
			values['table.speed.description'] = '[mm/s] Print head movement speed during printing (print) and travel moves (travel).<br>For print speed, it is recommended to set lower values to minimize the impact of acceleration and deceleration on the average volumetric flow rate during sample printing';
			values['table.flowrate_delta.title'] = 'Flow Rate Step';
			values['table.flowrate_delta.description'] = '[mm^3/s] Difference in volumetric flow rate between adjacent test elements. The smaller the step, the more samples will be printed';
			values['table.extrusion_length.title'] = 'Extrusion Length<br>per Sample';
			values['table.extrusion_length.description'] = '[mm] The length of filament to be extruded for printing each sample. Measured at the input to the feed mechanism';
			values['table.start_gcode.title'] = 'Start G-code';
			values['table.gcodes.description'] = 'For printing to be correctly initialized and finalized, you need to provide G-codes suitable for your printer. The easiest way is to copy them from your slicer. If the slicer\'s G-codes do not contain unsupported elements, you can copy them unchanged. If there are unsupported elements, you will have to remove them - manually replace unsupported placeholders with values, calculate results of mathematical functions, expand conditions, etc.<hr><u>Supported placeholders:</u><br>- <code>$HOTTEMP</code>, <code>{temperature}</code>, <code>{first_layer_temperature}</code>, <code>{nozzle_temperature}</code> and similar will be replaced with the hotend temperature;<br>- <code>$BEDTEMP</code>, <code>{bed_temperature}</code>, <code>{first_layer_bed_temperature}</code>, <code>{hot_plate_temp}</code>, <code>{textured_plate_temp_initial_layer}</code> and similar will be replaced with the bed temperature;<br>- <code>{chamber_temperature}</code>, <code>{chamber_minimal_temperature}</code> will be replaced with 0;<br>- Various modifications of the placeholders mentioned above. For example, using square brackets instead of curly braces, specifying an array element index, arithmetic inside curly braces, etc.<hr><u>Not supported:</u><br>- Other placeholders;<br>- Conditional statements (<code>if</code>);<br>- Ternary operators (<code>a ? b : c</code>);<br>- Mathematical functions and expressions';
			values['table.end_gcode.title'] = 'End G-code';

			values['generator.generate_and_download'] = 'Generate and Download';
			values['generator.generate_button_loading'] = 'Generator is loading...';
			values['generator.reset_to_default'] = 'Reset Settings';

			values['error.bed_size_x.format'] = 'Format Error: Print Area Size along X axis';
			values['error.bed_size_x.value'] = 'Value Error: Print Area Size along X axis (less than 100 or greater than 1000 mm)';
			values['error.bed_size_y.format'] = 'Format Error: Print Area Size along Y axis';
			values['error.bed_size_y.value'] = 'Value Error: Print Area Size along Y axis (less than 100 or greater than 1000 mm)';
			values['error.nozzle_diameter.format'] = 'Format Error: Nozzle Diameter';
			values['error.nozzle_diameter.value'] = 'Value Error: Nozzle Diameter (less than 0.1 or greater than 2.0 mm)';
			values['error.hotend_temp.format'] = 'Format Error: Hotend Temperature';
			values['error.hotend_temp.value'] = 'Value Error: Hotend Temperature (less than 100°C or greater than 500°C)';
			values['error.bed_temp.format'] = 'Format Error: Bed Temperature';
			values['error.bed_temp.value'] = 'Value Error: Bed Temperature (greater than 150°C)';
			values['error.init_flowrate.format'] = 'Format Error: Initial Volumetric Flow Rate value';
			values['error.init_flowrate.value'] = 'Value Error: Initial Volumetric Flow Rate value (less than 1.0 or greater than 50.0 mm^3/s)';
			values['error.end_flowrate.format'] = 'Format Error: Final Volumetric Flow Rate value';
			values['error.end_flowrate.value'] = 'Value Error: Final Volumetric Flow Rate value (less than 1.0 or greater than 50.0 mm^3/s)';
			values['error.flowrate_delta.format'] = 'Format Error: Flow Rate Step';
			values['error.flowrate_delta.value'] = 'Value Error: Flow Rate Step (less than 0.5 or greater than 10.0)';
			values['error.extrusion_length.format'] = 'Format Error: Extrusion Length per Sample';
			values['error.extrusion_length.value'] = 'Value Error: Extrusion Length per Sample (less than 50 or greater than 5000 mm)';
			values['error.z_offset.format'] = 'Format Error: Z-offset';
			values['error.z_offset.value'] = 'Value Error: Z-offset (less than -0.5 or greater than 0.5 mm)';
			values['error.speed.format'] = 'Format Error: Print Head Movement Speed during sample printing';
			values['error.speed.value'] = 'Value Error: Print Head Movement Speed during sample printing (less than 5 or greater than 500 mm/s)';
			values['error.travel_speed.format'] = 'Format Error: Print Head Movement Speed during travel moves';
			values['error.travel_speed.value'] = 'Value Error: Print Head Movement Speed during travel moves (less than 5 or greater than 1000 mm/s)';
			values['error.firmware.not_set'] = 'Error: Firmware not selected';
			values['error.unsupported_placeholders'] = 'Error: Unsupported placeholders detected in Start or End G-code:';
			values['error.unsupported_conditions'] = 'Error: Conditional statements (if) detected in Start or End G-code';
			values['error.unsupported_ternary_operators'] = 'Error: Ternary operators (a ? b : c) detected in Start or End G-code';
			values['error.unsupported_math'] = 'Error: Mathematical operations (min, max, round, etc.) detected in Start or End G-code';
			values['error.different_versions'] = 'Version mismatch: Please hard refresh (Ctrl + F5) or clear your browser cache.';

			values['info.out_of_print_area'] = 'Some samples did not fit on the bed. Number of skipped samples: %s. Increase the test value step or reduce the speed\n';
			values['info.cali_params_template'] = ';Sample: %d Flow Rate: %.3f mm^3/s\n';

			break;
		case 'ru':
			values['header.title'] = 'K3D калибровщик расхода';
			values['profile.select.label'] = 'Профиль: ';
			values['lang.label'] = 'Язык: ';

			values['table.bed_size.title'] = 'Размер области<br> печати';
			values['table.bed_size.description'] = '[мм] Для декартовых принтеров - размер области печати по осям X/Y<br>Для дельта-принтеров - диаметр стола';
			values['table.firmware.title'] = 'Прошивка';
			values['table.firmware.description'] = 'Прошивка, установленная на вашем принтере<br>Если не знаете, то, скорее всего, Klipper<br>Для принтеров Bambu Lab выберите Marlin';
			values['table.delta.title'] = 'Начало координат<br>в центре стола';
			values['table.delta.description'] = 'Для декартовых принтеров должно быть выключено<br>Для дельт и AD5M включено';
			values['table.nozzle_diameter.title'] = 'Диаметр сопла';
			values['table.nozzle_diameter.description'] = '[мм] Диаметр сопла, установленного на вашем принтере. Влияет на ширину линии и толщину слоя при печати образцов и вспомогательных элементов';
			values['table.temperatures.title'] = 'Температура';
			values['table.temperatures.description'] = '[°C] Температуры хотэнда и стола при печати калибровочной модели<br>Если в принтере есть активная термокамера, то её температура будет установлена равной 0°C';
			values['table.flowrate.title'] = 'Объёмный расход';
			values['table.flowrate.description'] = '[мм^3/с] Начальное и конечное значение объёмного расхода для проверки<br>- Для V6-подобных хотэндов рекомендуется от 4 до 20<br>- Для Volcano-подобных рекомендуется от 4 до 30<br>- Для Goliath-подобных от 4 до 55<br>При наличии CHT вставки в хотэнде, стоит увеличить максимальное проверяемое значение';
			values['table.z_offset.title'] = 'Z-offset';
			values['table.z_offset.description'] = '[мм] Смещение всей модели по вертикали для компенсации неправильного расстояния от сопла до стола на первом слое. Положительное значение смещает модель вверх, отрицательное - вниз.<br>Если в вашем стартовом G-коде уже есть установка Z-offset\'а, то оставьте 0';
			values['table.speed.title'] = 'Скорость движения<br>печатающей головы';
			values['table.speed.description'] = '[мм/с] Скорость движения печатающей головы при печати (print) и при холостых перемещениях (travel).<br>Для скорости печати рекомендуется ставить невысокие значения, чтобы разгон и торможение оказывали минимальное влияние на средний объёмный расход при печати образца';
			values['table.flowrate_delta.title'] = 'Шаг расхода';
			values['table.flowrate_delta.description'] = '[мм^3/с] Разница объёмного расхода между соседними проверяемыми элементами. Чем меньше, тем больше образцов будет напечатано';
			values['table.extrusion_length.title'] = 'Длина экструзии<br>на образец';
			values['table.extrusion_length.description'] = '[мм] Сколько мм прутка будет экструдировано для печати каждого образца. Считается по входу в подающий механизм';
			values['table.start_gcode.title'] = 'Начальный G-код';
			values['table.gcodes.description'] = 'Для того, чтобы печать была правильно инициализирована и завершена, необходимо указать корректные для вашего принтера G-коды. Проще всего скопировать их из вашего слайсера. Если G-коды в слайсере не содержат неподдерживаемых элементов, то их можно копировать без изменений. Если неподдерживаемые элементы есть, то придётся от них избавиться - вручную заменить неподдерживаемые плейсхолдеры на значения, посчитать результаты математических функций, раскрыть условия и т.д.<hr><u>Поддерживаемые плейсхолдеры:</u><br>- <code>$HOTTEMP</code>, <code>{temperature}</code>, <code>{first_layer_temperature}</code>, <code>{nozzle_temperature}</code> и подобные будут заменены на температуру хотэнда;<br>- <code>$BEDTEMP</code>, <code>{bed_temperature}</code>, <code>{first_layer_bed_temperature}</code>, <code>{hot_plate_temp}</code>, <code>{textured_plate_temp_initial_layer}</code> и подобные будут заменены на температуру стола;<br>- <code>{chamber_temperature}</code>, <code>{chamber_minimal_temperature}</code> будут заменены на 0;<br>- Разные вариации указанных выше плейсхолдеров. Например, в квадратных скобках вместо фигурных, с указанием номера элемента массива, с арифметикой внутри фигурных скобок и т.д.<hr><u>Не поддерживаются:</u><br>- Другие плейсхолдеры;<br>- Условные операторы (<code>if</code>);<br>- Тернарные операторы (<code>a ? b : c</code>);<br>- Математические функции и выражения';
			values['table.end_gcode.title'] = 'Конечный G-код';

			values['generator.generate_and_download'] = 'Генерировать и скачать';
			values['generator.generate_button_loading'] = 'Генератор загружается...';
			values['generator.reset_to_default'] = 'Сбросить настройки';

			values['error.bed_size_x.format'] = 'Ошибка формата: Размер области печати по оси Х';
			values['error.bed_size_x.value'] = 'Ошибка значения: Размер области печати по оси X (меньше 100 или больше 1000 мм)';
			values['error.bed_size_y.format'] = 'Ошибка формата: Размер области печати по оси Y';
			values['error.bed_size_y.value'] = 'Ошибка значения: Размер области печати по оси Y (меньше 100 или больше 1000 мм)';
			values['error.nozzle_diameter.format'] = 'Ошибка формата: Диаметр сопла';
			values['error.nozzle_diameter.value'] = 'Ошибка значения: Диаметр сопла (меньше 0.1 или больше 2.0 мм)';
			values['error.hotend_temp.format'] = 'Ошибка формата: Температура хотэнда';
			values['error.hotend_temp.value'] = 'Ошибка значения: Температура хотэнда (меньше 100°C или больше 500°C)';
			values['error.bed_temp.format'] = 'Ошибка формата: Температура стола';
			values['error.bed_temp.value'] = 'Ошибка значения: Температура стола (выше 150°C)';
			values['error.init_flowrate.format'] = 'Ошибка формата: начальное значение объёмного расхода';
			values['error.init_flowrate.value'] = 'Ошибка значения: начальное значение объёмного расхода (меньше 1.0 или больше 50.0 мм^3/с)';
			values['error.end_flowrate.format'] = 'Ошибка формата: конечное значение объёмного расхода';
			values['error.end_flowrate.value'] = 'Ошибка значения: конечное значение объёмного расхода (меньше 1.0 или больше 50.0 мм^3/с)';
			values['error.flowrate_delta.format'] = 'Ошибка формата: Шаг расхода';
			values['error.flowrate_delta.value'] = 'Ошибка значения: Шаг расхода (меньше 0.5 или больше 10.0)';
			values['error.extrusion_length.format'] = 'Ошибка формата: длина экструзии на образец';
			values['error.extrusion_length.value'] = 'Ошибка значения: длина экструзии на образец (меньше 50 или больше 5000 мм)';
			values['error.z_offset.format'] = 'Ошибка формата: Z-offset';
			values['error.z_offset.value'] = 'Ошибка значения: Z-offset (меньше -0.5 или больше 0.5 мм)';
			values['error.speed.format'] = 'Ошибка формата: скорость движения печатающей головы при печати образцов';
			values['error.speed.value'] = 'Ошибка значения: скорость движения печатающей головы при печати образцов (меньше 5 или больше 500 мм/с)';
			values['error.travel_speed.format'] = 'Ошибка формата: скорость движения печатающей головы при холостых перемещениях';
			values['error.travel_speed.value'] = 'Ошибка значения: скорость движения печатающей головы при холостых перемещениях (меньше 5 или больше 1000 мм/с)';
			values['error.firmware.not_set'] = 'Ошибка: прошивка не выбрана';
			values['error.unsupported_placeholders'] = 'Ошибка: В начальном или конечном G-коде замечены неподдерживаемые плейсхолдеры:';
			values['error.unsupported_conditions'] = 'Ошибка: В начальном или конечном G-коде замечены условные операторы (if)';
			values['error.unsupported_ternary_operators'] = 'Ошибка: В начальном или конечном G-коде замечены тернарные операторы (a ? b : c)';
			values['error.unsupported_math'] = 'Ошибка: В начальном или конечном G-коде замечены математические операции (min, max, round и подобные)';
			values['error.different_versions'] = 'Ошибка: версия движка и страницы отличаются. Обновите страницу с помощью ctrl + F5 или почистите кэш браузера вручную';

			values['info.out_of_print_area'] = 'Некоторые образцы не поместились на стол. Количество пропущенных образцов: %s. Увеличьте шаг проверяемых значений или снизьте скорость\n';
			values['info.cali_params_template'] = ';Образец: %d Расход: %.3f мм^3/с\n';
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

function reset() {
	for (var elementId of formFields) {
		localStorage.removeItem(elementId);
	}

	window.location.reload(false);
}

function init() {
	initForm();
	initProfileManager();

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
	initProfileLang(lang);

	setTimeout(function () {
		if (checkGo != undefined && window.lang != undefined) {
			checkGo();
		} else {
			setTimeout(this, 100);
		}
	}, 100);
}

function updateInputDisability() {

}

// ===================== Функции управления профилями =====================

let profileManager = null;

function initProfileManager() {
	// Инициализация менеджера профилей
	profileManager = new ProfileManager('k3d_la_profiles', 10);

	profileManager.setCallbacks({
		onProfileChange: function(profile) {
			if (profile) {
				loadProfileSettings(profile.settings);
				const message = getProfileMessage('profile.status.loaded', profile.name);
			}
		},
		onProfileSave: function(profile) {
			updateProfileSelect();
			const message = getProfileMessage('profile.status.saved', profile.name);
		},
		onProfileDelete: function(profile) {
			updateProfileSelect();
			const message = getProfileMessage('profile.status.deleted', profile.name);
		}
	});

	updateProfileSelect();

	// Загружаем последний активный профиль, если есть
	const currentProfile = profileManager.getCurrentProfile();
	if (currentProfile) {
		document.getElementById('profileSelect').value = currentProfile.id;
		loadProfileSettings(currentProfile.settings);
	}

	// Обработчик изменения выбора профиля
	document.getElementById('profileSelect').addEventListener('change', function(e) {
		const profileId = e.target.value;
		if (profileId) {
			profileManager.setCurrentProfile(profileId);
		} else {
			// Сброс на дефолтные настройки
			profileManager.setCurrentProfile(null);
			loadForm(); // Загружаем из localStorage
			const message = window.lang.values['profile.status.default'] || 'Using default settings';
		}
	});
}

function initProfileLang(key) {
	var values = window.lang.values;
	if (key === 'en') {
		// UI элементы
		values['profile.select.label'] = 'Profile: ';
		values['profile.default'] = 'Default';
		values['profile.save'] = 'Save';
		values['profile.save_as'] = 'Save As...';
		values['profile.delete'] = 'Delete';
		values['profile.export'] = 'Export';
		values['profile.import'] = 'Import';
		values['profile.modal.title'] = 'Save Profile';
		values['profile.modal.save'] = 'Save';
		values['profile.modal.cancel'] = 'Cancel';
		values['profile.modal.placeholder'] = 'Enter profile name';

		// Статусные сообщения
		values['profile.status.loaded'] = 'Profile loaded: %s';
		values['profile.status.saved'] = 'Profile "%s" saved';
		values['profile.status.deleted'] = 'Profile "%s" deleted';
		values['profile.status.default'] = 'Using default settings';
		values['profile.status.enterName'] = 'Please enter a profile name';
		values['profile.status.importError'] = 'Import error: %s';
		values['profile.status.exportedSettings'] = 'Exported settings';
		values['profile.status.exportedCurrent'] = 'Current settings exported';
		values['profile.status.exportedAll'] = 'Exported %s profiles';
		values['profile.status.importedMultiple'] = 'Imported %s profiles';
		values['profile.status.importedSingle'] = 'Profile "%s" imported';

		// Диалоги
		values['profile.confirm.delete'] = 'Delete profile "%s"?';
		values['profile.confirm.overwrite'] = 'Profile "%s" already exists. Overwrite?';
	} else {
		// UI элементы
		values['profile.select.label'] = 'Профиль: ';
		values['profile.default'] = 'По умолчанию';
		values['profile.modal.title'] = 'Сохранить профиль';
		values['profile.modal.save'] = 'Сохранить';
		values['profile.modal.cancel'] = 'Отмена';
		values['profile.modal.placeholder'] = 'Введите название профиля';

		// Диалоги
		values['profile.confirm.delete'] = 'Удалить профиль "%s"?';
		values['profile.confirm.overwrite'] = 'Профиль "%s" уже существует. Перезаписать?';
	}

	// Обновляем тексты в UI
	updateProfileTexts();
}

function updateProfileTexts() {
	// Обновляем все элементы с классом lang внутри профильной панели
	const profileElements = document.querySelectorAll('#profilePanel .lang, #profileModal .lang');
	profileElements.forEach(element => {
		const key = element.id;
		if (key && window.lang.values[key]) {
			element.textContent = window.lang.values[key];
		}
	});

	// Обновляем placeholder
	const nameInput = document.getElementById('profileNameInput');
	if (nameInput && window.lang.values['profile.modal.placeholder']) {
		nameInput.placeholder = window.lang.values['profile.modal.placeholder'];
	}

	// Обновляем опцию "По умолчанию" в select
	const select = document.getElementById('profileSelect');
	if (select && select.options[0] && window.lang.values['profile.default']) {
		select.options[0].textContent = window.lang.values['profile.default'];
	}
}

function getProfileMessage(key, ...args) {
	const template = window.lang.values[key] || key;
	let message = template;

	// Простая замена %s на аргументы
	args.forEach(arg => {
		message = message.replace('%s', arg);
	});

	return message;
}

function getCurrentSettings() {
	const settings = {};
	for (var elementId of formFields) {
		var element = document.getElementById(elementId);
		if (element) {
			if (checkboxes.includes(elementId)) {
				settings[elementId] = element.checked;
			} else {
				settings[elementId] = element.value;
			}
		}
	}
	return settings;
}

function loadProfileSettings(settings) {
	for (var elementId of formFields) {
		if (settings[elementId] !== undefined) {
			var element = document.getElementById(elementId);
			if (element) {
				if (checkboxes.includes(elementId)) {
					element.checked = settings[elementId];
				} else {
					element.value = settings[elementId];
				}
			}
		}
	}
	updateInputDisability();
}

function updateProfileSelect() {
	const select = document.getElementById('profileSelect');
	const currentValue = select.value;

	// Очищаем все опции кроме первой (По умолчанию)
	while (select.options.length > 1) {
		select.remove(1);
	}

	// Добавляем профили
	const profiles = profileManager.getAllProfiles();
	profiles.forEach(profile => {
		const option = document.createElement('option');
		option.value = profile.id;
		option.textContent = profile.name;
		select.appendChild(option);
	});

	// Восстанавливаем выбранное значение
	if (currentValue) {
		select.value = currentValue;
	}
}

function saveCurrentProfile() {
	const select = document.getElementById('profileSelect');
	const profileId = select.value;

	if (profileId) {
		// Обновляем существующий профиль
		const settings = getCurrentSettings();
		profileManager.updateProfile(profileId, settings);
	} else {
		// Создаем новый профиль
		saveAsNewProfile();
	}
}

function saveAsNewProfile() {
	document.getElementById('profileModal').style.display = 'flex';
	document.getElementById('profileNameInput').value = '';
	document.getElementById('profileNameInput').focus();
}

function confirmSaveProfile() {
	const name = document.getElementById('profileNameInput').value.trim();

	if (!name) {
		const message = window.lang.values['profile.status.enterName'] || 'Please enter a profile name';
		return;
	}

	try {
		const settings = getCurrentSettings();
		const profile = profileManager.createProfile(name, settings);
		profileManager.setCurrentProfile(profile.id);
		document.getElementById('profileSelect').value = profile.id;
		closeProfileModal();
	} catch (e) {
		console.log(e.message)
	}
}

function closeProfileModal() {
	document.getElementById('profileModal').style.display = 'none';
}

function deleteCurrentProfile() {
	const select = document.getElementById('profileSelect');
	const profileId = select.value;

	if (!profileId) return;

	const profile = profileManager.getProfile(profileId);
	const confirmMessage = getProfileMessage('profile.confirm.delete', profile.name);
	if (confirm(confirmMessage)) {
		profileManager.deleteProfile(profileId);
		select.value = '';
	}
}

function exportProfile() {
	const profiles = profileManager.getAllProfiles();

	if (profiles.length === 0) {
		// Если нет сохраненных профилей, экспортируем текущие настройки
		const settings = getCurrentSettings();
		const exportName = window.lang.values['profile.status.exportedSettings'] || 'Exported settings';
		const tempProfile = {
			id: profileManager.generateId(),
			name: exportName,
			createdAt: Date.now(),
			updatedAt: Date.now(),
			settings: settings
		};

		const exportData = {
			version: '1.0',
			exported: Date.now(),
			profiles: [tempProfile]
		};

		const json = JSON.stringify(exportData, null, 2);
		downloadJSON('la_calibrator_profiles.json', json);

		const message = window.lang.values['profile.status.exportedCurrent'] || 'Current settings exported';
	} else {
		// Экспортируем все сохраненные профили
		const json = profileManager.exportAllProfiles();
		const timestamp = new Date().toISOString().split('T')[0];
		downloadJSON(`la_calibrator_profiles_${timestamp}.json`, json);

		const message = getProfileMessage(
			window.lang.values['profile.status.exportedAll'] || 'Exported %s profiles',
			profiles.length
		);
	}
}

function importProfile() {
	document.getElementById('profileFileInput').click();
}

function handleProfileImport(event) {
	const file = event.target.files[0];
	if (!file) return;

	const reader = new FileReader();
	reader.onload = function(e) {
		try {
			const data = JSON.parse(e.target.result);

			// Проверяем формат файла
			if (data.profiles && Array.isArray(data.profiles)) {
				// Новый формат с множественными профилями
				const imported = profileManager.importAllProfiles(e.target.result);

				if (imported.length > 0) {
					// Активируем первый импортированный профиль
					profileManager.setCurrentProfile(imported[0].id);
					document.getElementById('profileSelect').value = imported[0].id;
					updateProfileSelect();

					const message = getProfileMessage(
						window.lang.values['profile.status.importedMultiple'] || 'Imported %s profiles',
						imported.length
					);
				} else {
					throw new Error('No profiles were imported');
				}
			} else {
				throw new Error('Invalid profile file format');
			}
		} catch (error) {
			const message = getProfileMessage('profile.status.importError', error.message);
		}
	};
	reader.readAsText(file);

	// Очищаем input для возможности повторного импорта того же файла
	event.target.value = '';
}

function downloadJSON(filename, content) {
	const blob = new Blob([content], { type: 'application/json' });
	const url = URL.createObjectURL(blob);
	const a = document.createElement('a');
	a.href = url;
	a.download = filename;
	document.body.appendChild(a);
	a.click();
	document.body.removeChild(a);
	URL.revokeObjectURL(url);
}