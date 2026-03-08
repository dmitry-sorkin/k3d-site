package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"syscall/js"
)

const calibratorVersion = "v0.2-beta" // Версия калибратора для сверки с lib.js
const retractSpeed = 35.0             // Скорость отката по умолчанию
const retractLength = 1.0             // Длина отката по умолчанию
const filamentDiameter = 1.75         // Диаметр филамента
const sampleSpacing = 5.0             // Расстояние между образцами
const log = true                      // Писать ли в консоль логи происходящего

// Переменные, которые должны быть доступны для записи и чтения из любой части программы
var (
	// Параметры принтера
	bedX, bedY     float64
	zOffset        float64
	nozzleDiameter float64
	firmware       int
	delta          bool

	// Параметры филамента
	hotendTemperature int
	bedTemperature    int

	// Параметры теста
	speed           int
	travelSpeed     int
	initFlowrate    float64
	endFlowrate     float64
	flowrateDelta   float64
	extrusionLength float64
	layerHeight     float64
	lineWidth       float64

	// Текущее состояние
	currentCoordinates Point
	retracted          bool
	currentE           float64
	currentSpeed       int

	// G-коды
	startGcode string
	endGcode   string
)

type Point struct {
	X float64
	Y float64
	Z float64
}

type sampleParameters struct {
	fullLineLength float64
	numLines       int
	speed          float64
	startX         float64
	endX           float64
}

func main() {
	c := make(chan struct{})
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("checkGo", js.FuncOf(checkJs))
	<-c
}

func setErrorDescription(doc js.Value, key string, curErr string, hasErr bool, allowModify bool) {
	// Эта функция используется для того, чтобы отображать ошибку в описании параметра
	if !allowModify {
		return
	}
	el := doc.Call("getElementById", key)
	el.Get("style").Set("display", "")
	el.Set("rowSpan", "1")
	if hasErr {
		el.Set("innerHTML", getLangString(key)+"<br><br><span class=\"inline-error\">"+curErr+"</span>")
	} else {
		el.Set("innerHTML", getLangString(key))
	}
}

func getLangString(key string) string {
	lang := js.Global().Get("lang")
	return lang.Call("getString", key).String()
}

func getElementValue(id string) string {
	return js.Global().Get("document").Call("getElementById", id).Get("value").String()
}

