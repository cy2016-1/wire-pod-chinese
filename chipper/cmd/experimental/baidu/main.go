package main

//luowei add the file

import (
	"github.com/kercre123/chipper/pkg/initwirepod"
	stt "github.com/kercre123/chipper/pkg/wirepod/stt/baidu"
)

func main() {
	initwirepod.StartFromProgramInit(stt.Init, stt.STT, stt.Name)
}
