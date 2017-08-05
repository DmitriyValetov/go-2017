package main

import (
	"fmt"
	"os"
	"strings"
)
func main() {

	if len(os.Args)<2 {
		return
	}

	var name string
	name = os.Args[1]

	count := new(int)
	*count = 0

	file, err := os.Open(name)

	if err != nil {
		// handle the error here
		return
	}

	defer file.Close()

	CountLines(file,count)

	fmt.Printf("%d",*count)

}


func CountLines(file *os.File, counts *int){
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}


	str := string(bs)
	var strs []string
	strs = strings.Split(str,"\n")
	for _, s := range strs{
		if len(s)> 0 {
			(*counts)++
		}
	}
}
