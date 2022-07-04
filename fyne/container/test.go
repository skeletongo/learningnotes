package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// Test1
// 水平布局 HBox
// 垂直布局 VBox
func Test1() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	// layout.NewSpacer() 垫片，占满剩余空间
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()
}

// Test2
// 网格布局 Grid
// 占满容器的网格布局
func Test2() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)

	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}

// Test3
// 网格布局 GridWrap
// 自定义网格大小，自动换行
func Test3() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")
	myWindow.SetPadded(false)

	text1 := canvas.NewText("1", color.White)
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 100)),
		text1, text2, text3)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(400, 100))
	myWindow.ShowAndRun()
}

// Test4
// 边框布局 Border
// container.New(layout.NewBorderLayout(top, down, left, right), top, down, left, right, middle)
//										上	  下	左		右	  上	下	 左		右		中
// layout.NewBorderLayout 中某个位置为nil, 但是后面参数位置不会nil，则会把元素放到中间位置
func Test4() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("up", color.White)
	down := canvas.NewText("down", color.White)
	left := canvas.NewText("left", color.White)
	right := canvas.NewText("right", color.White)
	middle := canvas.NewText("middle", color.White)
	content := container.New(layout.NewBorderLayout(top, down, left, right), top, down, left, right, middle)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// Test5
// 表单布局 Form
func Test5() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Layout")

	label1 := canvas.NewText("Label 1", color.Black)
	value1 := canvas.NewText("Value", color.White)
	label2 := canvas.NewText("Label 2", color.Black)
	value2 := canvas.NewText("Something", color.White)
	grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}

// Test6
// 中心布局 Center
// 所有元素放在容器中心位置
func Test6() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Center Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	// canvas绘图是使用size大小
	// 中心布局会将size设置为minSize,所以这里要修改minSize
	img.SetMinSize(fyne.NewSize(100, 100))
	text := canvas.NewText("Overlay", color.White)
	content := container.New(layout.NewCenterLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// Test7
// 最大布局 Max
// 元素大小和容器一样
func Test7() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")

	img := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Overlay", color.White)
	content := container.New(layout.NewMaxLayout(), img, text)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// Test8
// 标签布局
func Test8() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))
	// 设置标签所在位置
	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
