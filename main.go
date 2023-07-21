package main

import (
	"go-ppocr/service"
	"go-ppocr/util"
)

func main() {
	println("run demo")
	util.InitConfig()
	service.Test()
}
