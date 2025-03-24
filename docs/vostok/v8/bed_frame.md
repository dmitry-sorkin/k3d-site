---
icon: material/source-commit-local
title: K3D VOSTOK - Сборка рамы стола
description: Инструкция по сборке VOSTOK
---

# Сборка рамы стола

## Варианты рамы стола

Вариант рамы стола зависит от размера области печати, количества стоек и приводов оси Z, а также количества точек крепления стола к раме. Основными являются 2 варианта:

| Вариант с 1 поперечной балкой | Вариант с 2 поперечными балками |
| :---------------------------: | :-----------------------------: |
| ![variant 1](./pics/bed_frame/bed_frame_variant_1.png) | ![variant 2](./pics/bed_frame/bed_frame_variant_2.png) |

Вариант с 1 поперечной балкой обычно собирается с 4 или 6 точками крепления стола т.к. размещению центральных точек мешает поперечная балка. Также этот вариант может использоваться только при сборке VOSTOK'а с 2 стойками оси Z. Иными словами, это упрощенный вариант для принтеров с относительно небольшой областью печати по оси Y, примерно до 300мм.

Вариант с 2 поперечными балками является более универсальным. Он поддерживает любое количество точек крепления стола, от 4 до 9. Стол с такой рамой может быть установлен при любом количестве стоек и приводов оси Z. Но он является немного более сложным, поэтому предполагается, что он будет использоваться при габарите области печати более 300мм по оси Y.

## Сборка варианта с 1 поперечной балкой

### Подготовка печатных деталей

| Наименование | Ориентация | Количество | Предпочтительный материал |
|:------------ |:------:|:------:|:-----:|
| Bed arm end | ![bed_arm_end_orientation](./pics/bed_frame/bed_arm_end_orientation.png){ width=200 } | По кол-ву точек крепления стола | Любой жёсткий с термостойкостью выше 80°C |
| Bed arm knob | ![bed_arm_knob_orientation](./pics/bed_frame/bed_knob_orientation.png){ width=200 } | По кол-ву точек крепления стола | Любой жёсткий с термостойкостью выше 80°C |
| 2020 to 2060 mount | ![2020_to_2060_mount_orientation](./pics/bed_frame/2020_to_2060_mount_orientation.png){ width=200 } | По 2 на продольную балку | Любой жёсткий, кроме `PLA` |

Все детали печатаются без поддержек и не требуют какой-то постобработки, кроме удаления юбки.

### Сборка плеча

<figure markdown>
  ![](./pics/bed_frame/arm_end_install.png){ width="700" }
</figure>

Концевые детали (bed arm end) устанавливаются в 10мм от края профиля как показано на картинке выше, и крепятся на 1 винт М4х8 с потайной головкой и 1 пазовую гайку М4 каждая.

<figure markdown>
  ![](./pics/bed_frame/2020_to_2060_mount_install.png){ width="500" }
</figure>

Два крепления профилей 2020 к 2060 устанавливаются в центрах плеч и крепятся на 2 винта М4х8 с полукруглой или цилиндрической головкой и 2 пазовые гайки М4 каждое. На данном этапе винты затягивать не надо.

### Сборка рамы

<figure markdown>
  ![](./pics/bed_frame/bed_arm_install.png){ width="500" }
</figure>

Прикрепите плечи рамы стола к профилю 2060 на 6 винтов М4х8 с полукруглой или цилиндрической головкой и 6 пазовых гаек М4 на каждое. Затягивать крепления на данном этапе не надо.

![](./pics/bed_frame/bed_frame_install.png){ width="1000" }

1. Положите стол принтера на любую горизонтальную поверхность винтами вверх;
2. Наденьте пружины на винты;
3. Подстраивая положение плеч на поперечной балке, установите раму на винты и добейтесь того, чтобы поперечная балка была центрирована относительно стола;
4. Затяните все крепления.

![](./pics/bed_frame/bed_knob_install.png){ width="850" }

Заложите в барашки по 1 гайке М3 и закрутите получившуюся конструкцию на винты стола.

## Сборка варианта с 2 поперечными балками

### Подготовка печатных деталей

| Наименование | Ориентация | Количество | Предпочтительный материал |
|:------------ |:------:|:------:|:-----:|
| Bed arm end | ![bed_arm_end_orientation](./pics/bed_frame/bed_arm_end_orientation.png){ width=200 } | По кол-ву точек крепления стола | Любой жёсткий с термостойкостью выше 80°C |
| Bed arm knob | ![bed_arm_knob_orientation](./pics/bed_frame/bed_knob_orientation.png){ width=200 } | По кол-ву точек крепления стола | Любой жёсткий с термостойкостью выше 80°C |
| 2020 to 2020 mount | ![2020_to_2020_mount_orientation](./pics/bed_frame/2020_to_2020_mount_orientation.png){ width=200 } | По 2 на продольную балку | Любой жёсткий, кроме `PLA` |

