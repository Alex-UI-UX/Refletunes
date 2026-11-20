package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// QualityProfile define um conjunto fixo e validado de flags do scrcpy.
// Nunca montar flags a partir de entrada do usuário.
type QualityProfile struct {
	Name string
	Desc string
	Args []string
}

// Perfis baseados nas flags reais do scrcpy v5.x (--help da versão fixada).
var Profiles = []QualityProfile{
	{
		Name: "Padrão",
		Desc: "Qualidade equilibrada (recomendado)",
		Args: nil, // usa os defaults do scrcpy
	},
	{
		Name: "Desempenho",
		Desc: "Resolução e bitrate menores, mais leve",
		Args: []string{"--max-size=1024", "--video-bit-rate=2M", "--max-fps=30"},
	},
}

// FindScrcpy localiza o executável do scrcpy ao lado do launcher.
func FindScrcpy() (string, error) {
	name := "scrcpy"
	if runtime.GOOS == "windows" {
		name = "scrcpy.exe"
	}
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	candidate := filepath.Join(filepath.Dir(exe), name)
	if st, e := os.Stat(candidate); e == nil && !st.IsDir() {
		return candidate, nil
	}
	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}
	return "", fmt.Errorf("o componente de espelhamento não foi encontrado na pasta do %s", appName)
}

// StartMirror inicia o scrcpy para um dispositivo e perfil. Só argumentos da
// lista permitida são passados; o serial vem de `adb devices`, não de texto livre.
func StartMirror(scrcpyPath string, d Device, p QualityProfile) *exec.Cmd {
	args := append([]string{}, p.Args...)
	// -s seleciona o dispositivo explicitamente (evita ambiguidade com vários)
	if d.Serial != "" {
		args = append([]string{"-s", d.Serial}, args...)
	}
	cmd := exec.Command(scrcpyPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Start()
	return cmd
}
