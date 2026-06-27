---

authors:
  - sorkin
icon: octicons/gear-16
title: K3D VOSTOK - Сборка Parasolid
description: Скачать сборку Parasolid
show_created: false
show_updated: false
show_author: false

---

Сейчас вы будете перенаправлены на скачивание файла. Если этого не произошло автоматически, перейдите по ссылке вручную:

[:octicons-package-16: Скачать](https://k3d.tech/vostok/releases/latest/assembly.x_t.7z){ .download-button data-file="assembly.x_t.7z" }

<script>
  const fileUrl = "https://k3d.tech/vostok/releases/latest/assembly.x_t.7z";
  const fileName = "assembly.x_t.7z";
  fetch("https://k3d.tech/vostok/releases/latest/version.md")
    .then(r => r.text())
    .then(v => {
      const version = v.trim();
      const downloadName = `vostok_${version}_${fileName}`;
      const a = document.createElement("a");
      a.href = fileUrl;
      a.download = downloadName;
      document.body.appendChild(a);
      a.click();
      a.remove();
      const link = document.querySelector("a.download-button");
      if (link) link.setAttribute("download", downloadName);
    })
    .catch(() => {
      window.location.href = fileUrl;
    });
</script>
