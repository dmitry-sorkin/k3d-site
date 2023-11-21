# Релизы K3D Тест точности

## Последние версии файлов

- [:material-file-document: Таблица v2.0](./releases/calcs/k3d_accuracy_calibration_v2.0.ods){ download="k3d_accuracy_calibration_v2.0.ods" }
- [:material-video-3d: Модель v7 STEP](./releases/models/k3d_accuracy_test_v7.stp){ download="k3d_accuracy_test_v7.stp" }
- [:material-video-3d: Модель v7 STL](./releases/models/k3d_accuracy_test_v7.stl){ download="k3d_accuracy_test_v7.stl" }

## Релизы

### Таблица v2.0

[:material-file-document: Таблица v2.0](./releases/calcs/k3d_accuracy_calibration_v2.0.ods){ download="k3d_accuracy_calibration_v2.0.ods" }

Аналогична v2.0-alpha2, за исключением изначально введенных размеров. Смена версии связана с тем, что прошло большое время, таблица протестирована и больше не является альфа версией.

### Таблица v2.0-alpha2

[:material-file-document: Таблица v2.0-alpha2](./releases/calcs/k3d_accuracy_calibration_v2.0-alpha2.ods){ download="k3d_accuracy_calibration_v2.0-alpha2.ods" }

По сравнению с предыдущими версиями:

- Из ввода и расчёта убран поток;
- Расчёт полностью переделан. По предварительным тестам, новая версия работает лучше;
- Ввод замеров проверочной модели перенесен на лист "ввод";

### Модель v7

- [:material-video-3d: Модель v7 STEP](./releases/models/k3d_accuracy_test_v7.stp){ download="k3d_accuracy_test_v7.stp" }
- [:material-video-3d: Модель v7 STL](./releases/models/k3d_accuracy_test_v7.stl){ download="k3d_accuracy_test_v7.stl" }

По сравнению с предыдущей версией:

- Фаска на гранях изменена так, чтобы уменьшить требования к обдуву;
- Надписи сделаны крупнее и насквозь модели, чтобы их было видно с двух сторон.

### Модель v6

[:material-video-3d: Модель v7](./releases/models/k3d_accuracy_test_v6.stp){ download="k3d_accuracy_test_v6.stp" }

По сравнению с предыдущей версией:

- Добавлена "ловушка" для шва, чтобы уменьшить его влияние на снимаемые размеры;
- Добавлена фаска снизу и сверху, чтобы уменьшить влияние крышки и дна, а также слоновьей ноги на снимаемые размеры;
- Уменьшены скругления, чтобы даже при дефектах прохождения поворота, штангенциркуль при измерениях не попал рабочей поверхностью на скругление;
- Модель переведена в точный формат STEP.

### Таблица v1.1

[:material-file-document: Таблица v1.1](./releases/calcs/k3d_accuracy_calibration_v1.1.ods){ download="k3d_accuracy_calibration_v1.1.ods" }

По сравнению с v1.0 исправлены ошибки на листе "анализ".

### Таблица v1.0

Первоначальная версия таблицы из видео.

[:material-file-document: Таблица v1.0](./releases/calcs/k3d_accuracy_calibration_v1.0.ods){ download="k3d_accuracy_calibration_v1.0.ods" }

### Модель v5

Первоначальная версия модели из видео. Может применяться с более новыми версиями таблиц.

[:material-printer-3d-nozzle: Модель v5](releases/models/dimension_test_v5.stl){ download="dimension_test_v5.stl" }
