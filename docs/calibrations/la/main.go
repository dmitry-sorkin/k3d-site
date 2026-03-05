package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"syscall/js"
)

const caliVersion = "v2.3"     // Версия калибратора. Пишется в сгенерированном файле
const retractSpeed = 35.0      // Скорость отката по умолчанию
const retractLength = 1.0      // Длина отката по умолчанию
const log = false              // Писать ли в консоль логи происходящего
const maxSeparatorOffset = 0.2 // Ограничение выпирания линии, разделяющей сегменты

// Переменные, которые должны быть доступны для записи и чтения из любой части программы
var (
	bedX, bedY, zOffset, nozzleDiameter, firstLayerLineWidth, lineWidth, currentE, layerHeightLimit, filamentDiameter, defaultSlowLineLength, defaultModelWidth, defaultSegmentHeight                float64
	layerHeight, initKFactor, endKFactor, smoothTime, initAdditionalParameter, endAdditionalParameter, segmentHeight                                                                                 float64
	firmware, hotendTemperature, bedTemperature, cooling, maxPrintSpeed, minPrintSpeed, numSegments, flow, additionalTarget, currentSpeed, acceleration, numPerimeters, maxSpeedDelta, minSpeedDelta int
	retracted, delta                                                                                                                                                                                 bool
	startGcode, endGcode                                                                                                                                                                             string
	currentCoordinates                                                                                                                                                                               Point
)

type Point struct {
	X float64
	Y float64
	Z float64
}

