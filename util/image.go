package util

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"gocv.io/x/gocv"
	"gonum.org/v1/gonum/mat"
)

func ReadImage(image_path string) gocv.Mat {
	println("read image")
	img := gocv.IMRead(image_path, gocv.IMReadColor)
	if img.Empty() {
		log.Printf("Could not read image %s\n", image_path)
		os.Exit(1)
	}
	return img
}

func ReadImage1(image_path string) []float32 {
	// 打开图片文件
	f, err := os.Open(image_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 解码图片
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	// 获取图片的长宽
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// 创建一个长度为 width * height * 3 的 []float32 类型的向量
	vector := make([]float32, width*height*3)

	// 遍历图片的每个像素点，将其红、绿、蓝三个分量分别存入向量
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// 获取像素点的颜色
			color := img.At(x, y)
			// 将颜色转换为 uint8 类型
			r, g, b, _ := color.RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			// 将红、绿、蓝三个分量转换为 float32 类型并存入向量
			vector[(x+y*width)*3] = float32(r)
			vector[(x+y*width)*3+1] = float32(g)
			vector[(x+y*width)*3+2] = float32(b)
		}
	}
	return vector
}

func ReadImage2(image_path string) []float32 {
	// 打开图片文件
	f, err := os.Open(image_path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 解码图片
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	// 获取图片的长宽
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Create a new dense matrix.
	m := mat.NewDense(height, width, nil)

	// Copy the image data into the matrix.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			m.Set(y, x, float64((r+g+b)/3))
		}
	}
	data := make([]float32, m.RawMatrix().Stride)
	for i := 0; i < m.RawMatrix().Stride; i++ {
		data[i] = float32(m.At(i/m.RawMatrix().Cols, i%m.RawMatrix().Cols))
	}
	return data
}
