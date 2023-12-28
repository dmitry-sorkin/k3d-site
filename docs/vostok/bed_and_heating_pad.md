---
icon: material/source-commit-local
title: K3D VOSTOK - Изготовление и сборка стола
description: Инструкция по сборке VOSTOK
---

# Изготовление и сборка стола

## Изготовление теплораспределителя стола

Теплораспределитель стола в большинстве случаев приходится производить самостоятельно. Для этого используется алюминиевый лист толщиной 3мм при габарите наименьшей стороны до 300мм и 4мм при габарите более 300мм. Лист для производства стола лучше заказывать с небольшим припуском и отрезать точно в размер самостоятельно, т.к. рез листа на металлобазе может быть очень не точным или стороны стола могут быть не перпендикулярными.

Чтобы получить правильную геометрию стола, надо правильно разметить его перед резкой:

| Изображение | Описание |
|:-----------:|:-------- |
| ![step_1](./pics/bed_and_heating_pad/how_to_draw_rectangle_s1.png){ width=500} | Выберите базовую сторону (на рисунке выше отмечена зеленым). Обычно лучшим вариантом является длинная прямая сторона листа. В дальнейшем все размеры, включая угловые, будут откладываться только от этой стороны |
| ![step_2](./pics/bed_and_heating_pad/how_to_draw_rectangle_s2.png){ width=500 } | С помощью угольника и карандаша начертите линию, перпендикулярную базовой. Предпочтительно провести линию как можно ближе к одной из боковых сторон заготовки. Если одна из боковых сторон заготовки с достаточной точностью перпендикулярна базовой стороне, то можно использовать её |
| ![step_3](./pics/bed_and_heating_pad/how_to_draw_rectangle_s3.png){ width=500 } | Отмерьте габарит стола вдоль базовой линии (зеленой) и с помощью угольника проведите еще одну линию, перпендикулярную базовой |
| ![step_4](./pics/bed_and_heating_pad/how_to_draw_rectangle_s4.png){ width=500 } | Отмерьте второй габарит стола вдоль линий, перпендикулярных базовой (вдоль синей и розовой), после чего соедините две полученные точки |
| ![step_5](./pics/bed_and_heating_pad/how_to_draw_rectangle_s5.png){ width=500 } | Сравните длины противолежащих сторон полученного прямоугольника, они должны быть равны. Если это так, то сравните еще и длину диагоналей (красные стрелки). Если длина какой-то пары отрезков не равна, то в построение закралась ошибка. Придётся проверять все размеры и углы |
| ![step_6](./pics/bed_and_heating_pad/how_to_draw_rectangle_s6.png){ width=500 } | Отрежьте лишние участки заготовки по нанесенной разметке. Для резки алюминиевого листа можно использовать лобзик с полотном по металлу или ножовку. Чтобы рез был ровным, можно приклеить на заготовку направляющую из какого-нибудь бруска, в который будет упираться подошва лобзика или полотно ножовки. Далее края надо будет обработать от заусенцев, чтобы потом не порезать пальцы |

## Сверление отверстий в теплораспределителе

![](./pics/bed_and_heating_pad/bed_drawing_black.png#only-light)
![](./pics/bed_and_heating_pad/bed_drawing_white.png#only-dark)

Количество отверстий в столе зависит от ваших предпочтений. Для небольших столов, например, 200х200мм, хватит 4 точек крепления, и, соответственно, 4 отверстий по углам. Если габарит какой-то стороны превышает 300мм, то рекомендуется добавить центральные точки вдоль этой стороны. Если обе стороны превышают 300мм, то рекомендуется использовать все 9 точек крепления, включая центральную, так как это позволит бороться с линзовидным изгибом стола от нагрева.

Все отверстия должны быть с зенкованием под винты М3 с потайной головкой. Головка винтов обязательно должна утапливаться ниже поверхности стола т.к. на стол будет помещаться какое-либо покрытие.

## Заказ грелки

В случае, если планируется использовать стол с 4 или 6 крепежными точками, то грелка заказывается по размеру области печати (на 50мм меньше плиты стола по оси Y). Вырезы под винты в таком случае не нужны т.к. винты будут располагаться вокруг грелки.

Если предполагается использование 8 или 9 точек крепления, то необходимо будет написать продавцу грелок, чтобы он сделал вырезы под винты. Термистор при этом стоит располагать как можно ближе к центру грелки.

![](./pics/bed_and_heating_pad/heater_pad_black.png#only-light)
![](./pics/bed_and_heating_pad/heater_pad_white.png#only-dark)

## Cборка стола

![](pics/bed_and_heating_pad/glue_heater_pad.png)

Единственный не очевидный многим момент в наклейке грелки на теплораспределитель - то, что поверхность теплораспределителя надо предварительно обезжирить. Иначе велик шанс, что в каком-то месте грелка отойдёт и перегреется, что в лучшем случае приведет к необходимости переделывать стол, а в худшем - к пожару.

![](pics/bed_and_heating_pad/bed_screws_install.png)

Установите винты М3х50 с потайной головкой в отверстия в теплораспределителе стола и зафиксируйте их резьбовыми втулками. Желательно посадить втулки на резьбовой герметик т.к. если они открутятся, то придётся пересобирать пол стола. При этом случаи, когда эти винты надо было бы снять для обслуживания принтера очень маловероятны.

!!! note "Втулки удобны тем, что одновременно с фиксацией винтов позволяют еще и центрировать пружины. Если втулок нет, то можно зафиксировать винты обычными гайками М3, а пружины центрировать любым другим способом."

---

<table class="navitable">
    <tbody>
        <tr>
            <td><a class="md-button" href="../belts" style="width: 100%; padding-left: 0em; padding-right: 0em;"><span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M20 11v2H8l5.5 5.5-1.42 1.42L4.16 12l7.92-7.92L13.5 5.5 8 11h12Z"></path></svg></span> Установка ремней</a></td>
            <td><a class="md-button" href="https://forum.k3d.tech/t/obsuzhdenie-stati-izgotovlenie-stola/38" target="_blank" style="width: 100%; padding-left: 0em; padding-right: 0em;"><span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17 12V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v14l4-4h10a1 1 0 0 0 1-1m4-6h-2v9H6v2a1 1 0 0 0 1 1h11l4 4V7a1 1 0 0 0-1-1Z"></path></svg></span> Обсуждение статьи</a></td>
            <td><a class="md-button" href="../bed_frame" style="width: 100%; padding-left: 0em; padding-right: 0em;">Рама стола <span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 11v2h12l-5.5 5.5 1.42 1.42L19.84 12l-7.92-7.92L10.5 5.5 16 11H4Z"></path></svg></span></a></td>
        </tr>
    </tbody>
</table>

---

<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><a property="dct:title" rel="cc:attributionURL" href="https://k3d.tech/vostok/">K3D VOSTOK</a> by <span property="cc:attributionName">Dmitry Sorkin</span> is licensed under <a href="http://creativecommons.org/licenses/by/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">CC BY 4.0<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1"></a></p>