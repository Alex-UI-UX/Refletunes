package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// Device representa um aparelho Android visto pelo ADB.
type Device struct {
	Serial   string // serial ou ip:porta
	State    string // device | unauthorized | offline | authorizing | no permissions
	Model    string // model=...
	Transp   string // usb | tcp
	Selected bool
}

// StateLabel traduz o estado do adb para linguagem simples.
func (d Device) StateLabel() string {
	switch d.State {
	case "device":
		return "Pronto"
	case "unauthorized":
		return "Aguardando autorização no celular"
	case "authorizing":
		return "Autorizando..."
	case "offline":
		return "Offline — reconecte o cabo"
	case "no permissions":
		return "Sem permissão — veja a ajuda"
	default:
		return d.State
	}
}

// Ready indica se o dispositivo pode iniciar espelhamento.
func (d Device) Ready() bool { return d.State == "device" }

// ConnLabel descreve como o aparelho está conectado.
func (d Device) ConnLabel() string {
	if d.Transp == "usb" {
		return "Cabo USB"
	}
	return "Wi-Fi (rede)"
}

// FindADB localiza o binário adb: ao lado do executável, na pasta do pacote
// ou no PATH do sistema.
func FindADB() (string, error) {
	name := "adb"
	if runtime.GOOS == "windows" {
		name = "adb.exe"
	}
	exe, err := os.Executable()
	if err == nil {
		candidate := filepath.Join(filepath.Dir(exe), name)
		if st, e := os.Stat(candidate); e == nil && !st.IsDir() {
			return candidate, nil
		}
	}
	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}
	return "", fmt.Errorf("%s não encontrado", name)
}

// runADB executa o adb com argumentos de uma lista controlada (nunca shell).
func runADB(adbPath string, args ...string) (string, error) {
	cmd := exec.Command(adbPath, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

var devLineRe = regexp.MustCompile(`^(\S+)\s+(\S+)(.*)$`)

// ListDevices executa `adb devices -l` e interpreta o resultado.
func ListDevices(adbPath string) ([]Device, error) {
	out, err := runADB(adbPath, "devices", "-l")
	if err != nil {
		return nil, fmt.Errorf("falha ao listar dispositivos: %v\n%s", err, out)
	}
	var devices []Device
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "List of devices") || strings.HasPrefix(line, "*") {
			continue
		}
		m := devLineRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		d := Device{Serial: m[1], State: m[2], Transp: "tcp"}
		if !strings.Contains(d.Serial, ":") && !strings.Contains(d.Serial, ".") {
			d.Transp = "usb"
		}
		for _, kv := range strings.Fields(m[3]) {
			if strings.HasPrefix(kv, "model:") {
				d.Model = strings.TrimPrefix(kv, "model:")
			}
			if kv == "usb" {
				d.Transp = "usb"
			}
		}
		devices = append(devices, d)
	}
	return devices, nil
}

var ipPortRe = regexp.MustCompile(`^(\d{1,3}\.){3}\d{1,3}:\d{1,5}$`)

// Pair executa `adb pair ip:porta código` (Android 11+, pareamento Wi-Fi).
func Pair(adbPath, addr, code string) (string, error) {
	if !ipPortRe.MatchString(addr) {
		return "", fmt.Errorf("endereço inválido, use o formato IP:PORTA exibido no celular")
	}
	code = strings.TrimSpace(code)
	if len(code) != 6 {
		return "", fmt.Errorf("o código de pareamento tem 6 dígitos")
	}
	return runADB(adbPath, "pair", addr, code)
}

// Connect executa `adb connect ip:porta`.
func Connect(adbPath, addr string) (string, error) {
	if !ipPortRe.MatchString(addr) {
		return "", fmt.Errorf("endereço inválido, use o formato IP:PORTA")
	}
	return runADB(adbPath, "connect", addr)
}

// Disconnect desconecta um dispositivo Wi-Fi (ou todos se addr vazio).
func Disconnect(adbPath, addr string) (string, error) {
	if addr == "" {
		return runADB(adbPath, "disconnect")
	}
	if !ipPortRe.MatchString(addr) {
		return "", fmt.Errorf("endereço inválido")
	}
	return runADB(adbPath, "disconnect", addr)
}
