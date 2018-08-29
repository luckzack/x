package main

import (
	"fmt"
	//l4g "code.google.com/p/log4go"
	l4g "github.com/gogoods/x/logger"
)

var (
	//Logger = l4g.NewDefaultLogger(l4g.INFO)
	Logger = l4g.NewDefaultLogger(l4g.INFO)
)

const (
	LogInfo = l4g.INFO
	LogWarn = l4g.WARNING
	LogErr  = l4g.ERROR
)

func init() {

	flw := l4g.NewFileLogWriter("basic.log", false)

	//flw.SetFormat("%D %T [%L] (%S) %M")
	flw.SetFormat("%D %T %L (%S) %M")
	flw.SetRotate(true)
	//	flw.SetRotateSize(0)
	//	flw.SetRotateLines(0)
	flw.SetRotateDaily(true)

	Logger.AddFilter("file", l4g.FINE, flw)

	Logger.Info("@@@@@@@@@@@@@@@@")
	Logger.Error("!!!!!!!!!")
}

func main() {
	fmt.Println("ready")
}
