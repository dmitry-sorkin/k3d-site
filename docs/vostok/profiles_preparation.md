---
icon: material/source-commit-local
title: K3D VOSTOK - Подготовка профилей рамы
description: Инструкция по сборке VOSTOK
---

# Подготовка профилей

## Способы соединения профилей

![blind joint](./pics/profile_preparation/blind_joint.png)

Так как основную жесткость раме придают панели зашивки, профили в раме VOSTOK'а могут соединяться несколькими способами, но основным является blind joint. Это соединение предполагает, что винт М6 будет вкручен в торец профиля, а головка винта будет заведена в паз другого профиля, пристыкованного перпендикулярно. Такое соединение требует нарезки резьбы в торцах профилей, а также сверления профилей для доступа ключом к шлицу винта.

## Нарезка резьбы в торцах профилей

Резьбу М6 надо нарезать во всех торцевых отверстиях всех профилей, кроме:

- Балки оси Х;
- Балок рамы стола.

В целом, процесс нарезания резьбы не имеет каких-то подводных камней, так что описывать как это правильно делать не имеет большого смысла. Дам только пару рекомендаций:

1. Лучше использовать наборы из 2-3 метчиков, где каждый нарезает резьбу глубже предыдущего;
2. Ломайте стружку при нарезке резьбы. Один оборот вперед - половина назад, еще один вперед, еще половина назад. Иначе налипшая на метчик стружка помнёт профиль;
3. Не забудьте высыпать стружку из профилей после нарезки.

Процесс нарезки резьбы показан [в этом видео](https://youtu.be/7GxtSSZW7pc?t=55).

## Сверление отверcтий 

![Holes page](./pics/profile_preparation/holes_page.png)

Для доступа ключом к шлицам винтов в соединениях, некоторые профили надо будет просверлить. Состав этих профилей, а также все необходимые координаты для этого процесса рассчитываются автоматически на листе `Сверление профилей` в спецификации (BOM). 

- Отверстия схематично изображены окружностями;
- Диаметр всех отверстий 6мм;
- Выше схемы профиля находится размерная цепь, начинающаяся от торца профиля;
- Все размеры указаны в миллиметрах (в сантиметре 10мм, в метре 1000мм);
- Если между двумя отверстиями указано расстояние 0 мм, то это одно отверстие;
- Профили, которые не должны быть просверлены, не упоминаются на этом листе;
- Все профили одного типа сверлятся по одинаковой схеме.

Сверление лучше всего производить на сверлильном станке с предварительным кернением. Если такого оборудования нет, то можно обойтись шуруповертом и печатными кондукторами, которые можно найти [на странице с инструментами](./tools.md).

!!! warning ""
    Отклонение по положение или углу отверстий может привести к невозможности подлезть ключом к шлицу винта. В таком случае отверстие придётся расширять напильником или шарошкой.

---

<table class="navitable">
    <tbody>
        <tr>
            <td><a class="md-button" href="../guide_index" style="width: 100%; padding-left: 0em; padding-right: 0em;"><span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M20 11v2H8l5.5 5.5-1.42 1.42L4.16 12l7.92-7.92L13.5 5.5 8 11h12Z"></path></svg></span> С чего начать</a></td>
            <td><a class="md-button" href="../portal" style="width: 100%; padding-left: 0em; padding-right: 0em;">Сборка портала <span class="twemoji"><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M4 11v2h12l-5.5 5.5 1.42 1.42L19.84 12l-7.92-7.92L10.5 5.5 16 11H4Z"></path></svg></span></a></td>
        </tr>
    </tbody>
</table>

<script type="text/javascript">
  DiscourseEmbed = {
    discourseUrl: 'https://forum.k3d.tech/',
    discourseEmbedUrl: 'https://k3d.tech/vostok/profiles_preparation/',
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