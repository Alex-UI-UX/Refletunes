package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Log local mínimo para diagnóstico. Sem conteúdo de tela nem dados pessoais:
// apenas eventos do launcher (detecção, pareamento, erros).
var (
	logMu   sync.Mutex
	logFile *os.File
)

func logPath() string {
	exe, err := os.Executable()
	if err != nil {
		return "refletunes.log"
	}
	return filepath.Join(filepath.Dir(exe), "refletunes.log")
}

func logf(format string, args ...any) {
	logMu.Lock()
	defer logMu.Unlock()
	if logFile == nil {
		f, err := os.OpenFile(logPath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return
		}
		logFile = f
	}
	line := fmt.Sprintf(format, args...)
	fmt.Fprintf(logFile, "%s %s\n", time.Now().Format("2006-01-02 15:04:05"), line)
}

// readLogTail devolve as últimas n linhas do log para a tela de ajuda.
func readLogTail(n int) string {
	data, err := os.ReadFile(logPath())
	if err != nil {
		return "(sem registros ainda)"
	}
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	return strings.Join(lines, "\n")
}
