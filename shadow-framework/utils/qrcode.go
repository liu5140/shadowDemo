package utils

import (
	"bytes"
	"image"
	"image/png"

	qrcode "github.com/skip2/go-qrcode"
)

func MakeQRCode(instr string, imgSize int) (string, error) {

	var pngbyte []byte
	pngbyte, err := qrcode.Encode(instr, qrcode.Low, imgSize)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(pngbyte))
	if err != nil {
		return "", err
	}

	var subImg image.Image
	if rgbImg, ok := img.(*image.RGBA); ok {

		top := int(float64(imgSize) * 0.1)
		left := int(float64(imgSize) * 0.1)
		bottom := int(float64(imgSize) * 0.9)
		right := int(float64(imgSize) * 0.9)

		subImg = rgbImg.SubImage(image.Rect(left, top, right, bottom)).(*image.RGBA) //图片裁剪x0 y0 x1 y1

	}
	if subImg != nil {
		img = subImg
	}
	// _ = subImg
	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return "", err
	}

	pngbyte = buf.Bytes()

	imgb64 := Base64Encode(pngbyte)
	pngstr := "data:image/png;base64," + string(imgb64)

	return pngstr, nil
}
