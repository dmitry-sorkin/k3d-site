---
icon: material/printer-3d-nozzle
title: Тест хотэндов
subtitle: Flashforge AD5M/AD5X
description: PETG 235/255ºC
---

# Тест хотэндов AD5M / AD5X

## Список хотэндов

| Название | Цена | Ссылка | Примечание |
| :------- | :--: | :----: | :--------- |
| AD5M Stock | 28$ | [:material-shopping:](https://ali.click/0x6o31n?erid=2SDnjeTpmbF){ target="_blank" } | Только для AD5M |
| "бамбухот" | ~10$ | [:material-shopping:](https://ali.click/dz6o31e?erid=2SDnjemccyP){ target="_blank" } | Только для AD5M |
| Phaetus Conch (AD5M) | 31$ | [:material-shopping:](https://ali.click/t17o317?erid=2SDnjbrcNL1){ target="blank" } | Только для AD5M |
| AD5X Stock | 23$ | [:material-shopping:](https://ali.click/g97o31a?erid=2SDnjcsusbo){ target="_blank" } | Только для AD5X |
| Bambu Lab A1 hotend | 15$ | [:material-shopping:](https://ali.click/lf7o31s?erid=2SDnjdoHRkD){ target="_blank" } | Только для AD5X |
| Bambu Lab H2 hotend | 24$ | [:material-shopping:](https://ali.click/sp7o31s?erid=2SDnjf3vpZ7){ target="_blank" } | Только для AD5X |

## Методика тестирования

1. Тест проводился с помощью печати образцов одинакового объёма на разных значениях объёмного расхода, и последующего их взвешивания. Тестовый G-код был сгенерирован с помощью [K3D калибровщика расхода](../calibrations/max_flowrate/index.md);
2. В качестве материала использовался чёрный PETG от Bestfilament;
3. Температура 235ºС и 255ºС;
4. Для каждого хотэнда тест проводился по 3 раза. На графики внесены средние значения.

## Adventurer 5X

### PETG 235ºC

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "background": "transparent",
  "title": {
    "text": "K3D Hotend flowrate test",
    "subtitle": "Flashforge Adventurer 5X; PETG; 235ºC",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.04,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": 98.45,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": 97.92,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": 97.38,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": 97.22,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": 96.47,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": 95.55,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": 94.8,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": 93.62,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": 93.62,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": 90.29,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": 83.25,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": 76.43,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": 70.82,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": 65.26,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": 60.91,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": 56.48,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": 52.53,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": 49.07,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": 47.21,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.8,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.26,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 8,
        "underextrusion": 98.83,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 10,
        "underextrusion": 98.35,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 12,
        "underextrusion": 97.81,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 14,
        "underextrusion": 97.38,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 16,
        "underextrusion": 96.95,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 18,
        "underextrusion": 96.3,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 20,
        "underextrusion": 95.77,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 22,
        "underextrusion": 95.02,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 24,
        "underextrusion": 94.37,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 24,
        "underextrusion": 94.37,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 26,
        "underextrusion": 93.19,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 28,
        "underextrusion": 85.99,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 30,
        "underextrusion": 79.33,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 32,
        "underextrusion": 73.8,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 34,
        "underextrusion": 66.49,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 36,
        "underextrusion": 61.12,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 38,
        "underextrusion": 55.54,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 40,
        "underextrusion": 52.53,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 4,
        "underextrusion": 100.06,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.69,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.15,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 10,
        "underextrusion": 98.83,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 12,
        "underextrusion": 98.24,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 14,
        "underextrusion": 97.7,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 16,
        "underextrusion": 97,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 18,
        "underextrusion": 96.63,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 20,
        "underextrusion": 95.77,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 22,
        "underextrusion": 94.53,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 24,
        "underextrusion": 93.67,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 24,
        "underextrusion": 93.67,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 26,
        "underextrusion": 86.58,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 28,
        "underextrusion": 77.4,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 30,
        "underextrusion": 72.35,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 32,
        "underextrusion": 66.07,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 34,
        "underextrusion": 61.88,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 36,
        "underextrusion": 56.88,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 38,
        "underextrusion": 52.85,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 40,
        "underextrusion": 50.76,
        "skipping": 1
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
          "scale": {"domainMin": 4, "domainMax": 40},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 90, "domainMax": 101},
          "axis": {
            "title": "Relative extrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Hotend",
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
            "domain": ["Stock 0.4", "H2D 0.4", "A1 0.4"],
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
        "opacity": {"condition": {"param": "Hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "Hotend",
          "select": {"type": "point", "fields": ["Hotend"]},
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
          "scale": {"domainMin": 5, "domainMax": 30}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "Hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Hotend"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

### PETG 255ºC

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "background": "transparent",
  "title": {
    "text": "K3D Hotend flowrate test",
    "subtitle": "Flashforge Adventurer 5X; PETG; 255ºC",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.96,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.63,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.26,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": 99.04,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": 98.67,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": 97.92,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": 97.86,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": 97.16,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": 96.79,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": 96.41,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": 94.91,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": 94.91,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": 87.98,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": 82.45,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": 78.37,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": 72.94,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": 68.91,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": 64.56,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": 62.52,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": 57.93,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 4,
        "underextrusion": 100.23,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 6,
        "underextrusion": 100.01,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.53,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 10,
        "underextrusion": 99.58,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 12,
        "underextrusion": 99.42,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 14,
        "underextrusion": 98.45,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 16,
        "underextrusion": 98.45,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 18,
        "underextrusion": 97.70,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 20,
        "underextrusion": 97.54,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 22,
        "underextrusion": 97.54,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 24,
        "underextrusion": 97.27,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 26,
        "underextrusion": 96.84,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 28,
        "underextrusion": 96.52,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 30,
        "underextrusion": 96.04,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 32,
        "underextrusion": 95.39,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 34,
        "underextrusion": 94.10,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 34,
        "underextrusion": 94.10,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 36,
        "underextrusion": 88.57,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 38,
        "underextrusion": 84.06,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 40,
        "underextrusion": 79.44,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.98,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.58,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.42,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 10,
        "underextrusion": 99.02,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 12,
        "underextrusion": 99.26,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 14,
        "underextrusion": 98.86,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 16,
        "underextrusion": 98.53,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 18,
        "underextrusion": 98.13,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 20,
        "underextrusion": 97.33,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 22,
        "underextrusion": 97.33,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 24,
        "underextrusion": 96.60,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 26,
        "underextrusion": 95.96,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 28,
        "underextrusion": 94.83,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 28,
        "underextrusion": 94.83,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 30,
        "underextrusion": 89.83,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 32,
        "underextrusion": 84.68,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 34,
        "underextrusion": 79.60,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 36,
        "underextrusion": 75.41,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 38,
        "underextrusion": 71.14,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 40,
        "underextrusion": 67.19,
        "skipping": 1
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
          "scale": {"domainMin": 4, "domainMax": 40},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 90, "domainMax": 101},
          "axis": {
            "title": "Relative extrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Hotend",
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
            "domain": ["Stock 0.4", "H2D 0.4", "A1 0.4"],
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
        "opacity": {"condition": {"param": "Hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "Hotend",
          "select": {"type": "point", "fields": ["Hotend"]},
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
          "scale": {"domainMin": 5, "domainMax": 30}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "Hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Hotend"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }
## Adventurer 5M

### PETG 235ºC

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "background": "transparent",
  "title": {
    "text": "K3D Hotend flowrate test",
    "subtitle": "Flashforge Adventurer 5M; PETG; 235ºC",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.63,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.04,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.04,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": 98.78,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": 98.29,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": 98.24,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": 97.54,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": 97.38,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": 97.06,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": 96.25,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": 95.50,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": 95.50,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": 94.10,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": 83.41,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": 75.95,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": 77.29,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": 71.44,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": 66.12,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": 63.34,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": 58.22,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.37,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.58,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.63,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 10,
        "underextrusion": 98.99,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 12,
        "underextrusion": 98.88,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 14,
        "underextrusion": 97.81,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 16,
        "underextrusion": 97.86,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 18,
        "underextrusion": 96.41,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 18,
        "underextrusion": 96.41,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 20,
        "underextrusion": 91.73,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 22,
        "underextrusion": 82.98,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 24,
        "underextrusion": 73.91,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 26,
        "underextrusion": 65.85,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 28,
        "underextrusion": 59.67,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 30,
        "underextrusion": 56.99,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 32,
        "underextrusion": 53.87,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 34,
        "underextrusion": 51.19,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 36,
        "underextrusion": 48.39,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 38,
        "underextrusion": 45.39,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 40,
        "underextrusion": 42.97,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 4,
        "underextrusion": 98.40,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 6,
        "underextrusion": 98.45,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 8,
        "underextrusion": 98.51,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 10,
        "underextrusion": 98.08,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 12,
        "underextrusion": 97.89,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 14,
        "underextrusion": 97.43,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 16,
        "underextrusion": 97.27,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 18,
        "underextrusion": 96.60,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 20,
        "underextrusion": 96.84,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 22,
        "underextrusion": 96.57,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 24,
        "underextrusion": 95.87,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 24,
        "underextrusion": 95.87,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 26,
        "underextrusion": 94.85,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 28,
        "underextrusion": 92.28,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 30,
        "underextrusion": 85.56,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 32,
        "underextrusion": 88.25,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 34,
        "underextrusion": 80.51,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 36,
        "underextrusion": 74.82,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 38,
        "underextrusion": 69.82,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 40,
        "underextrusion": 65.58,
        "skipping": 1
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
          "scale": {"domainMin": 4, "domainMax": 40},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 90, "domainMax": 101},
          "axis": {
            "title": "Relative extrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Hotend",
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
            "domain": ["Stock 0.4", "Bambuhot 0.4", "Conch 0.4"],
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
        "opacity": {"condition": {"param": "Hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "Hotend",
          "select": {"type": "point", "fields": ["Hotend"]},
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
          "scale": {"domainMin": 5, "domainMax": 30}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "Hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Hotend"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }
### PETG 255ºC

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "background": "transparent",
  "title": {
    "text": "K3D Hotend flowrate test",
    "subtitle": "Flashforge Adventurer 5M; PETG; 255ºC",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "values": [
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.85,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.58,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.80,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": 99.80,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": 99.90,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": 99.47,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": 99.15,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": 99.21,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": 98.45,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": 98.35,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": 98.02,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": 97.43,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": 97.65,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": 97.06,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": 97.06,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": 94.18,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": 91.61,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": 85.29,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": 81.32,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": 76.65,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.63,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.37,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.37,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 10,
        "underextrusion": 98.94,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 12,
        "underextrusion": 98.51,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 14,
        "underextrusion": 98.18,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 16,
        "underextrusion": 98.02,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 18,
        "underextrusion": 96.68,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 20,
        "underextrusion": 96.68,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 20,
        "underextrusion": 96.68,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 22,
        "underextrusion": 95.34,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 24,
        "underextrusion": 80.08,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 26,
        "underextrusion": 73.96,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 28,
        "underextrusion": 69.02,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 30,
        "underextrusion": 65.85,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 32,
        "underextrusion": 62.14,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 34,
        "underextrusion": 58.60,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 36,
        "underextrusion": 55.27,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 38,
        "underextrusion": 52.69,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 40,
        "underextrusion": 52.69,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 4,
        "underextrusion": 99.85,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 6,
        "underextrusion": 99.58,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 8,
        "underextrusion": 99.58,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 10,
        "underextrusion": 98.71,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 12,
        "underextrusion": 99.21,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 14,
        "underextrusion": 99.04,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 16,
        "underextrusion": 99.04,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 18,
        "underextrusion": 98.99,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 20,
        "underextrusion": 99.15,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 22,
        "underextrusion": 97.65,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 24,
        "underextrusion": 97.70,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 26,
        "underextrusion": 97.33,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 28,
        "underextrusion": 96.36,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 30,
        "underextrusion": 96.60,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 32,
        "underextrusion": 95.82,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 34,
        "underextrusion": 94.75,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 36,
        "underextrusion": 94.53,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 36,
        "underextrusion": 94.53,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 38,
        "underextrusion": 89.54,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 40,
        "underextrusion": 85.78,
        "skipping": 1
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
          "scale": {"domainMin": 4, "domainMax": 40},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 90, "domainMax": 101},
          "axis": {
            "title": "Relative extrusion [%]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Hotend",
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
            "domain": ["Stock 0.4", "Bambuhot 0.4", "Conch 0.4"],
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
        "opacity": {"condition": {"param": "Hotend", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "Hotend",
          "select": {"type": "point", "fields": ["Hotend"]},
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
          "scale": {"domainMin": 5, "domainMax": 30}
        },
        "y": {
          "field": "underextrusion",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5}
        },
        "text": {"field": "underextrusion"},
        "opacity": {
          "condition": {"param": "Hotend", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Hotend"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

---

^`Реклама: ООО "АЛИБАБА.КОМ (РУ)" ИНН 7703380158`^ :material-information-outline:{ title="Статья написана по своему желанию и в интересах сообщества K3D. Нет человека или компании, которые заказали бы написание этой статьи, или каким-либо образом повлияли на информацию в ней. Данная приписка нужна так как на странице имеются ссылки, ведущие на товары на AliExpress" }