# Refletunes

Espelhe e controle seu Android pelo computador com uma configuração simples.

**Este é um projeto independente baseado no [scrcpy](https://github.com/Genymobile/scrcpy).**
Não é afiliado, patrocinado ou endossado pela Genymobile. Veja
[legal/THIRD_PARTY_NOTICES.md](legal/THIRD_PARTY_NOTICES.md).

## Estrutura

```
upstream/             scrcpy v5.0.1 (submodule)
desktop/launcher/     launcher Go + Fyne (refletunes.exe)
docs/                 documentação e decisões de arquitetura
legal/                licenças e avisos de terceiros
.github/workflows/    build Windows automatizado
```

## Desenvolvimento (macOS)

```bash
git submodule update --init
cd desktop/launcher
go build -o launcher .
```

Para rodar o launcher com espelhamento real, copie para `desktop/launcher/`
os binários de um release scrcpy: `scrcpy`, `scrcpy-server`, `adb`,
`scrcpy.png`, `disconnected.png`.

## Build Windows

Gerado pelo GitHub Actions (`.github/workflows/build-windows.yml`):
cross-compila o scrcpy com mingw (método oficial upstream), compila o
launcher, monta o ZIP portátil e gera SHA-256.
