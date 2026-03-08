---
icon: octicons/package-16
title: K3D Калибровщик расхода
hide:
    - toc
---

<h1 class="lang" id="header.title">K3D калибровщик расхода</h1>

<script src="../assets/js/lib.js"></script>
<script src="../assets/js/profileManager.js"></script>
<script src="../assets/js/wasm_exec.js"></script>
<script src="../assets/js/gwaloader.js"></script>
<script src="../assets/js/streamsavermin.js"></script>

<span class="lang" id="lang.label">Язык: </span>[:flag_gb: English](?lang=en) [:flag_ru: Русский](?lang=ru) | <span class="lang" id="profile.select.label">Профиль: </span><select id="profileSelect" class="profileSelect"><option value="" class="lang" id="profile.default">По умолчанию</option></select> :material-content-save:{class="twemoji inlineButton" title="Save" onclick="saveCurrentProfile()"} :material-content-save-edit:{class="twemoji inlineButton" title="Save as" onclick="saveAsNewProfile()"} :material-delete:{class="twemoji inlineButton" id="deleteProfileBtn" title="Delete" onclick="deleteCurrentProfile()"} :material-export:{class="twemoji inlineButton" title="Export" onclick="exportProfile()"} :material-import:{class="twemoji inlineButton" title="Import" onclick="importProfile()"}

<!-- Модальное окно для сохранения профиля -->
<div id="profileModal" class="profileModal" style="display: none;">
    <div class="profileModalContent">
        <h3 class="lang" id="profile.modal.title">Сохранить профиль</h3>
        <input type="text" id="profileNameInput" class="profileNameInput" placeholder="Введите название профиля">
        <div class="profileModalButtons">
            <button onclick="confirmSaveProfile()" class="caliButton">
                <span class="lang" id="profile.modal.save">Сохранить</span>
            </button>
            <button onclick="closeProfileModal()" class="caliButton caliButtonDanger">
                <span class="lang" id="profile.modal.cancel">Отмена</span>
            </button>
        </div>
    </div>
