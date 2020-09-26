package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	InfoColor    = "\033[1;34m"
	NoticeColor  = "\033[1;36m"
	WarningColor = "\033[1;33m"
	ErrorColor   = "\033[1;31m"
	DebugColor   = "\033[0;36m"
	SuccessColor = "\033[1;32m"
	ResetColor   = "\033[0;0m"
)

func logger (msg string, args... interface{}) {
	fmt.Println(">", fmt.Sprintf(msg, args...))
}

func success (msg string, args... interface{}) {
	fmt.Printf(SuccessColor)
	logger("🎉 " + msg, args...)
	fmt.Printf(ResetColor)
}

func loggerTitle (msg string, args... interface{}) {
	fmt.Println(NoticeColor, "====== ", fmt.Sprintf(msg, args...), " ======", ResetColor)
}

func loggerReset () {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func _error (msg string, args... interface{}) {
	fmt.Println(ErrorColor, "Fatal Error > ", fmt.Sprintf(msg, args...), ResetColor)
	os.Exit(0)
}

func loggerCount (i int, count int) {
	loggerReset()
	fmt.Println(WarningColor, "=============================")
	fmt.Println(WarningColor, "    🚀 NPM Booster 1.0.0    ")
	fmt.Println(WarningColor, "=============================")
	fmt.Printf("%sWe have actually build %d / %d package(s)%s\n", SuccessColor, i, count, ResetColor)
}