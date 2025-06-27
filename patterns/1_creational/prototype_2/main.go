package main

import (
	"fmt"
)

func main() {
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.mp4"}
	file3 := &File{name: "file3.pdf"}

	folder1 := &Folder{
		name:     "Files1",
		children: []Inode{file1},
	}

	folder2 := &Folder{
		name:     "Files2",
		children: []Inode{folder1, file2, file3},
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print(" ")

	cloneFolder2 := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder2")
	cloneFolder2.print(" ")
}
