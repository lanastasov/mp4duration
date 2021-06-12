package main

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/pillash/mp4util"
)

//video.mp4  -> 54 sec
//video2.mp4 -> 736 sec

// go get github.com/pillash/mp4util

func main() {
	files, _ := filepath.Glob("*.mp4")
	sum := 0
	for i := range files {
		fmt.Println(mp4util.Duration(files[i]))
		val, _ := mp4util.Duration(files[i])
		sum += val
	}
	fmt.Println("sum = ", sum)
	var h = int(sum / 3600)
	var m = (sum - int(sum/3600)*3600) / 60
	var s = sum - h*3600 - m*60
	fmt.Println("h=", h, "m=", m, "s=", s)
	fmt.Println(reflect.TypeOf(files))

	fmt.Println("hi")
	fmt.Println(mp4util.Duration("video2.mp4"))
}
