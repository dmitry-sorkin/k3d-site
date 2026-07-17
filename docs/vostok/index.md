---
authors:
  - sorkin
icon: material/printer-3d
title: VOSTOK
description: Общее описание конструкции
show_created: false
show_updated: false
show_author: false
---

#

![](./pics/vostok_logo_white.png#only-dark){ .notrounded width="900" }
![](./pics/vostok_logo_black.png#only-light){ .notrounded width="900" }

<video autoplay loop muted playsinline width="900">
  <source src="./pics/index_main.mp4" type="video/mp4">
</video>

<p align="center" style="font-size: 1.3em;">Проект быстрого крупногабаритного IDEX 3D-принтера с открытыми исходниками</p>

---

<div class="image-text-row" markdown>
<div class="image-text-img" markdown>

<video autoplay loop muted playsinline width="300">
  <source src="manual/pics/slicer_configuration/fast_tool_swaps.mp4" type="video/mp4">
</video>

</div>
<div class="image-text-content" markdown>

!!! time "Полное время смены инструмента — меньше секунды!"

Для сравнения, принтеры с AMS-подобными системами меняют материал за ~90 с, двойные экструдеры и тулченджеры за ~10 с, а RatRig V4 IDEX ~7 с[^1].

В отличие от других принтеров, в том числе других IDEX принтеров, VOSTOK умеет прочищать неактивную печатающую голову пока активная ещё печатает. Благодаря этому не надо тратить время на печать черновой башни после смены инструмента.

</div>
</div>

---

<div class="image-text-row reverse" markdown>
<div class="image-text-img" markdown>

<video autoplay loop muted playsinline width="300">
  <source src="./pics/hotend.mp4" type="video/mp4">
</video>

</div>
<div class="image-text-content" markdown>

!!! fire "Производительность хотэнда до 50 мм³/с"

Если хотэнд не тянет, то не важно, насколько быстрая будет механика — принтер будет медленным. Поэтому на VOSTOK используются одни из наиболее производительных хотэндов на рынке — Goliath и CHC XL.

Другим бутылочным горлышком может быть обдув. И с ним у VOSTOK'а тоже всё прекрасно. Компрессор ws7040 вкупе с оптимизированными соплами системы охлаждения настолько производительны, что почти никогда не требуется включать их более чем на 70%.

</div>
</div>

---

<div class="image-text-row" markdown>
<div class="image-text-img" markdown>

<video autoplay loop muted playsinline width="300">
  <source src="./pics/speed.mp4" type="video/mp4">
</video>

</div>
<div class="image-text-content" markdown>

!!! speed "Быстрый несмотря на большую область печати"

| Принтер | Область печати | Макс. ускорения[^2] |
| :------ | :------------: | :-------------: |
| VOSTOK M | 400x250x250 мм | ~11000 мм/с² |
| VOSTOK L | 400x400x400 мм | ~10000 мм/с² |
| VOSTOK XL | 500x500x500 мм | ~8000 мм/с² |
| RatRig V4 IDEX | 400x400x400 мм | ~5000 мм/с² |
| Bambu Lab H2D | 350×320x325 мм | ~5000 мм/с² |

</div>
</div>

---

<div class="image-text-row reverse" markdown>
<div class="image-text-img" markdown>

<video autoplay loop muted playsinline width="330">
  <source src="./pics/mirror.mp4" type="video/mp4">
</video>

</div>
<div class="image-text-content" markdown>

!!! clone "Удвоенная производительность при печати партий"

Мало кто сможет сравниться с VOSTOK в скорости печати партии деталей. Ведь вдобавок к быстрой механике, производительным хотэндам и мощному охлаждению модели, у VOSTOK'а ещё есть зеркальный и повторяющий режимы печати. В них обе печатающие головы работают одновременно, буквально удваивая и так высокую скорость печати принтера.

</div>
</div>

---

<div class="image-text-row" markdown>
<div class="image-text-img" markdown>

![](./pics/motors.png){ width="300" }

</div>
<div class="image-text-content" markdown>

!!! fire "Готовность под термокамеру до 80°C"

- Все моторы, кроме моторов экструдеров, вынесены из отсека печати;
- Тепловое расширение деталей механики не влияет на смещения голов;
- Есть штатные места под термистор термокамеры и отверстия для прохождения кабелей;
- Забор воздуха в систему охлаждения осуществляется из горячей точки зоны печати;
- В области печати используются только теплостойкие компоненты.

</div>
</div>

---

<div class="image-text-row reverse" markdown>
<div class="image-text-img" markdown>

![](pics/printhead_boards.png#only-light){ width="300" }
![](pics/printhead_boards_dark.png#only-dark){ width="300" }

</div>
<div class="image-text-content" markdown>

!!! info "Полная и подробная документация"

VOSTOK — это не набор файлов «как есть, дальше разбирайтесь сами». Это законченный проект с хорошей документацией:

- Спецификация с конкретной информацией о том, что и где купить, чего и сколько напечатать и т.д.;
- Для заказа профилей, деталей стола и панелей зашивки — подробные и понятные чертежи;
- STL-файлы, оптимизированные под печать;
- Полные STEP/Parasolid сборки;
- Конфигурации Klipper под разные платы;
- Пошаговые инструкции: от покупки деталей до настроенного слайсера.


</div>
</div>

---

<div class="image-text-row" markdown>
<div class="image-text-img" markdown>

![](./pics/github.png){ width="300" }

</div>
<div class="image-text-content" markdown>

!!! note "Минимум лицензионных ограничений"

VOSTOK выпускается под лицензией [Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/){ target="_blank" }. Единственное требование лицензии — указать авторство и ссылку на эту страницу. Всё остальное от использования и модификации, до коммерческого использования, не ограничено.

</div>
</div>

---

Уже несколько сотен VOSTOK'ов радуют своих владельцев. Соберите и свой:

<div class="grid cards" markdown>

-   :material-download-box:{ .lg .middle } __Начать сборку__

    ---

    Скачайте необходимые для сборки файлы — спецификацию, CAD сборки, STL файлы, чертежи.

    [:octicons-arrow-right-24: Скачать](./releases/index.md)

-   :lucide-book-open:{ .lg .middle } __Почитать__

    ---

    Узнать о структуре документации и первых шагах сборки ещё до сборки.

    [:octicons-arrow-right-24: С чего начать?](./manual/start.md)

-   :material-frequently-asked-questions:{ .lg .middle } __Остались вопросы__

    ---

    Многие из особенностей проекта лучше описаны в ЧаВо — частые вопросы и ответы на них.

    [:octicons-arrow-right-24: Почитать ЧаВо](./faq.md)

-   :simple-telegram:{ .lg .middle } __Пообщаться__

    ---

    Сообщество проекта наиболее активно в Telegram. Присоединяйтесь

    [:octicons-arrow-right-24: Обсуждение K3D VOSTOK](https://t.me/k3d_vostok){ target="_blank" }

</div>

[^1]: RatRig на стоковой прошивке не умеет прочищать неактивную голову пока активная ещё печатает, и ему требуется печать черновой башенки. Подробнее в [ЧаВо](./faq.md){ target="_blank" }
[^2]: Максимальные ускорения до падения детализации. Вычисляются при калибровке Input Shaping