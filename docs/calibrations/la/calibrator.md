---
icon: octicons/package-16
title: K3D Калибровщик Linear Advance (Pressure Advance)
hide:
    - toc
---

<h1 class="lang" id="header.title">K3D калибровщик Pressure Advance</h1>

<script src="../assets/js/lib.js"></script>
<script src="../assets/js/wasm_exec.js"></script>
<script src="../assets/js/gwaloader.js"></script>
<script src="../assets/js/streamsavermin.js"></script>

[:material-file-document: Инструкция](./index.md) | Язык: [:flag_gb: English](?lang=en) [:flag_ru: Русский](?lang=ru)

<table style="width: 100%; font-size: 0.7rem;">
    <tbody>
    <!-- Параметры принтера -->
        <!-- Размер стола -->
        <tr>
            <th class="lang" id="table.bed_size.title">Размер стола</td>
            <td style="text-align:center">
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">X</td>
                            <td class="innerTableTD"  style="text-align:center">Y</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_bedX" name="k3d_la_bedX" value="235"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_bedY" name="k3d_la_bedY" value="235"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.bed_size.description" style="text-align: justify; font-size: 0.9em;">[мм] Для декартовых принтеров - размер области печати по осям X/Y<br>Для дельта-принтеров - диаметр стола</td>
        </tr>
        <!-- Прошивка -->
        <tr>
            <th class="lang" id="table.firmware.title">Прошивка</td>
            <td align="center">
                <form style="text-align:left; width:fit-content;"><input type="radio" id="k3d_la_marlin" name="k3d_la_firmware" value="Marlin" checked><label for="k3d_la_marlin"> Marlin</label><br>
                <input type="radio" id="k3d_la_klipper" name="k3d_la_firmware" value="Klipper"><label for="k3d_la_klipper"> Klipper</label><br>
                <input type="radio" id="k3d_la_rrf" name="k3d_la_firmware" value="RRF"><label for="k3d_la_rrf"> RRF</label>
                </form>
            </td>
            <td class="lang" id="table.firmware.description" style="text-align: justify; font-size: 0.9em;">Прошивка, установленная на вашем принтере<br>Если не знаете, то, скорее всего, Klipper<br>Для принтеров Bambu Lab выберите Marlin</td>
        </tr>
        <!-- Начало координат в центре стола (режим дельты) -->
        <tr>
            <th class="lang" id="table.delta.title">Начало координат в<br>центре стола</td>
            <td style="text-align:center"><input type="checkbox" id="k3d_la_delta" name="k3d_la_delta"></td>
            <td class="lang" id="table.delta.description" style="text-align: justify; font-size: 0.9em;">Для декартовых принтеров должно быть выключено<br>Для дельт и AD5M включено</td>
        </tr>
        <!-- Диаметр сопла -->
        <tr>
            <th class="lang" id="table.nozzle_diameter.title">Диаметр сопла</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_nozzleDiameter" name="k3d_la_nozzleDiameter" value="0.4"></td>
            <td class="lang" id="table.nozzle_diameter.description" style="text-align: justify; font-size: 0.9em; font-size: 0.9em;">[мм] Диаметр сопла, установленного на вашем принтере. Влияет на:<br>
             - Ширина линии первого слоя = D * 1.5<br>
             - Ширина линии модели = D * 1.05<br>
             - Толщина слоя = D * 0.5<br>
             При использовании доп. цели калибровки "расход", ширина линии модели и толщина слоя будут увеличены до 1.5*D и 0.75*D соответственно<br>
             Если хотите использовать линии уже или слой тоньше, то можете указать заниженный диаметр сопла
            </td>
        </tr>
        <tr>
            <th class="lang" id="table.z_offset.title">Z offset</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_zOffset" name="k3d_la_zOffset" value="0.0"></td>
            <td class="lang" id="table.z_offset.description" style="text-align: justify; font-size: 0.9em;">[мм] Смещение всей модели по вертикали для компенсации неправильного расстояния от сопла до стола на первом слое. Положительное значение смещает модель вверх, отрицательное - вниз.<br>Если в вашем стартовом G-коде уже есть установка Z-offset'а, то оставьте 0</td>
        </tr>
    <!-- Параметры филамента -->
        <!-- Температуры -->
        <tr>
            <th class="lang" id="table.temperatures.title">Температуры</td>
            <td style="text-align:center">
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">Hotend</td>
                            <td class="innerTableTD"  style="text-align:center">Bed</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_hotendTemperature" name="k3d_la_hotendTemperature" value="210"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_bedTemperature" name="k3d_la_bedTemperature" value="60"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.temperatures.description" style="text-align: justify; font-size: 0.9em;">[°C] Температуры хотэнда и стола при печати калибровочной модели<br>Если в принтере есть активная термокамера, то её температура будет установлена равной 0°C</td>
        </tr>
        <!-- Скорость вентилятора -->
        <tr>
            <th class="lang" id="table.fan_speed.title">Скорость вентилятора</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_cooling" name="k3d_la_cooling" value="100"></td>
            <td class="lang" id="table.fan_speed.description" style="text-align: justify; font-size: 0.9em;">[%] Обороты вентилятора в процентах<br>Для предотвращения ошибок, связанных с охлаждением хотэнда, указанная скорость вентилятора будет достигнута только к 4 слою</td>
        </tr>
        <!-- Поток -->
        <tr>
            <th class="lang" id="table.flow.title">Поток</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_flow" name="k3d_la_flow" value="100"></td>
            <td class="lang" id="table.flow.description" style="text-align: justify; font-size: 0.9em;">[%] Поток в процентах. Нужен для компенсации общей пере- или недоэкструзии</td>
        </tr>
    <!-- Параметры калибровки -->
        <!-- Скорость движения печатающей головы -->
        <tr>
            <th class="lang" id="table.speed.title">Скорость движения<br>печатающей головы</td>
            <td style="text-align:center">
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">Slow</td>
                            <td class="innerTableTD"  style="text-align:center">Fast</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_minPrintSpeed" name="k3d_la_minPrintSpeed" value="20"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_maxPrintSpeed" name="k3d_la_maxPrintSpeed" value="100"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.speed.description" style="text-align: justify; font-size: 0.9em;">[мм/с] Скорость печати медленных (slow) и быстрых (fast) участков при проверке значения PA<br>
            Также, скорость холостых перемещений будет равна скорости печати быстрых участков, а скорость печати первого слоя - скорости печати медленных участков<br>
            Если выбрана цель "расход", то скорость быстрых участков будет переопределена, чтобы достичь указанного расхода</td>
        </tr>
        <!-- Ускорение -->
        <tr>
            <th class="lang" id="table.acceleration.title">Ускорение</th>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_acceleration" name="k3d_la_acceleration" value="5000"></td>
            <td class="lang" id="table.acceleration.description" style="text-align:justify; font-size: 0.9em;">[мм/с^2] Ускорение, с которым будет напечатана тестовая модель<br>В общем случае стоит поставить равной максимальному ускорению вашего принтера, полученному при калибровке input shaping'а или по критерию отстуствия эхо после углов<br>При калибровке адаптивного PA рекомендуется провести 3 калибровки: при ускорении 1000, половине от максимального ускорения и при максимальном ускорении<br>Не рекомендуется ставить значения ниже 1000 т.к. тогда методика может отработать некорректно</td>
        </tr>
        <tr>
            <th class="lang" id="table.num_perimeters.title">Количество<br>периметров</th>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_numPerimeters" name="k3d_la_numPerimeters" value="2"></td>
            <td class="lang" id="table.num_perimeters.description" style="text-align:justify; font-size: 0.9em;">Сколько периметров будет у основно части модели<br>В общем случае хорошо работает 2. Для определения перепада толщины на просвет поставьте 1. Для эластомеров можно попробовать 3-4, если модель ведёт при 2</td>
        </tr>
        <!-- Коэффициенты LA/PA -->
        <tr>
            <th class="lang" id="table.la_values.title">Коэффиценты LA/PA</th>
            <td>
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">Start</td>
                            <td class="innerTableTD"  style="text-align:center">End</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_initKFactor" name="k3d_la_initKFactor" value="0.0"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_endKFactor" name="k3d_la_endKFactor" value="0.1"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.la_values.description" style="text-align: justify; font-size: 0.9em;">Начальное и конечное значение коэффициентов LA/PA для проверки<br>
            На принтерах с директ экструдерами, при калибровке PA для твёрдого материала, стоит установить 0.0 - 0.1. Если этого диапазона не хватит, то проверить 0.0 - 0.2<br>
            При подборе PA для эластомеров, стоит увеличить конечное значение до 0.5-1.0<br>
            На принтерах с боуден экструдерами, начинать калибровку лучше с диапазона 0.0-1.0</td>
        </tr>
        <!-- Количество сегментов -->
        <tr>
            <th class="lang" id="table.num_segments.title">Количество сегментов</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_numSegments" name="k3d_la_numSegments" value="11"></td>
            <td class="lang" id="table.num_segments.description" style="text-align: justify; font-size: 0.9em;">Количество сегментов башенки. В течение сегмента коэффициент LA/PA остаётся неизменным. Сегменты визуально
                разделены для упрощения анализа модели<br>
                В общем случае рекомендуется установить 11 для быстрой проверки, или 21 для проверки с более широким диапазоном значений
            </td>
        </tr>
        <!-- Время сглаживания -->
        <tr>
            <th class="lang" id="table.smooth_time.title">Время сглаживания</th>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_la_smoothTime" name="k3d_la_smoothTime" value="0.04"></td>
            <td class="lang" id="table.smooth_time.description" style="text-align:justify; font-size: 0.9em;">[с] Время сглаживания PA<br>
                В прошивке Klipper уменьшает нагрузку на подающий механизм, но завышенное значение сглаживания может вызвать дефекты в виде ямок перед и после шва. Поэтому это значение стоит откалибровать с помощью установки дополнительной цели "время сглаживания"<br>
                В Marlin и RRF нет управления временем сглаживания PA, поэтому там этот параметр ничего не делает<br>
                При выборе доп. цели калибровки "время сглаживания", значение этого параметра будет проигнорировано
            </td>
        </tr>
        <!-- Дополнительная цель калибровки -->
        <tr>
            <th class="lang" id="table.additional_calibration_target.title">Дополнительная цель<br>калибровки</td>
            <td align="center">
                <form style="text-align:left; width:fit-content;"><input type="radio" id="k3d_la_targetNone" name="k3d_la_additionalCalibrationTarget" value="None" checked><label for="k3d_la_targetNone"> Нет</label><br>
                <input type="radio" id="k3d_la_targetSmoothTime" name="k3d_la_additionalCalibrationTarget" value="SmoothTime"><label for="k3d_la_targetSmoothTime"> Smooth time</label><br>
                <input type="radio" id="k3d_la_targetAcceleration" name="k3d_la_additionalCalibrationTarget" value="Acceleration"><label for="k3d_la_targetAcceleration"> Acceleration</label><br>
                <input type="radio" id="k3d_la_targetFlowrate" name="k3d_la_additionalCalibrationTarget" value="Flowrate"><label for="k3d_la_targetFlowrate"> Flowrate</label>
                </form>
            </td>
            <td class="lang" id="table.additional_calibration_target.description" style="text-align: justify; font-size: 0.9em;">При всех доп. целях, коэффициент PA меняется по высоте. А значение доп. параметра меняется от минимального на задней   стенке модели, до максимального на левой стенке (направление печати по часовой стрелке)<br>
            Нет - никакого доп. параметра меняться не будет, все стенки одинаковые<br>
            Время сглаживания - между стенками будет меняться время сглаживания PA. Работает только для прошивки Klipper<br>
            Ускорение - между стенками будет меняться ускорение печати<br>
            Расход - модель будет сгенерирована с увеличенными толщиной слой и шириной линии. Скорость быстрых участков будет меняться так, чтобы расход на них соответствовал значению доп. параметра            
            </td>
        </tr>
        <!-- Значения дополнительного параметра -->
        <tr id="table_additional_parameter_row">
            <th class="lang" id="table.additional_parameter.title">Значения доп.<br>параметра</th>
            <td>
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">Start</td>
                            <td class="innerTableTD"  style="text-align:center">End</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_initAdditionalParameter" name="k3d_la_initAdditionalParameter" value="0.02"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_la_endAdditionalParameter" name="k3d_la_endAdditionalParameter" value="0.05"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.additional_parameter.description" style="text-align: justify; font-size: 0.9em;">Минимальное и максимальное значение параметра дополнительной цели калибровки<br>Если дополнительная цель выключена, то не делают ничего</td>
        </tr>
    </tbody>
