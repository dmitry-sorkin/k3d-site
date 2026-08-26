---
authors:
  - sorkin
icon: lucide/package-open
title: VOSTOK - Подготовка
description: Гайд по подготовке деталей к сборке 3D принтера
---

# Подготовка к сборке

## Проверка деталей

Список того, что следует проверить **перед** началом сборки принтера:

1. Длина профилей не должна отличаться от заданной в чертежах более чем на 1мм. Особенно критично сравнить длину стоек оси Z. Если профили 2020 V-slot будут заметно отличаться по длине от других стоек, то это может привести к изгибу портала;
2. На профилях не должно быть замятий и прочих повреждений. Особенно это касается 2020 V-slot профилей т.к. замятие на них почти гарантированно приведёт к проблемам в работе оси Z;
3. Профили и труба оси X должны быть прямыми. Просто сложите их вместе, и сразу увидите если какие-то из них кривые;
4. Труба оси Х должна иметь правильную длину и **толщину стенки**. Очень много продавцов, включая соберизавод, любят занижать толщину стенки вплоть до -0.5мм или прислать трубу на несколько миллиметров длиннее\короче. Если такое случилось с вами, возвращайте эту трубу и ищите соответствующую спецификации;
5. Обязательно проведите аудит покупных, печатных и стандартных изделий перед началом сборки.

## Смятая резьба

<div class="image-text-row reverse" markdown>
<div class="image-text-img" markdown>

![](pics/parts_postprocessing/profile_threads.png){ width="250" .notrounded }

</div>
<div class="image-text-content" markdown>

Нарезать резьбу и сверлить отверстия надо аккуратно и не торопясь. Если вы поторопитесь и сомнёте резьбу в каком-то из профилей, то потом при эксплуатации принтера винт может вырвать из профиля, и рама в этой точке ослабнет. Если это произойдёт с какой-то из стоек, то заменить профиль будет достаточно легко. А вот заменить какой-либо из профилей портала после сборки - уже задача, требующая разборки половины принтера. Поэтому сорвали резьбу или каким-то другим образом повредили профиль портала - проще заменить его до сборки принтера.

</div>
</div>

## Искривлённые детали

<div class="image-text-row reverse" markdown>
<div class="image-text-img" markdown>

![](pics/parts_postprocessing/warped_part.png){ width="250" .notrounded }

</div>
<div class="image-text-content" markdown>

Часто при печати инженерными материалами, имеющими высокую усадку, детали получаются искривлёнными. Для некоторых деталей, например, декоративных панелей, это не критично, и при желании вы можете использовать их. Но несущие детали, в особенности каретки оси Y, приводы осей XWY, натяжители ремней, детали подающего механизма и т.д., должны быть ровными. Иначе после сборки могут начаться проблемы - закусывание рельс, поджирание ремня и т.д.

!!! tip "Можно попробовать спасти искривившуюся деталь. Положите деталь на стол принтера в том положении, в котором она печаталась, нагрейте стол до температуры печати, и продержите так несколько часов. Работает не всегда, но иногда помогает и деталь становится ровной"

</div>
</div>

## Привалочные плоскости

<div class="image-text-row reverse" markdown>
<div class="image-text-img" markdown>

![](pics/parts_postprocessing/seating_surface.png){ width="250" .notrounded }

</div>
<div class="image-text-content" markdown>

Даже очень хорошо настроенный принтер не может выдать идеально плоские поверхности. Поэтому перед сборкой рекомендуется осмотреть напечатанные детали на предмет слегка выступающих углов или других геометрических дефектов на привалочных плоскостях. Дефекты стоит обработать надфилем, но без фанатизма - достаточно просто ровной поверхности без сильно выпирающих углов и прочих подобных элементов.

</div>
</div>

## Жертвенные слои

Для того, чтобы цекованные отверстия на деталях печатались без необходимости ставить мосты под них, они закрыты мембранами толщиной в 1-2 слоя. Перед сборкой принтера эти мембраны необходимо удалить.

