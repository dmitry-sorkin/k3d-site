---
icon: octicons/package-16
title: K3D Калибровщик откатов
hide:
    - toc
---

<h1 class="lang" id="header.title">K3D калибровщик откатов</h1>

<script src="../assets/js/lib.js"></script>
<script src="../assets/js/wasm_exec.js"></script>
<script src="../assets/js/gwaloader.js"></script>
<script src="https://cdn.jsdelivr.net/npm/streamsaver@2.0.3/StreamSaver.min.js"></script>

[:material-file-document: Инструкция](./index.md) | Язык: [:flag_gb: English](?lang=en) [:flag_ru: Русский](?lang=ru)

<table style="width: 100%; font-size: 0.8rem;">
    <tbody>
      <tr>
        <td class="lang" id="table.bed_size_x.title">Размер стола по X</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="bedX" name="BedX" value="235"></td>
        <td class="lang" id="table.bed_size_x.description" style="text-align: justify;">[мм] Для декартовых принтеров - максимальная координата по оси X<br>Для дельта-принтеров - <b>диаметр стола</b></td>
      </tr>
      <tr>
        <td class="lang" id="table.bed_size_y.title">Размер стола по Y</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="bedY" name="BedY" value="235"></td>
        <td class="lang" id="table.bed_size_y.description" style="text-align: justify;">[мм] Для декартовых принтеров - максимальная координата по оси Y<br>Для дельта-принтеров - <b>диаметр стола</b></td>
      </tr>
      <tr>
        <td class="lang" id="table.firmware.title">Прошивка</td>
        <td align="center">
          <form style="text-align:left; width:fit-content;">
			<input type="radio" id="firmwareMarlin" name="firmware" value="Marlin" checked><label for="firmwareMarlin">Marlin</label><br>
            <input type="radio" id="firmwareKlipper" name="firmware" value="Klipper"><label for="firmwareKlipper">Klipper</label><br>
            <input type="radio" id="firmwareRRF" name="firmware" value="RRF"><label for="firmwareRRF">RRF</label>
          </form>
        </td>
        <td class="lang" id="table.firmware.description">Прошивка, установленная на вашем принтере. Если не знаете, то, скорее всего, Marlin</td>
      </tr>
      <tr>
        <td class="lang" id="table.z_offset.title">Z-offset</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="zOffset" name="zOffset" value="0.0"></td>
        <td class="lang" id="table.z_offset.description" style="text-align: justify;">[мм] Смещение всей модели по вертикали. Нужно чтобы компенсировать слишком тонкую/толстую калибровку первого слоя. В общем случае оставьте ноль</td>
      </tr>
      <tr>
        <td class="lang" id="table.delta.title">Начало координат в центре стола</td>
        <td style="text-align:center"><input type="checkbox" id="delta" name="delta"></td>
        <td class="lang" id="table.delta.description" style="text-align: justify;">Для декартовых принтеров должно быть выключено, для дельт включено</td>
      </tr>
      <tr>
        <td class="lang" id="table.bed_probe.title">Автокалибровка стола</td>
        <td style="text-align:center"><input type="checkbox" id="bedProbe" name="bedProbe"></td>
        <td class="lang" id="table.bed_probe.description" style="text-align: justify;">Надо ли делать автокалибровку стола перед печатью (G29)? Если у вас нет датчика автокалибровки, то оставляйте
          выключенным</td>
      </tr>
      <tr>
        <td class="lang" id="table.hotend_temp.title">Температура хотэнда</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="hotendTemperature" name="hotendTemperature" value="210"></td>
        <td class="lang" id="table.hotend_temp.description" style="text-align: justify;">[°C] Температура, до которой нагреть хотэнд перед печатью</td>
      </tr>
      <tr>
        <td class="lang" id="table.bed_temp.title">Температура стола</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="bedTemperature" name="bedTemperature" value="60"></td>
        <td class="lang" id="table.bed_temp.description" style="text-align: justify;">[°C] Температура, до которой нагреть стол перед печатью. Стол будет нагрет до выполнения парковки и автокалибровки стола</td>
      </tr>
	  <tr>
        <td class="lang" id="table.flow.title">Поток</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="flow" name="flow" value="100"></td>
        <td class="lang" id="table.flow.description" style="text-align: justify;">[%] Поток в процентах. Нужен для компенсации пере- или недоэкструзии</td>
      </tr>
      <tr>
        <td class="lang" id="table.fan_speed.title">Скорость вентилятора</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="cooling" name="cooling" value="100"></td>
        <td class="lang" id="table.fan_speed.description" style="text-align: justify;">[%] Обороты вентилятора в процентах. Для того, чтобы температура хотэнда резко не упала при включении вентилятора, на 1 слое он будет включен на 1/3 от заданного значения, на 2 слое на 2/3, на 4 слое на заданное значение</td>
      </tr>
      <tr>
        <td class="lang" id="table.line_width.title">Ширина линии</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="lineWidth" name="lineWidth" value="0.4"></td>
        <td class="lang" id="table.line_width.description" style="text-align: justify;">[мм] Ширина линий, с которой будут напечатаны башенки. В общем случае рекомендуется выставить равной диаметру сопла</td>
      </tr>
      <tr>
        <td class="lang" id="table.first_line_width.title">Ширина линии первого слоя</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="firstLayerLineWidth" name="firstLayerLineWidth" value="0.6"></td>
        <td class="lang" id="table.first_line_width.description" style="text-align: justify;">[мм] Ширина линий, с которой будет напечатана подложка под башенками. В общем случае рекомендуется выставить 150% от диаметра сопла</td>
      </tr>
      <tr>
        <td class="lang" id="table.layer_height.title">Толщина слоя</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="layerHeight" name="layerHeight" value="0.25"></td>
        <td class="lang" id="table.layer_height.description" style="text-align: justify;">[мм] Толщина слоёв всей модели. В общем случае 50% от ширины линии</td>
      </tr>
      <tr>
        <td class="lang" id="table.print_speed.title">Скорость печати</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="printSpeed" name="printSpeed" value="60"></td>
        <td class="lang" id="table.print_speed.description" style="text-align: justify;">[мм/с] Скорость, с которой будут напечатаны башенки</td>
      </tr>
      <tr>
        <td class="lang" id="table.first_print_speed.title">Скорость печати первого слоя</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="firstLayerPrintSpeed" name="firstLayerPrintSpeed" value="30"></td>
        <td class="lang" id="table.first_print_speed.description" style="text-align: justify;">[мм/с] Скорость, с которой будет напечатана подложка под башенки</td>
      </tr>
      <tr>
        <td class="lang" id="table.travel_speed.title">Скорость перемещений</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="travelSpeed" name="travelSpeed" value="150"></td>
        <td class="lang" id="table.travel_speed.description" style="text-align: justify;">[мм/с] Скорость перемещений между башенками</td>
      </tr>
      <tr>
        <td class="lang" id="table.init_retract_length.title">Начальная длина отката</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="initRetractLength" name="initRetractLength" value="1.0"></td>
        <td class="lang" id="table.init_retract_length.description" style="text-align: justify;">[мм] Длина отката, с которой будет напечатан нижний сегмент</td>
      </tr>
      <tr>
        <td class="lang" id="table.end_retract_length.title">Конечная длина отката</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="endRetractLength" name="endRetractLength" value="0.2"></td>
        <td class="lang" id="table.end_retract_length.description" style="text-align: justify;">[мм] Длина отката, с которой будет напечатан верхний сегмент. Между нижним и верхним сегментом длина отката
          будет изменяться ступенчато за каждый сегмент. Если хотите, чтобы длина отката не менялась, то укажите такое
          же значение, как начальное</td>
      </tr>
      <tr>
        <td class="lang" id="table.init_retract_speed.title">Начальная скорость отката</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="initRetractSpeed" name="initRetractSpeed" value="30"></td>
        <td class="lang" id="table.init_retract_speed.description" style="text-align: justify;">[мм/с] Скорость, с которой будут выполняться откаты</td>
      </tr>
      <tr>
        <td class="lang" id="table.end_retract_speed.title">Конечная скорость отката</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="endRetractSpeed" name="endRetractSpeed" value="30"></td>
        <td class="lang" id="table.end_retract_speed.description" style="text-align: justify;">[мм/с] Скорость отката, с которой будет напечатан верхний сегмент. Между нижним и верхним сегментом скорость отката будет изменяться ступенчато за каждый сегмент. Если хотите, чтобы скорость отката не менялась, то укажите такое же значение, как начальное</td>
      </tr>
      <tr>
        <td class="lang" id="table.num_segments.title">Количество сегментов</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="numSegments" name="numSegments" value="10"></td>
        <td class="lang" id="table.num_segments.description" style="text-align: justify;">Количество сегментов башенки. В течение сегмента длина и скорость отката остаются неизменными. Сегменты визуально разделены для упрощения анализа модели</td>
      </tr>
      <tr>
        <td class="lang" id="table.segment_height.title">Высота сегмента</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="segmentHeight" name="segmentHeight" value="3"></td>
        <td class="lang" id="table.segment_height.description" style="text-align: justify;">[мм] Высота одного сегмента башенки. К примеру, если высота сегмента 3мм, а количество сегментов 10, то
          высота всей башенки будет 30мм</td>
      </tr>
      <tr>
        <td class="lang" id="table.k_factor.title">k-фактор Linear Advance</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="kFactor2" name="kFactor2" value="0.0"></td>
        <td class="lang" id="table.k_factor.description" style="text-align: justify;">Введите сюда ваше значение для Linear/Pressure Advance. Если вы не пользуетесь Linear/Pressure Advance, то очистите поле или оставьте значение нулевым</td>
      </tr>
      <tr>
        <td class="lang" id="table.tower_spacing.title">Расстояние между башенками</td>
        <td style="text-align:center"><input class="calibratorInput" type="text" id="towerSpacing" name="towerSpacing" value="100"></td>
        <td class="lang" id="table.tower_spacing.description" style="text-align: justify;">[мм] Для проверки откатов, обычно, хватает около 100 мм. Для крупногабаритных принтеров, которые часто
          печатают большие модели, рекомендуется около половины длины большей стороны стола</td>
      </tr>
      <tr>
        <td class="lang" id="table.hardmode.title">Усложненный режим</td>
        <td style="text-align:center"><input type="checkbox" id="hardmode" name="hardmode"></td>
        <td class="lang" id="table.hardmode.description" style="text-align: justify;">В обычном режиме (параметр выключен) порядок печати башен генерируется с оптимизациями как в слайсере, в усложненном - неоптимальным методом. Рекомендуется включать только тогда, когда обычный режим показывает слишком оптимистичный результат. Подробнее в инструкции</td>
      </tr>
	  <tr>
    </tbody>
</table>

<table class="calibrationCodeTable" style="padding-top: 0; margin-top: 0;">
    <tbody>
        <tr>
            <th class="lang" id="table.start_gcode.title" style="text-align: center;">Начальный G-код</th>
            <th class="lang" id="table.end_gcode.title" style="text-align: center;">Конечный G-код</th>
        </tr>
        <tr>
            <td width="50%"><textarea type="text" id="startGcode" name="startGcode" rows="11">
M104 S150 ;прогреть хотэнд до 150 градусов
M190 S$BEDTEMP ;прогреть стол до температуры, указанной в настройках
M109 S$HOTTEMP ;прогреть хотэнд до температуры, указанной в настройках
G28 ;припарковать все оси
$G29 ;снять карту высот стола
G90 ;абсолютная система координат
G92 E0 ;сбросить координату экструдера
M220 S100 ;Множитель скорости 100%
M221 S$FLOW ;Множитель потока взять из настроек</textarea></td>
            <td width="50%"><textarea type="text" id="endGcode" name="endGcode" rows="11">
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