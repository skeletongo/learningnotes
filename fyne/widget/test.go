package widget

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Test1() {
	a := app.New()
	w := a.NewWindow("My Widget")

	// 标签
	label := widget.NewLabel("标签")
	// 按钮
	button := widget.NewButton("按钮", func() {
		log.Println("push button")
	})
	// 输入框
	input := widget.NewEntry()
	input.PlaceHolder = "please input"
	input.OnSubmitted = func(s string) {
		log.Printf("input submit: %v\n", s)
	}
	input.OnChanged = func(s string) { // 内容有变化时
		log.Printf("input change: %v\n", s)
	}
	input.OnCursorChanged = func() { // 获取焦点时
		log.Println("input cursor change")
	}
	// 选项
	check := widget.NewCheck("Optional", func(value bool) {
		log.Println("Check set to", value)
	})
	// 单选组
	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Radio set to", value)
	})
	// 下拉单选框
	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})
	// 表单提交
	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Entry", Widget: entry}},
		SubmitText: "提交",
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
		},
		CancelText: "取消",
		OnCancel: func() {
			log.Println("Form cancel")
		},
	}
	form.Append("Text", textArea)
	// 进度条
	progress := widget.NewProgressBar()         // 显示百分比
	infinite := widget.NewProgressBarInfinite() // 进行中
	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()
	// 工具栏
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(), // 分割线
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(), // 垫片，占满剩余空间
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	// 列表项
	itemLabel := widget.NewLabel("")
	var data = []string{"a", "bb", "string", "list"}
	list := widget.NewList(
		func() int { // 选项数量
			//log.Println("list length")
			return len(data)
		},
		func() fyne.CanvasObject { // 选项创建方法
			log.Println("list create")
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) { // 选项更新方法
			log.Println("list update")
			o.(*widget.Label).SetText(data[i])
		})
	list.OnSelected = func(id widget.ListItemID) { // 选中时执行
		log.Printf("list onSelected %v\n", id)
		itemLabel.SetText(data[id])
	}
	list.OnUnselected = func(id widget.ListItemID) { // 失去选中时执行
		log.Printf("list onUnselected %v\n", id)
	}
	list.Select(1) // 默认选中哪个
	listBox := container.New(layout.NewGridWrapLayout(
		fyne.NewSize(list.MinSize().Width, list.MinSize().Height*float32(list.Length()))),
		list, itemLabel)

	// 表格项
	itemLabel2 := widget.NewLabel("")
	var datas = [][]string{[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"}}
	listTable := widget.NewTable(
		func() (int, int) { // 行，列
			return len(datas), len(datas[0])
		},
		func() fyne.CanvasObject { // 创建一格
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) { // 刷新一格
			o.(*widget.Label).SetText(datas[i.Row][i.Col])
		})
	listTable.OnSelected = func(id widget.TableCellID) { // 选中时执行
		log.Printf("listTable onSelected %v\n", id)
		itemLabel2.SetText(datas[id.Row][id.Col])
	}
	listTable.OnUnselected = func(id widget.TableCellID) { // 失去选中时执行
		log.Printf("listTable onUnselected %v\n", id)
	}
	listTable.Select(widget.TableCellID{
		Row: 1,
		Col: 1,
	}) // 默认选中哪个
	r, c := listTable.Length()
	listTableBox := container.New(layout.NewGridWrapLayout(
		fyne.NewSize(listTable.MinSize().Width*float32(c), listTable.MinSize().Height*float32(r))),
		listTable, itemLabel2)

	w.SetContent(container.NewVBox(
		label, button, input,
		check, radio, combo,
		form,
		progress, infinite,
		toolbar,
		listBox, listTableBox))
	w.ShowAndRun()
}
