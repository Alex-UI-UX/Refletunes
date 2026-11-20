# Avisos de terceiros — Refletunes

O Refletunes é um projeto independente baseado no [scrcpy](https://github.com/Genymobile/scrcpy).
Não é afiliado, patrocinado ou endossado pela Genymobile.

## Componentes redistribuídos (pacote Windows)

| Componente | Versão | Licença | Fonte |
|---|---|---|---|
| scrcpy (client + server) | 5.0.1 | Apache License 2.0 | https://github.com/Genymobile/scrcpy |
| FFmpeg | 9.0.2 | LGPL 2.1+ | https://ffmpeg.org |
| SDL3 | 3.4.18 | zlib | https://libsdl.org |
| dav1d | 1.5.4 | BSD 2-Clause | https://code.videolan.org/videolan/dav1d |
| libusb | 1.0.30 | LGPL 2.1 | https://libusb.info |
| Android platform-tools (adb) | 37.0.1 | Android SDK License | https://developer.android.com/tools/releases/platform-tools |
| Fyne (launcher) | 2.8.2 | BSD 3-Clause | https://fyne.io |

## Obrigações

- **Apache 2.0 (scrcpy):** cópia da licença incluída em `LICENSE.txt`; arquivos
  modificados são identificados; marcas e logos da Genymobile não são usados
  como identidade do Refletunes.
- **LGPL (FFmpeg, libusb):** distribuídos como bibliotecas dinâmicas (DLL); o
  código-fonte correspondente está disponível nos repositórios oficiais acima,
  nas versões indicadas — obtidos e compilados pelos scripts em
  `upstream/app/deps/` sem `--enable-gpl`.
- **Android platform-tools:** redistribuído da mesma forma que o release
  oficial do scrcpy; revisar os termos do Android SDK antes do lançamento
  público.

Este documento é informativo e não substitui revisão jurídica antes do
lançamento público.
