package main

import (
	"fmt"
	"os"
)

// a lot of things are in this i have skipped as i felt that boring it is like basic file handling which we have used to do in python

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("File size:", fileInfo.Size())
	// fmt.Println("File name:", fileInfo.Name())
	// fmt.Println("File last modified:", fileInfo.ModTime())

	// read the file
	f, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 1024)
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Read bytes:", d)
	fmt.Println("Content:", string(buf[:d]))
}
