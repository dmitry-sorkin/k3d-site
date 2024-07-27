---
icon: fontawesome/solid/drum-steelpan
title: 3D рекомендатор - филамент
description: Пользовательский рейтинг производителей филамента
---

# 3D рекомендатор: филаменты и расходники

В этом разделе вы можете найти информацию по поводу того, какие филаменты от каких производителей, а также какие расходники стоит покупать, а какие нет.

- [:simple-telegram: Задать вопрос по поводу выбора комплектующих](https://t.me/K_3_D/1944033)
- [:simple-telegram: Сообщить об ошибке в telegram](https://t.me/K_3_D/1944075) - обязательно тегайте @dmitry_sorkin в сообщении, чтобы мне пришло уведомление
- [:octicons-mail-16: Сообщить об ошибке по e-mail](mailto:dbsorkin@gmail.com) - ⚠️ Только для сообщений об ошибках. На вопросы по e-mail не отвечаю

## Филаменты

### Тип филамента

<iframe width="900" height="506" src="https://www.youtube.com/embed/CzgOCkJbcxE?si=vsBjAMU5iFWQ6Lze" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

Для FDM 3D принтеров выпускают очень большое количество видов филамента, каждый из которых обладает своими уникальными физико-механическими свойствами. Иными словами, нет "хороших" и "плохих" филаментов, они все нужны для решения разных задач. Основные отличия самых распространенных видов филамента между собой описаны в соответствующем ролике.

Новичкам рекомендуется учиться на более простых в печати филаментах. Для начала хорошо подойдут PLA и PETG. Они хороши по механическим свойствам, но имеют низкую температуру эксплуатации. Если нужен более термостойкий материал, то оптимальным для новичков является композитный ABS т.к. стоит относительно недорого и имеет очень низкую усадку. Обратите внимание, что для печати композитными материалами нужно износостойкое сопло.

### Производители филамента

Даже один и тот же филамент, будучи выпущенным разными компаниями, может не только отличаться по цене, но и иметь разные физико-механические свойства и качество. Поэтому 2 раза в год проводится народный рейтинг производителей филамента. Оценка в этом рейтинге показывает то, насколько другие пользователи рекомендовали бы филамент от указанных компаний. 

```vegalite
{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "title": {
    "text": "Народный рейтинг производителей филамента 2024",
    "fontSize": 20,
    "subtitle": "k3d.tech/part-navi/filament",
    "subtitleFontSize": 16
  },
  "width": "container",
  "height": 1000,
  "padding": 10,
  "data": {
    "values": [
      {"a": "FDPlast", "b": 667, "c": 8.55, "d": "8.55 (668 голосов)"},
      {"a": "НИТ", "b": 526, "c": 7.04, "d": "7.04 (526 голосов}"},
      {"a": "Bestfilament", "b": 522, "c": 7.72, "d": "7.72 (522 голоса}"},
      {"a": "ESUN", "b": 499, "c": 8.39, "d": "8.39 (499 голосов}"},
      {"a": "Eryone", "b": 336, "c": 8.55, "d": "8.55 (336 голосов}"},
      {"a": "Mako", "b": 305, "c": 6.68, "d": "6.68 (305 голосов}"},
      {"a": "Некрасовский полимер", "b": 250, "c": 4.58, "d": "4.58 (250 голосов}"},
      {"a": "Syntech", "b": 236, "c": 7.43, "d": "7.43 (236 голосов}"},
      {"a": "Filamentarno", "b": 232, "c": 7.73, "d": "7.73 (232 голоса}"},
      {"a": "ABS Maker", "b": 227, "c": 7.69, "d": "7.69 (227 голосов}"},
      {"a": "Creality", "b": 198, "c": 7.85, "d": "7.85 (198 голосов}"},
      {"a": "Rec", "b": 191, "c": 7.31, "d": "7.31 (191 голос}"},
      {"a": "Greg", "b": 182, "c": 6.55, "d": "6.55 (182 голоса}"},
      {"a": "Hi Tech Plast", "b": 153, "c": 7.24, "d": "7.24 (153 голоса}"},
      {"a": "KINGROON", "b": 152, "c": 7.17, "d": "7.17 (152 голоса}"},
      {"a": "LIDER-3D", "b": 148, "c": 5.83, "d": "5.83 (148 голосов}"},
      {"a": "U3Print", "b": 146, "c": 7.4, "d": "7.40 (146 голосов}"},
      {"a": "Novaprint", "b": 131, "c": 8.21, "d": "8.21 (131 голос}"},
      {"a": "Geek Filament", "b": 127, "c": 6.95, "d": "6.95 (127 голосов}"},
      {"a": "3D-Club", "b": 117, "c": 8.33, "d": "8.33 (117 голосов}"},
      {"a": "DEXP", "b": 106, "c": 4.14, "d": "4.14 (106 голосов}"},
      {"a": "SUNLU", "b": 101, "c": 8.17, "d": "8.17 (101 голос}"},
      {"a": "KREMEN", "b": 101, "c": 7.01, "d": "7.01 (101 голос}"},
      {"a": "123 3D", "b": 101, "c": 6.73, "d": "6.73 (101 голос}"},
      {"a": "Anycubic", "b": 82, "c": 7.34, "d": "7.34 (82 голоса}"},
      {"a": "Bambu Lab", "b": 80, "c": 8.09, "d": "8.09 (80 голосов}"},
      {"a": "Elegoo", "b": 68, "c": 7.06, "d": "7.06 (68 голосов}"},
      {"a": "ПластикОФФ", "b": 62, "c": 6.73, "d": "6.73 (62 голоса}"},
      {"a": "Print Product", "b": 57, "c": 7.63, "d": "7.63 (57 голосов}"},
      {"a": "Plastiq", "b": 57, "c": 3.96, "d": "3.96 (57 голосов}"},
      {"a": "ATECO", "b": 49, "c": 8.58, "d": "8.58 (49 голосов*)"},
      {"a": "CYBERFIBER", "b": 41, "c": 8.24, "d": "8.24 (41 голос*)"}
    ]
  },
  "encoding": {
    "y": {
      "field": "a",
      "type": "nominal",
      "axis": {"title": "", "labelFontSize": 16},
      "sort": "-x"
    },
    "x": {
      "field": "c",
      "type": "quantitative",
      "scale": {"domain": [0, 10]},
      "axis": {
        "title": "Оценка",
        "titleFontSize": 16,
        "labelFontSize": 16,
        "tickCount": 10
      }
    }
  },
  "layer": [
    {"mark": {"type": "bar", "size": 28}},
    {
      "mark": {
        "type": "text",
        "align": "right",
        "baseline": "middle",
        "dx": -10,
        "fontSize": 14,
        "fontWeight": 700,
        "color": "white"
      },
      "encoding": {"text": {"field": "d", "type": "nominal"}}
    }
  ]
}
```
Голосование проводилось в июле 2024 года среди подписчиков K3D. В нем участвовало 1393 человека, которые оценили своё желание рекомендовать филамент разных производителей другим печатникам по 5-балльной шкале.

Из полученных голосов было отсеяно 15 шт. т.к. они нарушали правила голосования (вместо пропуска неизвестных производителей всем были проставлены одинаковые оценки, чаще всего минимальные). Также были отсеяны производители филамента, набравшие менее 50 голосов, кроме ATECO и CYBERFIBER. Для них пришлось сделать исключение т.к. уже после окончания голосования выяснилось, что они объединились и теперь являются одним производителем. Суммировать их голоса было бы некорректно т.к., скорее всего, были люди, которые голосовали одновременно за обоих производителей. Отсеивать этого производителя было бы некорректно т.к., скорее всего, они не набрали пороговое значение голосов ввиду моей ошибки. Так что самым справедливым вариантом мне показалось добавить их в голосование с пометкой и объяснением ситуации.

Список производителей, которые не прошли в рейтинг: (1)
{ .annotate }

1. INFILL (42 голоса)<br>STAR PLAST (35 голосов)<br>Flashforge (30 голосов)<br>Solid Filament (29 голосов)<br>Polymaker (28 голосов)<br>Wellhan (27 голосов)<br>OPY (26 голосов)<br>Element3D (24 голоса)<br>EXOFLEX (23 голоса)<br>Cactus (23 голоса)<br>FusRock (21 голос)<br>Tiger3D (20 голосов)<br>Funtastique (18 голосов)<br>PIC (17 голосов)<br>UNID (17 голосов)<br>ИКЦ (15 голосов)<br>Raise 3D (15 голосов)<br>Траектория 3D (12 голосов)<br>SEM (12 голосов)<br>JamgHe (11 голосов)<br>ChiTu Systems (10 голосов)<br>ОРФЕЙ (10 голосов)<br>YouSu (9 голосов)<br>NV Print (9 голосов)<br>Q3T (6 голосов)<br>Trianglelab (6 голосов)<br>Onlyplast (6 голосов)<br>ANYPLAST (6 голосов)<br>BIQU (6 голосов)<br>ColorFabb (5 голосов)<br>VolPrint (5 голосов)<br>Longer (5 голосов)<br>Fulcrum (5 голосов)<br>Sovol (5 голосов)<br>R-Filament (5 голосов)<br>RSPlast (4 голоса)<br>Sting 3D (4 голоса)<br>FL33 (4 голоса)<br>TWO TREES (4 голоса)<br>BEHANT (3 голоса)<br>Charm3D (3 голоса)<br>TOYAR (3 голоса)<br>Studia 3D (3 голоса)<br>Funtasy (3 голоса)<br>Mindao (2 голоса)<br>HELEOS (2 голоса)<br>3D Craft (1 голос)<br>Jorden (1 голос)<br>Sk-Plast (1 голос)

## Клей

`PLA`, `PETG`, `TPU`, `SBS`, `PVA`, `PVB` нормально липнут на стекло с покрытием Ultrabase или на стальной лист с покрытием PEI. Главное - обезжирить поверхность. `PP` липнет на стеклотекстолит FR-4. В других случаях требуется нанесение на стол клея.

| Наименование | Ссылки | Комментарий |
|:------------ |:------:|:----------- |
| 3D-клей | [:material-shopping:](https://www.ozon.ru/category/3d-kley/){ target="_blank" } [:material-shopping:](https://market.yandex.ru/search?text=3d%20%D0%BA%D0%BB%D0%B5%D0%B9&rs=eJwdUD0LQWEYfWUyyYy8o00p6eZeXf-ABQNFLP6BSddHSik2-UjXxmSifOUtTMpmkIE_otxzltPpOc9zznnfaMuddt1c4ps5OWiFj-DawUG10R2UOQMTH1TxBrf84HYBqtqDm9MEsIMr84QdOQSW5lBVhLwMZ3HHlZ2ip4dZC6aY9Blws8ImE8xFG552EXO1BgoNqqydyTGxc-xm8LZL5xh7vuBg_YBmj55bTuLcv6Kz9DJlxlbBJBx27M8U1We3Kq-e2LfyTFzy6kGHsYbEEd4i6_yBC9NXUNWHWTqbB8C_ITZpsGeTPKv_AQJfhpA%2C&allowCollapsing=1&local-offers-first=0){ target="_blank" } | Хорошо липнут `PLA`, `PETG`, `ABS`, `ASA`, `HIPS`, `TPU`, `PVA`, `PVB`, `SBS`, `SAN`. Средне липнут `PC`, `PA12`. Наносится тонким слоем, не влияет на внешний вид нижней поверхности |
| БФ-2 | [:material-shopping:](https://www.ozon.ru/product/kley-bf-2-solins-universalnyy-dlya-metalla-alyuminiya-keramiki-tkani-stali-plastikov-stekla-282539355/?asb=aDFtWa9TLph59Zk7AO3pf7DZXriRuM%252B9wwMXqXs7OVQ%253D&asb2=CXPORi2-s4-zMEmwM6TGgXFKWHiCaAVFM-aN9jCUN5sgVGt2s21HOekrUVhmejht&avtc=1&avte=2&avts=1698447793&keywords=%D0%B1%D1%84-2+solins){ target="_blank" } [:material-shopping:](https://www.ozon.ru/product/kley-bf-2-100-ml-0-08-kg-630977519/?advert=xHiVbuKE4It3U70aNbyoHILD-midKz5HVYVluxZcxqSnKSXm9RnKpv0LhRdujsQEFOgwaavRX_QJyI9fZZAf9WAe5wRIUu3SkRaF-SFVB4tuyoi7NeKUAU4pr8EjoVF6-KI5cusMhzyIaxUkrrB0IFZzJM_FX-RlGzvOXos17dE89jJWEtRia67_r8KSEJqhsuu4yFen3ZIZmhWhqWSFZFrhmEGh67I21JnFOg-cHNHjP0cx8jN_rJcB7glkZxoSFZsiFzyFUyGzMT7EcpLSBqpRKEW_WjdpPlv_EaWRAnqlhXn5R6eHoxHcOnKbLKn_jLXSM0ZO0VDR7mif4BpK_zqk&avtc=1&avte=4&avts=1698447853&keywords=%D0%B1%D1%84-2+solins){ target="_blank" } | Необходимо размешать с изопропанолом 1:1. Хорошо липнет всё, кроме `SEBS`. Слишком хорошо липнут нейлоны, могут быть проблемы с отделением от стола. Наносится кисточкой, может оставлять на модели следы |

## Смазка

Если производитель рекомендует использовать какую-то конкретную смазку, то следуйте этим рекомендациям. Если нигде не указано что и чем смазывать, то всё просто. Смазываются ходовые винты, а также направляющие кроме колёс. При этом, так как нагрузки и скорости в 3д печати низкие, то для этого подойдёт любая консистентная смазка для подшипников, например, литол-24. Проще всего купить в ближайшем автомагазине или на заправке.

!!! note "жидкие смазки в 3д принтерах не используются так как не встречаются системы подачи или удержания смазки"

| Наименование | Ссылки | Комментарий |
|:------------ |:------:|:----------- |
| Литол-24 | [:material-shopping:](https://www.ozon.ru/product/smazka-universalnaya-avtomobilnaya-vodostoykaya-gazpromneft-litol-24-150-gr-1046659010/?asb=NaMd2w6Wh4PC92ltjnz45wjObaO94DrNa0OS0eQ2Mog%253D&asb2=Bph6pI21MVLxpeNeZrqVtVfMIh4WxYi4a_rqE2_JLyl4TEDM7fuAgx7_3v51hfmE&avtc=1&avte=2&avts=1698448613&keywords=%D0%BB%D0%B8%D1%82%D0%BE%D0%BB+24){ target="_blank" } [:material-shopping:](https://market.yandex.ru/product--smazka-luxe-litol-24/673454000?nid=54800&show-uid=16984494965777696679016007&context=search&text=%D0%BB%D0%B8%D1%82%D0%BE%D0%BB-24&uniqueId=67601608&own-cards=67601608%3A100963245737&sku=100963245737&cpc=kZSAkpGfh0-OD3AFjzxADtzkbBxCBhYXaF8hg4CJr6XgTsXIpyYu9JqEDMIpvzAF88P-xsk8SwQtOmLMH2M6jSY2MjrSyfcSbtJCqMJGHm45pVScCkOYG7LMKUV08Ulal2YYWPcbe2ukBa3V6cossxRaUdleBQWlKBVzqCPzPrEXmGwkoQM9oVEXTLr4TdxE8o6JPK_zd95Jl3c92hiaNxB6Ar_y2wdwHXNL5LuKIbKVjhPMeXAe0kgr71onMOJXJ-Bp5vFZLgzCkhHFFchffQ%2C%2C&do-waremd5=8eqKNMjCfUMpTzoRaI9fdw&sponsored=1&rs=eJwzmsEcwFjFwvHzEOssRt4Luy_suNh0Yd-F3bpGJkcZGRSu7AWSCzbtAZINx7YDyQQLKyDJsM8GJL7aFsQuBLEfRILZTdYgNWdB7IReEKmwEyTScBFkAoMfiP2A3Q5IOkzfDZLtBOlt0AbZktAFUqPADBI5IL0PRJaBTGioBJsvBDZ_Atjk82Dxz2C9d0F6G9jAZMJ-kMmbwW47BhbRAJEMt0EiDFNBbIdPYNMyQC5hKAaZ82A1SCThGIi9oAmk5gDY5QmsYL3KIPYBiL0NYHdOAIfJb5AuBwGQXxY8AYfPdrD5QWBzZMCufQhWvwOkUoENbMJRsK91wSKcIBEHNnD4XAH7OhgcYmvA4nrWAATQkG8%2C){ target="_blank" } | Дешевая универсальная смазка |

## Растворители и обезжириватели

| Наименование | Ссылки | Комментарий |
|:------------ |:------:|:----------- |
| Изопропанол | [:material-shopping:](https://www.ozon.ru/product/spirt-izopropilovyy-absolyutirovannyy-99-97-dlya-dezinfektsii-lancet-1-l-1154450032/?asb=RU16isS5zWryRFcEBD88PyMXxj%252B%252BW%252FswaxyShRbL%252FRg%253D&asb2=yRxwGwbtV0-cuh9AHF5kNFVeCoudroMFFbImT7UgDy71QRqEx2p7wveYmrmyS_Mv&avtc=1&avte=2&avts=1698449291&keywords=%D0%B8%D0%B7%D0%BE%D0%BF%D1%80%D0%BE%D0%BF%D0%B0%D0%BD%D0%BE%D0%BB){ target="_blank" } [:material-shopping:](https://www.wildberries.ru/catalog/164292958/detail.aspx){ target="_blank" } | Используется для разбавления БФ-2, промывки направляющих и очистки деталей механики принтера. Для обезжиривания стола не лучший вариант т.к. оставляет тонкую плёнку, к которой липнет не так хорошо, как если стол хорошо обезжирен |
| Обезжириватель | [:material-shopping:](https://www.ozon.ru/product/obezzhirivatel-bystroisparyayushchiysya-1l-welltex-656712566/?asb=I8kWzx0VhHVrlkTHcRBhAJLnArY7fYt%252FZLI28DxzcW4%253D&asb2=vwJL80esrb_PWH9LcT6Zn-V-uuwGQQi5hFgNxIW1dYkKKHth12ujX-7vhzsZTh5T&avtc=1&avte=2&avts=1698449436&keywords=%D0%BE%D0%B1%D0%B5%D0%B7%D0%B6%D0%B8%D1%80%D0%B8%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C+welltex){ target="_blank" } [:material-shopping:](https://market.yandex.ru/search?text=%D0%BE%D0%B1%D0%B5%D0%B7%D0%B6%D0%B8%D1%80%D0%B8%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%20welltex%201%D0%BB&rs=eJwzamIOYDzKyMAgYAskF9zeAyQdVtgASYWGvSDxCSC2w3kQe4EiSE2DHEiNghKI_eAoSFwhCkwW7gKSB1aAZA98B8kqFIHNbAez54LIhMbdIBOY9oHULALpSsgDkQ7rwCZfBpvjYwfSdRWsiwOsSx1MzgC5pAHsHoag_SDbtUF2PdADyTKstAbJ-oH0KjCD3fYfbAIzyMwDvGDyENico2DfGYBtLAGJN2iBffcZZBrDJbAtk8F6GcAulAT5y4Ef5OYF2mBzesCy9WDZXLAfX4LdsAHkOwdzsO0TwXbFgNlHwOJX9sD9_uAkSCRhHziEi_cAAFzViVQ%2C&allowCollapsing=1&local-offers-first=0){ target="_blank" } | Обезжиривает не оставляя следов, может использоваться для промывки направляющих или чистки механики принтера |

[^1]: для печати композитными материалами требуется цельнометаллическое или биметаллическое горло + износостойкое сопло, обычно стальное закалёное или медное с калёной вставкой. Титановые и медные никелированные сопла износостойкими не являются.