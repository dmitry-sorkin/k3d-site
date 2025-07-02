---
icon: material/source-commit-local
title: K3D VOSTOK - Сборка приводов оси Y
description: Инструкция по сборке VOSTOK
---

# Сборка приводов оси Y

## Подготовка печатных деталей

| Наименование | Ориентация | Количество | Предпочтительный материал | Комментарий |
|:------------ |:------:|:------:|:-----:|:---- |
| Y motor mount | ![](./pics/y_drive/y_motor_mount_orientation.png){ width="600" } | 2 | `PA6 GF30`, `PP GF30` | Возможно напечатать из других материалов, например, на основе `ABS`, `ASA`, `PC` и подобных. Одну деталь нужно отзеркалить |
| Y motor strengthener | ![](./pics/y_drive/y_motor_strengthener_orientation.png) | 0-2 | Любой жесткий с термостойкостью ≥90°С | Опционально |

При печати деталей в указанной ориентации, поддержки не нужны. Все детали должны быть максимально прочными и жесткими. Поэтому рекомендуется печатать со следующими параметрами:

- Толщина стенок 2-3мм;
- 100% заполнения;
- Не высокая скорость;
- Ширина линий >120% от диаметра сопла;
- Толщина слоя <50% от ширины линий;
- Кайма 3+ мм.

## Сборка привода оси Y

### Подготовка мотора

<figure markdown>
  ![](./pics/y_drive/s9_pulley_install.png){ width="300" }
</figure>

Установите шкив на мотор как показано на изображении выше. Если используются шкив и мотор из спецификации, то можно просто совместить торец шкива и торец вала мотора.

### Сборка привода оси Y

<figure markdown>
  ![](./pics/y_drive/s9_motor_mount_assembly.png){ width="400" }
</figure>

Установите мотор в привод так, чтобы провод выходил вниз (на изображении показано синей стрелкой) и закрепите на 4 винта М3x12.

### Установка привода оси Y

<figure markdown>
  ![](./pics/y_drive/s9_nut_install.png){ width="400" }
</figure>

Заложите по 1 пазовой гайке М4 в указанные на изображении пазы профиля.

<figure markdown>
  ![](./pics/y_drive/s9_drive_install.png){ width="900" }
</figure>

1. Установите винт М4x8 с потайной головкой как показано на изображении;
2. Установите привод на место и закрутите заранее заложенный винт М4x8 с потайной головкой, а так же еще один такой же винт и винт М4x8 с полукруглой головкой.

### Установка усилителя (опционально)

<figure markdown>
  ![](./pics/y_drive/s9_optional_nut_install.png){ width="700" }
</figure>

1. Из мотора выкрутите 2 указанных винта;
2. В указанные на изображении пазы заложите по 1 гайке М4.

<figure markdown>
  ![](./pics/y_drive/s9_strengthener_install.png){ width="900" }
</figure>

1. Прикрутите усилитель к мотору на 2 винта М3х[длина мотора + 2] с потайной головкой;
2. Закрепите усилитель к раме на 2 винта М4x8.

## Сборка зеркального привода

Второй привод оси Y собирается и крепится зеркально описанному без каких-либо отличий.

---

<table class="navitable">
    <tbody>
        <tr>
            <td><a class="md-button" href="../xu_drives" style="width: 100%; padding-left: 0em; padding-right: 0em;"><span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M20 11v2H8l5.5 5.5-1.42 1.42L4.16 12l7.92-7.92L13.5 5.5 8 11h12Z"></path></svg></span> Приводы осей XU</a></td>
            <td><a class="md-button" href="../belts" style="width: 100%; padding-left: 0em; padding-right: 0em;">Установка ремней <span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 11v2h12l-5.5 5.5 1.42 1.42L19.84 12l-7.92-7.92L10.5 5.5 16 11H4Z"></path></svg></span></a></td>
        </tr>
    </tbody>
</table>

---

<div id='discourse-comments'></div>
<meta name='discourse-username' content='DISCOURSE_USERNAME'>

<script type="text/javascript">
  DiscourseEmbed = {
    discourseUrl: 'https://forum.k3d.tech/',
    discourseEmbedUrl: 'https://k3d.tech/vostok/y_drive/',
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