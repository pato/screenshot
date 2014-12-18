package screenshot

import (
	"image"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func Setup() (*xgb.Conn, error) {
	return xgb.NewConn()
}

func Close(c *xgb.Conn) {
	c.Close()
}

func ScreenRect(c *xgb.Conn) (image.Rectangle, error) {
	screen := xproto.Setup(c).DefaultScreen(c)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	return image.Rect(0, 0, int(x), int(y)), nil
}

func CaptureScreen(c *xgb.Conn) (*image.RGBA, error) {
	r, e := ScreenRect(c)
	if e != nil {
		return nil, e
	}
	return CaptureRect(c, r)
}

func CaptureRect(c *xgb.Conn, rect image.Rectangle) (*image.RGBA, error) {
	screen := xproto.Setup(c).DefaultScreen(c)
	x, y := rect.Dx(), rect.Dy()
	xImg, err := xproto.GetImage(c, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * x, image.Rect(0, 0, x, y)}
	return img, nil
}
