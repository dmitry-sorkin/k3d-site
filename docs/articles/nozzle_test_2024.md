---
icon: material/printer-3d-nozzle
title: Тест сопел
description: Volcano 0.4mm
hide:
   - navigation
   - toc
---

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "K3D Nozzle flowrate test",
    "subtitle": "PETG; 245ºC; K3D Minifeeder; Titanium heatbreak",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "description": "Volcano 0.4mm nozzle test - PETG 245ºC",
  "width": 1100,
  "height": 700,
  "data": {
    "url": "https://raw.githubusercontent.com/dmitry-sorkin/k3d-site/main/docs/assets/nozzle_test_2024_petg.csv"
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
          "scale": {"domainMin": 5, "domainMax": 30},
          "axis": {
            "title": "Flow rate [mm^3/s]",
            "titleFontSize": 16,
            "labelFontSize": 16
          }
        },
        "y": {
          "field": "sample_weight",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5},
          "axis": {
            "title": "Avg. sample weight [g]",
            "titleFontSize": 16,
            "labelFontSize": 16
          }
        },
        "color": {
          "field": "Nozzle",
          "type": "nominal",
          "legend": {
            "orient": "right",
            "titleFontSize": 16,
            "labelFontSize": 16
          }
        },
        "opacity": {"condition": {"param": "nozzle", "value": 1}, "value": 0.1},
        "strokeDash": {
          "field": "skipping",
          "type": "nominal",
          "legend": {
            "orient": "right",
            "titleFontSize": 16,
            "labelFontSize": 16
          }
        }
      },
      "params": [
        {
          "name": "nozzle",
          "select": {"type": "point", "fields": ["Nozzle"]},
          "bind": "legend"
        },
        {"name": "grid", "select": "interval", "bind": "scales"}
      ]
    },
    {
      "mark": {
        "type": "text",
        "dx": 5,
        "dy": -10,
        "fontSize": 16,
        "align": "left"
      },
      "encoding": {
        "x": {
          "field": "volumetric_flow",
          "type": "quantitative",
          "scale": {"domainMin": 5, "domainMax": 30}
        },
        "y": {
          "field": "sample_weight",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5}
        },
        "text": {"field": "sample_weight"},
        "opacity": {"condition": {"param": "nozzle", "empty": false, "value": 1}, "value": 0},
        "color": {"field": "Nozzle"}
      }
    }
  ]
}
```