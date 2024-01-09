---
icon: octicons/package-16
title: K3D Калибровщик Linear Advance
hide:
    - toc
---

<h1 class="lang" id="header.title">K3D калибровщик Linear Advance</h1>

<script src="../assets/js/lib.js"></script>
<script src="../assets/js/wasm_exec.js"></script>
<script src="../assets/js/gwaloader.js"></script>
<script src="https://cdn.jsdelivr.net/npm/streamsaver@2.0.3/StreamSaver.min.js"></script>

[:material-file-document: Инструкция](./index.md) | Язык: [:flag_gb: English](?lang=en) [:flag_ru: Русский](?lang=ru)

<table style="width: 100%; font-size: 0.8rem;">
    <tbody>
    <!-- Параметры принтера -->
        <tr>
            <th class="lang" id="table.bed_size_x.title">Размер стола по X</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_bedX" name="k3d_la_bedX" value="235"></td>
            <td class="lang" id="table.bed_size_x.description" style="text-align: justify;">[мм] Для декартовых принтеров - максимальная координата по оси X<br>Для дельта-принтеров - <b>диаметр стола</b></td>
        </tr>
        <tr>
            <th class="lang" id="table.bed_size_y.title">Размер стола по Y</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_bedY" name="k3d_la_bedY" value="235"></td>
            <td class="lang" id="table.bed_size_y.description" style="text-align: justify;">[мм] Для декартовых принтеров - максимальная координата по оси Y<br>Для дельта-принтеров - <b>диаметр стола</b></td>
        </tr>
        <tr>
            <th class="lang" id="table.firmware.title">Прошивка</td>
            <td align="center">
                <form style="text-align:left; width:fit-content;"><input type="radio" id="k3d_la_firmwareMarlin" name="k3d_la_firmware" value="Marlin" checked><label for="k3d_la_firmwareMarlin"> Marlin</label><br>
                <input type="radio" id="k3d_la_firmwareKlipper" name="k3d_la_firmware" value="Klipper"><label for="k3d_la_firmwareKlipper"> Klipper</label><br>
                <input type="radio" id="k3d_la_firmwareRRF" name="k3d_la_firmware" value="RRF"><label for="k3d_la_firmwareRRF"> RRF</label>
                </form>
            </td>
            <td class="lang" id="table.firmware.description" style="text-align: justify;">Прошивка, установленная на вашем принтере. Если не знаете, то, скорее всего, Marlin</td>
        </tr>
        <tr>
            <th class="lang" id="table.delta.title">Начало координат в центре стола</td>
            <td style="text-align:center"><input type="checkbox" id="k3d_la_delta" name="k3d_la_delta"></td>
            <td class="lang" id="table.delta.description" style="text-align: justify;">Для декартовых принтеров должно быть выключено, для дельт включено</td>
        </tr>
        <tr>
            <th class="lang" id="table.bed_probe.title">Автокалибровка стола</td>
            <td style="text-align:center"><input type="checkbox" id="k3d_la_g29" name="k3d_la_g29"></td>
            <td class="lang" id="table.bed_probe.description" style="text-align: justify;">Надо ли делать автокалибровку стола перед печатью (G29)? Если у вас нет датчика автокалибровки, то
                оставляйте выключенным</td>
        </tr>
        <tr>
            <th class="lang" id="table.travel_speed.title">Скорость перемещений</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_travelSpeed" name="k3d_la_travelSpeed" value="150"></td>
            <td class="lang" id="table.travel_speed.description" style="text-align: justify;">[мм/с] Скорость, с которой будут происходить перемещения без экструдирования</td>
        </tr>
    <!-- Параметры филамента -->
        <tr>
            <th class="lang" id="table.hotend_temp.title">Температура хотэнда</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_hotendTemperature" name="k3d_la_hotendTemperature" value="210"></td>
            <td class="lang" id="table.hotend_temp.description" style="text-align: justify;">[°C] До прогрева стола хотэнд будет нагрет до 150 градусов. После полного нагрева стола хотэнд догреется до
                указанной температуры</td>
        </tr>
        <tr>
            <th class="lang" id="table.bed_temp.title">Температура стола</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_bedTemperature" name="k3d_la_bedTemperature" value="60"></td>
            <td class="lang" id="table.bed_temp.description">[°C] Температура, до которой нагреть стол перед печатью. Стол будет нагрет до выполнения парковки и автокалибровки стола</td>
        </tr>
        <tr>
            <th class="lang" id="table.fan_speed.title">Скорость вентилятора</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_cooling" name="k3d_la_cooling" value="100"></td>
            <td class="lang" id="table.fan_speed.description" style="text-align: justify;">[%] Обороты вентилятора в процентах. Для того, чтобы температура хотэнда резко не упала при включении вентилятора, на 1 слое он будет включен на 1/3 от заданного значения, на 2 слое на 2/3, на 4 слое на заданное значение</td>
        </tr>
        <tr>
            <th class="lang" id="table.flow.title">Поток</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_flow" name="k3d_la_flow" value="100"></td>
            <td class="lang" id="table.flow.description" style="text-align: justify;">[%] Поток в процентах. Нужен для компенсации пере- или недоэкструзии</td>
        </tr>
    <!-- Параметры первого слоя -->
        <tr>
            <th class="lang" id="table.first_line_width.title">Ширина линии первого слоя</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_firstLayerLineWidth" name="k3d_la_firstLayerLineWidth" value="0.6"></td>
            <td class="lang" id="table.first_line_width.description" style="text-align: justify;">[мм] Ширина линий, с которой будет напечатана подложка под моделью. В общем случае рекомендуется выставить 150% от диаметра сопла</td>
        </tr>
        <tr>
            <th class="lang" id="table.first_print_speed.title">Скорость печати первого слоя</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_firstLayerSpeed" name="k3d_la_firstLayerSpeed" value="30"></td>
            <td class="lang" id="table.first_print_speed.description" style="text-align: justify;">[мм/с] Скорость, с которой будет напечатана подложка</td>
        </tr>
        <tr>
            <th class="lang" id="table.z_offset.title">Z-offset</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_zOffset" name="k3d_la_zOffset" value="0.0"></td>
            <td class="lang" id="table.z_offset.description" style="text-align: justify;">[мм] Смещение всей модели по вертикали. Нужно чтобы компенсировать слишком тонкую/толстую калибровку первого слоя. В общем случае оставьте ноль</td>
        </tr>
    <!-- Параметры модели -->
        <tr>
            <th class="lang" id="table.num_perimeters.title">Количество периметров</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_numPerimeters" name="k3d_la_numPerimeters" value="2"></td>
            <td class="lang" id="table.num_perimeters.description" style="text-align: justify;">Количество периметров для основного тела калибровочной модели. Для филаментов с околонулевой усадкой (PLA,
                некоторые композиты) 1-2. Для филаментов с сильной усадкой (ABS и подобные) 2+. Для флексов 2-4 в зависимости
                от их жесткости и желаемой высоты башенки</td>
        </tr>
        <tr>
            <th class="lang" id="table.line_width.title">Ширина линии</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_lineWidth" name="k3d_la_lineWidth" value="0.4"></td>
            <td class="lang" id="table.line_width.description" style="text-align: justify;">[мм] Ширина линий, с которой будут напечатана тестовая модель. В общем случае рекомендуется выставить равной диаметру сопла</td>
        </tr>
        <tr>
            <th class="lang" id="table.layer_height.title">Толщина слоя</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_layerHeight" name="k3d_la_layerHeight" value="0.2"></td>
            <td class="lang" id="table.layer_height.description" style="text-align: justify;">[мм] Толщина слоёв всей модели. В общем случае 50% от ширины линии</td>
        </tr>
    <!-- Параметры калибровки -->
        <tr>
            <th class="lang" id="table.fast_segment_speed.title">Скорость быстрых участков</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_fastPrintSpeed" name="k3d_la_fastPrintSpeed" value="100"></td>
            <td class="lang" id="table.fast_segment_speed.description" style="text-align: justify;">[мм/с] Скорость, с которой будут печататься быстрые участки. Лучше указать высокие значения (100-150)</td>
        </tr>
        <tr>
            <th class="lang" id="table.slow_segment_speed.title">Скорость медленных участков</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_slowPrintSpeed" name="k3d_la_slowPrintSpeed" value="20"></td>
            <td class="lang" id="table.slow_segment_speed.description" style="text-align: justify;">[мм/с] Скорость, с которой будут печататься медленные участки. Лучше указать низкие значения (10-30)</td>
        </tr>
        <tr>
            <th class="lang" id="table.init_la.title">Начальное значение коэффициента LA</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_initKFactor" name="k3d_la_initKFactor" value="0.0"></td>
            <td class="lang" id="table.init_la.description" style="text-align: justify;">С какого значения к-фактора начать калибровку. Округляется до 3 знака после разделителя</td>
        </tr>
        <tr>
            <th class="lang" id="table.end_la.title">Конечное значение коэффициента LA</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_endKFactor" name="k3d_la_endKFactor" value="0.2"></td>
            <td class="lang" id="table.end_la.description" style="text-align: justify;">До какого значения к-фактора проводить калибровку. Округляется до 3 знака после разделителя. Для директ
                экструдеров обычно хватает 0.2, для боуденов 1.5</td>
        </tr>
        <tr>
            <th class="lang" id="table.num_segments.title">Количество сегментов</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_numSegments" name="k3d_la_numSegments" value="10"></td>
            <td class="lang" id="table.num_segments.description" style="text-align: justify;">Количество сегментов башенки. В течение сегмента коэффициент LA остаётся неизменным. Сегменты визуально
                разделены для упрощения анализа модели</td>
        </tr>
        <tr>
            <th class="lang" id="table.segment_height.title">Высота сегмента</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_segmentHeight" name="k3d_la_segmentHeight" value="3.0"></td>
            <td class="lang" id="table.segment_height.description" style="text-align: justify;">[мм] Высота одного сегмента. К примеру, если высота сегмента 3мм, а количество сегментов 10, то
                высота всей модели будет 30мм</td>
        </tr>
    </tbody>
