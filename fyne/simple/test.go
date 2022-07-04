package simple

import (
	"image/color"
	"log"
	"math/rand"
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}

// 新建窗口和画布
func Test1() {
	myApp := app.New()                    // 新建应用
	myWindow := myApp.NewWindow("Canvas") // 新建窗口
	myWindow.SetMaster()                  // 主窗口，关闭主窗口程序会退出

	myCanvas := myWindow.Canvas() // 画布对象

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect) // 画一个矩形

	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		rect.FillColor = green
		rect.Refresh() // 修改矩形填充颜色

		time.Sleep(time.Second)
		setContentToText(myCanvas) // 标签

		time.Sleep(time.Second)
		setContentToCircle(myCanvas) // 圆
	}()

	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.ShowAndRun()
}

// 没有布局的布局
func Test2() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("There", green)
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewWithoutLayout(text1, text2) // 没有布局的容器

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// 折叠栏
func Test3() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	myWindow.SetContent(widget.NewAccordion( // 折叠栏
		widget.NewAccordionItem("A", widget.NewLabel("hello")),
	))

	myWindow.ShowAndRun()
}

type myEntry struct {
	widget.Entry
}

func NewMyEntry() *myEntry {
	e := new(myEntry)
	e.ExtendBaseWidget(e)
	return e
}

// 获取焦点后按快捷键会调用这个方法
func (m *myEntry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		m.Entry.TypedShortcut(s) // 优先使用组件本身的快捷键
		return
	}
	// 自定义快捷键，快捷键名称和自定义处理方法
	ctrlDown := &desktop.CustomShortcut{KeyName: fyne.KeyDown,
		Modifier: fyne.KeyModifierControl | fyne.KeyModifierShift}
	if s.ShortcutName() == ctrlDown.ShortcutName() { // 名称匹配
		log.Println("Ctrl+Shift+Down")
	} else {
		log.Println("Shortcut typed:", s)
	}
}

// 全局快捷键和组件快捷键
func Test4() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")

	// 全局快捷键
	ctrlUp := &desktop.CustomShortcut{KeyName: fyne.KeyUp, Modifier: fyne.KeyModifierControl}
	myWindow.Canvas().AddShortcut(ctrlUp, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Up")
	})

	// 组件获取焦点后的快捷键，组件快捷键
	e := NewMyEntry()
	myWindow.SetContent(e)
	myWindow.ShowAndRun()
}

// 本地存储
func Test5() {
	a := app.NewWithID("com.example.tutorial.preferences") // 1.唯一id
	a.Preferences().SetInt("one", 1)
	a.Preferences().SetFloat("float", 1.2)
	a.Preferences().SetBool("bool", true)
	a.Preferences().SetString("string", "sss")

	log.Println(a.Preferences().Int("one"))
	log.Println(a.Preferences().Int("two"))
	log.Println(a.Preferences().IntWithFallback("three", 3))
	log.Println(a.Preferences().Float("float"))
	log.Println(a.Preferences().Bool("bool"))
	log.Println(a.Preferences().String("string"))
}

// 系统托盘应用图标
func Test6() {
	a := app.New()
	w := a.NewWindow("SysTray")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show() // 显示窗口
			}))
		desk.SetSystemTrayMenu(m)
	}

	w.SetContent(widget.NewLabel("Fyne System Tray"))
	w.SetCloseIntercept(func() { // 重写窗口关闭处理方法
		w.Hide() // 隐藏窗口
	})
	w.ShowAndRun()
}

// 数据绑定
func Test7() {
	a := app.New()
	w := a.NewWindow("Hello")

	str := binding.NewString()
	go func() {
		dots := "....."
		for i := 5; i >= 0; i-- {
			str.Set("Count down" + dots[:i])
			time.Sleep(time.Second)
		}
		str.Set("Blast off!")
	}()

	w.SetContent(widget.NewLabelWithData(str)) // 随着str的修改，标签内容也会变化
	w.ShowAndRun()
}

// 简单图形
func Test8() {
	myApp := app.New()
	w := myApp.NewWindow("Line")

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1
	line.Position1 = fyne.NewPos(0, 0)
	line.Position2 = fyne.NewPos(50, 50)

	r := canvas.NewRectangle(color.White)
	r.Resize(fyne.NewSize(50, 50))
	r.Move(fyne.NewPos(50, 50))

	circle := canvas.NewCircle(color.White)
	circle.StrokeColor = color.Black
	circle.StrokeWidth = 1
	circle.Position1 = fyne.NewPos(50, 50)
	circle.Position2 = fyne.NewPos(100, 100)

	w.SetContent(container.NewWithoutLayout(line, r, circle))
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}

// 图片
func Test9() {
	myApp := app.New()
	w := myApp.NewWindow("Image")

	//image := canvas.NewImageFromResource(theme.FyneLogo())
	// image := canvas.NewImageFromURI(uri)
	// image := canvas.NewImageFromImage(src)
	// image := canvas.NewImageFromReader(reader, name)
	image := canvas.NewImageFromFile("Icon.png")
	image.FillMode = canvas.ImageFillStretch // 填满，自动拉伸，默认值
	//image.FillMode = canvas.ImageFillContain // 填充，自动拉伸，保持长宽比
	//image.FillMode = canvas.ImageFillOriginal //
	w.SetContent(image)

	w.ShowAndRun()
}

// 光栅
func Test10() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	raster := canvas.NewRasterWithPixels(
		func(_, _, w, h int) color.Color {
			return color.RGBA{uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255)), 0xff}
		})
	raster.Resize(fyne.NewSize(100, 100))

	gradient := canvas.NewHorizontalGradient(color.White, color.Transparent)
	gradient.Resize(fyne.NewSize(100, 100))
	gradient.Move(fyne.NewPos(100, 0))

	gradient1 := canvas.NewRadialGradient(color.White, color.Transparent)
	gradient1.Resize(fyne.NewSize(100, 100))
	gradient1.Move(fyne.NewPos(200, 0))

	w.SetContent(container.NewWithoutLayout(raster, gradient, gradient1))
	w.Resize(fyne.NewSize(300, 100))
	w.ShowAndRun()
}

// 动画
func Test11() {
	a := app.New()
	w := a.NewWindow("Hello")

	obj := canvas.NewRectangle(color.Black)
	obj.Resize(fyne.NewSize(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))

	red := color.NRGBA{R: 0xff, A: 0xff}
	blue := color.NRGBA{B: 0xff, A: 0xff}
	canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(c color.Color) {
		obj.FillColor = c                               // 修改颜色
		obj.Move(obj.Position().Add(fyne.NewPos(2, 0))) // 移动一点位置
		canvas.Refresh(obj)
	}).Start()

	w.Resize(fyne.NewSize(250, 50))
	w.SetPadded(false) // 去掉窗口内边距
	w.ShowAndRun()
}

func Test12() {
	a := app.New()

	l, err := url.Parse("https://www.baidu.com")
	if err != nil {
		log.Println(err)
	}
	if err = a.OpenURL(l); err != nil {
		log.Println(err)
	}
	a.Run()
}
