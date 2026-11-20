package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	appName    = "Refletunes"
	appVersion = "0.1.0"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(newTheme())
	w := a.NewWindow(appName)
	w.Resize(fyne.NewSize(620, 540))
	w.SetPadded(false)

	ui := &launcherUI{win: w, selectedIdx: -1}

	body := container.NewAppTabs(
		container.NewTabItem("Configurar", ui.buildSetup()),
		container.NewTabItem("Dispositivos", ui.buildDevices()),
		container.NewTabItem("Ajuda", ui.buildHelp()),
	)
	body.SetTabLocation(container.TabLocationLeading)
	body.OnSelected = func(t *container.TabItem) {
		// auto-refresh só na aba Dispositivos
		ui.autoRefresh = t.Text == "Dispositivos"
	}
	// aba inicial é "Configurar"; o polling liga ao entrar em "Dispositivos"
	go ui.pollDevices()

	header := ui.buildHeader()
	w.SetContent(container.NewBorder(header, nil, nil, nil, body))
	ui.detectADB(false)
	w.ShowAndRun()
}

type launcherUI struct {
	win fyne.Window

	adbPath   string
	adbDot    *canvas.Circle
	adbStatus *widget.Label

	deviceList   *fyne.Container
	statusLabel  *widget.Label
	profileSel   *widget.Select
	devices      []Device
	selectedIdx  int
	mirrorBtn    *widget.Button
	deviceChecks []*widget.Check
	autoRefresh  bool
	polling      bool
}

// ---------- Cabeçalho ----------

func (u *launcherUI) buildHeader() fyne.CanvasObject {
	bg := canvas.NewRectangle(colHeader)

	logo := canvas.NewText("◉", colAccent)
	logo.TextSize = 26

	name := canvas.NewText(appName, colText)
	name.TextStyle = fyne.TextStyle{Bold: true}
	name.TextSize = 18

	version := canvas.NewText("v"+appVersion, colTextDim)
	version.TextSize = 12

	titleBox := container.NewVBox(name, version)

	u.adbDot = canvas.NewCircle(colErr)
	u.adbDot.Resize(fyne.NewSize(10, 10))
	u.adbStatus = widget.NewLabel("verificando...")

	statusBox := container.NewHBox(
		container.NewCenter(u.adbDot),
		u.adbStatus,
	)

	row := container.NewBorder(nil, nil,
		container.NewHBox(logo, titleBox),
		statusBox,
	)
	return container.NewStack(bg, container.NewPadded(row))
}

func (u *launcherUI) setADBState(ok bool, msg string) {
	if ok {
		u.adbDot.FillColor = colOK
	} else {
		u.adbDot.FillColor = colErr
	}
	u.adbDot.Refresh()
	u.adbStatus.SetText(msg)
}

// ---------- Aba Configurar ----------

func (u *launcherUI) buildSetup() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Antes de conectar", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	intro := widget.NewLabel("Faça estas etapas uma vez no celular:")
	intro.Wrapping = fyne.TextWrapWord

	step := func(n, txt string) fyne.CanvasObject {
		num := canvas.NewText(n, colAccent)
		num.TextStyle = fyne.TextStyle{Bold: true}
		num.TextSize = 15
		lbl := widget.NewLabel(txt)
		lbl.Wrapping = fyne.TextWrapWord
		return container.NewBorder(nil, nil,
			container.NewVBox(container.NewCenter(num), layout.NewSpacer()), nil, lbl)
	}

	stepsBg := canvas.NewRectangle(colCard)
	stepsBg.CornerRadius = 10
	stepsCard := container.NewStack(stepsBg, container.NewPadded(container.NewVBox(
		step("1", "Configurações → Sobre o telefone → toque 7× em \"Número da versão\""),
		step("2", "Opções do desenvolvedor → ative \"Depuração USB\""),
		step("3", "Conecte o cabo e toque em PERMITIR no aviso do celular"),
		step("4", "Xiaomi/POCO: ative \"Depuração USB (configurações de segurança)\" para o controle funcionar"),
	)))

	wifiBg := canvas.NewRectangle(colCard)
	wifiBg.CornerRadius = 10
	wifiTitle := widget.NewLabelWithStyle("Prefere Wi-Fi?", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	wifiTxt := widget.NewLabel("Android 11+: ative \"Depuração sem fio\" nas Opções do desenvolvedor e vá para a aba Dispositivos.")
	wifiTxt.Wrapping = fyne.TextWrapWord
	wifiCard := container.NewStack(wifiBg, container.NewPadded(container.NewVBox(wifiTitle, wifiTxt)))

	checkBtn := widget.NewButton("Verificar conexão", func() { u.detectADB(true) })
	checkBtn.Importance = widget.MediumImportance

	return container.NewScroll(container.NewPadded(container.NewVBox(
		title, intro, stepsCard, wifiCard, checkBtn,
	)))
}

