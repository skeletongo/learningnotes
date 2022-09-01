package facade

import "fmt"

// 外观模式
// 简化功能调用,将复杂的接口或方法调用整合到一个对象中，由这个对象封装复杂度

type Sound struct {
}

func (s *Sound) PlaySound() {
	fmt.Println("播放声音")
}

type Video struct {
}

func (v *Video) PlayVideo() {
	fmt.Println("播放视频")
}

// Facade 外观对象
type Facade struct {
	Sound
	Video
}

func (f *Facade) PlayVideo() {
	f.Sound.PlaySound()
	f.Video.PlayVideo()
}