func check(showErrorBox bool, allowModify bool) bool {
	// Функция проверки введённых данных

	errorString := ""
	doc := js.Global().Get("document")
	doc.Call("getElementById", "resultContainer").Set("innerHTML", "")
	curErr := ""    // Текущее сообщение об ошибке
	hasErr := false // Флаг текущей ошибки. Обнуляется при выводе ошибки
	retErr := false // Флаг о наличии хотя бы 1 ошибки при проверке

	// Проверка версии калибратора
	libCaliVersion := js.Global().Get("calibrator_version").String()
	if libCaliVersion != calibratorVersion {
		errorMessage := getLangString("error.different_versions")
		println(errorMessage)
		js.Global().Call("showError", errorMessage)
		return false
	}

	// Параметры принтера
	// Размер стола

	docBedX, err := parseInputToFloat(getElementValue("k3d_mfr_bedX"))
	if err != nil {
		curErr, hasErr = getLangString("error.bed_size_x.format"), true
	} else if docBedX < 100 || docBedX > 1000 {
		curErr, hasErr = getLangString("error.bed_size_x.value"), true
	} else {
		bedX = docBedX
	}

	setErrorDescription(doc, "table.bed_size.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docBedY, err := parseInputToFloat(getElementValue("k3d_mfr_bedY"))
	if err != nil {
		curErr, hasErr = getLangString("error.bed_size_y.format"), true
	} else if docBedY < 100 || docBedY > 1000 {
		curErr, hasErr = getLangString("error.bed_size_y.value"), true
	} else {
		bedY = docBedY
	}

	setErrorDescription(doc, "table.bed_size.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Прошивка

	docMarlin := doc.Call("getElementById", "k3d_mfr_marlin").Get("checked").Bool()
	docKlipper := doc.Call("getElementById", "k3d_mfr_klipper").Get("checked").Bool()
	docRRF := doc.Call("getElementById", "k3d_mfr_rrf").Get("checked").Bool()
	if docMarlin {
		firmware = 0
	} else if docKlipper {
		firmware = 1
	} else if docRRF {
		firmware = 2
	} else {
		errorString = errorString + getLangString("error.firmware.not_set") + "\n"
	}

	// Начало координат в центре стола (режим дельты)

	delta = doc.Call("getElementById", "k3d_mfr_delta").Get("checked").Bool()

	// Диаметр сопла

	docNozzleDiameter, err := parseInputToFloat(getElementValue("k3d_mfr_nozzleDiameter"))
	if err != nil {
		curErr, hasErr = getLangString("error.nozzle_diameter.format"), true
	} else if docNozzleDiameter < 0.1 || docNozzleDiameter > 2.0 {
		curErr, hasErr = getLangString("error.nozzle_diameter.value"), true
	} else {
		nozzleDiameter = docNozzleDiameter
	}

	setErrorDescription(doc, "table.nozzle_diameter.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Z offset

	docZOffset, err := parseInputToFloat(getElementValue("k3d_mfr_zOffset"))
	if err != nil {
		curErr, hasErr = getLangString("error.z_offset.format"), true
	} else if docZOffset < -0.5 || docZOffset > 0.5 {
		curErr, hasErr = getLangString("error.z_offset.value"), true
	} else {
		zOffset = docZOffset
	}
	setErrorDescription(doc, "table.z_offset.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Параметры филамента
	// Температура хотэнда

	docHotTemp, err := parseInputToInt(getElementValue("k3d_mfr_hotendTemperature"))
	if err != nil {
		curErr, hasErr = getLangString("error.hotend_temp.format"), true
	} else if docHotTemp < 100 {
		curErr, hasErr = getLangString("error.hotend_temp.too_low"), true
	} else if docHotTemp > 500 {
		curErr, hasErr = getLangString("error.hotend_temp.too_high"), true
	} else {
		hotendTemperature = docHotTemp
	}

	setErrorDescription(doc, "table.temperatures.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Температура стола

	docBedTemp, err := parseInputToInt(getElementValue("k3d_mfr_bedTemperature"))
	if err != nil {
		curErr, hasErr = getLangString("error.bed_temp.format")+err.Error(), true
	} else if docBedTemp > 150 {
		curErr, hasErr = getLangString("error.bed_temp.too_high"), true
	} else {
		bedTemperature = docBedTemp
	}

	setErrorDescription(doc, "table.temperatures.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Параметры калибровки
	// Скорость движения печатающей головы при печати образцов
	docSpeed, err := parseInputToInt(getElementValue("k3d_mfr_speed"))
	if err != nil {
		curErr, hasErr = getLangString("error.speed.format"), true
	} else if docSpeed < 5 || docSpeed > 500 {
		curErr, hasErr = getLangString("error.speed.value"), true
	} else {
		speed = docSpeed
	}
	setErrorDescription(doc, "table.speed.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Скорость движения печатающей головы при холостых перемещениях
	docTravelSpeed, err := parseInputToInt(getElementValue("k3d_mfr_travelSpeed"))
	if err != nil {
		curErr, hasErr = getLangString("error.travel_speed.format"), true
	} else if docTravelSpeed < 5 || docTravelSpeed > 1000 {
		curErr, hasErr = getLangString("error.travel_speed.value"), true
	} else {
		travelSpeed = docTravelSpeed
	}
	setErrorDescription(doc, "table.speed.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Объёмный расход - начальное значение
	docInitFlowrate, err := parseInputToFloat(getElementValue("k3d_mfr_initFlowrate"))
	if err != nil {
		curErr, hasErr = getLangString("error.init_flowrate.format"), true
	} else if docInitFlowrate < 1.0 || docInitFlowrate > 50.0 {
		curErr, hasErr = getLangString("error.init_flowrate.value"), true
	} else {
		initFlowrate = docInitFlowrate
	}
	setErrorDescription(doc, "table.flowrate.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Объёмный расход - конечное значение
	docEndFlowrate, err := parseInputToFloat(getElementValue("k3d_mfr_endFlowrate"))
	if err != nil {
		curErr, hasErr = getLangString("error.end_flowrate.format"), true
	} else if docEndFlowrate < 1.0 || docEndFlowrate > 50.0 {
		curErr, hasErr = getLangString("error.end_flowrate.value"), true
	} else {
		endFlowrate = docEndFlowrate
	}
	setErrorDescription(doc, "table.flowrate.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Дельта изменения объёмного расхода
	docFlowrateDelta, err := parseInputToFloat(getElementValue("k3d_mfr_flowrateDelta"))
	if err != nil {
		curErr, hasErr = getLangString("error.flowrate_delta.format"), true
	} else if docFlowrateDelta < 0.01 || docFlowrateDelta > 10.0 {
		curErr, hasErr = getLangString("error.flowrate_delta.value"), true
	} else {
		flowrateDelta = docFlowrateDelta
	}
	setErrorDescription(doc, "table.flowrate_delta.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Длина экструзии
	docExtrusionLength, err := parseInputToFloat(getElementValue("k3d_mfr_extrusionLength"))
	if err != nil {
		curErr, hasErr = getLangString("error.extrusion_length.format"), true
	} else if docExtrusionLength < 50 || docExtrusionLength > 5000 {
		curErr, hasErr = getLangString("error.extrusion_length.value"), true
	} else {
		extrusionLength = docExtrusionLength
	}
	setErrorDescription(doc, "table.extrusion_length.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Начальный и конечный G-коды

	docStartGcode := getElementValue("k3d_mfr_startGcode")
	var errorMessage string
	docStartGcode, errorMessage = processGCode(docStartGcode)
	if errorMessage != "" {
		errorString += errorMessage
		hasErr = false
		retErr = true
	} else {
		startGcode = docStartGcode
	}

	docEndGcode := getElementValue("k3d_mfr_endGcode")
	errorMessage = ""
	docEndGcode, errorMessage = processGCode(docEndGcode)
	if errorMessage != "" {
		errorString += errorMessage
		hasErr = false
		retErr = true
	} else {
		// К конечному гкоду добавляем в конец перенос строки т.к.
		// klipper не выполняет команды на последней строке G-кода
		endGcode = docEndGcode + "\n"
	}

	// Окончание проверки параметров

	if !showErrorBox {
		return !retErr
	}

	if !retErr {
		println("OK")
		return true
	} else {
		println(errorString)
		js.Global().Call("showError", errorString)
		return false
	}
}

func write(str ...string) {
	for i := 0; i < len(str); i++ {
		js.Global().Call("writeToFile", str[i])
	}
}

func checkJs(this js.Value, i []js.Value) interface{} {
	check(false, true)
	return js.ValueOf(nil)
}

func processGCode(inputGCode string) (string, string) {
	// Функция, которая принимает на вход стартовый или конечный G-код,
	// заменяет в нём плейсхолдеры на значения переменных, а также
	// проверяет, нет ли неподдерживаемых плейсхолдеров. Если неподдерживаемые
	// плейсхолдеры есть, то возвращает еще и текст ошибки

	outputGCode := inputGCode
	errorMessage := ""

	// Замена плейсхолдеров температуры стола
	bedTempPatterns := []string{
		`\$BEDTEMP`,
		`\s*[\[{]\s*(?:bed_temperature|first_layer_bed_temperature|bed_temperature_initial_layer|bed_temperature_initial_layer_single|hot_plate_temp_initial_layer|hot_plate_temp|supertack_plate_temp_initial_layer|supertack_plate_temp|cool_plate_temp|cool_plate_temp_initial_layer|textured_cool_plate_temp|textured_cool_plate_temp_initial_layer|eng_plate_temp|eng_plate_temp_initial_layer|textured_plate_temp|textured_plate_temp_initial_layer)\s*(?:\[\s*\d+\s*\])?(?:\s*[+\-*/]\s*(?:\d+(?:\.\d+)?)|\d+(?:\.\d+)?)?\s*[\]}]`,
	}
	re := regexp.MustCompile(`(?i)` + strings.Join(bedTempPatterns, "|"))
	outputGCode = re.ReplaceAllString(outputGCode, strconv.Itoa(bedTemperature))

	// Замена плейсхолдеров температуры термокамеры
	re = regexp.MustCompile(`\s*[\[{]\s*(?:chamber_temperature|chamber_minimal_temperature)\s*(?:\[\s*\d+\s*\])?(?:\s*[+\-*/]\s*(?:\d+(?:\.\d+)?)|\d+(?:\.\d+)?)?\s*[\]}]`)
	outputGCode = re.ReplaceAllString(outputGCode, "0")

	// Замена плейсхолдеров температуры хотэнда
	hotendTempPatterns := []string{
		`\$HOTTEMP`,
		`\s*[\[{]\s*(?:temperature|first_layer_temperature|nozzle_temperature_initial_layer_single|nozzle_temperature|nozzle_temperature_initial_layer)\s*(?:\[\s*\d+\s*\])?(?:\s*[+\-*/]\s*(?:\d+(?:\.\d+)?)|\d+(?:\.\d+)?)?\s*[\]}]`,
	}
	re = regexp.MustCompile(`(?i)` + strings.Join(hotendTempPatterns, "|"))
	outputGCode = re.ReplaceAllString(outputGCode, strconv.Itoa(hotendTemperature))

	// Проверка на неподдерживаемые плейсхолдеры
	unsupportedPattern := regexp.MustCompile(`\{[^}]+\}|\[[^\]]+\]`)
	unsupportedPlaceholders := unsupportedPattern.FindAllString(outputGCode, -1)

	if len(unsupportedPlaceholders) > 0 {
		errorMessage += getLangString("error.unsupported_placeholders") + "\n"
		for _, m := range unsupportedPlaceholders {
			errorMessage += m + "\n"
		}
	}

	// Проверка на наличие условий
	conditionsPattern := regexp.MustCompile(`\{if|\{ if`)
	unsupportedConditions := conditionsPattern.FindAllString(outputGCode, -1)

	if len(unsupportedConditions) > 0 {
		errorMessage += getLangString("error.unsupported_conditions") + "\n"
	}

	// Проверка на наличие тернарных операторов
	ternaryPattern := regexp.MustCompile(`\s*([^\s?]+?)\s*\?\s*([^\s:]+?)\s*:\s*([^\s:]+?)`)
	ternaryMatches := ternaryPattern.FindAllString(outputGCode, -1)

	if len(ternaryMatches) > 0 {
		errorMessage += getLangString("error.unsupported_ternary_operators") + "\n"
	}

	// Проверка на наличие математических функций: min(), max(), int(), round(), digits(), zdigits(), is_nil()
	mathFuncPattern := regexp.MustCompile(`\b(min|max|int|round|digits|zdigits|is_nil)\s*\([^)]*\)`)
	mathFuncMatches := mathFuncPattern.FindAllString(outputGCode, -1)

	if len(mathFuncMatches) > 0 {
		errorMessage += getLangString("error.unsupported_math") + "\n"
	}

	return outputGCode, errorMessage
}

func (p Point) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p.X, p.Y, p.Z)
}

func generate(this js.Value, i []js.Value) interface{} {

	// Проверить данные, полученные со страницы калибратора
	if !check(true, false) {
		return js.ValueOf(nil)
	}

	// Контейнеры для информационных сообщений
	infoMessage := ""
	caliParams := ""

	// Расчёт неопределённых ранее базовых параметров
	currentCoordinates.X, currentCoordinates.Y, currentCoordinates.Z = 0, 0, 0
	currentE = 0
	retracted = false
	currentSpeed = 60

	// Вывод базовых параметров в консоль
	if log {
		// Параметры принтера
		println("bedX:", bedX)
		println("bedY:", bedY)
		println("zOffset:", zOffset)
		println("nozzleDiameter:", nozzleDiameter)
		println("firmware:", firmware)
		println("delta:", delta)

		// Параметры филамента
		println("hotendTemperature:", hotendTemperature)
		println("bedTemperature:", bedTemperature)

		// Параметры теста
		println("initFlowrate:", initFlowrate)
		println("endFlowrate:", endFlowrate)
		println("flowrateDelta:", flowrateDelta)
		println("extrusionLength:", extrusionLength)

		// Текущее состояние
		println("currentCoordinates:", currentCoordinates.String())
		println("retracted:", retracted)
		println("currentE:", currentE)
		println("currentSpeed:", currentSpeed)

		// G-коды
		println("startGcode:", startGcode)
		println("endGcode:", endGcode)
	}

	// Расчёт параметров генерации
	var bedCenter Point // Центр стола
	if delta {
		bedCenter.X, bedCenter.Y, bedCenter.Z = 0, 0, 0
	} else {
		bedCenter.X, bedCenter.Y, bedCenter.Z = bedX/2, bedY/2, 0
	}
	maxSingleLineLength := bedY - 15
	minFlowrate := math.Min(initFlowrate, endFlowrate)
	maxFlowrate := math.Max(initFlowrate, endFlowrate)
	numSamples := int(math.Ceil((maxFlowrate - minFlowrate) / flowrateDelta))

	// Генерируем название файла и начинаем его запись
	fileName := "K3D_MFC_H" + fmt.Sprint(hotendTemperature) + "_B" + fmt.Sprint(bedTemperature) + "_FR" + fmt.Sprint(roundFloat(initFlowrate, 1)) + "-" + fmt.Sprint(roundFloat(endFlowrate, 1)) + "_d" + fmt.Sprint(roundFloat(flowrateDelta, 1)) + ".gcode"
	js.Global().Call("beginSaveFile", fileName)

	//Комментарий со списком параметров в начале файла
	write(";generated by K3D maximum flowrate calibrator ", js.Global().Get("calibrator_version").String(), "\n",
		";written by Dmitry Sorkin @ k3d.tech, Kekht and YTKAB0BP\n",
		";Bed size: ", fmt.Sprint(roundFloat(bedX, 1)), ":", fmt.Sprint(roundFloat(bedY, 1)), " [mm]\n",
		";Firmware (0-Marlin, 1-Klipper, 2-RRF): ", fmt.Sprint(firmware), "\n",
		";Delta: ", fmt.Sprint(delta), "\n",
		";Nozzle diameter: ", fmt.Sprint(roundFloat(nozzleDiameter, 2)), "[mm]\n",
		";Z-offset: ", fmt.Sprint(roundFloat(zOffset, 3)), " [mm]\n",
		";Temperature: ", fmt.Sprint(hotendTemperature), "/", fmt.Sprint(bedTemperature), " [°C]\n",
		";Speed: ", fmt.Sprint(speed), "[mm/s]\n",
		";Travel speed: ", fmt.Sprint(travelSpeed), "[mm/s]\n",
		";Flowrate: ", fmt.Sprint(roundFloat(minFlowrate, 2)), "-", fmt.Sprint(roundFloat(maxFlowrate, 2)), " D", fmt.Sprint(roundFloat(flowrateDelta, 2)), " [mm^3/s]\n",
		";Extrusion length: ", fmt.Sprint(extrusionLength), " [mm]\n")

	// Стартовый G-код
	write(startGcode, "\n")

	// Переводим принтер в нужное нам состояние
	write("G90;Absolute coordinates\n")
	write("M82;Absolute extruder coordinates\n")
	write("G92 E0;Set E coordinate to zero\n")
	write("M106 S0;Disable cooling\n")
	write("M204 S5000;Set acceleration\n")
	write("M221 S100;Reset flow multiplier\n")

	// Расчёт данных для печати линии прочистки
	lineWidth, layerHeight = calcWidthAndHeight(2.0, 5.0, speed)
	purgeLineLength := calcLineLength(50.0, 1.75, lineWidth, layerHeight)
	purgeLineStart := Point{X: 5.0, Y: 5.0, Z: layerHeight + zOffset}

	// Метаданные для слайсеров
	write(";Purge line\n")
	write(";Z:", fmt.Sprint(roundFloat(layerHeight, 3)), "\n")
	write(";HEIGHT:" + fmt.Sprint(roundFloat(layerHeight, 3)) + "\n")
	write(";WIDTH:" + fmt.Sprint(roundFloat(lineWidth, 3)) + "\n")

	// Перемещаемся к началу линии прочистки
	write(generateMove(currentCoordinates, Point{currentCoordinates.X, currentCoordinates.Y, 5.0}, 0.0, travelSpeed)...)
	write(generateRetraction())
	write(generateMove(currentCoordinates, Point{purgeLineStart.X, purgeLineStart.X, currentCoordinates.Z}, 0.0, travelSpeed)...)
	write(generateMove(currentCoordinates, purgeLineStart, 0.0, travelSpeed)...)
	write(generateDeretraction())

	// Печатаем линию прочистки
	write(generateRelativeMove(purgeLineLength/2, 0, 0, lineWidth, speed)...)
	write(generateRelativeMove(0, lineWidth, 0, lineWidth, speed)...)
	write(generateRelativeMove(-purgeLineLength/2, 0, 0, lineWidth, speed)...)

	// Печать образцов
	for i := 0; i <= numSamples; i++ {
		// Расчёт параметров текущего образца
		sampleSpeed := speed
		var currentFlowrate float64
		if i == numSamples {
			// У последнего образца расход устанавливается равным максимальному расходу
			currentFlowrate = maxFlowrate
		} else {
			// У остальных рассчитывается
			currentFlowrate = minFlowrate + flowrateDelta*float64(i)
		}
		lineWidth, layerHeight = calcWidthAndHeight(2.0, currentFlowrate, speed)
		if lineWidth < nozzleDiameter {
			// Если ширина линии меньше диаметра сопла, то уменьшаем скорость
			lineWidth = nozzleDiameter
			layerHeight = nozzleDiameter / 2.0
			sampleSpeed = calcSpeed(lineWidth, layerHeight, currentFlowrate)
		}
		fullLineLength := calcLineLength(extrusionLength, filamentDiameter, lineWidth, layerHeight)
		numLines := int(math.Ceil(fullLineLength/(2*maxSingleLineLength))) * 2
		singleLineLength := fullLineLength / float64(numLines)
		var lineStartPoint Point
		if i == 0 {
			// У первого образца задаём конкретную точку начала
			lineStartPoint = Point{X: 5.0, Y: 10.0, Z: layerHeight}
		} else {
			// У последующих начало на откладывается относительно предыдущей позиции
			lineStartPoint = currentCoordinates
			lineStartPoint.X += sampleSpacing
		}
		// Проверка на то, не выходит ли образец за пределы области печати
		if lineStartPoint.X+lineWidth*float64(numLines) > bedX {
			// Образец выходит за пределы области печати. Закончить печать
			infoMessage += fmt.Sprintf(getLangString("info.out_of_print_area"), fmt.Sprint(numSamples-i+1))
			break
		}
		// Метаданные для предпросмотра в слайсере
		write(";LAYER_CHANGE\n")
		write(";Z:" + fmt.Sprint(roundFloat(layerHeight, 3)) + "\n")
		write(";HEIGHT:" + fmt.Sprint(roundFloat(layerHeight, 3)) + "\n")
		write(";TYPE:Inner wall\n")
		write(";WIDTH:" + fmt.Sprint(roundFloat(lineWidth, 3)) + "\n")
		// Перемещение к началу образца
		write(generateRetraction())
		write(generateMove(currentCoordinates, lineStartPoint, 0, travelSpeed)...)
		write(generateDeretraction())
		// Печать текущего образца
		for j := 0; j < numLines; j++ {
			if j%2 == 0 {
				// Для нечётных линий двигаемся по Y+
				write(generateRelativeMove(0, singleLineLength, 0, lineWidth, sampleSpeed)...)
				write(generateRelativeMove(lineWidth, 0, 0, 0, sampleSpeed)...)
			} else {
				// Для чётных двигаемся по Y-
				write(generateRelativeMove(0, -singleLineLength, 0, lineWidth, sampleSpeed)...)
				// Если это не последняя линия, то делаем переход к следующей
				if j != numLines-1 {
					write(generateRelativeMove(lineWidth, 0, 0, 0, sampleSpeed)...)
				}
			}
		}
		if log {
			println("Generated sample #", i+1)
			println("currentFlowrate: ", currentFlowrate)
			println("lineWidth: ", lineWidth)
			println("layerHeight: ", layerHeight)
			println("fullLineLength: ", fullLineLength)
			println("numLines: ", numLines)
			println("singleLineLength: ", singleLineLength)
			println("lineStartPoint: ", lineStartPoint.String())
		}
		// Добавляем в параметры калибровки информацию о текущем образце
		caliParams += fmt.Sprintf(getLangString("info.cali_params_template"), i+1, currentFlowrate)
	}

	// Конечный гкод
	write(";TYPE:Custom\n")
	write(endGcode)

	// Добавление информации о параметрах калибровки в G-код
	write(";----------\n")
	write(caliParams)

	// Вывод информационного сообщения, если оно есть
	if infoMessage != "" {
		js.Global().Call("showError", infoMessage)
	}

	// Вывод параметров калибровки в информационном окне
	js.Global().Call("showError", caliParams)

	// Окончание записи файла
	js.Global().Call("finishFile")

	return js.ValueOf(nil)
}

func generateLACommand(kFactor float64) string {
	if firmware == 0 {
		return fmt.Sprintf("M900 K%s\n", fmt.Sprint(roundFloat(kFactor, 3)))
	} else if firmware == 1 {
		return fmt.Sprintf("SET_PRESSURE_ADVANCE ADVANCE=%s\n", fmt.Sprint(roundFloat(kFactor, 3)))
	} else if firmware == 2 {
		return fmt.Sprintf("M572 D0 S%s\n", fmt.Sprint(roundFloat(kFactor, 3)))
	}

	return ";no firmware information\n"
}

func generateRelativeMove(x, y, z, width float64, speed int) []string {
	endPoint := currentCoordinates
	endPoint.X += x
	endPoint.Y += y
	endPoint.Z += z
	return generateMove(currentCoordinates, endPoint, width, speed)
}

func generateMove(start, end Point, width float64, speed int) []string {
	// create move
	move := make([]string, 0, 1)

	// create G1 command
	command := "G1"

	// add X
	if end.X != start.X {
		command += fmt.Sprintf(" X%s", fmt.Sprint(roundFloat(end.X, 2)))
	}

	// add Y
	if end.Y != start.Y {
		command += fmt.Sprintf(" Y%s", fmt.Sprint(roundFloat(end.Y, 2)))
	}

	// add Z
	if end.Z != start.Z {
		command += fmt.Sprintf(" Z%s", fmt.Sprint(roundFloat(end.Z, 2)))
	}

	// add E
	if width > 0 && math.Sqrt(float64(math.Pow((end.X-start.X), 2)+math.Pow((end.Y-start.Y), 2))) > 0.8 {
		newE := currentE + calcExtrusion(start, end, width)
		command += fmt.Sprintf(" E%s", fmt.Sprint(roundFloat(newE, 5)))
		currentE = newE
	}

	// add F
	if speed != 0 {
		command += fmt.Sprintf(" F" + fmt.Sprint(speed*60))
		currentSpeed = speed
	}

	// add G1 to move
	move = append(move, command+"\n")
	currentCoordinates = end

	return move
}

func calcExtrusion(start, end Point, width float64) float64 {
	lineLength := math.Sqrt(math.Pow((end.X-start.X), 2) + math.Pow((end.Y-start.Y), 2))
	// extrusion := width * layerHeight * lineLength * 4 / math.Pi / math.Pow(filamentDiameter, 2)
	extrusion := (4 * lineLength * ((width-layerHeight)*layerHeight + math.Pi*layerHeight*layerHeight/4)) / (math.Pi * filamentDiameter * filamentDiameter)
	return extrusion
}

func generateRetraction() string {
	if retracted {
		fmt.Println("Called retraction, but already retracted")
		return ""
	} else {
		retracted = true
		currentSpeed = retractSpeed
		return "G1 E" + fmt.Sprint(roundFloat(currentE-retractLength, 2)) + " F" + fmt.Sprint(retractSpeed*60) + "\n"
	}
}

func generateDeretraction() string {
	if retracted {
		retracted = false
		currentSpeed = retractSpeed
		return "G1 E" + fmt.Sprint(roundFloat(currentE, 2)) + " F" + fmt.Sprint(retractSpeed*60) + "\n"
	} else {
		fmt.Println("Called deretraction, but not retracted")
		return ""
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func parseInputToFloat(val string) (float64, error) {
	f, err := strconv.ParseFloat(strings.ReplaceAll(val, ",", "."), 64)
	if err != nil {
		println(err.Error())
	}
	return f, err
}

func parseInputToInt(val string) (int, error) {
	f, err := parseInputToFloat(val)
	return int(roundFloat(f, 0)), err
}

func calcWidthAndHeight(ratio, flowrate float64, speed int) (float64, float64) {
	// Принимает на вход отношение между шириной линии и толщиной слоя, расход и скорость движения
	// печатающей головы. Возвращает ширину линии и толщину слоя для этих параметров
	height := math.Sqrt(flowrate / (float64(speed) * (ratio - 1 + math.Pi/4)))
	width := height * ratio
	return width, height
}

func calcSpeed(width, height, flowrate float64) int {
	// Принимает на вход ширину линии, толщину слоя и расход. Возвращает скорость
	return int(roundFloat(flowrate/(height*(width-height+math.Pi*height/4)), 0))
}

func calcLineLength(extrLength, filDiam, width, height float64) float64 {
	// Принимает необходимую длину экструзии, диаметр филамента, ширину и высоту.
	// Возвращает длину линии
	numerator := math.Pi * filDiam * filDiam * float64(extrLength)
	denominator := 4*width*height + height*height*(math.Pi-4)
	if denominator != 0 {
		return numerator / denominator
	} else {
		println("calcLineLength() error: denominator = 0")
		return 0.0
	}
}