!!! note "Проще всего удалить мембрану сверлом подходящего диаметра или надфилем круглого сечения. Но в большинстве мест может справиться и обычный канцелярский нож"

### V9-1142

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/v9-1142_2.png){ .notrounded }
- ![](pics/parts_postprocessing/v9-1142.png){ .notrounded }

</div>

### V9-1143

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/v9-1143.png){ .notrounded }
- ![](pics/parts_postprocessing/v9-1143-2.png){ .notrounded }

</div>

### V9-11461

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/v9-11461.png){ .notrounded }

</div>

### V9-121 и V9-141

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/v9-121.png){ .notrounded }

</div>

### V9-141 и V9-151

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/v9-141.png){ .notrounded }

</div>

### V9-4441

!!! note "Особенностью этой детали является то, что, если вы не собираетесь устанавливать разъёмы в какое-то из отверстий, то мембрану из него убирать не надо т.к. она будет выступать в роли заглушки"

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/v9-4441.png){ .notrounded }

</div>

### V9-563

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/v9-563.png){ .notrounded }

</div>

## Резьбовые втулки

Те соединения, которые может потребоваться неоднократно размыкать после сборки принтера, выполнены на вплавляемых резьбовых втулках. Отверстия под втулки во всех деталях имеют диаметр 4.3мм, что должно хорошо подходить для резьбовых втулок, указанных в спецификации.

!!! warning "В большинстве мест критически важно, чтобы втулки после вплавления не выступали за пределы поверхности, и не создавали выступающих горбов из окружающего пластика. Поэтому после вплавления втулок проверяйте, всё ли ровно. Если нет, то обработайте поверхность ножом или надфилем"

### V9-1111

`M3x3 - 2шт. | M3x5 - 8шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-1110.png){ .notrounded }
- ![](pics/parts_postprocessing/inserts_v9-1110-2.png){ .notrounded }

</div>

### V9-1142

`M3x3 - 10шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-1142.png){ .notrounded }
- ![](pics/parts_postprocessing/inserts_v9-1142-2.png){ .notrounded }

</div>

### V9-1143

`M3x5 - 4шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-1143.png){ .notrounded }

</div>

### V9-141 и V9-151

`M3x3 - 2шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-140.png){ .notrounded }

</div>

### V9-3111

`M3x5 - 5шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-3111.png){ .notrounded }

</div>

### V9-3212

`M3x5 - 1шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-3212.png){ .notrounded }

</div>

### V9-4151

`M3x3 - 2шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-4151.png){ .notrounded }

</div>

### V9-4361

`M3x3 - 2шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-4361.png){ .notrounded }

</div>

### V9-4362

`M3x3 - 2шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-4362.png){ .notrounded }

</div>

### V9-623

`M3x3 - 2шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-623.png){ .notrounded }

</div>

### V9-625

`M3x3 - 3шт.`

<div class="grid cards no-border no-gap cols-2" markdown>

- ![](pics/parts_postprocessing/inserts_v9-625.png){ .notrounded }

</div>

## Рассверливание отверстий

Из-за особенностей 3D печати, диаметр отверстия на напечатанной детали всегда будет заметно отличаться от диаметра этого отверстия в модели. Для компенсации этого, в моделях диаметр отверстий увеличен на 0.2-0.3 миллиметра. В большинстве случаев это работает хорошо.

Если в вашем случае диаметр каких-то отверстий оказался заметно больше необходимого, то детали придётся перепечатать. Если меньше, то не зазорно рассверлить отверстия до необходимого диаметра. Только перед этим не забудьте перепроверить по модели, что отверстие действительно должно быть диаметром больше, чем получилось у вас.

Также есть ряд деталей, отверстия на которых необходимо пройти сверлом в обязательном порядке:

### V9-11411

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/holes_v9-11411.png){ .notrounded }

</div>

### V9-3211

<div class="grid cards no-border cols-2" markdown>

- ![](pics/parts_postprocessing/holes_v9-3211.png){ .notrounded }

</div>