package bind

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

// 数据绑定，修改数据后所有使用的其它地方也同步修改

func Test1() {
	// 创建新数据类型
	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)

	// 引用已有数据
	myInt := 5
	boundInt := binding.BindInt(&myInt)
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)

	// 同步修改

	boundInt.Set(6)
	i, _ = boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)

	myInt = 7
	i, _ = boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)
}

func Test2() {
	// 时钟显示
	a := app.New()
	w := a.NewWindow("binding")

	s := binding.NewString()
	label := widget.NewLabelWithData(s)
	w.SetContent(label)

	go func() {
		for i := 0; ; i++ {
			s.Set(time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(time.Second)
		}
	}()

	w.ShowAndRun()
}

func Test3() {
	// 双向绑定，就是调用Set方法
	a := app.New()
	w := a.NewWindow("binding")

	str := binding.NewString()
	str.Set("Hi!")

	w.SetContent(container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	))
	w.ShowAndRun()
}

func Test4() {
	// 数据转换
	myApp := app.New()
	w := myApp.NewWindow("Conversion")

	f := binding.NewFloat() // 源数据
	str := binding.FloatToString(f)
	short := binding.FloatToStringWithFormat(f, "%0.0f%%")
	f.Set(25.0)

	w.SetContent(container.NewVBox(
		widget.NewSliderWithData(0, 100.0, f),
		widget.NewLabelWithData(str),
		widget.NewLabelWithData(short),
	))

	w.ShowAndRun()
}

func Test5() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Append", func() {
		// list与data绑定，所以修改data就可以同步修改list
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))
	myWindow.ShowAndRun()
}

func Test6() {
	// 测试Test5另一种绑定方式
	myApp := app.New()
	myWindow := myApp.NewWindow("List Data")

	arr := make([]string, 0, 5)
	data := binding.BindStringList(&arr)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		//arr = append(arr, val) // 不行，需要调用对象方法
		data.Append(val)
		log.Println(arr)
	})
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))
	myWindow.ShowAndRun()
}
