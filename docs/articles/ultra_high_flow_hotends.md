---
icon: material/printer-3d-nozzle
title: Сверхпроизводительные хотэнды
subtitle: Goliath, CHC XL, Rapido ACE
description: Goliath, CHC XL, Rapido ACE
---

# Сверхпроизводительные хотэнды

В этой статье представлены результаты тестирования Goliath, CHC XL и Rapido ACE. Причём как в обычном исполнении, так и с удлинителем 8.5мм.

!!! note "Ссылки на хотэнды есть в [3D рекомендаторе](../part-navi/extruder.md)"

## Тарирование термисторов

Так как хотэнды имеют разную конструкцию и разные датчики температуры, а тестирование стоит проводить при максимально одинаковых температурах, то необходимо тарировать термисторы. Для этого в каждый хотэнд помещалась термопара К-типа через канал для прохождения филамента до упора в сопло. После этого хотэнд грелся до 200, 210, 220 ... 320ºC. Показания с термистора и термопары нанесены на диаграмму ниже. Показания с термопары усреднялись за 1 минуту после стабилизации температуры по термистору.

Колебания температуры сопла на Goliath, CHC XL и Volcano были небольшие, температура держалась в пределах [-0.2;+0.2]ºC от заданной. На Rapido ACE в обоих вариантах температура сопла колебалась в пределах [-1;+1]ºC от указанной, а температура по термистору колебалась в пределах [-3;+3]ºC от целевой. Поэтому я протестировал хотэнд также и с другим термитором (ATC Semitec 104GT-2), ситуация с которым не поменялась. Возможности установить термистор на другое место на хотэнде не предусмотрено. Поэтому все дальнейшие тесты проводились при сильных колебаниях температуры на этом хотэнде.

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "[K3D] Hotend thermistor taring",
    "subtitle": "",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "description": "V6 0.4 copper nozzle",
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "hotend": "Goliath",
        "thermistor_temperature": 200,
        "nozzle_temperature": 11,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 210,
        "nozzle_temperature": 11,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 220,
        "nozzle_temperature": 12,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 230,
        "nozzle_temperature": 12.5,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 240,
        "nozzle_temperature": 13,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 250,
        "nozzle_temperature": 13.5,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 260,
        "nozzle_temperature": 14,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 270,
        "nozzle_temperature": 14.5,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 280,
        "nozzle_temperature": 15,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 290,
        "nozzle_temperature": 15,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 300,
        "nozzle_temperature": 15.2,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 310,
        "nozzle_temperature": 16,
        "theory": false
      },
      {
        "hotend": "Goliath",
        "thermistor_temperature": 320,
        "nozzle_temperature": 16,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 200,
        "nozzle_temperature": 10.3,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 210,
        "nozzle_temperature": 11,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 220,
        "nozzle_temperature": 11.5,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 230,
        "nozzle_temperature": 12.3,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 240,
        "nozzle_temperature": 12.7,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 250,
        "nozzle_temperature": 13.1,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 260,
        "nozzle_temperature": 13.5,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 270,
        "nozzle_temperature": 14,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 280,
        "nozzle_temperature": 14.4,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 290,
        "nozzle_temperature": 14.9,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 300,
        "nozzle_temperature": 15.3,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 310,
        "nozzle_temperature": 15.6,
        "theory": false
      },
      {
        "hotend": "CHC XL",
        "thermistor_temperature": 320,
        "nozzle_temperature": 16.4,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 200,
        "nozzle_temperature": 1.7,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 210,
        "nozzle_temperature": 1.5,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 220,
        "nozzle_temperature": 1.4,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 230,
        "nozzle_temperature": 1.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 240,
        "nozzle_temperature": 1.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 250,
        "nozzle_temperature": 0.9,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 260,
        "nozzle_temperature": 0.8,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 270,
        "nozzle_temperature": 0.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 280,
        "nozzle_temperature": -0.3,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 290,
        "nozzle_temperature": -0.9,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 300,
        "nozzle_temperature": -1.3,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 310,
        "nozzle_temperature": -2.1,
        "theory": false
      },
      {
        "hotend": "Rapido ACE HF",
        "thermistor_temperature": 320,
        "nozzle_temperature": -2.7,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 200,
        "nozzle_temperature": -0.5,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 210,
        "nozzle_temperature": -0.5,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 220,
        "nozzle_temperature": -0.7,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 230,
        "nozzle_temperature": -1.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 240,
        "nozzle_temperature": -1.8,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 250,
        "nozzle_temperature": -2.3,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 260,
        "nozzle_temperature": -3.1,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 270,
        "nozzle_temperature": -4.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 280,
        "nozzle_temperature": -5.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 290,
        "nozzle_temperature": -5.6,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 300,
        "nozzle_temperature": -6.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 310,
        "nozzle_temperature": -6.2,
        "theory": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "thermistor_temperature": 320,
        "nozzle_temperature": -6.9,
        "theory": false
      },
      {
        "hotend": "Theory",
        "thermistor_temperature": 0,
        "nozzle_temperature": 0,
        "theory": true
      },
      {
        "hotend": "Theory",
        "thermistor_temperature": 1000,
        "nozzle_temperature": 0,
        "theory": true
      }
    ]
  },
  "layer": [
    {
      "mark": {
        "type": "line",
        "point": {"size": 100},
        "clip": true,
        "strokeWidth": 3,
        "tooltip": true
      },
      "encoding": {
        "x": {
          "field": "thermistor_temperature",
          "type": "quantitative",
          "scale": {"domainMin": 200, "domainMax": 320},
          "axis": {
            "title": "Thermistor temperature [ºC]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "nozzle_temperature",
          "type": "quantitative",
          "scale": {"domainMin": -8, "domainMax": 18},
          "axis": {
            "title": "Nozzle temperature delta [ºC]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "hotend",
          "type": "nominal",
          "legend": {
            "orient": "top",
            "titleFontSize": 14,
            "labelFontSize": 14,
            "rowPadding": 5,
            "padding": 3,
            "columns": {
              "expr": "floor(width / 155) == 0 ? 1 : floor(width / 155)"
            }
          },
          "scale": {
            "domain": [
              "Goliath",
              "CHC XL",
              "Rapido ACE HF",
              "Rapido ACE UHF",
              "Volcano",
              "Theory"
            ],
            "range": [
              "#1F78B5",
              "#FF7F0D",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#7f7f7f"
            ]
          }
        },
        "opacity": {"condition": {"param": "hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "theory", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "hotend",
          "select": {"type": "point", "fields": ["hotend"]},
          "bind": "legend"
        }
      ]
    },
    {
      "mark": {
        "type": "text",
        "dx": 0,
        "dy": -30,
        "fontSize": 16,
        "align": "center",
        "clip": true
      },
      "encoding": {
        "x": {
          "field": "thermistor_temperature",
          "type": "quantitative",
          "scale": {"domainMin": 200, "domainMax": 340}
        },
        "y": {
          "field": "nozzle_temperature",
          "type": "quantitative",
          "scale": {"domainMin": 200, "domainMax": 320}
        },
        "text": {"field": "nozzle_temperature"},
        "opacity": {
          "condition": {"param": "hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "hotend"}
      }
    }
  ]
}
```

## Тест объёмного расхода

### Методика тестирования

Тесты проводились на K3D VOSTOK v9.0. Подающий механизм - Microfeeder на основе BMG подающих колёс и мотора Nema 14 20мм с заявленным крутящим моментом 100 мН*м. Охлаждение радиатора хотэнда вентилятором 40x40x20мм, охлаждение модели на основе ws7040. Сопло на всех хотэндах использовалось медное никелированное ø0.4 мм от Mellow. На Volcano для установки этого сопла использовать адаптер.

Тестирование максимального объёмного расхода производилось экструдированием 1000 мм филамента при мгновенном объёмном расходе 5, 10, 15 ... мм³/с. Далее полученные кусочки филамента взвешиваются и наносятся на диаграмму, на которой по горизонтали отмечены значения мгновенного объёмного расхода, а по вертикали - недоэкструзия в процентах от массы 1000мм прутка. К примеру, недоэкструзия в 3.7% означает, что масса экструдированного образца оказалась на 3.7% ниже, чем масса 1000 мм филамента с этой же катушки.

Также при тестировании сопел фиксировалось на каком объёмном расходе начинались пропуски подающего механизма. На диаграммах участки после последнего образца без пропусков отмечены пунктиром. 

!!! warning "Важно отметить, что все данные из этих тестов сильно зависят от подающего механизма, радиатора, термобарьера и даже от настроек драйвера. Поэтому результаты этих тестов напрямую можно сравнивать только между собой. При сравнении с другими тестами, необходимо учитывать разницу в тестовом стенде"

### PETG 235ºC

PETG Bestfilament. Температуры хотэндов по собственным термисторам:

| Хотэнд | Тип термистора | Температура по термистору | Примечание |
| :----- | :------------: | :-----------------------: | :--------- |
| Goliath | PT1000 | 223ºC |  |
| CHC XL | PT1000 | 223ºC |  |
| Rapido  ACE HF | PT1000* | 234ºC | В стоке на этом хотэнде стояла термопара К-типа, но у меня не было возмонжости её подключить. Поэтому я заменил её на PT1000, которую снял с CHC XL |
| Rapido ACE UHF | PT1000 | 237ºC |  |
| Volcano | ATC Semitec 104GT-2 | 230ºC |  |

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "[K3D] Ultra-high-flow hotends flowrate test",
    "subtitle": "PETG; 235ºC; K3D Microfeeder",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "description": "V6 0.4 copper nozzle",
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "hotend": "Goliath",
        "volumetric_flow": 5,
        "underextrusion": -1.12,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 10,
        "underextrusion": -2.26,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 15,
        "underextrusion": -3.39,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 20,
        "underextrusion": -4.43,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 25,
        "underextrusion": -5.02,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 30,
        "underextrusion": -6.19,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 35,
        "underextrusion": -7.16,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 40,
        "underextrusion": -8.46,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 40,
        "underextrusion": -8.46,
        "skipping": true
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 45,
        "underextrusion": -10.99,
        "skipping": true
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 50,
        "underextrusion": -100,
        "skipping": true
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 5,
        "underextrusion": -1.43,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 10,
        "underextrusion": -1.74,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 15,
        "underextrusion": -2.84,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 20,
        "underextrusion": -4.81,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 25,
        "underextrusion": -5.01,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 30,
        "underextrusion": -5.95,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 30,
        "underextrusion": -5.95,
        "skipping": true
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 35,
        "underextrusion": -7.92,
        "skipping": true
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 40,
        "underextrusion": -7.69,
        "skipping": true
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 45,
        "underextrusion": -11.86,
        "skipping": true
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 5,
        "underextrusion": -1.26,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 10,
        "underextrusion": -2.49,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 15,
        "underextrusion": -3.94,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 20,
        "underextrusion": -4.65,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 25,
        "underextrusion": -5.46,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 30,
        "underextrusion": -7.24,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 35,
        "underextrusion": -8.53,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 40,
        "underextrusion": -9.85,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 40,
        "underextrusion": -9.85,
        "skipping": true
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 45,
        "underextrusion": -12.28,
        "skipping": true
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 50,
        "underextrusion": -18.35,
        "skipping": true
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 5,
        "underextrusion": -0.81,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 10,
        "underextrusion": -1.78,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 15,
        "underextrusion": -3.3,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 20,
        "underextrusion": -4.26,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 25,
        "underextrusion": -5.4,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 30,
        "underextrusion": -6.53,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 35,
        "underextrusion": -6.72,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 40,
        "underextrusion": -7.4,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 40,
        "underextrusion": -7.4,
        "skipping": true
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 45,
        "underextrusion": -9.47,
        "skipping": true
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 50,
        "underextrusion": -17.67,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 5,
        "underextrusion": -1.23,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 10,
        "underextrusion": -2.0,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 15,
        "underextrusion": -3.36,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 20,
        "underextrusion": -5.2,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 25,
        "underextrusion": -7.11,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 25,
        "underextrusion": -7.11,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 30,
        "underextrusion": -10.34,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 35,
        "underextrusion": -21.36,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 5,
        "underextrusion": -1.78,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 10,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 15,
        "underextrusion": -4.14,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 20,
        "underextrusion": -5.46,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 25,
        "underextrusion": -6.91,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 30,
        "underextrusion": -8.76,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 35,
        "underextrusion": -10.69,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 35,
        "underextrusion": -10.69,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 40,
        "underextrusion": -22.04,
        "skipping": true
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 5,
        "underextrusion": -2.26,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 10,
        "underextrusion": -3.55,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 15,
        "underextrusion": -4.23,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 20,
        "underextrusion": -4.78,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 25,
        "underextrusion": -9.27,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 25,
        "underextrusion": -9.27,
        "skipping": true
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 30,
        "underextrusion": -25.2,
        "skipping": true
      }
    ]
  },
  "layer": [
    {
      "mark": {
        "type": "line",
        "point": {"size": 100},
        "clip": true,
        "strokeWidth": 3,
        "tooltip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0},
          "axis": {
            "title": "Underextrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "hotend",
          "type": "nominal",
          "legend": {
            "orient": "top",
            "titleFontSize": 14,
            "labelFontSize": 14,
            "rowPadding": 5,
            "padding": 3,
            "columns": {
              "expr": "floor(width / 155) == 0 ? 1 : floor(width / 155)"
            }
          },
          "scale": {
            "domain": [
              "Goliath",
              "Goliath + ext",
              "CHC XL",
              "CHC XL + ext",
              "Rapido ACE HF",
              "Rapido ACE UHF",
              "Volcano"
            ],
            "range": [
              "#1F78B5",
              "#AEC7E8",
              "#FF7F0D",
              "#FFBB78",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#ff9896",
              "#9467bd",
              "#ce6dbd",
              "#bcbd22"
            ]
          }
        },
        "opacity": {"condition": {"param": "hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "hotend",
          "select": {"type": "point", "fields": ["hotend"]},
          "bind": "legend"
        }
      ]
    },
    {
      "mark": {
        "type": "text",
        "dx": 5,
        "dy": -10,
        "fontSize": 16,
        "align": "left",
        "clip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "hotend"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

### PETG 248ºC

Точно такой же тест, но с завышенной на ~5% температурой. PETG Bestfilament. Температуры хотэндов по собственным термисторам:

| Хотэнд | Тип термистора | Температура по термистору | Примечание |
| :----- | :------------: | :-----------------------: | :--------- |
| Goliath | PT1000 | 235ºC |  |
| CHC XL | PT1000 | 235ºC |  |
| Rapido  ACE HF | PT1000* | 247ºC | В стоке на этом хотэнде стояла термопара К-типа, но у меня не было возмонжости её подключить. Поэтому я заменил её на PT1000, которую снял с CHC XL |
| Rapido ACE UHF | PT1000 | 250ºC |  |
| Volcano | ATC Semitec 104GT-2 | 242ºC |  |

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "[K3D] Ultra-high-flow hotends flowrate test",
    "subtitle": "PETG; 248ºC; K3D Microfeeder",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "description": "V6 0.4 copper nozzle",
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "hotend": "Goliath",
        "volumetric_flow": 5,
        "underextrusion": -1.22,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 10,
        "underextrusion": -1.32,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 15,
        "underextrusion": -2.1,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 20,
        "underextrusion": -3.2,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 25,
        "underextrusion": -4.14,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 30,
        "underextrusion": -4.89,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 35,
        "underextrusion": -5.57,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 40,
        "underextrusion": -6.87,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 45,
        "underextrusion": -7.62,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 50,
        "underextrusion": -9.69,
        "skipping": false
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 50,
        "underextrusion": -9.69,
        "skipping": true
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 55,
        "underextrusion": -12.03,
        "skipping": true
      },
      {
        "hotend": "Goliath",
        "volumetric_flow": 60,
        "underextrusion": -20.09,
        "skipping": true
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 5,
        "underextrusion": -0.19,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 10,
        "underextrusion": -1.2,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 15,
        "underextrusion": -1.94,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 20,
        "underextrusion": -2.88,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 25,
        "underextrusion": -4.04,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 30,
        "underextrusion": -4.85,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 35,
        "underextrusion": -5.91,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 40,
        "underextrusion": -6.3,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 45,
        "underextrusion": -6.75,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 50,
        "underextrusion": -7.79,
        "skipping": false
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 50,
        "underextrusion": -7.79,
        "skipping": true
      },
      {
        "hotend": "Goliath + ext",
        "volumetric_flow": 55,
        "underextrusion": -70.47,
        "skipping": true
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 5,
        "underextrusion": -1.55,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 10,
        "underextrusion": -2.42,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 15,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 20,
        "underextrusion": -3.65,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 25,
        "underextrusion": -4.49,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 30,
        "underextrusion": -5.62,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 35,
        "underextrusion": -7.01,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 40,
        "underextrusion": -7.92,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 45,
        "underextrusion": -9.14,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 50,
        "underextrusion": -10.92,
        "skipping": false
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 50,
        "underextrusion": -10.92,
        "skipping": true
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 55,
        "underextrusion": -15.35,
        "skipping": true
      },
      {
        "hotend": "CHC XL",
        "volumetric_flow": 60,
        "underextrusion": -22.13,
        "skipping": true
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 5,
        "underextrusion": -0.55,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 10,
        "underextrusion": -1.78,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 15,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 20,
        "underextrusion": -4.1,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 25,
        "underextrusion": -5.01,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 30,
        "underextrusion": -5.62,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 35,
        "underextrusion": -6.2,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 40,
        "underextrusion": -6.98,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 45,
        "underextrusion": -8.05,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 50,
        "underextrusion": -9.08,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 55,
        "underextrusion": -10.76,
        "skipping": false
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 55,
        "underextrusion": -10.76,
        "skipping": true
      },
      {
        "hotend": "CHC XL + ext",
        "volumetric_flow": 60,
        "underextrusion": -18.19,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 5,
        "underextrusion": -1.71,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 10,
        "underextrusion": -2.94,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 15,
        "underextrusion": -3.62,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 20,
        "underextrusion": -4.39,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 25,
        "underextrusion": -5.78,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 30,
        "underextrusion": -7.14,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 30,
        "underextrusion": -7.14,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 35,
        "underextrusion": -13.6,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE HF",
        "volumetric_flow": 40,
        "underextrusion": -24.78,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 5,
        "underextrusion": -0.32,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 10,
        "underextrusion": -2.0,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 15,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 20,
        "underextrusion": -3.81,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 25,
        "underextrusion": -5.04,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 30,
        "underextrusion": -6.43,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 35,
        "underextrusion": -7.88,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 40,
        "underextrusion": -11.02,
        "skipping": false
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 40,
        "underextrusion": -11.02,
        "skipping": true
      },
      {
        "hotend": "Rapido ACE UHF",
        "volumetric_flow": 45,
        "underextrusion": -20.32,
        "skipping": true
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 5,
        "underextrusion": -0.61,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 10,
        "underextrusion": -1.55,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 15,
        "underextrusion": -2.91,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 20,
        "underextrusion": -3.65,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 25,
        "underextrusion": -6.17,
        "skipping": false
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 25,
        "underextrusion": -6.17,
        "skipping": true
      },
      {
        "hotend": "Volcano",
        "volumetric_flow": 30,
        "underextrusion": -17.25,
        "skipping": true
      }
    ]
  },
  "layer": [
    {
      "mark": {
        "type": "line",
        "point": {"size": 100},
        "clip": true,
        "strokeWidth": 3,
        "tooltip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0},
          "axis": {
            "title": "Underextrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "hotend",
          "type": "nominal",
          "legend": {
            "orient": "top",
            "titleFontSize": 14,
            "labelFontSize": 14,
            "rowPadding": 5,
            "padding": 3,
            "columns": {
              "expr": "floor(width / 155) == 0 ? 1 : floor(width / 155)"
            }
          },
          "scale": {
            "domain": [
              "Goliath",
              "Goliath + ext",
              "CHC XL",
              "CHC XL + ext",
              "Rapido ACE HF",
              "Rapido ACE UHF",
              "Volcano"
            ],
            "range": [
              "#1F78B5",
              "#AEC7E8",
              "#FF7F0D",
              "#FFBB78",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#ff9896",
              "#9467bd",
              "#ce6dbd",
              "#bcbd22"
            ]
          }
        },
        "opacity": {"condition": {"param": "hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "hotend",
          "select": {"type": "point", "fields": ["hotend"]},
          "bind": "legend"
        }
      ]
    },
    {
      "mark": {
        "type": "text",
        "dx": 5,
        "dy": -10,
        "fontSize": 16,
        "align": "left",
        "clip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15.0, "domainMax": 0.0}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "hotend"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

### По хотэндам

#### Goliath

Bestfilament PETG и Anycubic PLA, протестированы на типичных для печати температурах + с перегревом на ~5%.

eSun ePLA HF протестирован на 300ºC по термопаре (285ºC по термистору хотэнда) в поиске того, как некоторые "тестировщики" получают трёхзначные цифры расхода. На такой температуре PLA HF становится настолько жидким, что протестировать его на расходах до 20 мм^3/с не получилось. Пропуски на расходах выше 60 мм^3/с, скорее всего, связаны с тем, что подающий механизм на основе обычного Nema14 мотора и BMG подающих колёс при такой скорости вращения значимо теряет крутящий момент.

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "[K3D] Goliath flowrate test",
    "subtitle": "PETG; K3D Microfeeder",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "description": "V6 0.4 copper nozzle",
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 5,
        "underextrusion": -1.12,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 10,
        "underextrusion": -2.26,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 15,
        "underextrusion": -3.39,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 20,
        "underextrusion": -4.43,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 25,
        "underextrusion": -5.02,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 30,
        "underextrusion": -6.19,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 35,
        "underextrusion": -7.16,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 40,
        "underextrusion": -8.46,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 40,
        "underextrusion": -8.46,
        "skipping": true
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 45,
        "underextrusion": -10.99,
        "skipping": true
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 50,
        "underextrusion": -100,
        "skipping": true
      },{
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 5,
        "underextrusion": -1.22,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 10,
        "underextrusion": -1.32,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 15,
        "underextrusion": -2.1,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 20,
        "underextrusion": -3.2,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 25,
        "underextrusion": -4.14,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 30,
        "underextrusion": -4.89,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 35,
        "underextrusion": -5.57,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 40,
        "underextrusion": -6.87,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 45,
        "underextrusion": -7.62,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 50,
        "underextrusion": -9.69,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 50,
        "underextrusion": -9.69,
        "skipping": true
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 55,
        "underextrusion": -12.03,
        "skipping": true
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 60,
        "underextrusion": -20.09,
        "skipping": true
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 5,
        "underextrusion": -0.61,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 10,
        "underextrusion": -1.93,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 15,
        "underextrusion": -2.41,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 20,
        "underextrusion": -3.42,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 25,
        "underextrusion": -4.34,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 30,
        "underextrusion": -5.59,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 35,
        "underextrusion": -6.5,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 40,
        "underextrusion": -8.43,
        "skipping": false
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 40,
        "underextrusion": -8.43,
        "skipping": true
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 45,
        "underextrusion": -14.52,
        "skipping": true
      },
      {
        "hotend": "PLA 210 (200)",
        "volumetric_flow": 50,
        "underextrusion": -100,
        "skipping": true
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 5,
        "underextrusion": -1.09,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 10,
        "underextrusion": -1.7,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 15,
        "underextrusion": -2.17,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 20,
        "underextrusion": -2.48,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 25,
        "underextrusion": -3.02,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 30,
        "underextrusion": -3.69,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 35,
        "underextrusion": -4.61,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 40,
        "underextrusion": -5.11,
        "skipping": false
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 40,
        "underextrusion": -5.11,
        "skipping": true
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 45,
        "underextrusion": -8.26,
        "skipping": true
      },
      {
        "hotend": "PLA 221 (210)",
        "volumetric_flow": 50,
        "underextrusion": -100,
        "skipping": true
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 20,
        "underextrusion": -1.32,
        "skipping": false
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 30,
        "underextrusion": -1.29,
        "skipping": false
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 40,
        "underextrusion": -2.03,
        "skipping": false
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 50,
        "underextrusion": -3.35,
        "skipping": false
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 60,
        "underextrusion": -4.4,
        "skipping": false
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 60,
        "underextrusion": -4.4,
        "skipping": true
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 70,
        "underextrusion": -6.77,
        "skipping": true
      },
      {
        "hotend": "PLA HF 300 (285)",
        "volumetric_flow": 80,
        "underextrusion": -100,
        "skipping": true
      }
    ]
  },
  "layer": [
    {
      "mark": {
        "type": "line",
        "point": {"size": 100},
        "clip": true,
        "strokeWidth": 3,
        "tooltip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 80},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0},
          "axis": {
            "title": "Underextrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "hotend",
          "type": "nominal",
          "legend": {
            "orient": "top",
            "titleFontSize": 14,
            "labelFontSize": 14,
            "rowPadding": 5,
            "padding": 3,
            "columns": {
              "expr": "floor(width / 155) == 0 ? 1 : floor(width / 155)"
            }
          },
          "scale": {
            "domain": [
              "PETG 235 (223)",
              "PETG 248 (235)",
              "PLA 210 (200)",
              "PLA 221 (210)",
              "PLA HF 300 (285)"
            ],
            "range": [
              "#1F78B5",
              "#AEC7E8",
              "#FF7F0D",
              "#FFBB78",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#ff9896",
              "#9467bd",
              "#ce6dbd",
              "#bcbd22"
            ]
          }
        },
        "opacity": {"condition": {"param": "hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "hotend",
          "select": {"type": "point", "fields": ["hotend"]},
          "bind": "legend"
        }
      ]
    },
    {
      "mark": {
        "type": "text",
        "dx": 5,
        "dy": -10,
        "fontSize": 16,
        "align": "left",
        "clip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 80}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "hotend"}
      }
    }
  ]
}
```

#### CHC XL

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "[K3D] Goliath flowrate test",
    "subtitle": "PETG; K3D Microfeeder",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "description": "V6 0.4 copper nozzle",
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 5,
        "underextrusion": -1.26,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 10,
        "underextrusion": -2.49,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 15,
        "underextrusion": -3.94,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 20,
        "underextrusion": -4.65,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 25,
        "underextrusion": -5.46,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 30,
        "underextrusion": -7.24,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 35,
        "underextrusion": -8.53,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 40,
        "underextrusion": -9.85,
        "skipping": false
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 40,
        "underextrusion": -9.85,
        "skipping": true
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 45,
        "underextrusion": -12.28,
        "skipping": true
      },
      {
        "hotend": "PETG 235 (223)",
        "volumetric_flow": 50,
        "underextrusion": -18.35,
        "skipping": true
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 5,
        "underextrusion": -0.81,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 10,
        "underextrusion": -1.78,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 15,
        "underextrusion": -3.3,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 20,
        "underextrusion": -4.26,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 25,
        "underextrusion": -5.4,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 30,
        "underextrusion": -6.53,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 35,
        "underextrusion": -6.72,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 40,
        "underextrusion": -7.4,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 40,
        "underextrusion": -7.4,
        "skipping": true
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 45,
        "underextrusion": -9.47,
        "skipping": true
      },
      {
        "hotend": "+ ext PETG 235 (223)",
        "volumetric_flow": 50,
        "underextrusion": -17.67,
        "skipping": true
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 5,
        "underextrusion": -1.55,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 10,
        "underextrusion": -2.42,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 15,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 20,
        "underextrusion": -3.65,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 25,
        "underextrusion": -4.49,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 30,
        "underextrusion": -5.62,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 35,
        "underextrusion": -7.01,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 40,
        "underextrusion": -7.92,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 45,
        "underextrusion": -9.14,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 50,
        "underextrusion": -10.92,
        "skipping": false
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 50,
        "underextrusion": -10.92,
        "skipping": true
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 55,
        "underextrusion": -15.35,
        "skipping": true
      },
      {
        "hotend": "PETG 248 (235)",
        "volumetric_flow": 60,
        "underextrusion": -22.13,
        "skipping": true
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 5,
        "underextrusion": -0.55,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 10,
        "underextrusion": -1.78,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 15,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 20,
        "underextrusion": -4.1,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 25,
        "underextrusion": -5.01,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 30,
        "underextrusion": -5.62,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 35,
        "underextrusion": -6.2,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 40,
        "underextrusion": -6.98,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 45,
        "underextrusion": -8.05,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 50,
        "underextrusion": -9.08,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 55,
        "underextrusion": -10.76,
        "skipping": false
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 55,
        "underextrusion": -10.76,
        "skipping": true
      },
      {
        "hotend": "+ ext PETG 248 (235)",
        "volumetric_flow": 60,
        "underextrusion": -18.19,
        "skipping": true
      }
    ]
  },
  "layer": [
    {
      "mark": {
        "type": "line",
        "point": {"size": 100},
        "clip": true,
        "strokeWidth": 3,
        "tooltip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0},
          "axis": {
            "title": "Underextrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "hotend",
          "type": "nominal",
          "legend": {
            "orient": "top",
            "titleFontSize": 14,
            "labelFontSize": 14,
            "rowPadding": 5,
            "padding": 3,
            "columns": {
              "expr": "floor(width / 155) == 0 ? 1 : floor(width / 155)"
            }
          },
          "scale": {
            "domain": [
              "PETG 235 (223)",
              "+ ext PETG 235 (223)",
              "PETG 248 (235)",
              "+ ext PETG 248 (235)"
            ],
            "range": [
              "#1F78B5",
              "#AEC7E8",
              "#FF7F0D",
              "#FFBB78",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#ff9896",
              "#9467bd",
              "#ce6dbd",
              "#bcbd22"
            ]
          }
        },
        "opacity": {"condition": {"param": "hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "hotend",
          "select": {"type": "point", "fields": ["hotend"]},
          "bind": "legend"
        }
      ]
    },
    {
      "mark": {
        "type": "text",
        "dx": 5,
        "dy": -10,
        "fontSize": 16,
        "align": "left",
        "clip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "hotend"}
      }
    }
  ]
}
```

#### Rapido ACE

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "[K3D] Goliath flowrate test",
    "subtitle": "PETG; K3D Microfeeder",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "description": "V6 0.4 copper nozzle",
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 5,
        "underextrusion": -1.23,
        "skipping": false
      },
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 10,
        "underextrusion": -2.0,
        "skipping": false
      },
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 15,
        "underextrusion": -3.36,
        "skipping": false
      },
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 20,
        "underextrusion": -5.2,
        "skipping": false
      },
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 25,
        "underextrusion": -7.11,
        "skipping": false
      },
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 25,
        "underextrusion": -7.11,
        "skipping": true
      },
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 30,
        "underextrusion": -10.34,
        "skipping": true
      },
      {
        "hotend": "HF PETG 235 (234)",
        "volumetric_flow": 35,
        "underextrusion": -21.36,
        "skipping": true
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 5,
        "underextrusion": -1.78,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 10,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 15,
        "underextrusion": -4.14,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 20,
        "underextrusion": -5.46,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 25,
        "underextrusion": -6.91,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 30,
        "underextrusion": -8.76,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 35,
        "underextrusion": -10.69,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 35,
        "underextrusion": -10.69,
        "skipping": true
      },
      {
        "hotend": "UHF PETG 235 (237)",
        "volumetric_flow": 40,
        "underextrusion": -22.04,
        "skipping": true
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 5,
        "underextrusion": -1.71,
        "skipping": false
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 10,
        "underextrusion": -2.94,
        "skipping": false
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 15,
        "underextrusion": -3.62,
        "skipping": false
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 20,
        "underextrusion": -4.39,
        "skipping": false
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 25,
        "underextrusion": -5.78,
        "skipping": false
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 30,
        "underextrusion": -7.14,
        "skipping": false
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 30,
        "underextrusion": -7.14,
        "skipping": true
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 35,
        "underextrusion": -13.6,
        "skipping": true
      },
      {
        "hotend": "HF PETG 248 (247)",
        "volumetric_flow": 40,
        "underextrusion": -24.78,
        "skipping": true
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 5,
        "underextrusion": -0.32,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 10,
        "underextrusion": -2.0,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 15,
        "underextrusion": -3.07,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 20,
        "underextrusion": -3.81,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 25,
        "underextrusion": -5.04,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 30,
        "underextrusion": -6.43,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 35,
        "underextrusion": -7.88,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 40,
        "underextrusion": -11.02,
        "skipping": false
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 40,
        "underextrusion": -11.02,
        "skipping": true
      },
      {
        "hotend": "UHF PETG 248 (250)",
        "volumetric_flow": 45,
        "underextrusion": -20.32,
        "skipping": true
      }
    ]
  },
  "layer": [
    {
      "mark": {
        "type": "line",
        "point": {"size": 100},
        "clip": true,
        "strokeWidth": 3,
        "tooltip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0},
          "axis": {
            "title": "Underextrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "hotend",
          "type": "nominal",
          "legend": {
            "orient": "top",
            "titleFontSize": 14,
            "labelFontSize": 14,
            "rowPadding": 5,
            "padding": 3,
            "columns": {
              "expr": "floor(width / 155) == 0 ? 1 : floor(width / 155)"
            }
          },
          "scale": {
            "domain": [
              "HF PETG 235 (234)",
              "UHF PETG 235 (237)",
              "HF PETG 248 (247)",
              "UHF PETG 248 (250)"
            ],
            "range": [
              "#1F78B5",
              "#AEC7E8",
              "#FF7F0D",
              "#FFBB78",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#ff9896",
              "#9467bd",
              "#ce6dbd",
              "#bcbd22"
            ]
          }
        },
        "opacity": {"condition": {"param": "hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "hotend",
          "select": {"type": "point", "fields": ["hotend"]},
          "bind": "legend"
        }
      ]
    },
    {
      "mark": {
        "type": "text",
        "dx": 5,
        "dy": -10,
        "fontSize": 16,
        "align": "left",
        "clip": true
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 60}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": -15, "domainMax": 0}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "hotend"}
      }
    }
  ]
}
```

## Тесты стабильности температуры

### Методика тестирования

На CHC XL и Rapido ACE используются керамические нагреватели, которые заметно теряют мощность при повышении температуры. Поэтому было решено проверить, хватит ли им производительности для стабильной работы при самых сложных из распространённых условий эксплуатации. Для этого включалось охлаждение модели на 100% оборотов, после чего хотэнд разогревался до 320ºC по термистору и снимался график скважности ШИМ при этом. Чем больше требуется средняя скважность для поддержания заданной температуры, тем меньше запас по мощности => тем хуже.

### Goliath

При выключеном охлаждении средняя скважность ШИМ ~30%. При включенном охлаждении ~60%.

![](./pics/uhf_hotends_goliath_stability.png)

### CHC XL

При выключенном охлаждении средняя скважность ШИМ ~40%. При включенном охлаждении ~60%.

![](./pics/uhf_hotends_chc-xl_stability.png)

### Rapido ACE

Ввиду нестабильного снятия температуры на этом хотэнде, не получается нормально откалибровать PID. Возникающие вследствие этого сильные колебания температуры приводят к тому, что скважность ШИМ резко колеблется в очень широком диапазоне значений. К примеру, при тестировании объёмного расхода на 248 (250)ºC скважность ШИМ на некоторых расходах колебалась от 0 до 100%, что видно в левой части графика. Из-за этого определить среднюю скважность ШИМ не представляется возможным.

Вместо теста стабильности поддержания высокой температуры было проверено, а способен ли UHF версия этого хотэнда в целом поддерживать температуру 320ºC под сильным обдувом. Это отражено в правой части графика. Температура, хоть и колеблется в пределах нескольких градусов от целевой, но всё-таки остаются плюс-минус близкой к целевой. То есть, формально, хотэнд способен работать на этой температуре. Скважность ШИМ при этом близится к 100%, то есть запаса по мощности нет.

![](./pics/uhf_hotends_rapido-ace_stability.png)

## Выводы

Goliath и CHC XL показали очень высокую производительность по сравнению с более простыми и дешевыми хотэндами наподобие Volcano. При этом CHC XL немного уступает Goliath по производительности, но разница может быть связана не с прогревом материала, а с гидравлическим сопротивлением. На это указывает то, что, несмотря на значимую разницу по величине недоэкструзии, пропуски подающего механизма начинаются с одинакового расхода. Плюс к этому, у этих хотэндов разная величина недоэкструзии даже на очень низких расходах, где оба точно полностью и равномерно прогревают филамент.

Тем не менее, Goliath всё равно показывает себя лучше, так что, если вы гонитесь за максимальными характеристиками вне зависимости от цены, то он будет лучшим выбором. CHC XL отстаёт, но не сильно. При этом стоит значительно дешевле. Так что он является оптимальным вариантом по отношению цена/производительность.

Rapido ACE HF, хоть и имеет значительно большую по сравнению с Volcano длину горячей зоны, но это не даёт ему значимого прироста производительности без перегрева. С перегревом на ~5% ситуация улучшается и Rapido ACE HF уже показывает себя лучше Volcano, но не настолько, чтобы ради этого покупать настолько дорогой хотэнд. Усугубляется ситуация проблемами со стабильностью температуры.

Rapido ACE UHF, ожидаемо, не дотягивает ни до Goliath, ни до CHC XL. Поэтому рассматриваться как полноценная альтернатива этим двум хотэндам не может.

Установка удлинителей горячей зоны на Goliath не даёт улучшения производительности, если не прибегать к небольшому перегреву филамента. Для CHC XL удлинитель даёт прирост к производительности, но без перегрева им можно пренебречь. Значимым прирост производительности становится только если немного перегревать филамент.

!!! note "Ссылки на хотэнды есть в [3D рекомендаторе](../part-navi/extruder.md)"