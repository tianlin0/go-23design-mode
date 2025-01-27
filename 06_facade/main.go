package main

import "fmt"

/*
定义两个及以上子系统，合成为一个整体系统，整体系统可以独立运行
- 定义修复器，将多个子系统组合在一起
  - AudioMixer：修复音频
  - VideoMixer：修复视频
*/
type fixer interface {
	Fix(name string)
}

// AudioMixer Subsystem 1：修复音频
type AudioMixer struct {
}

func (a *AudioMixer) Fix(name string) {
	fmt.Println(fmt.Sprintf("%s (audio fixed)", name))
}

// VideoMixer Subsystem 2：修复视频
type VideoMixer struct {
}

func (v *VideoMixer) Fix(name string) {
	fmt.Println(fmt.Sprintf("%s (video fixed)", name))
}

// MediaMixer Facade：将多个子系统组合在一起
type MediaMixer struct {
	fixerList []fixer
}

func (m *MediaMixer) Fix(name string) {
	for _, one := range m.fixerList {
		one.Fix(name)
	}
}

func NewMediaMixer() *MediaMixer {
	return &MediaMixer{
		fixerList: []fixer{&AudioMixer{}, &VideoMixer{}},
	}
}

func main() {
	mixer := NewMediaMixer()
	mixer.Fix("电视机")
}