func main() {
	c := make(chan struct{})
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("checkGo", js.FuncOf(checkJs))
	js.Global().Set("checkSegments", js.FuncOf(checkSegments))
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

	// Параметры принтера
	// Размер стола

	docBedX, err := parseInputToFloat(getElementValue("k3d_la_bedX"))
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

	docBedY, err := parseInputToFloat(getElementValue("k3d_la_bedY"))
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

	docMarlin := doc.Call("getElementById", "k3d_la_marlin").Get("checked").Bool()
	docKlipper := doc.Call("getElementById", "k3d_la_klipper").Get("checked").Bool()
	docRRF := doc.Call("getElementById", "k3d_la_rrf").Get("checked").Bool()
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

	delta = doc.Call("getElementById", "k3d_la_delta").Get("checked").Bool()

	// Диаметр сопла

	docNozzleDiameter, err := parseInputToFloat(getElementValue("k3d_la_nozzleDiameter"))
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

	docZOffset, err := parseInputToFloat(getElementValue("k3d_la_zOffset"))
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

	docHotTemp, err := parseInputToInt(getElementValue("k3d_la_hotendTemperature"))
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

	docBedTemp, err := parseInputToInt(getElementValue("k3d_la_bedTemperature"))
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

	// Скорость вентилятора

	docCooling, err := parseInputToInt(getElementValue("k3d_la_cooling"))
	if err != nil {
		curErr, hasErr = getLangString("error.fan_speed.format"), true
	} else {
		docCooling = int(float64(docCooling) * 2.55)
		if docCooling < 0 {
			docCooling = 0
		} else if docCooling > 255 {
			docCooling = 255
		}
		cooling = docCooling
	}

	setErrorDescription(doc, "table.fan_speed.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Поток

	docFlow, err := parseInputToInt(getElementValue("k3d_la_flow"))
	if err != nil {
		curErr, hasErr = getLangString("error.flow.format"), true
	} else if docFlow < 50 || docFlow > 150 {
		curErr, hasErr = getLangString("error.flow.value"), true
	} else {
		flow = docFlow
	}
	setErrorDescription(doc, "table.flow.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Параметры калибровки
	// Скорости движения печатающей головы

	docMinPrintSpeed, err := parseInputToInt(getElementValue("k3d_la_minPrintSpeed"))
	if err != nil {
		curErr, hasErr = getLangString("error.min_print_speed.format"), true
	} else if docMinPrintSpeed < 5 || docMinPrintSpeed > 1000 {
		curErr, hasErr = getLangString("error.min_print_speed.value"), true
	} else {
		minPrintSpeed = docMinPrintSpeed
	}

	setErrorDescription(doc, "table.speed.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docMaxPrintSpeed, err := parseInputToInt(getElementValue("k3d_la_maxPrintSpeed"))
	if err != nil {
		curErr, hasErr = getLangString("error.max_print_speed.format"), true
	} else if docMinPrintSpeed < 5 || docMinPrintSpeed > 1000 {
		curErr, hasErr = getLangString("error.max_print_speed.value"), true
	} else {
		maxPrintSpeed = docMaxPrintSpeed
	}

	setErrorDescription(doc, "table.speed.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Ускорение

	docAcceleration, err := parseInputToInt(getElementValue("k3d_la_acceleration"))
	if err != nil {
		curErr, hasErr = getLangString("error.acceleration.format"), true
	} else if docAcceleration < 500 || docAcceleration > 50000 {
		curErr, hasErr = getLangString("error.acceleration.value"), true
	} else {
		acceleration = docAcceleration
	}

	setErrorDescription(doc, "table.acceleration.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Количество периметров

	docNumPerimeters, err := parseInputToInt(getElementValue("k3d_la_numPerimeters"))
	if err != nil {
		curErr, hasErr = getLangString("error.num_perimeters.format"), true
	} else if docAcceleration < 500 || docAcceleration > 50000 {
		curErr, hasErr = getLangString("error.num_perimters.value"), true
	} else {
		numPerimeters = docNumPerimeters
	}

	setErrorDescription(doc, "table.num_perimeters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Коэффициенты LA/PA

	docInitKFactor, err := parseInputToFloat(getElementValue("k3d_la_initKFactor"))
	if err != nil {
		curErr, hasErr = getLangString("error.init_la.format"), true
	} else if docInitKFactor < 0.0 || docInitKFactor > 2.0 {
		curErr, hasErr = getLangString("error.init_la.value"), true
	} else {
		initKFactor = docInitKFactor
	}

	setErrorDescription(doc, "table.la_values.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docEndKFactor, err := parseInputToFloat(getElementValue("k3d_la_endKFactor"))
	if err != nil {
		curErr, hasErr = getLangString("error.end_la.format"), true
	} else if docEndKFactor < 0.0 || docEndKFactor > 2.0 {
		curErr, hasErr = getLangString("error.end_la.value"), true
	} else {
		endKFactor = docEndKFactor
	}
	setErrorDescription(doc, "table.la_values.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Время сглаживания

	docSmoothTime, err := parseInputToFloat(getElementValue("k3d_la_smoothTime"))
	if err != nil {
		curErr, hasErr = getLangString("error.smooth_time.format"), true
	} else if docSmoothTime < 0.005 || docSmoothTime > 0.2 {
		curErr, hasErr = getLangString("error.smooth_time.value"), true
	} else {
		smoothTime = docSmoothTime
	}

	setErrorDescription(doc, "table.smooth_time.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Количество сегментов

	docNumSegment, err := parseInputToInt(getElementValue("k3d_la_numSegments"))
	if err != nil {
		curErr, hasErr = getLangString("error.num_segments.format"), true
	} else if docNumSegment < 2 || docNumSegment > 100 {
		curErr, hasErr = getLangString("error.num_segments.value"), true
	} else {
		numSegments = docNumSegment
	}

	setErrorDescription(doc, "table.num_segments.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Дополнительная цель калибровки

	docTargetNone := doc.Call("getElementById", "k3d_la_targetNone").Get("checked").Bool()
	docTargetSmoothTime := doc.Call("getElementById", "k3d_la_targetSmoothTime").Get("checked").Bool()
	docTargetAcceleration := doc.Call("getElementById", "k3d_la_targetAcceleration").Get("checked").Bool()
	docTargetFlowrate := doc.Call("getElementById", "k3d_la_targetFlowrate").Get("checked").Bool()
	if docTargetNone {
		additionalTarget = 0
	} else if docTargetSmoothTime {
		additionalTarget = 1
	} else if docTargetAcceleration {
		additionalTarget = 2
	} else if docTargetFlowrate {
		additionalTarget = 3
	} else {
		errorString = errorString + getLangString("error.additional_target.not_set") + "\n"
	}

	// Значения дополнительного параметра

	docInitAdditionalParameter, err := parseInputToFloat(getElementValue("k3d_la_initAdditionalParameter"))
	if err != nil {
		curErr, hasErr = getLangString("error.init_additional_parameter.format"), true
	} else if additionalTarget == 1 && (docInitAdditionalParameter < 0.001 || docInitAdditionalParameter > 1.0) {
		// Время сглаживания
		curErr, hasErr = getLangString("error.init_additional_parameter.value"), true
	} else if additionalTarget == 2 && (docInitAdditionalParameter < 500.0 || docInitAdditionalParameter > 50000.0) {
		// Ускорение
		curErr, hasErr = getLangString("error.init_additional_parameter.value"), true
	} else if additionalTarget == 3 && (docInitAdditionalParameter < 1.0 || docInitAdditionalParameter > 200.0) {
		// Расход
		curErr, hasErr = getLangString("error.init_additional_parameter.value"), true
	} else {
		initAdditionalParameter = docInitAdditionalParameter
	}

	setErrorDescription(doc, "table.additional_parameter.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docEndAdditionalParameter, err := parseInputToFloat(getElementValue("k3d_la_endAdditionalParameter"))
	if err != nil {
		curErr, hasErr = getLangString("error.end_additional_parameter.format"), true
	} else if additionalTarget == 1 && (docEndAdditionalParameter < 0.001 || docEndAdditionalParameter > 1.0) {
		// Время сглаживания
		curErr, hasErr = getLangString("error.end_additional_parameter.value"), true
	} else if additionalTarget == 2 && (docEndAdditionalParameter < 500.0 || docEndAdditionalParameter > 50000.0) {
		// Ускорение
		curErr, hasErr = getLangString("error.end_additional_parameter.value"), true
	} else if additionalTarget == 3 && (docEndAdditionalParameter < 1.0 || docEndAdditionalParameter > 200.0) {
		// Расход
		curErr, hasErr = getLangString("error.end_additional_parameter.value"), true
	} else {
		endAdditionalParameter = docEndAdditionalParameter
	}

	setErrorDescription(doc, "table.additional_parameter.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Начальный и конечный G-коды

	docStartGcode := getElementValue("k3d_la_startGcode")
	var errorMessage string
	docStartGcode, errorMessage = processGCode(docStartGcode)
	if errorMessage != "" {
		errorString += errorMessage
		hasErr = false
		retErr = true
	} else {
		startGcode = docStartGcode
	}

	docEndGcode := getElementValue("k3d_la_endGcode")
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
	// Расширенные параметры

	docLayerHeightLimit, err := parseInputToFloat(getElementValue("k3d_la_layerHeightLimit"))
	if docLayerHeightLimit != docLayerHeightLimit {
		docLayerHeightLimit = 0.3
	} else if err != nil || docLayerHeightLimit < 0.1 || docLayerHeightLimit > 10.0 {
		curErr, hasErr = "Error: layer height limit", true
	} else {
		layerHeightLimit = docLayerHeightLimit
	}

	setErrorDescription(doc, "table.advanced_parameters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docFilamentDiameter, err := parseInputToFloat(getElementValue("k3d_la_filamentDiameter"))
	if docFilamentDiameter != docFilamentDiameter {
		docFilamentDiameter = 1.75
	} else if err != nil || docFilamentDiameter < 0.1 || docLayerHeightLimit > 5.0 {
		curErr, hasErr = "Error: filament diameter", true
	} else {
		filamentDiameter = docFilamentDiameter
	}

	setErrorDescription(doc, "table.advanced_parameters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docDefaultSegmentHeight, err := parseInputToFloat(getElementValue("k3d_la_defaultSegmentHeight"))
	if docDefaultSegmentHeight != docDefaultSegmentHeight {
		docDefaultSegmentHeight = 3.0
	} else if err != nil || docDefaultSegmentHeight < 1.0 || docDefaultSegmentHeight > 100.0 {
		curErr, hasErr = "Error: default segment height", true
	} else {
		defaultSegmentHeight = docDefaultSegmentHeight
	}

	setErrorDescription(doc, "table.advanced_parameters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docMinSpeedDelta, err := parseInputToInt(getElementValue("k3d_la_minSpeedDelta"))
	if docMinSpeedDelta != docMinSpeedDelta {
		docMinSpeedDelta = 10
	} else if err != nil || docMinSpeedDelta > 1000 {
		curErr, hasErr = "Error: min speed delta", true
	} else {
		minSpeedDelta = docMinSpeedDelta
	}

	setErrorDescription(doc, "table.advanced_parameters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docMaxSpeedDelta, err := parseInputToInt(getElementValue("k3d_la_maxSpeedDelta"))
	if docMaxSpeedDelta != docMaxSpeedDelta {
		docMaxSpeedDelta = 100
	} else if err != nil || docMaxSpeedDelta < 10 || docMaxSpeedDelta > 1000 {
		curErr, hasErr = "Error: max speed delta", true
	} else {
		maxSpeedDelta = docMaxSpeedDelta
	}

	setErrorDescription(doc, "table.advanced_parameters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docDefaultModelWidth, err := parseInputToFloat(getElementValue("k3d_la_defaultModelWidth"))
	if docDefaultModelWidth != docDefaultModelWidth {
		docDefaultModelWidth = 40.0
	} else if err != nil || docDefaultModelWidth < 10.0 || docDefaultModelWidth > 1000.0 {
		curErr, hasErr = "Error: default model width", true
	} else {
		defaultModelWidth = docDefaultModelWidth
	}

	setErrorDescription(doc, "table.advanced_parameters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	docDefaultSlowLineLength, err := parseInputToFloat(getElementValue("k3d_la_defaultSlowLineLength"))
	if docDefaultSlowLineLength != docDefaultSlowLineLength {
		docDefaultSlowLineLength = 2.0
	} else if err != nil || docDefaultSlowLineLength < 0.1 || docDefaultSlowLineLength > 1000.0 {
		curErr, hasErr = "Error: default slow line length", true
	} else {
		defaultSlowLineLength = docDefaultSlowLineLength
	}

	setErrorDescription(doc, "table.advanced_parameters.description", curErr, hasErr, allowModify)
	if hasErr {
		errorString = errorString + curErr + "\n"
		hasErr = false
		retErr = true
	}

	// Применяем ограничения

	if layerHeight > layerHeightLimit {
		layerHeight = layerHeightLimit
	}

	// Подгоняем высоту сегмента таким образом, чтобы она делилась на высоту слоя без остатка
	// Если в сегменте менее 5 слоёв, то увеличиваем количество слоёв до 5

	segmentHeight = math.Max(layerHeight*(math.Floor(defaultSegmentHeight/layerHeight)), 5*layerHeight)

	// Переопределение ширин линий и толщины слой для разных целей калибровки

	if additionalTarget == 3 {
		lineWidth = roundFloat(nozzleDiameter*1.5, 2)
		firstLayerLineWidth = lineWidth
		layerHeight = roundFloat(nozzleDiameter*0.75, 2)
	} else {
		lineWidth = roundFloat(nozzleDiameter*1.05, 2)
		firstLayerLineWidth = roundFloat(nozzleDiameter*1.5, 2)
		layerHeight = roundFloat(nozzleDiameter*0.5, 2)
	}

	// Если включено логирование, то выводим значение всех параметров в консоль
	if log {
		println("bedX=" + fmt.Sprint(bedX))
		println("bedY=" + fmt.Sprint(bedY))
		println("nozzleDiameter=" + fmt.Sprint(nozzleDiameter))
		println("zOffset=" + fmt.Sprint(zOffset))
		println("hotendTemperature=" + fmt.Sprint(hotendTemperature))
		println("bedTemperature=" + fmt.Sprint(bedTemperature))
		println("cooling=" + fmt.Sprint(cooling))
		println("flow=" + fmt.Sprint(flow))
		println("minPrintSpeed=" + fmt.Sprint(minPrintSpeed))
		println("maxPrintSpeed=" + fmt.Sprint(maxPrintSpeed))
		println("initKFactor=" + fmt.Sprint(initKFactor))
		println("endKFactor=" + fmt.Sprint(endKFactor))
		println("smoothTime=" + fmt.Sprint(smoothTime))
		println("numSegments=" + fmt.Sprint(numSegments))
		println("additionalTarget=" + fmt.Sprint(additionalTarget))
		println("lineWidth=" + fmt.Sprint(lineWidth))
		println("firstLayerLineWidth=" + fmt.Sprint(firstLayerLineWidth))
		println("layerHeight=" + fmt.Sprint(layerHeight))
		println("segmentHeight=" + fmt.Sprint(segmentHeight))
		println("initAdditionalParameter=" + fmt.Sprint(initAdditionalParameter))
		println("endAdditionalParameter=" + fmt.Sprint(endAdditionalParameter))
		println("delta=" + fmt.Sprint(delta))
		println("firmware=" + fmt.Sprint(firmware))
		println("defaultSegmentHeight=" + fmt.Sprint(defaultSegmentHeight))
		println("layerHeightLimit=" + fmt.Sprint(layerHeightLimit))
		println("filamentDiameter=" + fmt.Sprint(filamentDiameter))
		println("minSpeedDelta=" + fmt.Sprint(minSpeedDelta))
		println("maxSpeedDelta=" + fmt.Sprint(maxSpeedDelta))
		println("defaultModelWidth=" + fmt.Sprint(defaultModelWidth))
		println("defaultSlowLineLength=" + fmt.Sprint(defaultSlowLineLength))
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

func checkSegments(this js.Value, i []js.Value) interface{} {
	if check(false, false) {
		segmentStr := getLangString("generator.segment")

		deltaKFactor := math.Abs((endKFactor - initKFactor) / float64(numSegments-1))
		maxKFactor := math.Max(initKFactor, endKFactor)
		deltaAdditionalParameter := math.Abs(((endAdditionalParameter - initAdditionalParameter) / 3))
		minAdditionalParameter := math.Min(initAdditionalParameter, endAdditionalParameter)

		caliParams := ""
		for i := 0; i < numSegments; i++ {
			caliParams += fmt.Sprintf(segmentStr, numSegments-i, fmt.Sprint(roundFloat(maxKFactor-deltaKFactor*float64(i), 3)))
		}
		if additionalTarget != 0 {
			// Значения дополнительного параметра выводятся только если указана дополнительная цель калибровки
			if additionalTarget == 1 {
				// Дополнительная цель - время сглаживания
				caliParams += fmt.Sprintf(getLangString("generator.additional_target.smooth_time"))
			} else if additionalTarget == 2 {
				// Дополнительная цель - ускорение
				caliParams += fmt.Sprintf(getLangString("generator.additional_target.acceleration"))
			} else if additionalTarget == 3 {
				// Дополнительная цель - расход
				caliParams += fmt.Sprintf(getLangString("generator.additional_target.flowrate"))
			}
			caliParams += fmt.Sprintf(getLangString("generator.additional_target.side"),
				fmt.Sprint(roundFloat(minAdditionalParameter, 3)),
				fmt.Sprint(roundFloat(minAdditionalParameter+deltaAdditionalParameter, 3)),
				fmt.Sprint(roundFloat(minAdditionalParameter+2*deltaAdditionalParameter, 3)),
				fmt.Sprint(roundFloat(minAdditionalParameter+3*deltaAdditionalParameter, 3)))
		}

		js.Global().Call("setSegmentsPreview", caliParams)
	} else {
		js.Global().Call("setSegmentsPreview", js.ValueOf(nil))
		check(false, true)
	}
	return js.ValueOf(nil)
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

	// Замена плейсхолдеров потока
	re = regexp.MustCompile(`\$FLOW`)
	outputGCode = re.ReplaceAllString(outputGCode, strconv.Itoa(flow))

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

func generate(this js.Value, i []js.Value) interface{} {

	// Проверить данные, полученные со страницы калибратора
	if !check(true, false) {
		return js.ValueOf(nil)
	}

	// Проверяем рассчитываемые значения, которые могут выдать ошибку
	modelWidth := defaultModelWidth
	minAdditionalParameter := math.Min(initAdditionalParameter, endAdditionalParameter)
	maxAdditionalParameter := math.Max(initAdditionalParameter, endAdditionalParameter)
	if additionalTarget == 3 {
		// Для тест объёмного расхода необходимо рассчитать размер модели так,
		// чтобы при указанном ускорении и рассчитанной максимальной скорости
		// движения, эта максимальная скорость достигалась.
		maxVelocity := int(roundFloat(maxAdditionalParameter/((lineWidth-layerHeight)*layerHeight+math.Pi*layerHeight*layerHeight/4), 0))
		accelerationTime := float64(maxVelocity) / float64(acceleration)
		accelerationTravel := float64(acceleration)*accelerationTime*accelerationTime/2 + 1
		modelWidth = math.Max(float64(accelerationTravel)*4+defaultSlowLineLength, defaultModelWidth)
		if log {
			println("maxVelocity: " + fmt.Sprint(maxVelocity))
			println("accelerationTime: " + fmt.Sprint(accelerationTime))
			println("accelerationTravel: " + fmt.Sprint(accelerationTravel))
			println("modelWidthComputed: " + fmt.Sprint(roundFloat(float64(accelerationTravel)*4+2+defaultSlowLineLength, 3)))
		}
		// Если размер модели больше штатного, выводим информационное сообщение
		if modelWidth > defaultModelWidth {
			if modelWidth > math.Min(bedX, bedY)-20 {
				modelWidth = math.Min(bedX, bedY) - 20
				js.Global().Call("showError", getLangString("error.model_size_limited")+"\n")
				return js.ValueOf(nil)
			} else {
				js.Global().Call("showError", getLangString("info.model_size_changed")+fmt.Sprint(roundFloat(modelWidth, 2))+" мм\n")
			}
		}
	}

	// Контейнеры для информационного сообщения
	infoSpeedDeltaTooSmall := false

	// Вычисление параметров калибровки
	segmentStr := getLangString("generator.segment")

	caliParams := ""
	deltaKFactor := math.Abs((endKFactor - initKFactor) / float64(numSegments-1))
	maxKFactor := math.Max(initKFactor, endKFactor)
	minKFactor := math.Min(initKFactor, endKFactor)
	currentKFactor := minKFactor
	for i := 0; i < numSegments; i++ {
		caliParams += fmt.Sprintf(segmentStr, numSegments-i, fmt.Sprint(roundFloat(maxKFactor-deltaKFactor*float64(i), 3)))
	}
	deltaAdditionalParameter := math.Abs((endAdditionalParameter - initAdditionalParameter) / 3)
	if additionalTarget != 0 {
		// Значения дополнительного параметра выводятся только если указана дополнительная цель калибровки
		if additionalTarget == 1 {
			// Дополнительная цель - время сглаживания
			caliParams += fmt.Sprintf(getLangString("generator.additional_target.smooth_time"))
		} else if additionalTarget == 2 {
			// Дополнительная цель - ускорение
			caliParams += fmt.Sprintf(getLangString("generator.additional_target.acceleration"))
		} else if additionalTarget == 3 {
			// Дополнительная цель - расход
			caliParams += fmt.Sprintf(getLangString("generator.additional_target.flowrate"))
		}
		caliParams += fmt.Sprintf(getLangString("generator.additional_target.side"),
			fmt.Sprint(roundFloat(minAdditionalParameter, 3)),
			fmt.Sprint(roundFloat(minAdditionalParameter+deltaAdditionalParameter, 3)),
			fmt.Sprint(roundFloat(minAdditionalParameter+2*deltaAdditionalParameter, 3)),
			fmt.Sprint(roundFloat(minAdditionalParameter+3*deltaAdditionalParameter, 3)))
	}

	// Название файла
	fileName := "K3D_PA_H" + fmt.Sprint(hotendTemperature) + "_B" + fmt.Sprint(bedTemperature) + "_PA" + fmt.Sprint(roundFloat(initKFactor, 2)) + "-" + fmt.Sprint(roundFloat(endKFactor, 2)) + "_d" + fmt.Sprint(roundFloat(deltaKFactor, 3))
	if additionalTarget == 1 {
		// Время сглаживания
		fileName += "+smooth_time_" + fmt.Sprint(roundFloat(minAdditionalParameter, 3)) + "-" + fmt.Sprint(roundFloat(maxAdditionalParameter, 3)) + "_d" + fmt.Sprint(roundFloat(deltaAdditionalParameter, 3))
	} else if additionalTarget == 2 {
		// Ускорение
		fileName += "+accel_" + fmt.Sprint(roundFloat(minAdditionalParameter, 0)) + "-" + fmt.Sprint(roundFloat(maxAdditionalParameter, 0)) + "_d" + fmt.Sprint(roundFloat(deltaAdditionalParameter, 0))
	} else if additionalTarget == 3 {
		// Расход
		fileName += "+flowrate_" + fmt.Sprint(roundFloat(minAdditionalParameter, 1)) + "-" + fmt.Sprint(roundFloat(maxAdditionalParameter, 1)) + "_d" + fmt.Sprint(roundFloat(deltaAdditionalParameter, 1))
	}
	fileName += ".gcode"

	js.Global().Call("beginSaveFile", fileName)

	// Комментарий со списком параметров в начале файла
	write("; generated by K3D LA calibration ", js.Global().Get("calibrator_version").String(), "\n",
		"; Written by Dmitry Sorkin @ http://k3d.tech/, Kekht and YTKAB0BP\n",
		fmt.Sprintf(";Bedsize: %s:%s [mm]\n", fmt.Sprint(roundFloat(bedX, 1)), fmt.Sprint(roundFloat(bedY, 1))),
		fmt.Sprintf(";Firmware (0-Marlin, 1-Klipper, 2-RRF): %d\n", firmware),
		fmt.Sprintf(";Z-offset: %s [mm]\n", fmt.Sprint(roundFloat(zOffset, 3))),
		fmt.Sprintf(";Delta: %s\n", strconv.FormatBool(delta)),
		fmt.Sprintf(";Temp: %d/%d [°C]\n", hotendTemperature, bedTemperature),
		fmt.Sprintf(";Flow: %d\n", flow),
		fmt.Sprintf(";Fan: %s\n", fmt.Sprint(roundFloat(float64(cooling)/2.55, 1))),
		fmt.Sprintf(";Nozzle diameter: %s [mm]\n", fmt.Sprint(roundFloat(nozzleDiameter, 2))),
		fmt.Sprintf(";Max print speed: %d [mm/s]\n", maxPrintSpeed),
		fmt.Sprintf(";Min print speed: %d [mm/s]\n", minPrintSpeed),
		fmt.Sprintf(";Acceleration: "+fmt.Sprint(acceleration))+" [mm/s^2]\n",
		fmt.Sprintf(";K factor: %s-%s\n", fmt.Sprint(roundFloat(initKFactor, 3)), fmt.Sprint(roundFloat(endKFactor, 3))),
		fmt.Sprintf(";Number of segments: %d\n", numSegments),
		fmt.Sprintf(";Smooth time: %s [s]\n", fmt.Sprint(roundFloat(smoothTime, 3))),
		fmt.Sprintf(";Additional target (0 - None, 1 - Smooth time, 2 - Acceleration, 3 - Flowrate): %d\n", additionalTarget),
		fmt.Sprintf(";Additional parameter: %s-%s\n", fmt.Sprint(roundFloat(initAdditionalParameter, 3)), fmt.Sprint(roundFloat(endAdditionalParameter, 3)))+"\n",
		fmt.Sprintf(";---------")+"\n",
		fmt.Sprintf(";First layer line lidth: "+fmt.Sprint(roundFloat(firstLayerLineWidth, 3)))+"\n",
		fmt.Sprintf(";Model line width: "+fmt.Sprint(roundFloat(lineWidth, 3)))+"\n",
		fmt.Sprintf(";Layer height: "+fmt.Sprint(roundFloat(layerHeight, 3)))+"\n",
		fmt.Sprintf(";Segment height: "+fmt.Sprint(roundFloat(segmentHeight/layerHeight, 0))+" layers ("+fmt.Sprint(roundFloat(segmentHeight, 3))+" mm)")+"\n",
		caliParams)

	// Объявления объекта для работы адаптивной карты высот в klipper
	if firmware == 1 {
		minX := int(roundFloat(bedX/2-55.0, 0))
		minY := int(roundFloat(bedY/2-modelWidth/2-25.0, 0))
		maxX := int(roundFloat(bedX/2+55.0, 0))
		maxY := int(roundFloat(bedY/2+modelWidth/2+15, 0))
		write(fmt.Sprintf("EXCLUDE_OBJECT_DEFINE NAME='Cali_id_0_copy_0' CENTER=%s,%s POLYGON=[[%d,%d],[%d,%d],[%d,%d],[%d,%d],[%d,%d]]\n",
			fmt.Sprint(roundFloat(bedX/2, 0)), fmt.Sprint(roundFloat(bedY/2, 0)),
			minX, minY,
			maxX, minY,
			maxX, maxY,
			minX, maxY,
			minX, minY))
	}

	// Стартовый G-код
	write(startGcode, "\n")

	// Переводим принтер в нужное нам состояние
	write("G90\n")
	write("M82\n")
	write("G92 E0\n")

	// Устанавливаем значение обдува
	write("M106 S0\n")

	// Устанавливаем значение ускорения
	write("M204 S" + fmt.Sprint(acceleration) + "\n")

	// Устанавливаем значение потока
	write("M221 S" + fmt.Sprint(flow) + "\n")

	// Информация о слое для предпросмотрщиков G-кода
	write(";LAYER_CHANGE\n")
	write(";Z:" + fmt.Sprint(roundFloat(layerHeight, 3)) + "\n")
	write(";HEIGHT:" + fmt.Sprint(roundFloat(layerHeight, 3)) + "\n")
	write(";TYPE:Support interface\n")
	write(";WIDTH:" + fmt.Sprint(roundFloat(firstLayerLineWidth, 3)) + "\n")

	// Генерируем первый слой
	var bedCenter Point
	if delta {
		bedCenter.X, bedCenter.Y, bedCenter.Z = 0, 0, layerHeight
	} else {
		bedCenter.X, bedCenter.Y, bedCenter.Z = bedX/2, bedY/2, layerHeight
	}
	currentE = 0
	currentSpeed = minPrintSpeed
	currentCoordinates.X, currentCoordinates.Y, currentCoordinates.Z = 0, 0, 0

	// Передвигаемся по Z на высоту слоя и применяем Z-оффсет
	write(fmt.Sprintf("G1 Z%s\n", fmt.Sprint(roundFloat(layerHeight+zOffset, 2))))
	write(fmt.Sprintf("G92 Z%s\n", fmt.Sprint(roundFloat(layerHeight, 2))))
	currentCoordinates.Z = layerHeight

	// Расчёт точек для печати линии для прочистки сопла
	raftWidth := modelWidth + 10
	var purgeStart Point
	var purgeLineLength float64
	if bedX < 110.0 {
		purgeLineLength = bedX - 10.0
	} else {
		purgeLineLength = 100.0
	}
	purgeStart.X, purgeStart.Y, purgeStart.Z = bedCenter.X-purgeLineLength/2, bedCenter.Y-modelWidth/2-15.0, currentCoordinates.Z
	purgeTwo := purgeStart
	purgeTwo.X = bedCenter.X + purgeLineLength/2
	purgeThree := purgeTwo
	purgeThree.Y += firstLayerLineWidth
	purgeEnd := purgeThree
	purgeEnd.X = purgeStart.X

	// Холостое перемещение к началу линии для прочистки сопла
	write(generateRetraction())
	write(generateMove(currentCoordinates, purgeStart, 0.0, maxPrintSpeed)...)
	write(generateDeretraction())

	// Генерировать команды для печати линии прочистки сопла
	write(generateMove(currentCoordinates, purgeTwo, firstLayerLineWidth, minPrintSpeed)...)
	write(generateMove(currentCoordinates, purgeThree, firstLayerLineWidth, minPrintSpeed)...)
	write(generateMove(currentCoordinates, purgeEnd, firstLayerLineWidth, minPrintSpeed)...)

	// Расчёт точек зигзага на первом слое
	trajectory := generateZigZagTrajectory(bedCenter, firstLayerLineWidth, raftWidth-firstLayerLineWidth*0.5)

	// Перемещение к началу периметра подложки
	raftStart := trajectory[0]
	raftStart.X -= firstLayerLineWidth * 0.75
	raftStart.Y += firstLayerLineWidth * 0.75
	write(generateRetraction())
	write(generateMove(currentCoordinates, raftStart, 0.0, maxPrintSpeed)...)
	write(generateDeretraction())

	// Печать периметра подложки
	write(generateRelativeMove(raftWidth, 0.0, 0.0, firstLayerLineWidth, minPrintSpeed)...)
	write(generateRelativeMove(0.0, -raftWidth, 0.0, firstLayerLineWidth, minPrintSpeed)...)
	write(generateRelativeMove(-raftWidth, 0.0, 0.0, firstLayerLineWidth, minPrintSpeed)...)
	write(generateRelativeMove(0.0, raftWidth, 0.0, firstLayerLineWidth, minPrintSpeed)...)

	// Холостое перемещение к началу печати подложки
	// делается без отката т.к. расстояние перемещения очень маленькое
	write(generateMove(currentCoordinates, trajectory[0], 0.0, maxPrintSpeed)...)

	// Печать подложки
	for i := 1; i < len(trajectory); i++ {
		write(generateMove(currentCoordinates, trajectory[i], firstLayerLineWidth, minPrintSpeed)...)
	}

	// Установить значение LA/PA для первого сегмента
	write(generateLACommand(currentKFactor))

	// Установить значение времени сглаживания PA для всей модели. Если дополнительная цель
	// калибровки будет "время сглаживания", то это значение всё равно переопределится
	if firmware == 1 {
		write("SET_PRESSURE_ADVANCE SMOOTH_TIME=" + fmt.Sprint(roundFloat(smoothTime, 3)) + "\n")
	} else {
		write(";" + getLangString("info.smooth_time_not_supported_by_firmware") + "\n")
	}

	// Тело калибровочной модели
	layersPerSegment := int(segmentHeight / layerHeight)
	for i := 1; i < numSegments*layersPerSegment; i++ {
		// Комментарий в начале каждого слоя
		write(fmt.Sprintf(getLangString("info.layer_change_calibration_info"), fmt.Sprint(roundFloat(currentCoordinates.Z/layerHeight, 0)), fmt.Sprint(roundFloat(currentKFactor, 3))) + "\n")

		// Информация о слое для предпросмотрщиков G-кода
		write(";LAYER_CHANGE\n")
		write(";Z:" + fmt.Sprint(roundFloat(layerHeight*float64(i+1), 3)) + "\n")
		write(";HEIGHT:" + fmt.Sprint(roundFloat(layerHeight, 3)) + "\n")

		// На первых 4 слоях плавно повышаем скорость вращения вентилятора до cooling
		if i < 4 {
			write(fmt.Sprintf("M106 S%s\n", fmt.Sprint(roundFloat(float64(cooling*i/3), 0))))
		}

		// На слоях между сегментов временно увеличиваем размер модели на
		// половину ширины линии, а также меняем значение PA
		wallOffset := 0.0
		if i%layersPerSegment == 0 {
			currentKFactor += deltaKFactor
			write(generateLACommand(currentKFactor))
			wallOffset = lineWidth / 2
		} else {
			wallOffset = 0
		}
		// Ограничение wallOffset
		if wallOffset > maxSeparatorOffset {
			wallOffset = maxSeparatorOffset
		}

		// Перемещение к точке начала печати следующего слоя тела калибровочной модели. Модель
		// печатается начиная с внешнего периметра с левого-заднего угла по часовой стрелке
		var layerStart Point
		layerStart.X = bedCenter.X - (modelWidth+wallOffset-lineWidth/2)/2
		layerStart.Y = bedCenter.Y + (modelWidth+wallOffset-lineWidth/2)/2
		layerStart.Z = layerHeight * float64(i+1)
		if i == 1 {
			// При перемещнии к началу первого слоя, необходимо сделать откат
			write(generateRetraction())
			write(generateMove(currentCoordinates, layerStart, 0.0, maxPrintSpeed)...)
			write(generateDeretraction())
		} else {
			// При перемещении к началу любого другого слоя, откаты не нужны
			write(generateMove(currentCoordinates, layerStart, 0.0, maxPrintSpeed)...)
		}

		// Генерируем тело калибровочной модели
		for j := 0; j < numPerimeters; j++ {
			// Генерируем параметры линий
			currentModelWidth := modelWidth + wallOffset - lineWidth*float64(1+2*j)
			currentSlowLineLength := defaultSlowLineLength
			fastLineLength := (currentModelWidth - currentSlowLineLength) / 2
			fastPrintSpeed := maxPrintSpeed
			slowPrintSpeed := minPrintSpeed

			// Метаданные для слайсера
			if j == 0 {
				write(";TYPE:Outer wall\n")
			} else {
				write(";TYPE:Inner wall\n")
			}
			write(";WIDTH:" + fmt.Sprint(roundFloat(lineWidth, 3)) + "\n")

			// Выводим текущие параметры
			if log {
				println("currentModelWidth: " + fmt.Sprint(currentModelWidth))
				println("currentSlowLineLength: " + fmt.Sprint(currentSlowLineLength))
				println("fastLineLength: " + fmt.Sprint(fastLineLength))
				println("fastPrintSpeed: " + fmt.Sprint(fastPrintSpeed))
				println("slowPrintSpeed: " + fmt.Sprint(slowPrintSpeed))
			}

			// Печать задней стенки
			// Расчёт и установка дополнительного параметра
			if additionalTarget == 1 {
				// Время сглаживания
				write("SET_PRESSURE_ADVANCE SMOOTH_TIME=" + fmt.Sprint(roundFloat(minAdditionalParameter, 3)) + "\n")
			} else if additionalTarget == 2 {
				// Ускорение
				write("M204 S" + fmt.Sprint(roundFloat(minAdditionalParameter, 0)) + "\n")
			} else if additionalTarget == 3 {
				// Объёмный расход
				// Для достижения нужного расхода, необходимо пересчитать скорости быстрых и медленных участков
				// V = Q / (w -h) * h + Pi * h * h / 4
				fastPrintSpeed = int(roundFloat(minAdditionalParameter/((lineWidth-layerHeight)*layerHeight+math.Pi*layerHeight*layerHeight/4), 0))
				slowPrintSpeed = int(math.Max(float64(minPrintSpeed), (float64(fastPrintSpeed - maxSpeedDelta))))
				// Проверка на минимальную разницу между скоростью печати сегментов
				if fastPrintSpeed-slowPrintSpeed < minSpeedDelta {
					infoSpeedDeltaTooSmall = true
				}
			}
			// Печать линии
			write(generateRelativeMove(fastLineLength, 0, 0, lineWidth, fastPrintSpeed)...)
			write(generateRelativeMove(currentSlowLineLength, 0, 0, lineWidth, slowPrintSpeed)...)
			write(generateRelativeMove(fastLineLength, 0, 0, lineWidth, fastPrintSpeed)...)

			// Печать правой стенки
			// Рассчёт и установка дополнительного параметра
			if additionalTarget == 0 {
				// Если не выбрано никакой цели, то изменяем ширину медленного участка
				currentSlowLineLength = 2 * defaultSlowLineLength
				if currentSlowLineLength > modelWidth/2 {
					currentSlowLineLength = modelWidth / 2
				}
				fastLineLength = (currentModelWidth - currentSlowLineLength) / 2
				if log {
					println("currentSlowLineLength: " + fmt.Sprint(currentSlowLineLength))
					println("fastLineLength: " + fmt.Sprint(fastLineLength))
				}
			} else if additionalTarget == 1 {
				// Время сглаживания
				write("SET_PRESSURE_ADVANCE SMOOTH_TIME=" + fmt.Sprint(roundFloat(minAdditionalParameter+deltaAdditionalParameter, 3)) + "\n")
			} else if additionalTarget == 2 {
				// Ускорение
				write("M204 S" + fmt.Sprint(roundFloat(minAdditionalParameter+deltaAdditionalParameter, 0)) + "\n")
			} else if additionalTarget == 3 {
				// Объёмный расход
				// Для достижения нужного расхода, необходимо пересчитать скорости быстрых и медленных участков
				// V = Q / (w -h) * h + Pi * h * h / 4
				fastPrintSpeed = int(roundFloat((minAdditionalParameter+deltaAdditionalParameter)/((lineWidth-layerHeight)*layerHeight+math.Pi*layerHeight*layerHeight/4), 0))
				slowPrintSpeed = int(math.Max(float64(minPrintSpeed), (float64(fastPrintSpeed - maxSpeedDelta))))
				// Проверка на минимальную разницу между скоростью печати сегментов
				if fastPrintSpeed-slowPrintSpeed < minSpeedDelta {
					infoSpeedDeltaTooSmall = true
				}
			}
			// Печать линии
			write(generateRelativeMove(0, -fastLineLength, 0, lineWidth, fastPrintSpeed)...)
			write(generateRelativeMove(0, -currentSlowLineLength, 0, lineWidth, slowPrintSpeed)...)
			write(generateRelativeMove(0, -fastLineLength, 0, lineWidth, fastPrintSpeed)...)

			// Печать передней стенки
			if additionalTarget == 0 {
				// Если не выбрано никакой цели, то изменяем ширину медленного участка
				currentSlowLineLength = 4 * defaultSlowLineLength
				if currentSlowLineLength > modelWidth/2 {
					currentSlowLineLength = modelWidth / 2
				}
				fastLineLength = (currentModelWidth - currentSlowLineLength) / 2
				if log {
					println("currentSlowLineLength: " + fmt.Sprint(currentSlowLineLength))
					println("fastLineLength: " + fmt.Sprint(fastLineLength))
				}
			} else if additionalTarget == 1 {
				// Время сглаживания
				write("SET_PRESSURE_ADVANCE SMOOTH_TIME=" + fmt.Sprint(roundFloat(minAdditionalParameter+2*deltaAdditionalParameter, 3)) + "\n")
			} else if additionalTarget == 2 {
				// Ускорение
				write("M204 S" + fmt.Sprint(roundFloat(minAdditionalParameter+2*deltaAdditionalParameter, 0)) + "\n")
			} else if additionalTarget == 3 {
				// Объёмный расход
				// Для достижения нужного расхода, необходимо пересчитать скорости быстрых и медленных участков
				// V = Q / (w -h) * h + Pi * h * h / 4
				fastPrintSpeed = int(roundFloat((minAdditionalParameter+2*deltaAdditionalParameter)/((lineWidth-layerHeight)*layerHeight+math.Pi*layerHeight*layerHeight/4), 0))
				slowPrintSpeed = int(math.Max(float64(minPrintSpeed), (float64(fastPrintSpeed - maxSpeedDelta))))
				// Проверка на минимальную разницу между скоростью печати сегментов
				if fastPrintSpeed-slowPrintSpeed < minSpeedDelta {
					infoSpeedDeltaTooSmall = true
				}
			}
			// Печать линии
			write(generateRelativeMove(-fastLineLength, 0, 0, lineWidth, fastPrintSpeed)...)
			write(generateRelativeMove(-currentSlowLineLength, 0, 0, lineWidth, slowPrintSpeed)...)
			write(generateRelativeMove(-fastLineLength, 0, 0, lineWidth, fastPrintSpeed)...)

			// Печать левой стенки
			if additionalTarget == 0 {
				// Если не выбрано никакой цели, то изменяем ширину медленного участка
				currentSlowLineLength = 8 * defaultSlowLineLength
				if currentSlowLineLength > modelWidth/2 {
					currentSlowLineLength = modelWidth / 2
				}
				fastLineLength = (currentModelWidth - currentSlowLineLength) / 2
				if log {
					println("currentSlowLineLength: " + fmt.Sprint(currentSlowLineLength))
					println("fastLineLength: " + fmt.Sprint(fastLineLength))
				}
			} else if additionalTarget == 1 {
				// Время сглаживания
				write("SET_PRESSURE_ADVANCE SMOOTH_TIME=" + fmt.Sprint(roundFloat(maxAdditionalParameter, 3)) + "\n")
			} else if additionalTarget == 2 {
				// Ускорение
				write("M204 S" + fmt.Sprint(roundFloat(maxAdditionalParameter, 0)) + "\n")
			} else if additionalTarget == 3 {
				// Объёмный расход
				// Для достижения нужного расхода, необходимо пересчитать скорости быстрых и медленных участков
				// V = Q / (w -h) * h + Pi * h * h / 4
				fastPrintSpeed = int(roundFloat(maxAdditionalParameter/((lineWidth-layerHeight)*layerHeight+math.Pi*layerHeight*layerHeight/4), 0))
				slowPrintSpeed = int(math.Max(float64(minPrintSpeed), (float64(fastPrintSpeed - maxSpeedDelta))))
				// Проверка на минимальную разницу между скоростью печати сегментов
				if fastPrintSpeed-slowPrintSpeed < minSpeedDelta {
					infoSpeedDeltaTooSmall = true
				}
			}
			// Печать линии
			write(generateRelativeMove(0, fastLineLength, 0, lineWidth, fastPrintSpeed)...)
			write(generateRelativeMove(0, currentSlowLineLength, 0, lineWidth, slowPrintSpeed)...)
			write(generateRelativeMove(0, fastLineLength, 0, lineWidth, fastPrintSpeed)...)

			// Если это не последний периметр, то перемещаемся к началу следующего периметра

			if j != numPerimeters-1 {
				write(generateRelativeMove(lineWidth, -lineWidth, 0, 0.0, maxPrintSpeed)...)
			}
		}
	}

	// Конечный гкод
	write(endGcode)

	// Вывод информационного сообщения
	infoMessage := ""
	if infoSpeedDeltaTooSmall {
		infoMessage += getLangString("info.delta_print_speed_too_small")
	}
	if infoMessage != "" {
		js.Global().Call("showError", infoMessage)
	}

	// Вывод параметров калибровки в информационное окно
	js.Global().Call("showError", caliParams)

	// save file
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

func generateZigZagTrajectory(towerCenter Point, lineWidth, raftWidth float64) []Point {
	sideLength := raftWidth - lineWidth
	pointsOnOneSide := int(sideLength / (lineWidth * math.Sqrt(2)))
	pointsOnOneSide = pointsOnOneSide - (pointsOnOneSide-1)%2
	pointSpacing := sideLength / float64(pointsOnOneSide-1)
	firstLayerLineWidth = pointSpacing / math.Sqrt(2)

	totalPoints := pointsOnOneSide*4 - 4
	unsortedPoints := make([]Point, totalPoints)

	minX := towerCenter.X - sideLength/2
	minY := towerCenter.Y - sideLength/2
	maxX := towerCenter.X + sideLength/2
	maxY := towerCenter.Y + sideLength/2

	// Generate unsorted slice of points clockwise
	for i := 0; i <= pointsOnOneSide-1; i++ {
		unsortedPoints[i].X = minX + pointSpacing*float64(i)
		unsortedPoints[i].Y = maxY
	}
	for i := 1; i <= pointsOnOneSide-1; i++ {
		unsortedPoints[pointsOnOneSide+i-1].X = maxX
		unsortedPoints[pointsOnOneSide+i-1].Y = maxY - pointSpacing*float64(i)
	}
	for i := 1; i <= pointsOnOneSide-1; i++ {
		unsortedPoints[pointsOnOneSide*2+i-2].X = maxX - pointSpacing*float64(i)
		unsortedPoints[pointsOnOneSide*2+i-2].Y = minY
	}
	for i := 1; i < pointsOnOneSide-1; i++ {
		unsortedPoints[pointsOnOneSide*3+i-3].X = minX
		unsortedPoints[pointsOnOneSide*3+i-3].Y = minY + pointSpacing*float64(i)
	}

	// add Z coordinates

	for i := 1; i < len(unsortedPoints); i++ {
		unsortedPoints[i].Z = currentCoordinates.Z
	}

	// Sort points to make zigzag moves
	trajectory := make([]Point, len(unsortedPoints))

	trajectory[0] = unsortedPoints[0]
	trajectory[1] = unsortedPoints[len(unsortedPoints)-1]
	trajectory[2] = unsortedPoints[1]
	trajectory[3] = unsortedPoints[2]
	for i := 4; i < len(unsortedPoints); i = i + 4 {
		j := int(i / 2)
		trajectory[i] = unsortedPoints[len(unsortedPoints)-j]
		trajectory[i+1] = unsortedPoints[len(unsortedPoints)-j-1]
		trajectory[i+2] = unsortedPoints[j+1]
		trajectory[i+3] = unsortedPoints[j+2]
	}

	for i := 0; i < len(trajectory); i++ {
		trajectory[i].Z = currentCoordinates.Z
	}

	return trajectory
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