func (u *launcherUI) detectADB(showDialog bool) {
	path, err := FindADB()
	if err != nil {
		u.adbPath = ""
		u.setADBState(false, "componente ADB ausente")
		logf("ERRO: adb nao encontrado: %v", err)
		if showDialog {
			dialog.ShowError(fmt.Errorf(
				"O componente adb não foi encontrado.\n\n"+
					"Ele deve estar na mesma pasta do %s.\n"+
					"Baixe novamente o pacote completo no site do Refletunes.", appName), u.win)
		}
		return
	}
	u.adbPath = path
	u.setADBState(true, "ADB pronto")
	logf("ADB localizado em %s", path)
	u.refreshDevices()
}

// ---------- Aba Dispositivos ----------

func (u *launcherUI) buildDevices() fyne.CanvasObject {
	u.statusLabel = widget.NewLabel("")
	u.deviceList = container.NewVBox()

	refreshBtn := widget.NewButtonWithIcon("Verificar", theme.ViewRefreshIcon(), func() {
		if u.adbPath == "" {
			u.detectADB(true)
			return
		}
		u.refreshDevices()
	})

	pairBtn := widget.NewButtonWithIcon("Parear Wi-Fi", theme.ComputerIcon(), func() { u.showPairDialog() })
	discBtn := widget.NewButtonWithIcon("Desconectar", theme.LogoutIcon(), func() {
		if u.adbPath == "" {
			return
		}
		out, _ := Disconnect(u.adbPath, "")
		u.statusLabel.SetText(strings.TrimSpace(out))
		u.refreshDevices()
	})
	discBtn.Importance = widget.LowImportance

	names := make([]string, len(Profiles))
	for i, p := range Profiles {
		names[i] = p.Name
	}
	u.profileSel = widget.NewSelect(names, func(s string) {
		for _, p := range Profiles {
			if p.Name == s {
				u.statusLabel.SetText(p.Desc)
			}
		}
	})
	u.profileSel.SetSelected(Profiles[0].Name)

	u.mirrorBtn = widget.NewButtonWithIcon("Iniciar espelhamento", theme.MediaPlayIcon(), func() { u.startMirror() })
	u.mirrorBtn.Importance = widget.HighImportance
	u.mirrorBtn.Disable()

	actions := container.NewHBox(refreshBtn, pairBtn, layout.NewSpacer(), discBtn)
	footer := container.NewBorder(nil, nil,
		container.NewHBox(widget.NewLabel("Qualidade"), u.profileSel), u.mirrorBtn,
	)

	top := container.NewVBox(actions, container.NewPadded(u.statusLabel), widget.NewSeparator())
	return container.NewBorder(
		container.NewPadded(top),
		container.NewPadded(footer),
		nil, nil,
		container.NewScroll(u.deviceList),
	)
}

