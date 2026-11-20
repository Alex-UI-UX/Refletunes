# AUDIT.md — Auditoria Fase 0 (scrcpy upstream)

**Data:** 2026-11-20
**Repositório auditado:** `upstream/` = clone de https://github.com/Genymobile/scrcpy
**Status:** somente leitura — nenhum arquivo do upstream foi modificado.

---

## 1. Resumo do projeto upstream

O scrcpy ("screen copy") espelha vídeo e áudio de dispositivos Android via USB ou
TCP/IP e permite controle por teclado/mouse, sem root e sem app instalado no
aparelho. Arquitetura em duas partes:

- **client** (`app/`): programa nativo em C (meson/ninja), usa SDL3 para janela/input,
  FFmpeg para decodificar vídeo/áudio, libusb para modo OTG/HID, e se comunica com
  o dispositivo via ADB.
- **server** (`server/`): componente Java/Android (Gradle), compilado para
  `scrcpy-server`, que é *enviado* ao dispositivo em tempo de execução. Não fica
  instalado no aparelho.

Requisitos do dispositivo: Android 5.0+ (API 21); encaminhamento de áudio exige
Android 11+ (API 30); câmera exige Android 12+.

## 2. Versão/commit recomendado

- **Tag estável mais recente:** `v5.0.1` (HEAD do `master` já está em 5.0.1).
- O projeto mantém `master` = última release e `dev` = desenvolvimento.
- Recomendação: **fixar a tag `v5.0.1`** (`git checkout v5.0.1`) como base da
  distribuição Refletunes. O server prebuilt oficial para v5.0.1 existe e tem
  SHA-256 publicado (`764eb6f7...`), permitindo build do client sem Android SDK.

## 3. Processo de build — macOS (iMac Intel)

Duas opções documentadas em `doc/macos.md` / `doc/build.md`:

### A. Usar o release oficial (validação rápida, recomendado p/ primeiro teste)

```bash
# baixar scrcpy-macos-x86_64-v5.0.1.tar.gz do GitHub Releases e extrair
brew install --cask android-platform-tools   # fornece adb no PATH
```

### B. Compilar do fonte

```bash
brew install sdl3 ffmpeg libusb pkg-config meson
# para compilar também o server:
brew install openjdk@17
export JAVA_HOME="$(/usr/libexec/java_home --version 1.17)"
export PATH="$JAVA_HOME/bin:$PATH"
export ANDROID_SDK_ROOT=~/Library/Android/sdk   # requer Android SDK instalado

meson setup x --buildtype=release --strip -Db_lto=true
ninja -Cx        # NUNCA como root/sudo
./run x          # executa sem instalar
```

Alternativa sem Android SDK nem Java: baixar o `scrcpy-server-v5.0.1` prebuilt e
usar `-Dprebuilt_server=/caminho/scrcpy-server` no `meson setup`.

## 4. Processo de build — Windows

Método oficial (e usado nas releases upstream): **cross-compile a partir de
Linux Debian** com mingw-w64 — não é necessário Windows para compilar.

```bash
sudo apt install mingw-w64 mingw-w64-tools libz-mingw-w64-dev openjdk-17-jdk
./release.sh        # gera win32 e win64 em dist/ (release/release.sh)
# ou apenas win64:
./release/build_windows.sh 64
```

O script `build_windows.sh`:
1. `app/deps/adb_windows.sh` — baixa **platform-tools 37.0.1** do Google
   (SHA-256 verificado) e extrai `adb.exe`, `AdbWinApi.dll`, `AdbWinUsbApi.dll`.
2. Compila as deps do zero via `app/deps/*.sh`:
   - FFmpeg 9.0.2 (build customizado, ver §5)
   - SDL3 3.4.18
   - dav1d 1.5.4 (decoder AV1)
   - libusb 1.0.30
3. `meson setup` com `cross_win64.txt`, `-Dportable=true`, `-Dcompile_server=false`.
4. Monta `dist/` com `scrcpy.exe`, DLLs, adb, `scrcpy.png`, `icon.ico`,
   `scrcpy-noconsole.vbs`, `open_a_terminal_here.bat` e `LICENSE.txt`.

Resultado: um **ZIP portátil** (sem instalador) — basta extrair e rodar. Isso se
alinha perfeitamente com o plano: o launcher Refletunes pode ser adicionado
dentro desse ZIP como um `.exe`/script separado que chama `scrcpy.exe` com
argumentos validados.

### GitHub Actions

O upstream já possui `.github/workflows/release.yml` que roda exatamente esse
processo em `ubuntu-latest` — **confirmado que compila em runners gratuitos do
GitHub**. Podemos adaptar esse workflow em vez de escrever um do zero.
Observações:
- Repositório público: minutos de Actions ilimitados. Privado: cota de
  ~2.000 min/mês no plano gratuito — um build completo consome dezenas de
  minutos (compila FFmpeg etc.), convém cachear `app/deps/work/`.
- O server exige JDK 17 + Android SDK/Gradle no CI (upstream usa zulu 17 +
  gradle nativo). Alternativa: usar o `scrcpy-server` prebuilt da release
  v5.0.1 com SHA-256 verificado, simplificando o workflow (sem SDK Android).

## 5. Licenças e obrigações

