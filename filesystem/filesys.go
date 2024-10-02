package main

import (
	"fmt"
	"os"
)

func main() {

	// Deleting the file

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	err := os.Remove("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Deleted Successfully")

	// read from one file to store it into another file
	// read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("test.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// // reader :

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		//  EOF Error
	// 		break
	// 	}

	// 	// writebyte

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush()

	// fmt.Println("Done")

	// write in file and creating the file

	// f, err := os.Create("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("Hi go i am here")
	// f.WriteString("Nice language")

	// bytes := []byte("hello golang this is nhit ")

	// f.Write(bytes)

	// reading the folders / directory

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(0)

	// for _, fi := range fileInfo {
	// 	fmt.Println("f", fi.Name(), fi.IsDir())
	// }

	// fileInfo, err := dir.ReadDir()

	// READ THE FILE

	// f, err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(f))

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	// something with error
	// 	// log the error

	// 	panic(err)

	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("filename : ", fileInfo.Name())
	// fmt.Println("fileOrFolder", fileInfo.IsDir())
	// fmt.Println("filesize", fileInfo.Size())
	// fmt.Println("fileMode", fileInfo.Mode())
	// fmt.Println("filemodifiedat", fileInfo.ModTime())

}
