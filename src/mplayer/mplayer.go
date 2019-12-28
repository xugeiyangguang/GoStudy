package main

import (
	"bufio"
	"fmt"
	"milb"
	"mp"
	"os"
	"strconv"
	"strings"
)

var lib *milb.MusicManager
var id int

//var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&milb.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("增加音乐的时候出错了！")
		}
	case "remove":
		if len(tokens) == 3 {
			tmp, _ := strconv.Atoi(tokens[2])
			lib.Remove(tmp)
		} else {
			fmt.Println("删除音乐出错了！")
		}
	default:
		fmt.Println("没有此音乐操作命令！")
	}

}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("播放命令出错！")
		return
	}

	e := lib.Find(tokens[1])

	if e == nil {
		fmt.Println("播放的歌曲不存在！")
		return
	}
	mp.Play(e.Source, e.Type)

}

func main() {
	lib = milb.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("输入命令：")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("首命令不认识：", tokens[0])
		}
	}
}