</table>

<table class="calibrationCodeTable" style="padding-top: 0; margin-top: 0;">
    <tbody>
        <tr>
            <th class="lang" id="table.start_gcode.title" style="text-align: center;">Начальный G-код</th>
            <th class="lang" id="table.end_gcode.title" style="text-align: center;">Конечный G-код</th>
        </tr>
        <tr>
            <td width="50%"><textarea type="text" id="k3d_la_startGcode" name="k3d_la_startGcode" rows="11">
M104 S150 ;прогреть хотэнд до 150 градусов
M190 S$BEDTEMP ;прогреть стол до температуры, указанной в настройках
M109 S$HOTTEMP ;прогреть хотэнд до температуры, указанной в настройках
G28 ;припарковать все оси
$G29 ;снять карту высот стола
G90 ;абсолютная система координат
G92 E0 ;сбросить координату экструдера
M220 S100 ;Множитель скорости 100%
M221 S$FLOW ;Множитель потока взять из настроек</textarea></td>
            <td width="50%"><textarea type="text" id="k3d_la_endGcode" name="k3d_la_endGcode" rows="11">
M104 S0 ;выключить хотэнд
M140 S0 ;выключить нагрев стола
M106 S0 ;выключить вентилятор модели
G91 ;относительная система координат
G1 E-5 F600 ;сделать откат на 5мм
G1 Z1 F300 ;поднять голову на 1мм</textarea></td>
        </tr>
        <tr>
            <td class="lang" id="table.start_gcode.description" style="text-align: justify;" colspan="2">
                Начальный код выполняется перед печатью модели и служит для правильной инициализации принтера. Конечный код выполняется после окончания печати. На большинстве принтеров стандартные коды работают хорошо и менять их не надо. <a href="../#g-">Список возможных плейсхолдеров</a>
            </td>
        </tr>
    </tbody>
</table>

<table class="caliButtonTable">
    <tbody>
        <tr>
            <td align="right" width="50%">
                <button class="caliButton" onclick="generate();" id="generateButton" style="display:none">Генерировать и скачать</button>
                <p id="generateButtonLoading"> Генератор загружается...</p>
            </td>
            <td align="left" width="50%">
                <button class="caliButton" onclick="reset();" id="resetButton">Сбросить настройки</button>
            </td>
        </tr>
        <tr>
            <td align="center" colspan="2">
                <br><div id="resultContainer"></div>
            </td>
        </tr>
    </tbody>
</table>

<script>document.body.onload = init();</script>