| Componente | Licença | Obrigação ao redistribuir |
|---|---|---|
| scrcpy (client+server) | **Apache 2.0** | incluir LICENSE; indicar modificações; sem uso de marcas Genymobile como endosso |
| FFmpeg 9.0.2 (build do script) | **LGPL 2.1+** — o configure NÃO usa `--enable-gpl` e linka como DLL compartilhada | incluir texto LGPL; oferta de fontes/código correspondente; DLLs dinâmicas ok |
| SDL3 3.4.18 | zlib | incluir aviso de copyright |
| dav1d 1.5.4 | BSD 2-Clause | incluir aviso |
| libusb 1.0.30 | LGPL 2.1 | mesmas obrigações do FFmpeg |
| adb / platform-tools 37.0.1 | Termos do Android SDK (Google) | o upstream já redistribui; **verificar os termos vigentes** antes da nossa redistribuição — ponto de revisão legal |
| Ícones do scrcpy (`app/data/scrcpy.svg`, `.png`, `icon.ico`, `disconnected.png`) | Apache 2.0 (parte do repo) | podemos usar como base, mas **é marca visual do upstream** — devemos criar identidade própria e não apresentar o ícone do scrcpy como nosso |
| `assets/screenshot-debian-600.jpg` | parte do repo (Apache 2.0 presume-se, sem aviso separado) | usar apenas screenshots próprios no site |

Não existe arquivo `NOTICE` no repositório → não há obrigação de reproduzir
NOTICE, mas a Apache 2.0 exige preservar copyright e indicar arquivos
modificados.

**Entregáveis legais obrigatórios no pacote:**
- `LICENSE.txt` (Apache 2.0) — já copiado pelo script de release;
- `THIRD_PARTY_NOTICES.md` listando FFmpeg, SDL3, dav1d, libusb, platform-tools
  com versões e licenças;
- textos das licenças LGPL/BSD/zlib (as DLLs trazem obrigação de oferta de
  fontes — linkar para os repositórios/tags exatos);
- aviso de independência (Genymobile).

## 6. Requisitos para o launcher (viabilidade confirmada)

- O pacote Windows é portátil e o `scrcpy.exe` aceita argumentos de linha de
  comando (`scrcpy --help`, `doc/*`). O launcher pode ser um processo externo
  que: executa `adb devices` (o `adb.exe` já vem no pacote), interpreta o estado
  (`device` / `unauthorized` / nenhum) e lança `scrcpy.exe` com flags fixas de
  uma lista permitida.
- Flags de qualidade reais na v5.x (confirmar no `--help` da versão fixada antes
  de usar): `-m/--max-size`, `-b/--video-bit-rate`, `--max-fps`.
- Nenhuma modificação no código C é necessária para o MVP do launcher → mantém
  compatibilidade total com upstream.
- Opções simples de launcher: script `.bat` (mínimo, feio), **Python+Tkinter
  empacotado** ou **ferramenta nativa leve** (Go/Rust/C# WinForms). Decisão em
  aberto (§8) — registrar em `docs/ARCHITECTURE.md` quando decidirmos.

## 7. Riscos

| Risco | Severidade | Mitigação |
|---|---|---|
| Termos de redistribuição do adb/platform-tools | Médio | revisar termos do Android SDK; alternativa: instruir usuário a baixar platform-tools (pior UX) |
| Build Windows exige Linux/mingw — não compila no iMac diretamente | Baixo | GitHub Actions (workflow upstream já prova que funciona); opcionalmente Docker/UTM local |
| Usuário final não sabe ativar depuração USB | Alto (UX) | tutoriais com imagens por fabricante — core do produto |
| Executável sem assinatura Authenticode → aviso SmartScreen | Médio | comunicar honestamente; avaliar certificado só com orçamento |
| FFmpeg build customizado pode quebrar ao atualizar versão | Baixo | fixar tag v5.0.1; atualizar upstream em commits separados |
| SmartScreen/antivírus em launcher novo (desconhecido) | Médio | manter launcher simples, assinar no futuro, VirusTotal |
| Teste funcional Windows exige VM/real | Médio | UTM no Intel Mac ou Windows físico emprestado |

## 8. Decisões que dependem de você

1. **Nome definitivo** ("Refletunes" precisa de busca de domínio/marca).
2. **Fork público vs. referência ao upstream**: recomendo **nosso próprio repo**
   com `upstream/` como remote/submodule e customizações em commits separados —
   facilita rebase quando o scrcpy atualizar.
3. **Tecnologia do launcher** (a decidir na Fase 2, após protótipo): candidatos
   Go (binário único, sem runtime), C# WinForms (nativo Windows) ou Python
   empacotado (mais pesado). Recomendo avaliar Go primeiro.
4. **ZIP portátil vs instalador** — upstream já gera ZIP; instalador (Inno
   Setup) é possível depois. MVP: ZIP + launcher.
5. **Server prebuilt vs compilado no CI** — prebuilt (com SHA-256) simplifica o
   workflow e elimina a necessidade do Android SDK no CI; compilado dá
   rastreabilidade total. Recomendo compilar no CI para integridade da
   distribuição.
6. **Escopo da conexão no MVP**: somente USB (Wi-Fi fica para depois — como no
   plano).

## 9. Próximos passos (Fase 1)

1. `cd upstream && git checkout v5.0.1` (fixar versão)
2. Instalar no iMac: `brew install --cask android-platform-tools` + deps do
   §3-B **ou** baixar o tarball macOS x86_64 oficial para o primeiro teste.
3. Conectar o Android físico, ativar depuração USB, `adb devices`, aceitar RSA.
4. Rodar `scrcpy` e registrar resultado do teste funcional.
5. Configurar repo Refletunes no GitHub + adaptar workflow de release para
   reproduzir o build win64 sem customizações.
