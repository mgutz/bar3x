package base

import (
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/shimmerglass/bar3x/ui"
	"github.com/ungerik/go-cairo"
)

const heightChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Text struct {
	Base

	defaultColorKey    string
	defaultFontKey     string
	defaultFontSizeKey string

	setFont     string
	setFontSize float64
	setColor    color.Color
	text        string
	maxWidth    int
}

func NewText(p ui.ParentDrawable) *Text {
	return &Text{
		Base:               NewBase(p),
		defaultColorKey:    "text_color",
		defaultFontKey:     "text_font",
		defaultFontSizeKey: "text_font_size",
	}
}

func NewIcon(p ui.ParentDrawable) *Text {
	return &Text{
		Base:               NewBase(p),
		defaultColorKey:    "accent_color",
		defaultFontKey:     "icon_font",
		defaultFontSizeKey: "icon_font_size",
	}
}

func (t *Text) Text() string {
	return t.text
}
func (t *Text) SetText(s string) {
	t.text = s
	t.updateSize()
}

func (t *Text) Font() string {
	if t.setFont != "" {
		return t.setFont
	}
	return t.ctx.MustString(t.defaultFontKey)
}
func (t *Text) SetFont(f string) {
	t.setFont = f
	t.updateSize()
}

func (t *Text) FontSize() float64 {
	if t.setFontSize > 0 {
		return t.setFontSize
	}
	return t.ctx.MustFloat(t.defaultFontSizeKey)
}
func (t *Text) SetFontSize(f float64) {
	t.setFontSize = f
	t.updateSize()
}

func (t *Text) Color() color.Color {
	if t.setColor != nil {
		return t.setColor
	}
	return t.ctx.MustColor(t.defaultColorKey)
}
func (t *Text) SetColor(c color.Color) {
	t.setColor = c
}

func (t *Text) MaxWidth() int {
	return t.maxWidth
}
func (t *Text) SetMaxWidth(i int) {
	t.maxWidth = i
	t.updateSize()
}

func (t *Text) updateSize() {
	font := t.Font()
	fontSize := t.FontSize()

	if t.text == "" || fontSize == 0 {
		t.width.set(0)
		t.height.set(0)
		return
	}

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 0, 0)
	defer surface.Destroy()

	surface.SelectFontFace(font, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	surface.SetFontSize(fontSize)

	exX := surface.TextExtents(t.text)
	exY := surface.TextExtents(heightChars)
	surface.Finish()

	t.width.set(int(math.Ceil(exX.Xadvance)))
	t.height.set(int(math.Ceil(fontSize + (exY.Height + exY.Ybearing))))
}

func (t *Text) Draw(x, y int, im draw.Image) {
	font := t.Font()
	fontSize := t.FontSize()
	col := t.Color()

	if t.text == "" || fontSize == 0 {
		return
	}
	w, h := t.width.v, t.height.v

	rgba := color.RGBAModel.Convert(col).(color.RGBA)
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, w, h)
	defer surface.Destroy()

	surface.SelectFontFace(font, cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	surface.SetFontSize(fontSize)
	surface.SetSourceRGBA(
		float64(rgba.R)/255,
		float64(rgba.G)/255,
		float64(rgba.B)/255,
		float64(rgba.A)/255,
	)
	surface.MoveTo(0, fontSize)
	surface.ShowText(t.text)

	draw.Draw(
		im,
		image.Rect(x, y, x+w, y+h),
		surface.GetImage(),
		image.ZP,
		draw.Over,
	)
}