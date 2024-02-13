---
icon: material/printer-3d-nozzle
title: Тест сопел
description: Volcano 0.4mm
---

# Тест сопел

В этой статье представлены результаты тестирования следующих видов сопел:

| Материал | Маркетинговое название | Тип | Диаметр | Производитель | Примечание |
|:-------- |:---:|:--- |:-------------:|:-------:|:---------- |
| Латунь | - | Volcano | 0.4мм | Trianglelab | Б\У износ незначителен. Прочищено перед тестом |
| Закалённая сталь | - | Volcano | 0.4мм | Trianglelab | Б\У износ незначителен. Прочищено перед тестом |
| Никелированная медь, вставка из закалённой стали | Bimetal HC | Volcano | 0.4мм | Mellow | 2 экземпляра. Новые |
| Никелированная медь, вставка из закалённой стали | ZS | Volcano | 0.4мм | Trianglelab | Б\У износ незначителен. Прочищено перед тестом |
| Закалённая сталь с CHT медной вставкой | - | Volcano | 0.4мм | ? | Новое |
| Латунное с CHT медной вставкой | - | Volcano | 0.4мм | ? | Новое |
| Никелированная медь, вставка из карбида вольфрама | ZSTC | Volcano | 0.4мм | Trianglelab | Новое |
| Карбид вольфрама[^1] | - | Volcano | 0.4мм | Oston Carbide | Б\У износ незначителен. Протестирован и без прочистки, и после прочистки |
| Никелированная медь, вставка из поликристаллического алмаза | ZSD | Volcano | 0.4мм | Trianglelab | Новое |

## Тест объёмного расхода

### Методика тестирования

Испытания проводятся экструдированием 500мм филамента - по 3 раза на каждом мгновенном объёмном расходе: 5, 10, 12, 14 ... 30 мм³/с. Далее полученные кусочки филамента взвешиваются и наносятся на диаграмму, на которой по горизонтали наносятся значения мгновенного объёмного расхода, а по вертикали масса образца.

Тестирование проводилось на 2 филаментах: PETG от Bestfilament при температуре 245ºC и PLA от H-T-P при температуре 220ºC. Это сделано т.к. существующие на данный момент методики калибровки температуры печати как критерий берут внешний вид деталей, что слишком субъективно, а более объективные критерии требуют оборудования, которого у меня нет. Поэтому было принято решение тесты всех сопел проводить на одинаковых температурах, примерно равных средней температуре из предыдущих тестов.

Также при тестировании сопел фиксировалось на каком объёмном расходе начинались пропуски подающего механизма (на диаграммах участки после последнего образца без пропусков отмечены пунктиром). 

!!! warning "Важно отметить, что все данные из этих тестов сильно зависят от подающего механизма, радиатора, термобарьера и даже от настроек драйвера. Поэтому результаты этих тестов напрямую можно сравнивать только между собой. При сравнении с другими тестами, необходимо учитывать разницу в тестовом стенде"

### Объёмный расход на PETG

