package adapter

import "fmt"

// 适配器模式
// 当被使用对象的行为不符合使用者的规范时，通过一个适配器将被使用者对象的行为转换成调用者可以使用的规范
// 1.确定被调用者行为
// 2.确认调用者需要的规范
// 3.创建适配器使被调用者行为符合调用者的规范

// FileList 打印当前目录所有文件接口
type FileList interface {
	Ls()
}

type Linux struct {
}

func (l *Linux) Ls() { // linux 是 ls 命令
	fmt.Println("linux ls")
}

type Windows struct {
}

func (w *Windows) Dir() { //  windows 是 dir 命令
	fmt.Println("windows dir")
}

// List 打印当前目录所有文件
func List(list FileList) {
	list.Ls()
}

type WindowsAdapter struct {
	*Windows
}

func (w *WindowsAdapter) Ls() {
	w.Dir()
}
