package main

import (
	"fmt"
	"os"
	"strings"
)
func main() {
	// check the command line args
	if len(os.Args)<2 {
		return
	}
	
	// counter of strings
	counter := new(int)
	*counter = 0

	file, err := os.Open(os.Args[1])
	// check the file existance
	if err != nil {
		fmt.Fprintf(os.Stderr,"Error: %v\n",err)
		return
	}
	defer file.Close()

	CountLines(file,counter)

	fmt.Printf("%d",*counter)
}


func CountLines(file *os.File, counts *int){
	// size of file
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

	// convert bs to string and
	// split it by '\n'
	str := string(bs)
	var strs []string
	strs = strings.Split(str,"\n")
	for _, s := range strs{
		if len(s)> 0 {
			(*counts)++
		}
	}
}