Все детали печатаются без поддержек и не требуют какой-то постобработки, кроме удаления юбки.

### Сборка плеча

<figure markdown>
  ![](./pics/bed_frame/bed_arm_assembled.png){ width="700" }

  ![](./pics/bed_frame/arm_end_install.png){ width="700" }
</figure>

Концевые детали (bed arm end) устанавливаются в 10мм от края профиля как показано на картинке выше, а также в центре плеча. Каждая деталь крепится на 1 винт М4х8 с потайной головкой и 1 пазовую гайку М4.

### Сборка основания рамы

<figure markdown>
  ![](./pics/bed_frame/bed_frame_square_assembly.png){ width="850" }
</figure>

Расположите профили рамы стола на ровной горизонтальной поверхности и соедините между собой на винты М6х12 как показано на изображении выше. Чтобы проверить, что всё собрано ровно, промерьте диагонали полученного портала, они должны быть равны.

<figure markdown>
  ![](./pics/bed_frame/2020_to_2020_mount_install.png){ width="1000" }
</figure>

Установите крепления профилей 2020 друг к другу по 3 на каждую поперечную балку основания рамы. Закрепите каждое крепление на 4 винта М4х8 и 4 пазовые гайки М4. Не затягивайте винты на данном этапе, детали должны сохранять подвижность.

### Сборка рамы

<figure markdown>
  ![](./pics/bed_frame/bed_arm_install_v2.png){ width="1000" }
</figure>

Установите плечи на основание рамы и закрепите их на винты М4х8 и пазовые гайки М4. Не затягивайте винты на данном этапе, детали должны сохранять подвижность.

<figure markdown>
  ![](./pics/bed_frame/bed_frame_install_v2.png){ width="850" }
</figure>

1. Положите стол принтера на любую горизонтальную поверхность винтами вверх;
2. Наденьте пружины на винты;
3. Подстраивая положение плеч на поперечной балке, установите раму на винты и добейтесь того, чтобы поперечная балка была центрирована относительно стола;
4. Затяните все крепления.

!!! note "В текущей версии конструкции концевая деталь плеча может блокировать доступ к одному из винтов крепления 2020 профилей друг к другу. Пока это не исправлено, вы можете либо не устанавливать эти винты т.к. 3 оставшихся будет более чем достаточно, либо частично разобрать раму чтобы получить доступ для этих винтов, затянуть их и собрать всё в обратном порядке"

<figure markdown>
  ![](./pics/bed_frame/bed_knob_install_v2.png){ width="1000" }
</figure>

Заложите в барашки по 1 гайке М3 и закрутите получившуюся конструкцию на винты стола.

---

<table class="navitable">
    <tbody>
        <tr>
            <td><a class="md-button" href="../bed_and_heating_pad/" style="width: 100%; padding-left: 0em; padding-right: 0em;"><span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M20 11v2H8l5.5 5.5-1.42 1.42L4.16 12l7.92-7.92L13.5 5.5 8 11h12Z"></path></svg></span> Изготовление стола</a></td>
            <td><a class="md-button" href="../z_drive" style="width: 100%; padding-left: 0em; padding-right: 0em;">Ось Z <span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 11v2h12l-5.5 5.5 1.42 1.42L19.84 12l-7.92-7.92L10.5 5.5 16 11H4Z"></path></svg></span></a></td>
        </tr>
    </tbody>
</table>

---

<div id='discourse-comments'></div>
<meta name='discourse-username' content='DISCOURSE_USERNAME'>

<script type="text/javascript">
  DiscourseEmbed = {
    discourseUrl: 'https://forum.k3d.tech/',
    discourseEmbedUrl: 'https://k3d.tech/vostok/bed_frame/',
    // className: 'CLASS_NAME',
  };

  (function() {
    var d = document.createElement('script'); d.type = 'text/javascript'; d.async = true;
    d.src = DiscourseEmbed.discourseUrl + 'javascripts/embed.js';
    (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(d);
  })();
</script>

---

<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><a property="dct:title" rel="cc:attributionURL" href="https://k3d.tech/vostok/">K3D VOSTOK</a> by <span property="cc:attributionName">Dmitry Sorkin</span> is licensed under <a href="http://creativecommons.org/licenses/by/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">CC BY 4.0<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1"></a></p>