#### В абсолютных значениях

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
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "url": "https://raw.githubusercontent.com/dmitry-sorkin/k3d-site/main/docs/articles/data/nozzle_test_2024_petg.csv"
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
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "sample_weight",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5},
          "axis": {
            "title": "Avg. sample weight [g]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Nozzle",
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
          "scale": {"domain": ["Trianglelab brass","Trianglelab HS","Mellow Bimetal HC 1", "Mellow Bimetal HC 2","Trianglelab ZS","Noname brass CHT","Noname HS CHT","Trianglelab ZSTC","Oston carbide (used)","Oston carbide (clean)","Trianglelab ZSD"], "range": ["#1F78B5","#AEC7E8","#FF7F0D","#FFBB78","#2CA02C","#98DF8A","#D62729","#ff9896","#9467bd","#ce6dbd", "#bcbd22"]}
        },
        "opacity": {"condition": {"param": "nozzle", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "nozzle",
          "select": {"type": "point", "fields": ["Nozzle"]},
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
          "field": "sample_weight",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5}
        },
        "text": {"field": "sample_weight"},
        "opacity": {
          "condition": {"param": "nozzle", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Nozzle"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

Примечания:

- Сопло от Oston Carbide с пометкой "used" - после обычной эксплуатации на различных видах филамента. С пометкой "clean" - после atomic pull;
- Mellow Bimetal HC протестирован в 2 экземплярах т.к. были жалобы на нестабильное качество этих сопел, и у меня как раз два было в наличии.

#### В относительных значениях

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
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "url": "https://raw.githubusercontent.com/dmitry-sorkin/k3d-site/main/docs/articles/data/nozzle_test_2024_petg_relative.csv"
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
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "sample_weight",
          "type": "quantitative",
          "scale": {"domainMin": 90, "domainMax": 100},
          "axis": {
            "title": "Avg. sample weight [g]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Nozzle",
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
              "Trianglelab brass",
              "Trianglelab HS",
              "Mellow Bimetal HC 1",
              "Mellow Bimetal HC 2",
              "Trianglelab ZS",
              "Noname brass CHT",
              "Noname HS CHT",
              "Trianglelab ZSTC",
              "Oston carbide (used)",
              "Oston carbide (clean)",
              "Trianglelab ZSD"
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
        "opacity": {"condition": {"param": "nozzle", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "nozzle",
          "select": {"type": "point", "fields": ["Nozzle"]},
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
        "y": {"field": "sample_weight", "type": "quantitative"},
        "text": {"field": "sample_weight"},
        "opacity": {
          "condition": {"param": "nozzle", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Nozzle"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

Данная диаграмма построена на основе диаграммы в абсолютных значениях, за исключением того, что тут первые значения в каждом ряду приравнены к 100%, а последующие высчитываются от них. Это приводит к потере значимой части информации, то есть использовать эту диаграмму в качестве основной или единственной некорректно. Тем не менее, она предоставляет возможность хотя бы очень приблизительно сравнить результаты с другими тестами, где данные приведены только в относительных значениях.

### Объёмный расход на PLA

#### В абсолютных значениях

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
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "url": "https://raw.githubusercontent.com/dmitry-sorkin/k3d-site/main/docs/articles/data/nozzle_test_2024_pla.csv"
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
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "sample_weight",
          "type": "quantitative",
          "scale": {"domainMin": 1.35, "domainMax": 1.5},
          "axis": {
            "title": "Avg. sample weight [g]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Nozzle",
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
              "Trianglelab brass",
              "Trianglelab HS",
              "Mellow Bimetal HC",
              "Trianglelab ZS",
              "Noname brass CHT",
              "Noname HS CHT",
              "Trianglelab ZSTC",
              "Oston carbide (clean)",
              "Trianglelab ZSD"
            ],
            "range": [
              "#1F78B5",
              "#AEC7E8",
              "#FF7F0D",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#ff9896",
              "#ce6dbd",
              "#bcbd22"
            ]
          }
        },
        "opacity": {"condition": {"param": "nozzle", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "nozzle",
          "select": {"type": "point", "fields": ["Nozzle"]},
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
        "y": {"field": "sample_weight", "type": "quantitative"},
        "text": {"field": "sample_weight"},
        "opacity": {
          "condition": {"param": "nozzle", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Nozzle"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

Примечания: 

- Тот факт, что масса образцов при тесте PLA в среднем заметно меньше массы образцов при тесте PETG обусловлен разницей в плотности этих филаментов;
- Из двух сопел Mellow Bimetal HC было протестировано только одно, в диаграммах для PETG филамента обозначенное под цифрой 1;
- Сопло от Oston Carbide тестировалось после всех тестов на PETG, поэтому оно не могло быть протестировано в варианте "до чистки";
- Все сопла, кроме Trianglelab ZS тестировались на филаменте PLA от H-T-P. К сожалению, он закончился прямо перед последним тестом, и Trianglelab ZS пришлось тестировать на PLA от FDPlast. Поэтому заметно отличающиеся результаты на этом сопле могут быть объяснены просто другим филаментом.

#### В относительных значениях

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
  "width": "container",
  "height": 600,
  "autosize": "pad",
  "data": {
    "url": "https://raw.githubusercontent.com/dmitry-sorkin/k3d-site/main/docs/articles/data/nozzle_test_2024_pla_relative.csv"
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
            "labelFontSize": 14
          }
        },
        "y": {
          "field": "sample_weight",
          "type": "quantitative",
          "scale": {"domainMin": 90, "domainMax": 100},
          "axis": {
            "title": "Avg. sample weight [g]",
            "titleFontSize": 16,
            "labelFontSize": 14
          }
        },
        "color": {
          "field": "Nozzle",
          "type": "nominal",
          "legend": {
            "orient": "top",
            "titleFontSize": 14,
            "labelFontSize": 14,
            "rowPadding": 5,
            "padding": 3,
            "columns": {"expr": "floor(width / 155) == 0 ? 1 : floor(width / 155)"}
          },
          "scale": {
            "domain": [
              "Trianglelab brass",
              "Trianglelab HS",
              "Mellow Bimetal HC",
              "Trianglelab ZS",
              "Noname brass CHT",
              "Noname HS CHT",
              "Trianglelab ZSTC",
              "Oston carbide (clean)",
              "Trianglelab ZSD"
            ],
            "range": [
              "#1F78B5",
              "#AEC7E8",
              "#FF7F0D",
              "#2CA02C",
              "#98DF8A",
              "#D62729",
              "#ff9896",
              "#ce6dbd",
              "#bcbd22"
            ]
          }
        },
        "opacity": {"condition": {"param": "nozzle", "value": 1}, "value": 0.2},
        "strokeDash": {"field": "skipping", "type": "nominal", "legend": null}
      },
      "params": [
        {
          "name": "nozzle",
          "select": {"type": "point", "fields": ["Nozzle"]},
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
          "field": "sample_weight",
          "type": "quantitative"
        },
        "text": {"field": "sample_weight"},
        "opacity": {
          "condition": {"param": "nozzle", "empty": false, "value": 1},
          "value": 0
        },
        "color": {"field": "Nozzle"}
      }
    }
  ]
}
```

Управление диаграммой: :material-information-outline:{ title="Выделить график - ЛКМ по названию в легенде<br>Выделить несколько графиков - Shift + ЛКМ<br>Сбросить выделение - ЛКМ по свободному месту в легенде" }

Данная диаграмма построена на основе диаграммы в абсолютных значениях, за исключением того, что тут первые значения в каждом ряду приравнены к 100%, а последующие высчитываются от них. Это приводит к потере значимой части информации, то есть использовать эту диаграмму в качестве основной или единственной некорректно. Тем не менее, она предоставляет возможность хотя бы очень приблизительно сравнить результаты с другими тестами, где данные приведены только в относительных значениях.

## Тест Linear Advance

### Методика тестирования

Все сопла тестировались только на PETG по модели, сгенерированной [калибровщиком LA](../calibrations/la/calibrator.md). Сначала тест проходил при времени сглаживания Pressure Advance 0.02с. Если на этой модели в районе k-фактора, на котором были наименьшие перепады толщины линий, были сильные дефекты, то тест признавался неудачным и далее запускался тот же самый G-код, но вручную выставлялось время сглаживания 0.03с. Если и этот тест был неудачным, то выставлялось время сглаживания 0.04с.

Важно отметить, что, некоторые тестовые модели печатались матовыми, в то время, как другие - глянцевыми. Причём в рамках этого теста то, насколько модель была матовой, в точности коррелировало с тем, насколько большое время сглаживания приходилось ставить в итоге. Учитывая, что матовая поверхность на ПЭТГ получается только если он заметно недогрет, то можно предположить, что на соплах, для которых хватило времени сглаживания 0.02с, филамент прогревается лучше, чем на соплах, где требуется 0.03 или 0.04с.

### Результаты тестов

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "K3D Nozzle pressure advance test",
    "subtitle": "PETG; 245ºC; K3D Minifeeder; Titanium heatbreak",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "width": "container",
  "height": 350,
  "description": "Volcano 0.4mm nozzle test - retractions - PETG 245ºC",
  "padding": 10,
  "data": {
    "values": [
      {"a": "Trianglelab brass", "b": 0.05, "c": "0.05 @ 0.02s"},
      {"a": "Trianglelab HS", "b": 0.06, "c": "0.06 @ 0.04s"},
      {"a": "Mellow Bimetal HC", "b": 0.06, "c": "0.06 @ 0.02s"},
      {"a": "Trianglelab ZS", "b": 0.06, "c": "0.06 @ 0.03s"},
      {"a": "Noname HS CHT", "b": 0.09, "c": "0.09 @ 0.03s"},
      {"a": "Noname brass CHT", "b": 0.07, "c": "0.07 @ 0.03s"},
      {"a": "Trianglelab ZSTC", "b": 0.05, "c": "0.05 @ 0.02s"},
      {"a": "Oston carbide", "b": 0.13, "c": "0.13 @ 0.04s"},
      {"a": "Trianglelab ZSD", "b": 0.06, "c": "0.06 @ 0.02s"}
    ]
  },
  "encoding": {
    "y": {
      "field": "a",
      "type": "nominal",
      "axis": {"title": "", "labelFontSize": 16},
      "sort": "x"
    },
    "x": {
      "field": "b",
      "type": "quantitative",
      "scale": {"domain": [0, 0.15]},
      "axis": {
        "title": "PA [s]",
        "titleFontSize": 16,
        "labelFontSize": 16,
        "tickCount": 3
      }
    }
  },
  "layer": [
    {"mark": {"type": "bar"}},
    {
      "mark": {
        "type": "text",
        "align": "right",
        "baseline": "middle",
        "dx": -5,
        "fontSize": 14,
        "fontWeight": 700,
        "color": "white"
      },
      "encoding": {"text": {"field": "c", "type": "nominal"}}
    }
  ]
}
```

Примечания:

- При тестах сопел Oston Carbide, стального закалённого и стального закалённого с СНТ вставкой оставались значимые дефекты (разрывы) даже при времени сглаживания 0.04с;
- На тесте Trianglelab ZS при времени сглаживания 0.02с дефекты были, но не очень большие. Скорее всего, значение сглаживания 0.025с сработало бы хорошо, но в рамках этого теста я для всех сопел поднимал по 0.01с, так что и тут сделал так же.

## Тест откатов

### Методика тестирования

Все сопла тестировались только на PETG по модели, сгенерированной [калибровщиком откатов](../calibrations/retractions/calibrator.md). Все тесты выполнялись одним и тем же G-кодом, за исключением того, что для каждого сопла в печатаемый файл вручную вписывался коэффициент PA и время сглаживания PA из предыдущего теста. Скорость откатов для всех тестов была неизменной - 30 мм/с.

### Результаты тестов

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "K3D Nozzle retraction test",
    "subtitle": "PETG; 245ºC; K3D Minifeeder; Titanium heatbreak; 30mm/s",
    "fontSize": 20,
    "subtitleFontSize": 16
  },
  "width": "container",
  "height": 350,
  "description": "Volcano 0.4mm nozzle test - retractions - PETG 245ºC",
  "padding": 10,
  "data": {
    "values": [
      {"a": "Trianglelab brass", "b": 0.3, "c": "0.3mm"},
      {"a": "Trianglelab HS", "b": 0.3, "c": "0.3mm"},
      {"a": "Mellow Bimetal HC", "b": 0.4, "c": "0.4mm"},
      {"a": "Trianglelab ZS", "b": 0.3, "c": "0.3mm"},
      {"a": "Noname HS CHT", "b": 0.5, "c": "0.5mm"},
      {"a": "Noname brass CHT", "b": 0.3, "c": "0.3mm"},
      {"a": "Trianglelab ZSTC", "b": 0.4, "c": "0.4mm"},
      {"a": "Oston carbide", "b": 0.4, "c": "0.4mm"},
      {"a": "Trianglelab ZSD", "b": 0.3, "c": "0.3mm"}
    ]
  },
  "encoding": {
    "y": {
      "field": "a",
      "type": "nominal",
      "axis": {"title": "", "labelFontSize": 16},
      "sort": "x"
    },
    "x": {
      "field": "b",
      "type": "quantitative",
      "scale": {"domain": [0, 0.5]},
      "axis": {
        "title": "Retraction length [mm]",
        "titleFontSize": 16,
        "labelFontSize": 16,
        "tickCount": 5
      }
    }
  },
  "layer": [
    {"mark": {"type": "bar"}},
    {
      "mark": {
        "type": "text",
        "align": "right",
        "baseline": "middle",
        "dx": -5,
        "fontSize": 14,
        "fontWeight": 700,
        "color": "white"
      },
      "encoding": {"text": {"field": "c", "type": "nominal"}}
    }
  ]
}
```

[^1]: Судя по результатам тестов, по-настоящему сопло изготовлено из полукарбида вольфрама.