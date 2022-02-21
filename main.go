package main

import (
	"github.com/Akitora-R/gg"
	"github.com/disintegration/imaging"
	"image"
	"io/ioutil"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func gen(txt string, images []image.Image) {
	start := time.Now().UnixMilli()
	fontFace, err := gg.LoadFontFace("fonts/msyh.ttc", 18)
	if err != nil {
		panic(err)
	}
	const w float64 = 1000
	const gap float64 = 50
	h := gap * 3

	strHeight := gg.MeasureStringHeight(txt, 900, 18, 1.75, fontFace, true)
	h += strHeight
	for _, i2 := range images {
		h += float64(i2.Bounds().Max.Y)
	}
	h += float64(len(images)) * gap

	dc := gg.NewContext(int(w), int(h))
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	log.Println(strHeight)
	dc.SetFontFace(fontFace)
	dc.DrawStringWrapped(txt, gap, gap, 0, 0, 900, 1.75, gg.AlignLeft, true)
	y := gap*2 + strHeight
	for _, i := range images {
		dc.DrawImage(i, int(gap), int(y))
		y += gap + float64(i.Bounds().Max.Y)
	}
	if err = dc.SaveJPG("out.jpg", 85); err != nil {
		panic(err)
	}
	log.Println("耗时ms:", time.Now().UnixMilli()-start)
}

func main() {
	gen(getTxt(), getImg())
}

func getTxt() string {
	t, err := ioutil.ReadFile("texts/text1.txt")
	if err != nil {
		panic(err)
	}
	return string(t)
}

func getImg() []image.Image {
	var imgs []image.Image
	dir, _ := ioutil.ReadDir("images")
	for _, fileInfo := range dir {
		jpg, _ := gg.LoadJPG("images/" + fileInfo.Name())
		imgs = append(imgs, imaging.Resize(jpg, 900, 0, imaging.Linear))
	}
	return imgs
}