</table>

<table class="calibrationCodeTable">
    <tbody>
        <tr>
            <th class="lang" id="table.start_gcode.title">Начальный G-код</th>
            <td rowspan=5 style="text-align: left; font-size: 0.9em; padding-right: 1.0em; padding-top: 1.0em; vertical-align: middle;" class="lang" id="table.gcodes.description">Для того, чтобы печать была правильно инициализирована и завершена, необходимо указать корректные для вашего принтера G-коды. Проще всего скопировать их из вашего слайсера. Если G-коды в слайсере не содержат неподдерживаемых элементов, то их можно копировать без изменений. Если неподдерживаемые элементы есть, то придётся от них избавиться - вручную заменить неподдерживаемые плейсхолдеры на значения, посчитать результаты математических функций, раскрыть условия и т.д.<hr><u>Поддерживаемые плейсхолдеры:</u><br>- <code>$HOTTEMP</code>, <code>{temperature}</code>, <code>{first_layer_temperature}</code> будут заменены на температуру хотэнда;<br>- <code>$BEDTEMP</code>, <code>{bed_temperature}</code>, <code>{first_layer_bed_temperature}</code> будут заменены на температуру стола;<br>- <code>{chamber_temperature}</code>, <code>{chamber_minimal_temperature}</code> будут заменены на 0;- <code>$FLOW</code> будет заменён на значение потока<br>- Разные вариации указанных выше плейсхолдеров. Например, в квадратных скобках вместо фигурных, с указанием номера элемента массива, с арифметикой внутри фигурных скобок и т.д.<hr><u>Не поддерживаются:</u><br>- Другие плейсхолдеры;<br>- Условные операторы (<code>if</code>);<br>- Тернарные операторы (<code>a ? b : c</code>);<br>- Математические функции и выражения</td>
        </tr>
        <tr>
            <td width="50%"><textarea type="text" id="k3d_la_startGcode" name="k3d_la_startGcode" rows="15">
M104 S150 ;прогреть хотэнд до 150 градусов
M190 S$BEDTEMP ;прогреть стол до температуры, указанной в настройках
M109 S$HOTTEMP ;прогреть хотэнд до температуры, указанной в настройках
G28 ;припарковать все оси
$G29 ;снять карту высот стола
M109 S$HOTTEMP ;обход дебильных макросов Creality
M190 S$BEDTEMP ;обход дебильных макросов Creality
G90 ;абсолютная система координат
G92 E0 ;сбросить координату экструдера
M220 S100 ;Множитель скорости 100%
M221 S$FLOW ;Множитель потока взять из настроек</textarea></td>
        </tr>
        <tr>
            <th class="lang" id="table.end_gcode.title" style="padding-top: 0.5em;">Конечный G-код</th>
        </tr>
        <tr>
            <td width="50%"><textarea type="text" id="k3d_la_endGcode" name="k3d_la_endGcode" rows="15">
M104 S0 ;выключить хотэнд
M140 S0 ;выключить нагрев стола
M106 S0 ;выключить вентилятор модели
G91 ;относительная система координат
G1 E-5 F600 ;сделать откат на 5мм
G1 Z1 F300 ;поднять голову на 1мм</textarea></td>
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