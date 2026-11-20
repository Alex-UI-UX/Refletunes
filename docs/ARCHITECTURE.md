# Arquitetura — Refletunes

## Decisões registradas

| Decisão | Escolha | Motivo |
|---|---|---|
| Base | scrcpy **v5.0.1** como submodule em `upstream/` | rastreável, facilita atualização |
| Launcher | **Go + Fyne** | binário .exe único, cross-compila no GitHub Actions com mingw (mesmo toolchain do build do scrcpy), sem runtime extra para o usuário |
| Pacote | ZIP portátil | alinhado ao formato oficial upstream; launcher.exe + scrcpy.exe + adb.exe + scrcpy-server + DLLs na mesma pasta (modo portátil do scrcpy) |
| Escopo MVP | USB + pareamento Wi-Fi | `adb pair` (Android 11+) implementado no launcher |
| Perfis | Padrão (defaults) / Desempenho (`--max-size=1024 --video-bit-rate=2M --max-fps=30`) | flags reais da v5.0.1, testadas em POCO Android 14 |

## Layout do pacote Windows

```
refletunes.exe        # launcher Go/Fyne
scrcpy.exe            # client upstream
scrcpy-server         # server upstream (enviado ao Android em runtime)
adb.exe, AdbWinApi.dll, AdbWinUsbApi.dll
*.dll                 # SDL3, FFmpeg, libusb, dav1d
scrcpy.png, disconnected.png, icon.ico
LICENSE.txt, THIRD_PARTY_NOTICES.md
```

## Regras de segurança do launcher

- Argumentos do scrcpy vêm **somente** de `Profiles` (lista fixa) + `-s <serial>`
  obtido de `adb devices`. Nunca de texto digitado pelo usuário.
- `adb pair/connect` validam formato `IP:PORTA` antes de executar.
- Sem telemetria, sem rede além do ADB local.
