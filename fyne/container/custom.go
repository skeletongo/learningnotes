package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

// LeftDownLayout 左下角布局
// 元素都放左下角
type LeftDownLayout struct {
}

// Layout 摆放元素
// objects 包含的所有元素
// size 容器大小
func (l *LeftDownLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(0, size.Height-l.MinSize(objects).Height)
	for _, v := range objects {
		v.Move(pos)
		v.Resize(v.MinSize()) // canvas绘制图形时是使用size绘图的
		pos = pos.AddXY(v.MinSize().Width, v.MinSize().Height)
	}
}

// MinSize 元素占用大小
func (l *LeftDownLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	h, w := float32(0), float32(0)
	for _, v := range objects {
		w += v.MinSize().Width
		h += v.MinSize().Height
	}
	return fyne.NewSize(w, h)
}

func TestCustomLayout() {
	a := app.New()
	w := a.NewWindow("Left Down Layout")
	w.SetPadded(false)

	r := canvas.NewRectangle(color.White)
	r.SetMinSize(fyne.NewSize(20, 20))
	r2 := canvas.NewRectangle(color.White)
	r2.SetMinSize(fyne.NewSize(30, 30))
	r3 := canvas.NewRectangle(color.White)
	r3.SetMinSize(fyne.NewSize(40, 40))

	w.SetContent(container.New(new(LeftDownLayout), r, r2, r3))
	w.Resize(fyne.NewSize(150, 150))
	w.ShowAndRun()
}
