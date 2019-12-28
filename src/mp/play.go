package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mytype string) {
	var p Player
	switch mytype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("没有支持该音乐的播放器！")
		return
	}

	p.Play(source)
}
