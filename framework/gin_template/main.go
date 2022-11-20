package main

import (
	"gin_template/core"
	"gin_template/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	global.GLA_VIPER = core.Viper()
	core.RunWindowsServer()
}
