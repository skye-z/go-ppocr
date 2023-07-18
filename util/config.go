/*
全局配置服务

BetaX Micro OCR
Copyright © 2023 SkyeZhang <skai-zhang@hotmail.com>
*/

package util

import (
	"fmt"

	"github.com/spf13/viper"
)

const Version = "1.0.0"

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefault()
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println(err)
		}
	}
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	viper.WriteConfig()
}

func GetString(key string) string {
	return viper.GetString(key)
}

func createDefault() {
	// 是否启用GPU
	viper.SetDefault("basic.gpu", "false")
	// GPU的设备编号
	viper.SetDefault("basic.gpu.id", "0")
	// GPU初始分配(最低)显存大小(单位MB)
	viper.SetDefault("basic.gpu.min-memory", "100")
	// 是否启用TensorRt加速(仅N卡)
	viper.SetDefault("basic.gpu.trt", "false")
	// 是否启用动态子图
	viper.SetDefault("basic.gpu.trt-dynamic-shape", "false")
	// 动态子图大小
	viper.SetDefault("basic.gpu.trt-dynamic-shape-size", "1")

	// CPU运算线程数
	viper.SetDefault("basic.cpu.thread", "1")
	// 是否启用MKLDNN加速
	viper.SetDefault("basic.cpu.mkldnn", "false")
	// 是否启用MKLDNN BF16加速(需CPU支持AVX512)
	viper.SetDefault("basic.cpu.mkldnn-bf16", "false")

	// 文本检测模型
	viper.SetDefault("model.det.path", "inference/ch_PP-OCRv3_det_infer")
	// 方向分类模型
	viper.SetDefault("model.cls.path", "inference/ch_ppocr_mobile_v2.0_cls_infer")
	// 文本识别模型
	viper.SetDefault("model.rec.path", "inference/ch_PP-OCRv3_rec_infer")

	viper.SafeWriteConfig()
}
