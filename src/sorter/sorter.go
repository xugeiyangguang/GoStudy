package main

import (
	"algorithms/bubblesort"
	"algorithms/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "包含待排序的文件")
var outfile *string = flag.String("o", "outfile", "存放排序结果的文件")
var algorithm = flag.String("a", "qsort", "排序算法")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("打开文件失败：", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("此行太长，看起来不是需要的")
			return
		}

		str := string(line) //将字符数组转换为字符串

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return

}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("创建文件失败")
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")

	}
	return nil
}
func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}
	values, err := readValues(*infile)

	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("算法不确定")
		}
		t2 := time.Now()
		fmt.Println("算法排序时间：", t2.Sub(t1))
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
