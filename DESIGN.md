---
name: Refletunes
description: Espelhador de Android no computador — simples, direto, para leigos.
colors:
  bg-abyss: "#080b14"
  bg-surface: "#101522"
  accent-purple: "#a12bff"
  accent-blue: "#386bff"
  accent-solid: "#2d5de6"
  accent-cyan: "#00c8e8"
  text-primary: "#ffffff"
  text-muted: "#9299aa"
  text-faint: "#7d8496"
  tag-suggested: "#8ea5ff"
  border-hairline: "#242b3a"
typography:
  display:
    fontFamily: '"Inter Variable", Inter, system-ui, -apple-system, "Segoe UI", sans-serif'
    fontSize: "clamp(38px, 5.2vw, 62px)"
    fontWeight: 650
    lineHeight: 1.04
    letterSpacing: "-0.035em"
  headline:
    fontFamily: '"Inter Variable", Inter, system-ui, -apple-system, "Segoe UI", sans-serif'
    fontSize: "clamp(28px, 3.4vw, 38px)"
    fontWeight: 650
    lineHeight: 1.1
    letterSpacing: "-0.03em"
  title:
    fontFamily: '"Inter Variable", Inter, system-ui, -apple-system, "Segoe UI", sans-serif'
    fontSize: "19px"
    fontWeight: 600
    letterSpacing: "-0.02em"
  body:
    fontFamily: '"Inter Variable", Inter, system-ui, -apple-system, "Segoe UI", sans-serif'
    fontSize: "15px"
    fontWeight: 400
    lineHeight: 1.55
  label:
    fontFamily: '"Inter Variable", Inter, system-ui, -apple-system, "Segoe UI", sans-serif'
    fontSize: "13.5px"
    fontWeight: 550
rounded:
  sm: "6px"
  md: "8px"
  lg: "10px"
  pill: "99px"
spacing:
  xs: "8px"
  sm: "12px"
  md: "24px"
  lg: "56px"
  section: "88px"
components:
  button-primary:
    backgroundColor: "{colors.accent-solid}"
    textColor: "{colors.text-primary}"
    rounded: "{rounded.md}"
    height: "44px"
    padding: "0 20px"
  button-primary-hover:
    backgroundColor: "#4a79ff"
  button-ghost:
    backgroundColor: "transparent"
    textColor: "#e6e8ee"
    rounded: "{rounded.md}"
  platform-card:
    backgroundColor: "{colors.bg-abyss}"
    textColor: "{colors.text-primary}"
    rounded: "{rounded.lg}"
    padding: "24px"
---

# Design System: Refletunes

## Overview

**Creative North Star: "O Espelho Nítido"**

Clareza e fidelidade acima de tudo: a página existe para levar a tela do celular
ao computador sem ruído. O design é escuro porque o produto vive à noite, na
mesa, ao lado do monitor — e porque a moldura escura faz a tela espelhada brilhar.
Componentes são de ornamento mínimo absoluto: tipografia e espaço carregam a
hierarquia; a cor aparece só onde a ação acontece.

Anti-referência confirmada: aparência genérica de template/IA — gradientes
decisivos em larga escala, glow neon, cards iguais empilhados, ilustrações
abstratas, texto de marketing inflado.

**Key Characteristics:**
- Fundo quase-preto azulado com uma única superfície elevada
- Azul sólido como voz de ação; roxo e ciano apenas em micro-detalhes
- Inter como única família, com tracking apertado nos títulos
- Bordas hairline, raios moderados, nada flutuando
- Espaço negativo generoso — a página é curta de propósito

## Colors

Uma paleta escura e contida: a cor existe para sinalizar ação e estado, nunca
como atmosfera.

