package service

import (
	"fmt"
	"go-ppocr/paddle"
	"go-ppocr/util"
)

func Test() {
	img := util.ReadImage("test.jpg")

	config := paddle.NewConfig()
	config.SetModel(util.GetString("model.det.path")+"/inference.pdmodel", util.GetString("model.det.path")+"/inference.pdiparams")
	predictor := paddle.NewPredictor(config)
	inNames := predictor.GetInputNames()
	inHandle := predictor.GetInputHandle(inNames[0])

	outNames := predictor.GetOutputNames()
	outHandle := predictor.GetOutputHandle(outNames[0])

	inHandle.CopyFromCpu(img)
	predictor.Run()

	outData := make([]float32, numElements(outHandle.Shape()))
	outHandle.CopyToCpu(outData)
	fmt.Println(outHandle.Lod())
}

func numElements(shape []int32) int32 {
	n := int32(1)
	for _, v := range shape {
		n *= v
	}
	return n
}
