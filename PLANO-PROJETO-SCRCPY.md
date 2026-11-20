# Projeto Refletunes: distribuição independente baseada no scrcpy

nome do app: Refletunes

**Documento de execução para Windsurf**  
**Máquina principal:** iMac Intel + dispositivo Android físico  
**Plataforma inicial de distribuição:** Windows 10/11 64-bit  
**Orçamento inicial:** R$ 0, usando ferramentas e cotas gratuitas  
**Objetivo:** criar uma distribuição derivada, mais simples de utilizar, acompanhada de site de downloads e tutoriais monetizado futuramente com publicidade.

---

## 1. Visão do projeto

O projeto parte do [scrcpy](https://github.com/Genymobile/scrcpy), software open source da Genymobile que permite espelhar e controlar dispositivos Android conectados ao computador, sem exigir root no aparelho.

A proposta não é apenas trocar o nome e republicar o executável. Queremos criar uma distribuição independente com:

- nome e identidade visual próprios;
- experiência de instalação e inicialização simplificada para Windows;
- configuração guiada para conexão USB;
- detecção de dispositivos Android e mensagens de erro compreensíveis;
- opções de qualidade simples para usuários iniciantes;
- downloads verificáveis, versões identificadas e changelog;
- site mobile-first com download, tutoriais e solução de problemas;
- manutenção responsável, acompanhando atualizações de segurança do projeto original.

**Princípio:** preservar créditos, avisos e licenças do scrcpy e de todos os componentes redistribuídos. O produto não pode parecer oficial nem sugerir associação, patrocínio ou endosso da Genymobile.

### Resultado esperado do MVP

1. Pacote Windows funcional, testado em Windows real ou máquina virtual.
2. Launcher simples para orientar a conexão Android e iniciar o espelhamento.
3. Site responsivo com download e documentação.
4. Build Windows automatizado por GitHub Actions.
5. Processo documentado de atualização, verificação e publicação.

Não anunciar compatibilidade universal nem publicar builds sem teste funcional.

---

## 2. Escopo e prioridades

### MVP

- Windows 10/11 64-bit, condicionado à compatibilidade da versão upstream escolhida.
- Conexão USB via ADB.
- Detecção de dispositivo conectado.
- Orientação sobre modo desenvolvedor, depuração USB e autorização RSA.
- Inicialização do scrcpy com opções validadas.
- Launcher com estado da conexão, botão de verificar, iniciar e abrir ajuda.
- Dois perfis simples de qualidade, caso suportados pela versão fixada.
- Site com download, instalação, solução de problemas, changelog, créditos e licenças.
- Build automatizado, SHA-256 e release identificada.

### Depois do MVP

- Conexão Wi-Fi guiada.
- Seleção e gerenciamento de múltiplos dispositivos.
- Perfis avançados de resolução, FPS e bitrate.
- Atualizador automático seguro.
- Builds macOS/Linux da distribuição própria.
- Telemetria opcional, somente se houver necessidade, transparência e consentimento.

### Não fazer

- Não reescrever a tecnologia de espelhamento.
- Não criar login, contas, sincronização ou backend sem necessidade real.
- Não capturar, armazenar ou transmitir a tela do usuário.
- Não adicionar analytics ou telemetria ao aplicativo desktop no MVP.
- Não solicitar senha do Android.
- Não orientar a desativar antivírus, SmartScreen, firewall ou proteções do sistema.
- Não executar comandos arbitrários recebidos da internet ou de arquivos não confiáveis.
- Não remover créditos, avisos de copyright ou licenças.
- Não empacotar bibliotecas, drivers, fontes, ícones ou imagens sem verificar suas licenças.

---

## 3. Marca e posicionamento

Usar um nome independente, sem apresentar o produto como uma versão oficial do scrcpy. O nome final precisa passar por pesquisa de disponibilidade de domínio, marca, aplicativo e repositório.

Usar `MirrorDesk` apenas como placeholder interno nos exemplos deste plano. **Não assumir que esse nome está disponível para uso comercial.**

Posicionamento provisório:

> Espelhe e controle seu Android pelo computador com uma configuração simples, instruções claras e downloads verificáveis.

O site e o pacote devem explicar que se trata de uma distribuição independente baseada no scrcpy, sem associação oficial com a Genymobile.

Antes do lançamento:
- pesquisar o nome em mecanismos de busca;
- verificar marcas e domínios;
- evitar nomes, logos e elementos visuais que causem confusão com scrcpy ou Genymobile;
- substituir todos os placeholders por uma marca validada.

---

## 4. Licenças e conformidade — etapa obrigatória

O Windsurf deve tratar licenças como requisito técnico bloqueador. Não iniciar redistribuição antes de concluir esta auditoria.

### 4.1 scrcpy

O scrcpy utiliza Apache License 2.0. Em geral, essa licença permite uso comercial, modificação e redistribuição, desde que suas condições sejam cumpridas. Para a versão e commit efetivamente utilizados:

1. Ler a licença e os avisos presentes no repositório upstream.
2. Preservar avisos de copyright, licença e atribuição exigidos.
3. Incluir cópia da Apache License 2.0 na distribuição.
4. Criar `THIRD_PARTY_NOTICES.md` com componentes, autores, versões e licenças.
5. Identificar arquivos modificados de forma clara.
6. Não utilizar marcas ou logotipos da Genymobile como se fossem identidade ou endosso do produto.

### 4.2 Assets e dependências

Não presumir que tudo no repositório possui a mesma licença do código. Inspecionar separadamente:

- imagens, animações, vídeos e capturas de tela;
- ícones, fontes e logos;
- dependências e bibliotecas empacotadas;
- binários ADB / Android platform-tools;
- bibliotecas auxiliares;
- ferramentas de instalação e empacotamento.

Não redistribuir assets com licença incompatível ou não esclarecida. Quando necessário, substituir por assets próprios ou omitir.

### 4.3 Aviso de independência

Incluir no site e no pacote um aviso semelhante a:

> Este é um projeto independente baseado no scrcpy. Não é um produto oficial, nem é afiliado, patrocinado ou endossado pela Genymobile. O scrcpy é desenvolvido e mantido por seus respectivos autores. Consulte os créditos e licenças incluídos nesta distribuição.

Solicitar revisão humana das obrigações de licença antes do primeiro lançamento público. Este documento é um plano técnico, não aconselhamento jurídico.

---

## 5. Arquitetura recomendada

Separar o aplicativo desktop e o site. Evitar infraestrutura desnecessária.

### 5.1 Aplicativo desktop

No MVP, criar uma camada fina sobre o scrcpy. O launcher deve detectar ADB e dispositivos, validar pré-requisitos e iniciar o binário com argumentos permitidos.

Não começar com Electron ou uma aplicação pesada sem justificar a necessidade. Primeiro, o Windsurf deve auditar a estrutura upstream e comparar opções simples para o launcher. A decisão de linguagem/framework deve ser registrada em `ARCHITECTURE.md`, considerando manutenção, tamanho, compatibilidade Windows e facilidade de build.

Requisitos do launcher:
- interface simples e clara;
- botão “Verificar celular”;
- estado de conexão legível;
- instruções para ativar depuração USB;
- detecção de dispositivo via ADB;
- aviso quando o aparelho ainda não autorizou a chave RSA;
- botão “Iniciar espelhamento” habilitado somente quando apropriado;
- ação para encerrar sessão;
- perfis limitados de qualidade;
- acesso à ajuda e solução de problemas;
- logs locais mínimos para diagnóstico, sem conteúdo da tela e sem dados pessoais desnecessários.

Os argumentos do processo devem vir de uma lista permitida e validada. Nunca executar comandos arbitrários recebidos de arquivos, URLs ou conteúdo remoto.

### 5.2 Site

Stack inicial:
- React + Vite + TypeScript;
- Cloudflare Pages para frontend estático;
- GitHub para código, releases e automação;
- GitHub Releases ou Cloudflare R2 para arquivos de download;
- Cloudflare Workers somente se surgir necessidade real de API;
- D1 não é necessário no MVP.

Não criar backend, autenticação ou banco de dados sem uma necessidade comprovada.

### 5.3 Estrutura sugerida

O Windsurf deve adaptar a estrutura depois de inspecionar o repositório, preservando o processo upstream de build.

```text
projeto/
├── upstream/                 # referência rastreável ao scrcpy
├── desktop/
│   ├── launcher/
│   ├── packaging/
│   └── README.md
├── website/
│   ├── public/
│   └── src/
│       ├── components/
│       ├── pages/
│       ├── content/
│       └── styles/
├── docs/
│   ├── INSTALL-WINDOWS.md
│   ├── USB-CONNECTION.md
│   ├── WIFI-CONNECTION.md
│   ├── TROUBLESHOOTING.md
│   ├── FAQ.md
│   ├── CHANGELOG.md
│   └── ARCHITECTURE.md
├── legal/
│   ├── LICENSE
│   └── THIRD_PARTY_NOTICES.md
├── .github/workflows/
│   ├── build-windows.yml
│   └── deploy-site.yml
├── README.md
└── CONTRIBUTING.md
```

Avaliar se é melhor manter um fork público ou uma referência/submodule ao upstream. Manter customizações em commits separados para facilitar atualizações futuras. Não copiar cegamente o projeto para uma estrutura nova.

---

## 6. Desenvolvimento no iMac Intel

O iMac Intel será a máquina principal de desenvolvimento e testes macOS. O scrcpy oferece suporte ao macOS, mas o produto inicial será distribuído para Windows.

### 6.1 Levantamento do ambiente

Antes de instalar dependências, registrar:

- versão do macOS;
- modelo/ano do iMac;
- memória RAM;
- espaço livre;
- versão do Xcode Command Line Tools;
- versão do Homebrew.

Executar no Terminal:

```bash
uname -m
sw_vers
system_profiler SPHardwareDataType
df -h
```

Instalar ferramentas de linha de comando, se necessário:

```bash
xcode-select --install
```

Instalar Homebrew apenas se ainda não estiver instalado, seguindo a documentação oficial: https://brew.sh/.

Para ADB e scrcpy, seguir as instruções oficiais correspondentes à versão escolhida. Atenção: instalar `scrcpy` pelo Homebrew fornece o pacote publicado pelo Homebrew, não automaticamente a versão customizada deste projeto.

### 6.2 Teste com Android físico

É necessário um Android físico e cabo USB com suporte a transferência de dados.

No Android:
1. Abrir Configurações.
2. Ativar Opções do desenvolvedor, seguindo as instruções do fabricante.
3. Ativar Depuração USB.
4. Conectar ao iMac.
5. Aceitar a solicitação de autorização RSA no aparelho.

Verificar a conexão:

```bash
adb devices
```

Executar a versão upstream no macOS e verificar:
- imagem e latência;
- controle com mouse;
- controle com teclado;
- rotação;
- áudio, quando suportado pelo dispositivo e pela versão;
- encerramento limpo da sessão;
- desconexão e reconexão USB.

Não alegar compatibilidade com todos os aparelhos com base em um único telefone.

---

## 7. Build e testes Windows sem comprar outro computador

Usar GitHub Actions para compilar artefatos Windows em runners hospedados pelo GitHub, respeitando as cotas e condições gratuitas vigentes da conta.

Fluxo:
1. O iMac envia commits ao GitHub.
2. GitHub Actions compila a versão Windows.
3. O workflow gera os arquivos de distribuição.
4. O workflow calcula SHA-256.
5. Artefatos são anexados a uma release de teste.
6. Baixar o artefato no iMac.
7. Transferir para Windows real ou VM.
8. Executar testes funcionais.
9. Aprovar a release pública somente após os testes.

**Build verde não substitui teste funcional em Windows.**

### 7.1 Máquina virtual

Avaliar UTM ou outra solução compatível com Intel Mac. Compatibilidade depende da versão do macOS, recursos do iMac e licença/ISO do Windows.

Não presumir que USB passthrough e aceleração gráfica funcionarão perfeitamente. Quando possível, testar também em Windows físico.

### 7.2 Matriz mínima de testes

| Cenário | Resultado esperado |
|---|---|
| Windows sem celular conectado | Instrução clara; não travar |
| Android conectado sem depuração USB | Tutorial claro |
| Depuração ativa, sem autorização RSA | Explicar como autorizar no telefone |
| Um dispositivo autorizado | Permitir iniciar |
| Mais de um dispositivo | Pedir seleção explícita |
| Cabo removido durante sessão | Encerrar ou informar falha sem travar |
| ADB ausente/corrompido | Erro compreensível |
| Caminho de instalação com espaços | Funcionar |
| Usuário Windows com caracteres acentuados | Funcionar |
| Windows Defender/SmartScreen | Não orientar a desativar proteções |

Se não houver Windows real nem VM funcional, manter a versão como beta privada. Não anunciar como estável.

---

## 8. Funcionalidades do MVP desktop

### Tela inicial

Exibir:
- marca própria;
- estado “Nenhum dispositivo conectado” ou “Dispositivo conectado”;
- botão “Como conectar”;
- botão “Verificar novamente”;
- botão “Iniciar espelhamento”;
- link para documentação;
- versão do programa.

### Fluxo de conexão

1. Usuário conecta o Android.
2. Launcher verifica ADB.
3. Launcher identifica dispositivos.
4. Se houver um único dispositivo autorizado, seleciona-o.
5. Se houver vários, exibe uma lista e pede seleção.
6. Se não houver autorização, explica a confirmação RSA no celular.
7. Usuário escolhe perfil de qualidade.
8. Launcher inicia o scrcpy.
9. Usuário encerra a sessão pelo launcher ou pela janela de espelhamento.

### Perfis iniciais

- **Padrão:** qualidade equilibrada.
- **Desempenho:** menor resolução/bitrate para computadores modestos.
- **Personalizado:** deixar para uma versão futura.

Usar apenas opções suportadas pela versão upstream fixada. Consultar `scrcpy --help` e a documentação da mesma versão. Não inventar flags.

### Segurança

- Não solicitar senha do Android.
- Não capturar nem enviar a tela para servidores.
- Não armazenar conteúdo espelhado.
- Não executar comandos vindos de URLs, arquivos remotos ou conteúdo web.
- Não elevar privilégios sem necessidade.
- Explicar que a depuração USB concede capacidade de controle ao computador autorizado.
- Recomendar revogar autorizações USB em computadores compartilhados.

---

## 9. Site: páginas, conteúdo e UX

O site deve funcionar muito bem em celular, mesmo que o software seja destinado ao Windows.

### Páginas obrigatórias

1. `/` — página inicial e explicação do produto.
2. `/download` — versão atual, requisitos, arquitetura e checksum.
3. `/como-instalar` — instalação passo a passo.
4. `/conectar-usb` — conexão USB.
5. `/conectar-wifi` — publicar somente depois de validar o procedimento.
6. `/solucao-de-problemas` — erros comuns.
7. `/compatibilidade` — requisitos e limitações conhecidas.
8. `/changelog` — histórico de versões.
9. `/licencas` — licença, créditos e terceiros.
10. `/sobre` — projeto independente e contato.
11. `/privacidade` — política de privacidade.
12. `/termos` — termos de uso.

### Conteúdo SEO inicial

Criar artigos originais, úteis e testados:
- Como espelhar Android no PC pelo cabo USB.
- Como controlar Android com teclado e mouse.
- Como ativar depuração USB em diferentes fabricantes.
- Como resolver `device unauthorized`.
- Como resolver `device not found`.
- Como conectar Android ao computador por Wi-Fi.
- Como reduzir atraso e melhorar desempenho.
- Como espelhar Android no Windows 10 e Windows 11.
- Como desconectar e revogar autorização de depuração USB.

Não criar dezenas de páginas quase idênticas trocando apenas o fabricante. Cada artigo precisa conter instruções verificadas, imagens próprias/licenciadas e informação específica.

### Diretrizes visuais

- Design claro, moderno e confiável.
- Mobile-first e responsivo.
- Download objetivo, sem botões falsos.
- Mostrar versão, data, arquitetura e tamanho do arquivo.
- Diferenciar visualmente anúncios e botões de download.
- Não usar pop-ups agressivos, redirecionadores ou downloads automáticos.
- Contraste adequado, foco visível, navegação por teclado e textos legíveis.

---

## 10. Publicidade e monetização

A monetização inicial será no site, não dentro do aplicativo desktop.

### Antes de solicitar AdSense

- Publicar conteúdo original e útil.
- Criar páginas Sobre, Contato, Privacidade e Termos.
- Identificar claramente o responsável pelo site.
- Garantir navegação funcional e download seguro.
- Não comprar tráfego inválido nem clicar nos próprios anúncios.
- Solicitar aprovação quando o site estiver completo e em conformidade.

A aprovação não é garantida. Um site que funciona apenas como espelho de downloads pode oferecer pouco valor editorial. O conteúdo deve ser original, útil e não se limitar a copiar a documentação upstream.

### Posicionamento de anúncios

- Espaço contextual em tutoriais.
- Espaço ao final de artigos longos.
- Nunca fazer anúncio parecer botão de download.
- Não inserir anúncios dentro do launcher.
- Respeitar políticas de publicidade, privacidade e consentimento aplicáveis.

Possibilidades futuras: links afiliados para cabos USB de dados, hubs e acessórios, sempre identificados como publicidade/afiliado. Não priorizar isso antes de validar tráfego.

---

## 11. Distribuição e releases

Cada versão pública deve incluir:
- executável/pacote;
- versão e data de publicação;
- changelog;
- SHA-256;
- licenças e avisos de terceiros;
- requisitos mínimos;
- instruções de instalação e remoção;
- versão/tag e commit upstream utilizados.

### Segurança da distribuição

- Usar GitHub Releases inicialmente ou R2 após validação.
- Não substituir arquivos de uma release publicada; criar nova versão.
- Manter histórico de releases.
- Gerar SHA-256 automaticamente.
- Proteger a branch principal.
- Revisar workflows antes de permitir publicação.
- Usar permissões mínimas em GitHub Actions.
- Não colocar tokens ou secrets no código.
- Não alegar assinatura digital se não existir certificado válido.
- Informar honestamente se o executável não possui assinatura Authenticode.

A ausência de assinatura pode gerar avisos do Windows. Não instruir usuários a ignorar alertas indiscriminadamente. Avaliar assinatura somente quando houver orçamento e demanda.

---

## 12. Cloudflare: site e downloads

### Site

- Criar projeto Cloudflare Pages conectado ao GitHub.
- Configurar comando de build de acordo com o projeto Vite, normalmente `npm run build`.
- Diretório de saída normalmente `dist`.
- Usar subdomínio gratuito `*.pages.dev` no início.
- Testar rotas, fallback SPA, sitemap e metadados.

### Downloads

Para validar o MVP, GitHub Releases pode hospedar os primeiros artefatos. Se usar Cloudflare R2:
- configurar bucket e domínio de download corretamente;
- usar nomes versionados;
- não substituir silenciosamente binários existentes;
- testar download, tamanho e checksum;
- configurar cache coerente com releases imutáveis.

Não é necessário usar Workers ou D1 para um site estático de documentação e downloads.

### Custos

A meta é permanecer dentro das franquias gratuitas, mas limites e condições podem mudar. Antes do lançamento, verificar os limites vigentes de Cloudflare Pages, R2 e GitHub Actions. Não prometer gratuidade ilimitada ou permanente.

---

## 13. Plano de execução por fases

### Fase 0 — auditoria (não modificar o upstream)

1. Clonar o repositório oficial.
2. Ler README e documentação de build.
3. Identificar versão/tag estável apropriada.
4. Identificar licença principal e licenças de componentes.
5. Identificar assets com licenças próprias.
6. Entender como o build oficial Windows é gerado.
7. Registrar dependências, flags e requisitos.
8. Criar `AUDIT.md` com achados e riscos.

**Critério de saída:** processo de build e redistribuição compreendidos antes de customizar.

### Fase 1 — compilar e testar upstream

1. Preparar o iMac Intel.
2. Compilar/testar upstream no macOS conforme documentação oficial.
3. Conectar Android físico.
4. Validar USB, teclado, mouse, rotação e estabilidade.
5. Fixar versão/tag/commit upstream.
6. Reproduzir build Windows via GitHub Actions sem customizações.

**Critério de saída:** build reproduzível e teste funcional documentado.

### Fase 2 — distribuição Windows mínima

1. Criar pacote Windows usando o método oficial.
2. Preservar estrutura e avisos legais.
3. Criar launcher mínimo, se a auditoria confirmar viabilidade.
4. Implementar detecção ADB e de dispositivos.
5. Criar mensagens de erro claras.
6. Criar perfis simples usando flags validadas.
7. Gerar checksum e changelog.
8. Testar em Windows real/VM.

**Critério de saída:** uma pessoa sem conhecimento técnico consegue seguir o tutorial e iniciar o espelhamento.

### Fase 3 — site

1. Criar frontend React/Vite/TypeScript.
2. Criar identidade visual independente.
3. Criar páginas essenciais.
4. Criar página de download ligada à release real.
5. Criar guias e solução de problemas.
6. Adicionar páginas legais e créditos.
7. Implantar em Cloudflare Pages.
8. Testar mobile, desktop, SEO técnico e acessibilidade.

**Critério de saída:** visitante entende o produto, baixa o pacote correto e segue o tutorial.

### Fase 4 — beta fechado

1. Compartilhar com 5–10 pessoas que usem Android e Windows.
2. Observar sem coletar dados sensíveis.
3. Pedir feedback sobre instalação, clareza e falhas.
4. Registrar problemas reproduzíveis.
5. Corrigir bloqueadores.
6. Atualizar documentação.

**Critério de saída:** usuários conseguem instalar sem assistência individual na maioria dos casos.

### Fase 5 — lançamento público

1. Revisar licenças e atribuições.
2. Revisar instalador e binários.
3. Publicar release e checksum.
4. Publicar site e conteúdo.
5. Configurar Search Console e sitemap.
6. Solicitar AdSense quando houver conteúdo suficiente.
7. Monitorar erros, downloads e páginas visitadas sem rastreamento invasivo.

### Fase 6 — evolução

Priorizar com base no feedback:
- Wi-Fi guiado;
- seleção de múltiplos dispositivos;
- perfis avançados;
- suporte a mais versões Windows;
- atualizador seguro;
- builds macOS/Linux, se houver demanda.

---

## 14. Como trabalhar com o Windsurf

O usuário tem experiência em UI/UX, mas conhecimento limitado de programação. O agente deve fornecer instruções precisas, caminhos de arquivos, comandos completos e explicações simples.

### Regras obrigatórias para o agente

- Não fazer mudanças grandes antes da auditoria.
- Não reescrever o scrcpy inteiro.
- Não inventar APIs, comandos, flags ou caminhos.
- Consultar documentação da versão upstream fixada.
- Preservar compatibilidade com upstream sempre que possível.
- Separar customizações em commits claros.
- Explicar comandos e informar o diretório em que devem ser executados.
- Ao alterar um arquivo, informar o caminho exato e fornecer conteúdo completo quando for mais seguro.
- Não apagar arquivos sem explicar o motivo e o impacto.
- Não publicar, criar release ou fazer deploy sem autorização explícita.
- Não afirmar que um build foi testado em Windows se apenas compilou no GitHub Actions.
- Não colocar secrets no repositório.
- Atualizar changelog, README e documentação afetada.
- Executar testes após cada etapa e relatar resultados reais.
- Não avançar automaticamente para a próxima fase sem apresentar o resultado da fase atual.

Dividir o trabalho em tarefas pequenas: auditoria, ambiente local, build upstream, testes Android/macOS, build Windows, launcher, testes Windows, site, conteúdo/SEO e release. Não pedir ao Windsurf para executar todas as fases em um único prompt.

---

## 15. Prompt inicial para o Windsurf

Copie o prompt abaixo para iniciar o trabalho. A primeira tarefa é somente auditar; não modificar o projeto.

```text
Estamos iniciando um projeto de distribuição independente baseada no scrcpy:
https://github.com/Genymobile/scrcpy

Leia integralmente o arquivo PLANO-PROJETO-SCRCPY.md antes de executar qualquer tarefa.

CONTEXTO:
- Minha máquina principal é um iMac Intel.
- Tenho um dispositivo Android físico para testes.
- Quero inicialmente distribuir uma versão para Windows 10/11 64-bit.
- Quero usar GitHub Actions para builds Windows e Cloudflare Pages/R2 para site e downloads.
- O orçamento inicial é zero.
- Tenho experiência em UI/UX, mas conhecimento limitado de programação. Preciso de instruções passo a passo, caminhos e comandos completos.

SUA PRIMEIRA TAREFA É SOMENTE AUDITORIA. NÃO MODIFIQUE ARQUIVOS AINDA.

1. Inspecione o repositório oficial e identifique uma versão/tag estável adequada.
2. Leia README, documentação de compilação e instruções para Windows/macOS.
3. Identifique a licença principal e as licenças de componentes, binários, assets e dependências que possam ser redistribuídos.
4. Explique o método oficial de build para Windows.
5. Explique como testar o upstream no iMac Intel com Android físico.
6. Identifique requisitos para criar um launcher simplificado sem comprometer o funcionamento do scrcpy.
7. Identifique riscos técnicos, legais e de distribuição.
8. Recomende a estrutura mais segura para manter customizações separadas do upstream.
9. Verifique se o projeto pode ser compilado no GitHub Actions usando runners gratuitos e quais limitações precisam ser consideradas.
10. Não presuma que os assets possuem a mesma licença do código.

ENTREGÁVEL:
Crie AUDIT.md com:
- resumo do projeto;
- versão/commit recomendado;
- arquitetura atual;
- processo de build macOS;
- processo de build Windows;
- licenças e obrigações;
- dependências;
- riscos;
- plano técnico recomendado;
- decisões que ainda dependem de mim.

IMPORTANTE:
- Não faça fork, commit, release ou deploy por conta própria.
- Não remova créditos ou avisos de licença.
- Não invente comandos nem opções. Confirme-os na documentação da versão analisada.
- Ao final, apresente um resumo simples do que encontrou e os próximos comandos que devo executar.
```

---

## 16. Checklist de lançamento

- [ ] Build Windows reproduzível.
- [ ] Teste funcional em Windows real ou VM.
- [ ] Conexão USB validada com Android físico.
- [ ] Mensagens de erro compreensíveis.
- [ ] Launcher não executa comandos arbitrários.
- [ ] Licenças e créditos incluídos.
- [ ] Assets redistribuídos com autorização.
- [ ] SHA-256 publicado.
- [ ] Download testado.
- [ ] Site responsivo e acessível.
- [ ] Páginas de privacidade, termos, contato e sobre.
- [ ] Tutoriais testados.
- [ ] Changelog e versão visíveis.
- [ ] Sem telemetria invasiva no MVP.
- [ ] Release identificada como independente.
- [ ] Revisão de conformidade antes do lançamento.

---

## 17. Decisões em aberto

Não bloquear a auditoria por estas decisões. Registrar no `AUDIT.md` e recomendar opções:

1. Nome e identidade visual definitivos.
2. Linguagem/framework do launcher.
3. Distribuição portable ZIP ou instalador Windows.
4. Suporte inicial apenas USB ou USB + Wi-Fi.
5. Hospedar releases no GitHub ou Cloudflare R2.
6. Momento adequado para solicitar aprovação do AdSense.
7. Possibilidade de builds macOS/Linux no futuro.

**Direção final:** primeiro provar que conseguimos compilar, executar e distribuir corretamente o scrcpy no Windows. Depois simplificar a experiência. Só então investir em branding, SEO e monetização. A prioridade é entregar um software confiável, não apenas um site com um botão de download.