// deviceCard desenha um aparelho como cartão selecionável.
func (u *launcherUI) deviceCard(i int, d Device) fyne.CanvasObject {
	bg := canvas.NewRectangle(colCard)
	bg.CornerRadius = 10
	if i == u.selectedIdx {
		bg.FillColor = colCardHover
	}

	icon := widget.NewIcon(theme.ComputerIcon())
	if d.Transp == "usb" {
		icon = widget.NewIcon(theme.StorageIcon())
	}

	name := widget.NewLabelWithStyle(deviceTitle(d), fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	conn := canvas.NewText(d.ConnLabel(), colTextDim)
	conn.TextSize = 12

	chip := statusChip(d)

	radio := widget.NewCheck("", func(checked bool) {
		if checked {
			u.selectedIdx = i
			bg.FillColor = colCardHover
			for j, other := range u.deviceChecks {
				if j != i && other != nil && other.Checked {
					other.SetChecked(false)
				}
			}
		} else {
			if u.selectedIdx == i {
				u.selectedIdx = -1
			}
			bg.FillColor = colCard
		}
		bg.Refresh()
		u.updateMirrorBtn()
	})
	radio.SetChecked(i == u.selectedIdx)
	u.deviceChecks[i] = radio

	row := container.NewBorder(nil, nil,
		container.NewHBox(container.NewPadded(icon), container.NewVBox(name, conn)),
		container.NewHBox(chip, radio),
	)
	return container.NewStack(bg, container.NewPadded(row))
}

func deviceTitle(d Device) string {
	if d.Model != "" {
		return strings.ReplaceAll(d.Model, "_", " ")
	}
	return d.Serial
}

// statusChip mostra o estado com cor semântica.
func statusChip(d Device) fyne.CanvasObject {
	var col = colTextDim
	switch d.State {
	case "device":
		col = colOK
	case "unauthorized", "authorizing":
		col = colWarn
	case "offline", "no permissions":
		col = colErr
	}
	dot := canvas.NewCircle(col)
	dot.Resize(fyne.NewSize(8, 8))
	txt := canvas.NewText(d.StateLabel(), col)
	txt.TextSize = 12
	return container.NewHBox(container.NewCenter(dot), txt)
}

func (u *launcherUI) refreshDevices() {
	u.refreshDevicesQuiet(false)
}

// pollDevices atualiza a lista em segundo plano enquanto a aba está visível.
func (u *launcherUI) pollDevices() {
	for range time.Tick(3 * time.Second) {
		if u.autoRefresh && u.adbPath != "" && !u.polling {
			u.refreshDevicesQuiet(true)
		}
	}
}

func (u *launcherUI) refreshDevicesQuiet(silent bool) {
	u.polling = true
	if !silent {
		u.statusLabel.SetText("Verificando dispositivos...")
	}
	go func() {
		devices, err := ListDevices(u.adbPath)
		fyne.Do(func() {
			u.polling = false
			if err != nil {
				if !silent {
					u.statusLabel.SetText("Erro: " + err.Error())
				}
				return
			}
			u.devices = devices
			u.renderDeviceList()
		})
	}()
}

func (u *launcherUI) renderDeviceList() {
	u.deviceList.Objects = nil
	u.deviceChecks = make([]*widget.Check, len(u.devices))
	ready := 0
	for i, d := range u.devices {
		if d.Ready() {
			ready++
		}
		u.deviceList.Add(u.deviceCard(i, d))
	}
	if len(u.devices) == 0 {
		empty := widget.NewLabel("Nenhum dispositivo encontrado.\nConecte o cabo USB e aceite a autorização no celular.")
		empty.Alignment = fyne.TextAlignCenter
		empty.Wrapping = fyne.TextWrapWord
		u.deviceList.Add(container.NewPadded(empty))
	}
	switch {
	case len(u.devices) == 0:
		u.statusLabel.SetText("Aguardando conexão")
	case ready == 0:
		u.statusLabel.SetText("Autorize a depuração na tela do celular")
	default:
		u.statusLabel.SetText(fmt.Sprintf("%d aparelho(s) pronto(s)", ready))
	}
	u.deviceList.Refresh()
	u.updateMirrorBtn()
}

func (u *launcherUI) updateMirrorBtn() {
	if u.selectedIdx >= 0 && u.selectedIdx < len(u.devices) && u.devices[u.selectedIdx].Ready() {
		u.mirrorBtn.Enable()
	} else {
		u.mirrorBtn.Disable()
	}
}

func (u *launcherUI) startMirror() {
	if u.selectedIdx < 0 || u.selectedIdx >= len(u.devices) {
		return
	}
	scrcpyPath, err := FindScrcpy()
	if err != nil {
		dialog.ShowError(err, u.win)
		return
	}
	var profile QualityProfile
	for _, p := range Profiles {
		if p.Name == u.profileSel.Selected {
			profile = p
		}
	}
	d := u.devices[u.selectedIdx]
	StartMirror(scrcpyPath, d, profile)
	logf("espelhamento iniciado: %s perfil=%s", d.Serial, profile.Name)
	u.statusLabel.SetText("Espelhamento iniciado — " + deviceTitle(d))
	go func() {
		time.Sleep(2 * time.Second)
		fyne.Do(u.refreshDevices)
	}()
}

// ---------- Aba Ajuda ----------

func (u *launcherUI) buildHelp() fyne.CanvasObject {
	faq := widget.NewAccordion(
		widget.NewAccordionItem("O celular não aparece na lista",
			widget.NewLabel("Confira: depuração USB ativada, cabo com dados (não só carga), e o aviso de autorização aceito na tela do celular. Depois toque em \"Verificar\".")),
		widget.NewAccordionItem("Aparece \"aguardando autorização\"",
			widget.NewLabel("Desbloqueie o celular e procure o aviso \"Permitir depuração USB?\". Marque \"Sempre permitir\" e toque em Permitir.")),
		widget.NewAccordionItem("Espelha mas o mouse não funciona (Xiaomi/POCO)",
			widget.NewLabel("Ative \"Depuração USB (configurações de segurança)\" nas Opções do desenvolvedor e reinicie o celular. Exige conta Mi e chip SIM.")),
		widget.NewAccordionItem("Dispositivo ficou offline",
			widget.NewLabel("Reconecte o cabo. Se persistir, feche o app, reconecte e verifique novamente.")),
		widget.NewAccordionItem("Pareamento Wi-Fi falhou",
			widget.NewLabel("O código expira em segundos: gere um novo em \"Parear dispositivo com código\". Confirme que celular e PC estão na mesma rede Wi-Fi.")),
	)
	for _, item := range faq.Items {
		if lbl, ok := item.Detail.(*widget.Label); ok {
			lbl.Wrapping = fyne.TextWrapWord
		}
	}

	logTitle := widget.NewLabelWithStyle("Registro de diagnóstico", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	logView := widget.NewLabel(readLogTail(15))
	logView.TextStyle = fyne.TextStyle{Monospace: true}
	logBg := canvas.NewRectangle(colCard)
	logBg.CornerRadius = 10
	logCard := container.NewStack(logBg, container.NewPadded(logView))

	copyBtn := widget.NewButtonWithIcon("Copiar registro", theme.ContentCopyIcon(), func() {
		u.win.Clipboard().SetContent(readLogTail(50))
	})
	refreshLogBtn := widget.NewButton("Atualizar", func() {
		logView.SetText(readLogTail(15))
	})

	return container.NewScroll(container.NewPadded(container.NewVBox(
		widget.NewLabelWithStyle("Problemas comuns", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		faq,
		widget.NewSeparator(),
		logTitle, logCard,
		container.NewHBox(copyBtn, refreshLogBtn),
	)))
}

// ---------- Pareamento Wi-Fi ----------

func (u *launcherUI) showPairDialog() {
	if u.adbPath == "" {
		u.detectADB(true)
		return
	}
	addr := widget.NewEntry()
	addr.SetPlaceHolder("192.168.0.10:37123")
	code := widget.NewEntry()
	code.SetPlaceHolder("6 dígitos")

	help := widget.NewLabel("No celular: Opções do desenvolvedor → Depuração sem fio → \"Parear dispositivo com código\".")
	help.Wrapping = fyne.TextWrapWord

	dialog.ShowForm("Parear por Wi-Fi", "Parear", "Cancelar",
		[]*widget.FormItem{
			widget.NewFormItem("", help),
			widget.NewFormItem("IP:PORTA de pareamento", addr),
			widget.NewFormItem("Código", code),
		},
		func(ok bool) {
			if !ok {
				return
			}
			u.statusLabel.SetText("Pareando...")
			go func() {
				out, err := Pair(u.adbPath, strings.TrimSpace(addr.Text), strings.TrimSpace(code.Text))
				logf("pareamento %s: %s err=%v", addr.Text, strings.TrimSpace(out), err)
				fyne.Do(func() {
					if err != nil {
						u.statusLabel.SetText("Falha no pareamento: " + strings.TrimSpace(out))
						return
					}
					u.statusLabel.SetText(strings.TrimSpace(out))
					u.askConnectAddress()
				})
			}()
		}, u.win)
}

func (u *launcherUI) askConnectAddress() {
	addr := widget.NewEntry()
	addr.SetPlaceHolder("192.168.0.10:5555")
	dialog.ShowForm("Conectar por Wi-Fi", "Conectar", "Cancelar",
		[]*widget.FormItem{widget.NewFormItem("IP:PORTA de conexão", addr)},
		func(ok bool) {
			if !ok {
				return
			}
			out, err := Connect(u.adbPath, strings.TrimSpace(addr.Text))
			logf("conexao wifi %s: %s err=%v", addr.Text, strings.TrimSpace(out), err)
			if err != nil {
				u.statusLabel.SetText("Falha: " + strings.TrimSpace(out))
			} else {
				u.statusLabel.SetText(strings.TrimSpace(out))
			}
			u.refreshDevices()
		}, u.win)
}
