package proxy

import (
	"fmt"
	"log"
)

const drwaing = "Drawing image "

type Image interface {
	Draw() string
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() string {
	return drwaing + b.filename
}

func NewBitmap(filename string) *Bitmap {
	log.Println("loading image from ", filename)
	return &Bitmap{filename: filename}
}

// LazyBitmap become the virtual proxy
// Draw function in LazyBitmap add additionally function to check if the bitmap is still empty
// so its just load the image file its empty unlike
// the naked Bitmap will always load the image, although it's not call the draw function

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (l *LazyBitmap) Draw() string {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	return l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	fmt.Println("loading image from ", filename)
	return &LazyBitmap{filename: filename}
}

func DrawOnImage(image Image) string {
	return image.Draw()
}
