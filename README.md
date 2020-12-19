# go-filex
golang file扩展库
此库旨在简化golang文件操作，api与java.lang.File类似 

安装:
```
go get github.com/weili71/go-filex
```
例子：
```golang
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
```