</div>

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
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_bedX" value="235"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_bedY" value="235"></td>
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
                <form style="text-align:left; width:fit-content;">
                    <input type="radio" id="k3d_mfr_marlin" value="Marlin" checked><label for="k3d_mfr_marlin"> Marlin</label><br>
                    <input type="radio" id="k3d_mfr_klipper" value="Klipper"><label for="k3d_mfr_klipper"> Klipper</label><br>
                    <input type="radio" id="k3d_mfr_rrf" value="RRF"><label for="k3d_mfr_rrf"> RRF</label>
                </form>
            </td>
            <td class="lang" id="table.firmware.description" style="text-align: justify; font-size: 0.9em;">Прошивка, установленная на вашем принтере<br>Если не знаете, то, скорее всего, Klipper<br>Для принтеров Bambu Lab выберите Marlin</td>
        </tr>
        <!-- Начало координат в центре стола (режим дельты) -->
        <tr>
            <th class="lang" id="table.delta.title">Начало координат<br>в центре стола</td>
            <td style="text-align:center"><input type="checkbox" id="k3d_mfr_delta"></td>
            <td class="lang" id="table.delta.description" style="text-align: justify; font-size: 0.9em;">Для декартовых принтеров должно быть выключено<br>Для дельт и AD5M включено</td>
        </tr>
        <!-- Диаметр сопла -->
        <tr>
            <th class="lang" id="table.nozzle_diameter.title">Диаметр сопла</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_mfr_nozzleDiameter" value="0.4"></td>
            <td class="lang" id="table.nozzle_diameter.description" style="text-align: justify; font-size: 0.9em; font-size: 0.9em;">[мм] Диаметр сопла, установленного на вашем принтере. Влияет на ширину линии и толщину слоя при печати образцов и вспомогательных элементов</td>
        </tr>
        <!-- Z offset -->
        <tr>
            <th class="lang" id="table.z_offset.title">Z offset</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_mfr_zOffset" value="0.0"></td>
            <td class="lang" id="table.z_offset.description" style="text-align: justify; font-size: 0.9em;">[мм] Смещение всей модели по вертикали для компенсации неправильного расстояния от сопла до стола на первом слое. Положительное значение смещает модель вверх, отрицательное - вниз.<br>Если в вашем стартовом G-коде уже есть установка Z-offset'а, то оставьте 0</td>
        </tr>
    <!-- Параметры филамента -->
        <!-- Температуры -->
        <tr>
            <th class="lang" id="table.temperatures.title">Температура</td>
            <td style="text-align:center">
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">Hotend</td>
                            <td class="innerTableTD"  style="text-align:center">Bed</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_hotendTemperature" value="210"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_bedTemperature" value="60"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.temperatures.description" style="text-align: justify; font-size: 0.9em;">[°C] Температуры хотэнда и стола при печати калибровочной модели<br>Если в принтере есть активная термокамера, то её температура будет установлена равной 0°C</td>
        </tr>
    <!-- Параметры калибровки -->
        <!-- Скорости -->
        <tr>
            <th class="lang" id="table.speed.title">Скорость движения<br>печатающей головы</th>
            <td>
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">Print</td>
                            <td class="innerTableTD"  style="text-align:center">Travel</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_speed" value="20"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_travelSpeed" value="200"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.speed.description" style="text-align: justify; font-size: 0.9em;">[мм/с] Скорость движения печатающей головы при печати (print) и при холостых перемещениях (travel).<br>Для скорости печати рекомендуется ставить невысокие значения, чтобы разгон и торможение оказывали минимальное влияние на средний объёмный расход при печати образца</td>
        </tr>
        <!-- Проверяемый расход -->
        <tr>
            <th class="lang" id="table.flowrate.title">Объёмный расход</th>
            <td>
                <table class="innerTable">
                    <tbody>
                        <tr class="innerTableTR">
                            <td class="innerTableTD" style="text-align:center">Start</td>
                            <td class="innerTableTD"  style="text-align:center">End</td>
                        </tr>
                        <tr class="innerTableTR">
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_initFlowrate" value="5.0"></td>
                            <td class="innerTableTD"  style="text-align:center"><input class="calibratorInputSmall" type="text" id="k3d_mfr_endFlowrate" value="25.0"></td>
                        </tr>
                    </tbody>
                </table>
            </td>
            <td class="lang" id="table.flowrate.description" style="text-align: justify; font-size: 0.9em;">[мм^3/с] Начальное и конечное значение объёмного расхода для проверки<br>- Для V6-подобных хотэндов рекомендуется от 4 до 20<br>- Для Volcano-подобных рекомендуется от 4 до 30<br>- Для Goliath-подобных от 4 до 55<br>При наличии CHT вставки в хотэнде, стоит увеличить максимальное проверяемое значение</td>
        </tr>
        <!-- Шаг расхода -->
        <tr>
            <th class="lang" id="table.flowrate_delta.title">Шаг расхода</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_mfr_flowrateDelta" value="2"></td>
            <td class="lang" id="table.flowrate_delta.description" style="text-align: justify; font-size: 0.9em;">[мм^3/с] Разница объёмного расхода между соседними проверяемыми элементами. Чем меньше, тем больше образцов будет напечатано</td>
        </tr>
        <!-- Длина экструзии на образец -->
        <tr>
            <th class="lang" id="table.extrusion_length.title">Длина экструзии<br>на образец</td>
            <td style="text-align:center"><input class="calibratorInput" type="text" id="k3d_mfr_extrusionLength" value="200"></td>
            <td class="lang" id="table.extrusion_length.description" style="text-align: justify; font-size: 0.9em;">[мм] Сколько мм прутка будет экструдировано для печати каждого образца. Считается по входу в подающий механизм</td>
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
            <td width="50%"><textarea type="text" id="k3d_mfr_startGcode" name="k3d_mfr_startGcode" rows="15">
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
            <td width="50%"><textarea type="text" id="k3d_mfr_endGcode" name="k3d_mfr_endGcode" rows="15">
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