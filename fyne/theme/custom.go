package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 修改字库

// 方法一：切换本地字库
//func init() {
//	// 修改字符集支持中文
//	fontPaths := findfont.List()
//	for _, path := range fontPaths {
//		if strings.Contains(path, "simkai.ttf") {
//			os.Setenv("FYNE_FONT", path)
//			break
//		}
//	}
//}

// 方法二：使用自己提供的字库
//go:generate fyne bundle -package theme -o bundled.go assets

type MyTheme struct {
}

func (m *MyTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

// 修改字库，支持中文
func (m *MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return resourceTtf
}

func (m *MyTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *MyTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func TestMyTheme() {
	a := app.New()
	a.Settings().SetTheme(&MyTheme{})
	w := a.NewWindow("my theme")
	entry := widget.NewMultiLineEntry()
	entry.PlaceHolder = "Hello，你好。"
	w.SetContent(entry)
	w.ShowAndRun()
}
