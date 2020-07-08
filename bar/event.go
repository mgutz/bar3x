package bar

import (
	"image"

	"github.com/shimmerglass/bar3x/ui"
)

func (b *Bar) displatchEvent(ev ui.Event) {
	{
		w, h := b.LeftRoot.Width(), b.LeftRoot.Height()
		x, y := b.padding, (b.height-h)/2

		if ev.At.In(image.Rect(x, y, x+w, y+h)) {
			ev.At = ev.At.Sub(image.Pt(x, y))
			b.LeftRoot.SendEvent(ev)
		}
	}

	{
		w, h := b.CenterRoot.Width(), b.CenterRoot.Height()
		x, y := (b.screen.Width-w)/2, (b.height-h)/2

		if ev.At.In(image.Rect(x, y, x+w, y+h)) {
			ev.At = ev.At.Sub(image.Pt(x, y))
			b.CenterRoot.SendEvent(ev)
		}
	}

	{
		w, h := b.RightRoot.Width(), b.RightRoot.Height()
		x, y := b.screen.Width-w-b.padding, (b.height-h)/2

		if ev.At.In(image.Rect(x, y, x+w, y+h)) {
			ev.At = ev.At.Sub(image.Pt(x, y))
			b.RightRoot.SendEvent(ev)
		}
	}
}