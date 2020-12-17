package main

import (
	"filex"
	"fmt"
	"os"
)

func main() {
	file:=filex.NewFile("c://")
	for _, f := range file.ListFiles() {
		fmt.Println(f.CanonicalPath())
	}
	textFile:=filex.NewFile2(file,"test.txt")
	osFile, _ :=textFile.OpenFile(os.O_RDWR,0755)
	textFile.AppendText(osFile,"hello!\n")
	osFile.Close()
}
