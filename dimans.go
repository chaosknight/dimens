package main

import "fmt"
import "io/ioutil"
import "regexp"

const ldpi float32 = 0.75
const mdpi float32 = 1
const xhdip float32 = 1.5
const xhdpi float32 = 2
const xxhdpi float32 = 3
const xxxhdpi float32 = 4

func main() {
	//cc 57dp
	re := regexp.MustCompile("\\d")
	bype, error := ioutil.ReadFile("dimans.go")

	if nil != error {
		fmt.Printf("ccc\n")
		return
	}
	s := string(bype)
	fmt.Println(re.ReplaceAllStringFunc(s, func(ss string) {
		return ss
		}
	})

}
