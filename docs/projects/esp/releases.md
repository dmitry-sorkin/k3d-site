---
icon: material/download
title: K3D ESP - релизы
description: Проект печатающей головы для Ender-3 на основе экструдера Sprite
---

# Релизы K3D ESP

## Последние версии файлов

!!! warning "Для Ender-3S1 и Sprite Pro Kit (с переходной платой и кареткой) печатать кронштейн не надо!"

- Крепление Sprite к каретке Ender-3:
    - [:material-printer-3d-nozzle: Под 3010 радиальный вентилятор](./releases/v1/k3d_sprite_mount_(ender-3%2C_3010_fan).stl){ download="k3d_sprite_mount_(ender-3%2C_3010_fan).stl" }
    - [:material-printer-3d-nozzle: Под 4010 радиальный вентилятор](./releases/v2/k3d_sprite_ender-3_mount_(4010_fan).stl){ download="k3d_sprite_ender-3_mount_(4010_fan).stl" }
- Система охлаждения (подходит к Ender-3 S1/S1 Pro/S1 Plus):
    - [:material-printer-3d-nozzle: Для стокового хотэнда](./releases/v2/k3d_sprite_dual_5015_fan_duct_(stock_hotend).stl){ download="k3d_sprite_dual_5015_fan_duct_(stock_hotend).stl" }
    - [:material-printer-3d-nozzle: Для Volcano хотэнда](./releases/v2/k3d_sprite_dual_5015_fan_duct_(volcano_hotend).stl){ download="k3d_sprite_dual_5015_fan_duct_(volcano_hotend).stl" }
- Крепление BLTouch/CRTouch:
    - [:material-printer-3d-nozzle: Для стокового хотэнда](./releases/v2/k3d_sprite_crtouch_mount_(stock_hotend).stl){ download="k3d_sprite_crtouch_mount_(stock_hotend).stl" }
    - [:material-printer-3d-nozzle: Для Volcano хотэнда](./releases/v2/k3d_sprite_crtouch_mount_(volcano_hotend).stl){ download="k3d_sprite_crtouch_mount_(volcano_hotend).stl" }
- [:material-printer-3d-nozzle: Проставка CRTouch под хотэнд Volcano для E3S1](./releases/v2/e3s1_crtouch_volcano_spacer.stl){ download="e3s1_crtouch_volcano_spacer.stl" }
- [:material-video-3d: Сборка в формате .STEP](./releases/v2/k3d_esp_v2.1.stp){ download="k3d_esp_v2.1.stp" }

## V2.1 - Исправления крепления экструдера под 4010 вентилятор

Про разработке кронштейна экструдера под 4010 вентилятор была использована неправильная модель вентлятора, что привело к ошибке в форме воздуховода. Кроме этого, выяснилось, что у разных 4010 вентиляторов очень разная форма крепежных ушек. Поэтому кронштейн был перемоделирован так, чтобы подходить ко всем вариантам вентляторов

- [:material-printer-3d-nozzle: Крепление Sprite к каретке Ender-3 под 4010 вентилятор](./releases/v2/k3d_sprite_ender-3_mount_(4010_fan).stl){ download="k3d_sprite_ender-3_mount_(4010_fan).stl" }
- [:material-video-3d: Сборка в формате .STEP](./releases/v2/k3d_esp_v2.1.stp){ download="k3d_esp_v2.1.stp" }

## V2 - Полная совместимость с Volcano хотэндом

В предыдущей версии Volcano хотэнд нельзя было установить Volcano хотэнд с утолщенным носком от 5А (с обычным лезло). Обдув был сильно переработан и теперь не должен конфликтовать ни с какими хотэндами. Из-за изменения ширины обдува, пришлось сместить крепление датчика CRTouch. Крепления не изменились. 

- Система охлаждения (подходит к Ender-3 S1/S1 Pro/S1 Plus):
    - [:material-printer-3d-nozzle: Для стокового хотэнда](./releases/v2/k3d_sprite_dual_5015_fan_duct_(stock_hotend).stl){ download="k3d_sprite_dual_5015_fan_duct_(stock_hotend).stl" }
    - [:material-printer-3d-nozzle: Для Volcano хотэнда](./releases/v2/k3d_sprite_dual_5015_fan_duct_(volcano_hotend).stl){ download="k3d_sprite_dual_5015_fan_duct_(volcano_hotend).stl" }
- Крепление BLTouch/CRTouch:
    - [:material-printer-3d-nozzle: Для стокового хотэнда](./releases/v2/k3d_sprite_crtouch_mount_(stock_hotend).stl){ download="k3d_sprite_crtouch_mount_(stock_hotend).stl" }
    - [:material-printer-3d-nozzle: Для Volcano хотэнда](./releases/v2/k3d_sprite_crtouch_mount_(volcano_hotend).stl){ download="k3d_sprite_crtouch_mount_(volcano_hotend).stl" }
- [:material-video-3d: Сборка в формате .STEP](./releases/v2/k3d_esp_v2.stp){ download="k3d_esp_v2.stp" }

## V1 - первоначальный релиз

- Крепление Sprite к каретке Ender-3:
    - [:material-printer-3d-nozzle: Под 3010 радиальный вентилятор](./releases/v1/k3d_sprite_mount_(ender-3%2C_3010_fan).stl){ download="k3d_sprite_mount_(ender-3%2C_3010_fan).stl" }
    - [:material-printer-3d-nozzle: Под 4010 радиальный вентилятор](./releases/v1/k3d_sprite_mount_(ender-3%2C_4010_fan).stl){ download="k3d_sprite_mount_(ender-3%2C_4010_fan).stl" }
- Система охлаждения (подходит к Ender-3 S1/S1 Pro/S1 Plus):
    - [:material-printer-3d-nozzle: Для стокового хотэнда](./releases/v1/k3d_sprite_dual_5015_fan_duct_(stock_hotend).stl){ download="k3d_sprite_dual_5015_fan_duct_(stock_hotend).stl" }
    - [:material-printer-3d-nozzle: Для Volcano хотэнда](./releases/v1//k3d_sprite_dual_5015_fan_duct_(volcano_hotend).stl){ download="k3d_sprite_dual_5015_fan_duct_(volcano_hotend).stl" }
- Крепление BLTouch/CRTouch:
    - [:material-printer-3d-nozzle: Для стокового хотэнда](./releases/v1/k3d_sprite_bltouch_mount_(stock_hotend).stl){ download="k3d_sprite_bltouch_mount_(stock_hotend).stl" }
    - [:material-printer-3d-nozzle: Для Volcano хотэнда](./releases/v1/k3d_sprite_bltouch_mount_(volcano_hotend).stl){ download="k3d_sprite_bltouch_mount_(volcano_hotend).stl" }
- [:material-video-3d: Сборка в формате .STEP](./releases/v1/k3d_ender-3_sprite_printhead.stp){ download="k3d_ender-3_sprite_printhead.stp" }
- 