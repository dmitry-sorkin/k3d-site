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
        "underextrusion": -0.96,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": -1.55,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": -2.08,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": -2.62,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": -2.78,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": -3.53,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": -4.45,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": -5.2,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": -6.38,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": -6.38,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": -9.71,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": -16.75,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": -23.57,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": -29.18,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": -34.74,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": -39.09,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": -43.52,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": -47.47,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": -50.93,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": -52.79,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 4,
        "underextrusion": -0.2,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.74,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 8,
        "underextrusion": -1.17,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 10,
        "underextrusion": -1.65,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 12,
        "underextrusion": -2.19,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 14,
        "underextrusion": -2.62,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 16,
        "underextrusion": -3.05,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 18,
        "underextrusion": -3.7,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 20,
        "underextrusion": -4.23,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 22,
        "underextrusion": -4.98,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 22,
        "underextrusion": -4.98,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 24,
        "underextrusion": -5.63,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 26,
        "underextrusion": -6.81,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 28,
        "underextrusion": -14.01,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 30,
        "underextrusion": -20.67,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 32,
        "underextrusion": -26.2,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 34,
        "underextrusion": -33.51,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 36,
        "underextrusion": -38.88,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 38,
        "underextrusion": -44.46,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 40,
        "underextrusion": -47.47,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 4,
        "underextrusion": 0.06,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.31,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.85,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 10,
        "underextrusion": -1.17,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 12,
        "underextrusion": -1.76,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 14,
        "underextrusion": -2.3,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 16,
        "underextrusion": -3,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 18,
        "underextrusion": -3.37,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 20,
        "underextrusion": -4.23,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 22,
        "underextrusion": -5.47,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 22,
        "underextrusion": -5.47,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 24,
        "underextrusion": -6.33,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 26,
        "underextrusion": -13.42,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 28,
        "underextrusion": -22.6,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 30,
        "underextrusion": -27.65,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 32,
        "underextrusion": -33.93,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 34,
        "underextrusion": -38.12,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 36,
        "underextrusion": -43.12,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 38,
        "underextrusion": -47.15,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 40,
        "underextrusion": -49.24,
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
          "scale": {"domainMin": -10, "domainMax": 1},
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
        "underextrusion": -0.04,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.37,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.74,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": -0.96,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": -1.33,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": -2.08,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": -2.14,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": -2.84,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": -3.21,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": -3.59,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": -5.09,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": -5.09,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": -12.02,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": -17.55,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": -21.63,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": -27.06,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": -31.09,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": -35.44,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": -37.48,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": -42.07,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 4,
        "underextrusion": 0.23,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 6,
        "underextrusion": 0.01,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.47,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 10,
        "underextrusion": -0.42,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 12,
        "underextrusion": -0.58,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 14,
        "underextrusion": -1.55,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 16,
        "underextrusion": -1.55,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 18,
        "underextrusion": -2.30,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 20,
        "underextrusion": -2.46,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 22,
        "underextrusion": -2.46,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 24,
        "underextrusion": -2.73,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 26,
        "underextrusion": -3.16,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 28,
        "underextrusion": -3.48,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 30,
        "underextrusion": -3.96,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 32,
        "underextrusion": -4.61,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 34,
        "underextrusion": -5.90,
        "skipping": 0
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 34,
        "underextrusion": -5.90,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 36,
        "underextrusion": -11.43,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 38,
        "underextrusion": -15.94,
        "skipping": 1
      },
      {
        "Hotend": "H2D 0.4",
        "volumetric_flow": 40,
        "underextrusion": -20.56,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 4,
        "underextrusion": -0.02,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.42,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.58,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 10,
        "underextrusion": -0.98,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 12,
        "underextrusion": -0.74,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 14,
        "underextrusion": -1.14,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 16,
        "underextrusion": -1.47,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 18,
        "underextrusion": -1.87,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 20,
        "underextrusion": -2.67,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 22,
        "underextrusion": -2.67,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 24,
        "underextrusion": -3.40,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 26,
        "underextrusion": -4.04,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 28,
        "underextrusion": -5.17,
        "skipping": 0
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 28,
        "underextrusion": -5.17,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 30,
        "underextrusion": -10.17,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 32,
        "underextrusion": -15.32,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 34,
        "underextrusion": -20.40,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 36,
        "underextrusion": -24.59,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 38,
        "underextrusion": -28.86,
        "skipping": 1
      },
      {
        "Hotend": "A1 0.4",
        "volumetric_flow": 40,
        "underextrusion": -32.81,
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
          "scale": {"domainMin": -10, "domainMax": 1},
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
        "underextrusion": -0.37,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.96,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.96,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": -1.22,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": -1.71,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": -1.76,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": -2.46,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": -2.62,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": -2.94,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": -3.75,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": -4.50,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": -4.50,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": -5.90,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": -16.59,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": -24.05,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": -22.71,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": -28.56,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": -33.88,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": -36.66,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": -41.78,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 4,
        "underextrusion": -0.63,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.42,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.37,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 10,
        "underextrusion": -1.01,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 12,
        "underextrusion": -1.12,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 14,
        "underextrusion": -2.19,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 16,
        "underextrusion": -2.14,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 18,
        "underextrusion": -3.59,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 18,
        "underextrusion": -3.59,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 20,
        "underextrusion": -8.27,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 22,
        "underextrusion": -17.02,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 24,
        "underextrusion": -26.09,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 26,
        "underextrusion": -34.15,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 28,
        "underextrusion": -40.33,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 30,
        "underextrusion": -43.01,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 32,
        "underextrusion": -46.13,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 34,
        "underextrusion": -48.81,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 36,
        "underextrusion": -51.61,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 38,
        "underextrusion": -54.61,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 40,
        "underextrusion": -57.03,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 4,
        "underextrusion": -1.60,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 6,
        "underextrusion": -1.55,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 8,
        "underextrusion": -1.49,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 10,
        "underextrusion": -1.92,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 12,
        "underextrusion": -2.11,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 14,
        "underextrusion": -2.57,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 16,
        "underextrusion": -2.73,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 18,
        "underextrusion": -3.40,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 20,
        "underextrusion": -3.16,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 22,
        "underextrusion": -3.43,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 24,
        "underextrusion": -4.13,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 24,
        "underextrusion": -4.13,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 26,
        "underextrusion": -5.15,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 28,
        "underextrusion": -7.72,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 30,
        "underextrusion": -14.44,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 32,
        "underextrusion": -11.75,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 34,
        "underextrusion": -19.49,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 36,
        "underextrusion": -25.18,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 38,
        "underextrusion": -30.18,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 40,
        "underextrusion": -34.42,
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
          "scale": {"domainMin": -10, "domainMax": 1},
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
        "underextrusion": -0.15,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.42,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.2,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 10,
        "underextrusion": -0.2,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 12,
        "underextrusion": -0.1,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 14,
        "underextrusion": -0.53,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 16,
        "underextrusion": -0.85,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 18,
        "underextrusion": -0.79,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 20,
        "underextrusion": -1.55,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 22,
        "underextrusion": -1.65,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 24,
        "underextrusion": -1.98,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 26,
        "underextrusion": -2.57,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 28,
        "underextrusion": -2.35,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": -2.94,
        "skipping": 0
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 30,
        "underextrusion": -2.94,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 32,
        "underextrusion": -5.82,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 34,
        "underextrusion": -8.39,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 36,
        "underextrusion": -14.71,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 38,
        "underextrusion": -18.68,
        "skipping": 1
      },
      {
        "Hotend": "Stock 0.4",
        "volumetric_flow": 40,
        "underextrusion": -23.35,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 4,
        "underextrusion": -0.37,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.63,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.63,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 10,
        "underextrusion": -1.06,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 12,
        "underextrusion": -1.49,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 14,
        "underextrusion": -1.82,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 16,
        "underextrusion": -1.98,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 18,
        "underextrusion": -3.32,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 20,
        "underextrusion": -3.32,
        "skipping": 0
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 20,
        "underextrusion": -3.32,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 22,
        "underextrusion": -4.66,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 24,
        "underextrusion": -19.92,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 26,
        "underextrusion": -26.04,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 28,
        "underextrusion": -30.98,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 30,
        "underextrusion": -34.15,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 32,
        "underextrusion": -37.86,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 34,
        "underextrusion": -41.4,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 36,
        "underextrusion": -44.73,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 38,
        "underextrusion": -47.31,
        "skipping": 1
      },
      {
        "Hotend": "Bambuhot 0.4",
        "volumetric_flow": 40,
        "underextrusion": -47.31,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 4,
        "underextrusion": -0.15,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 6,
        "underextrusion": -0.42,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 8,
        "underextrusion": -0.42,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 10,
        "underextrusion": -1.29,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 12,
        "underextrusion": -0.79,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 14,
        "underextrusion": -0.96,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 16,
        "underextrusion": -0.96,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 18,
        "underextrusion": -1.01,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 20,
        "underextrusion": -0.85,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 22,
        "underextrusion": -2.35,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 24,
        "underextrusion": -2.3,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 26,
        "underextrusion": -2.67,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 28,
        "underextrusion": -3.64,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 30,
        "underextrusion": -3.4,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 32,
        "underextrusion": -4.18,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 34,
        "underextrusion": -5.25,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 36,
        "underextrusion": -5.47,
        "skipping": 0
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 36,
        "underextrusion": -5.47,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 38,
        "underextrusion": -10.46,
        "skipping": 1
      },
      {
        "Hotend": "Conch 0.4",
        "volumetric_flow": 40,
        "underextrusion": -14.22,
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
          "scale": {"domainMin": -10, "domainMax": 1},
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