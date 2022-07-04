package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

type Directions struct {
	widget.BaseWidget

	Up, Down, Left, Right         *widget.Button
	OnUp, OnDown, OnLeft, OnRight func()
}

func NewDirections() *Directions {
	var ret *Directions
	ret = &Directions{
		Up: widget.NewButtonWithIcon("", theme.DefaultTheme().Icon(theme.IconNameMoveUp), func() {
			ret.OnUp()
		}),
		Down: widget.NewButtonWithIcon("", theme.DefaultTheme().Icon(theme.IconNameMoveDown), func() {
			ret.OnDown()
		}),
		Left: widget.NewButtonWithIcon("", theme.DefaultTheme().Icon(theme.IconNameNavigateBack), func() {
			ret.OnLeft()
		}),
		Right: widget.NewButtonWithIcon("", theme.DefaultTheme().Icon(theme.IconNameNavigateNext), func() {
			ret.OnRight()
		}),
	}
	ret.ExtendBaseWidget(ret)
	return ret
}

func (w *Directions) CreateRenderer() fyne.WidgetRenderer {
	return &DirectionsRenderer{directions: w}
}

type DirectionsRenderer struct {
	directions *Directions
}

func (d *DirectionsRenderer) Destroy() {
}

// canvas绘图是使用size大小
func (d *DirectionsRenderer) Layout(size fyne.Size) {
	w := size.Width / 3
	h := size.Height / 3
	if d.directions.Up != nil {
		d.directions.Up.Resize(fyne.NewSize(w, h))
		d.directions.Up.Move(fyne.NewPos(size.Width/2-d.directions.Up.Size().Width/2, 0))
	}
	if d.directions.Down != nil {
		d.directions.Down.Resize(fyne.NewSize(w, h))
		d.directions.Down.Move(fyne.NewPos(
			size.Width/2-d.directions.Down.Size().Width/2, size.Height-d.directions.Down.Size().Height))
	}
	if d.directions.Left != nil {
		d.directions.Left.Resize(fyne.NewSize(w, h))
		d.directions.Left.Move(fyne.NewPos(0, size.Height/2-d.directions.Left.Size().Height/2))
	}
	if d.directions.Right != nil {
		d.directions.Right.Resize(fyne.NewSize(w, h))
		d.directions.Right.Move(fyne.NewPos(
			size.Width-d.directions.Right.Size().Width, size.Height/2-d.directions.Right.Size().Height/2))
	}
}

func (d *DirectionsRenderer) width() float32 {
	ret := float32(0)
	if d.directions.Left != nil {
		ret += d.directions.Left.MinSize().Width
	}
	if d.directions.Right != nil {
		ret += d.directions.Right.MinSize().Width
	}
	if d.directions.Up != nil {
		ret += d.directions.Up.MinSize().Width
	} else if d.directions.Down != nil {
		ret += d.directions.Down.MinSize().Width
	}
	return ret
}

func (d *DirectionsRenderer) height() float32 {
	ret := float32(0)
	if d.directions.Up != nil {
		ret += d.directions.Up.MinSize().Height
	}
	if d.directions.Down != nil {
		ret += d.directions.Down.MinSize().Height
	}
	if d.directions.Left != nil {
		ret += d.directions.Left.MinSize().Height
	} else if d.directions.Right != nil {
		ret += d.directions.Right.MinSize().Height
	}
	return ret
}

func (d *DirectionsRenderer) MinSize() fyne.Size {
	return fyne.NewSize(d.width(), d.height())
}

func (d *DirectionsRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{d.directions.Up, d.directions.Down, d.directions.Left, d.directions.Right}
}

func (d *DirectionsRenderer) Refresh() {
	for _, v := range d.Objects() {
		v.Refresh()
	}
}

func TestDirectionsRenderer() {
	a := app.New()
	w := a.NewWindow("custom widget")
	w.SetPadded(false)

	d := NewDirections()
	d.OnUp = func() {
		log.Println("on up")
	}
	d.OnDown = func() {
		log.Println("on down")
	}
	d.OnLeft = func() {
		log.Println("on left")
	}
	d.OnRight = func() {
		log.Println("on right")
	}

	w.SetContent(d)
	w.Resize(fyne.NewSize(150, 150))
	w.ShowAndRun()
}