### Primary
- **Azul Ação** (#2d5de6): o único preenchimento de CTA. Sólido — gradiente de
  marca nunca preenche controles.
- **Azul Marca** (#386bff): detalhes de estado (ponto de detecção, check,
  seleção de card via borda gradiente fina), ícones de passo.

### Tertiary
- **Roxo Marca** (#a12bff): reservado a micro-acentos (metade do gradiente de
  borda de seleção, brilho sutil do dispositivo no mockup). Nunca como fundo ou
  texto funcional.
- **Ciano Sinal** (#00c8e8): apenas o ponto de transmissão no mockup e o dot de
  detecção — "sinal" literal, não decoração.
- **Azul Sugestão** (#8ea5ff): tag "Sugerido para você" — sugestão, não ordem.

### Neutral
- **Abismo** (#080b14): fundo da página e dos cards.
- **Superfície** (#101522): faixa da seção de download e hover discretos.
- **Hairline** (#242b3a): todas as bordas e divisores.
- **Texto** (#ffffff) / **Mutado** (#9299aa) / **Fraco** (#7d8496): hierarquia
  de leitura; o fraco é o piso de contraste AA (≥4.5:1).

### Named Rules

**The One Voice Rule.** O azul sólido (#2d5de6) cobre no máximo um controle por
viewport. Botões primários competem por atenção; só um ganha.

**No Functional Glow.** Sombras têm offset e blur reais. Halos coloridos de
offset zero são decoração e não entram.

## Typography

**Display/Body/Label Font:** Inter Variable (Inter, system-ui fallback).

**Character:** uma só família fazendo todo o trabalho — peso e tracking criam
a hierarquia, não troca de fonte.

### Hierarchy
- **Display** (650, clamp 38–62px, lh 1.04, −0.035em): título do hero, único.
- **Headline** (650, clamp 28–38px, −0.03em): títulos de seção.
- **Title** (600, 19–20px, −0.02em): nomes de plataforma e passos.
- **Body** (400, 15–18px, lh 1.55): textos de apoio, `text-wrap: pretty`.
- **Label** (550, 13.5–14px): botões, tags, meta-informação.

### Named Rules

**The Quiet Label Rule.** Nada de kicker/eyebrow acima de títulos — o título
carrega o próprio peso.

## Layout

Container central de 1160px, padding lateral 24px (20px no mobile). Ritmo de
seção: 88px de respiro vertical (64px no mobile). Hero em grid 5fr/6fr com o
mockup à direita; colapsa para coluna única abaixo de 920px. A seção de
download é a única faixa de superfície elevada — a quebra visual sinaliza que
ali está a ação. Cards de plataforma lado a lado em 2 colunas, empilhados no
mobile. Passos em trilha dividida por hairlines, não em cards.

## Elevation & Depth

Híbrido: profundidade por camadas tonais (#080b14 → #101522) mais uma sombra
suave real onde o estado a pede (CTA, hover de card). O único efeito luminoso é
o dot de transmissão no mockup — literal, não decorativo.

### Shadow Vocabulary
- **CTA** (`0 1px 0 rgb(255 255 255 / .12) inset, 0 12px 28px -12px rgb(0 0 0 / .8)`):
  botão primário; sombra de contato, não de luz.
- **Seleção** (`0 0 0 4px rgb(56 107 255 / .12)`): anel de foco suave ao redor
  do card selecionado; state, não ornamento.
- **Dispositivo** (`0 30px 60px -20px rgb(0 0 0 / .9)`): o telefone no mockup —
  a única "flutuação" permitida porque é um objeto físico ilustrado.

### Named Rules

**Flat by Default.** Superfícies repousam planas; elevação é resposta a estado
(seleção, hover, foco), nunca ambientação.

## Shapes

Raios moderados e consistentes: 6px em elementos pequenos (tags, skip-link),
8px em botões e campos, 10px em cards de plataforma. O `99px` existe só para o
check circular e o "notch" do telefone no mockup. Bordas sempre hairline 1px
em `#242b3a`; o card selecionado troca a hairline por uma borda em gradiente
fina (roxo→azul) — a única licença do gradiente de marca.

## Components

### Buttons
- **Shape:** raio moderado (8px), altura 44px (52px na variante grande).
- **Primary:** Azul Ação sólido, texto branco, sombra de contato. Hover clareia
  um tom (#4a79ff); press desloca 1px para baixo — resposta tátil mínima.
- **Ghost:** transparente com hairline; hover enche com a superfície.
- **Disabled:** superfície + texto mutado, cursor `not-allowed`, sem sombra.

### Platform Card (componente assinatura)
- Card clicável `role="radio"`: ícone do SO, nome, descrição, linha de seleção.
- Estados: repouso hairline → hover levanta 2px e aclara a borda → selecionado
  ganha borda gradiente + anel azul + check circular animado.
- Indisponível: "Esta versão ainda não está disponível." em texto mutado; o
  card continua selecionável (a spec manda) mas a ação final permanece
  desabilitada.

### Navigation
- Header sticky de 56px com blur de vidro e hairline inferior. Logo
  tipográfico, links mutados que aclarão no hover, "Download" como link com
  borda (não botão preenchido — o CTA real vive no hero e na seção).

### Detection Banner
- Pill discreta com dot ciano: "Detectamos que você está usando macOS" —
  informa, nunca seleciona.

## Do's and Don'ts

### Do:
- **Do** manter o CTA em azul sólido; um por viewport.
- **Do** usar `#7d8496` como piso de texto secundário (≥4.5:1).
- **Do** reservar `#a12bff` e `#00c8e8` para micro-acentos e o mockup.
- **Do** manter copy em pt-BR direta, sem jargão (nunca "ADB", "scrcpy" na UI).
- **Do** tratar o mockup monitor+telefone como ilustração única e literal.

### Don't:
- **Don't** aplicar o gradiente roxo→azul como fundo de botão ou texto.
- **Don't** usar glow/halo colorido sem offset real.
- **Don't** criar eyebrow/kicker, seção de métricas, ou cards iguais de
  ícone+título+texto como estrutura de página.
- **Don't** prometer macOS ou iOS — plataformas indisponíveis dizem "ainda não
  está disponível", nunca "em breve" como CTA.
- **Don't** adicionar elementos flutuantes, emojis como ícones, ou ilustrações
  abstratas.
