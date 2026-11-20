# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

Leigos brasileiros (pt-BR) com Windows 10/11 que querem ver e controlar o celular
Android na tela do computador. Não conhecem terminal, ADB ou "depuração USB" —
esperam um app que guia cada passo. Audiência secundária futura: usuários de macOS
(build ainda não existe).

## Product Purpose

Espelhar a tela de um celular Android no computador (vídeo, controle por mouse e
teclado, áudio) através de um aplicativo com interface guiada — download →
conectar → espelhar — sem exigir conhecimento técnico. Sucesso = um leigo
espelha o celular sem consultar tutorial externo.

## Positioning

Um espelhador Android simples para quem não é técnico. O produto é o Refletunes
com identidade própria; a base técnica (scrcpy) pertence aos créditos e à
documentação legal, não ao marketing. O diferencial concreto: setup assistido
(guia de depuração USB, pareamento Wi-Fi pela própria interface) e tratamento dos
erros que travam iniciantes (ex.: "Depuração USB — Configurações de segurança"
em Xiaomi/MIUI).

## Operating Context

- Distribuição: ZIP portátil Windows (extrair e executar, sem instalador) via
  releases do GitHub; site React/Vite é a porta de download e documentação.
- Fluxo do usuário: site → baixar ZIP → extrair → abrir `refletunes.exe` →
  aba Configurar/Dispositivos → espelhar via cabo ou Wi-Fi.
- Estado atual: beta v0.1.0, build gerado por GitHub Actions, ainda sem teste
  funcional em Windows real (não declarar "testado em Windows").
- Sem telemetria, contas, login, pagamentos ou backend no MVP.

## Capabilities and Constraints

- Somente Android (5.0+). iPhone/iOS não é suportado — o mecanismo (ADB) não
  existe no iOS; o site não deve sugerir compatibilidade com iOS.
- Requisitos do dispositivo: depuração USB (e "Depuração USB — Configurações de
  segurança" em Xiaomi/MIUI para controle); Android 11+ para pareamento Wi-Fi.
- Base: scrcpy v5.0.1 (Apache 2.0) + launcher próprio em Go/Fyne. Preservar
  licenças e créditos; não insinuar vínculo com a Genymobile.
- Não prometer compatibilidade universal; não publicar builds sem teste real.
- Decisões em aberto: versão macOS; assinatura de código (hoje o download é
  não-assinado — documentar, sem orientar bypass de segurança do Windows).

## Brand Commitments

- Nome: Refletunes. Logotipo tipográfico (sem ícones de música/áudio).
- Identidade visual fixada pelo proprietário: fundo #080B14 / #101522, roxo
  #A12BFF, azul #386BFF, ciano #00C8E8 (usados com moderação), texto #FFFFFF /
  #9299AA, bordas #242B3A, tipografia Inter. Referências: Linear, Raycast,
  Vercel, Arc — sem aparência genérica de template ou de "IA".
- Publicidade futura no site é compromisso confirmado: o layout deve reservar
  espaço para blocos de anúncio discretos, sem prejudicar a página atual.
- Tom: confiança e simplicidade; copy em pt-BR direto, sem jargão técnico e sem
  promessas exageradas ("sem complicação" só onde o produto entrega).

## Evidence on Hand

- Teste real concluído (macOS + Xiaomi POCO Android 14): espelhamento e controle
  por mouse funcionam; caso documentado de "espelha mas não controla" resolvido
  pela chave "Depuração USB (Configurações de segurança)".
- Workflow de build Windows reproduzido no GitHub Actions; pacote
  `refletunes-win64.zip` gerado — sem teste funcional em Windows ainda.
- NÃO existem ainda: screenshots reais do launcher em Windows, vídeo de
  demonstração, depoimentos, métricas de uso. Não fabricar nada disso.

## Product Principles

1. O usuário não precisa saber o que é ADB — a interface resolve e explica.
2. Honestidade de maturidade: beta declarado, plataformas indisponíveis ditas
   como tal, compatibilidade real (só Android) sempre explícita.
3. Download exige consentimento: o site nunca inicia download sem dois cliques
   (selecionar plataforma, confirmar) e links vazios nunca habilitam o botão.
4. Créditos e licenças upstream preservados e acessíveis — independência do
   scrcpy/Genymobile sempre declarada.
