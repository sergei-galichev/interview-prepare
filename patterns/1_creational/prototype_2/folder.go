package main

import (
	"fmt"
)

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(identation string) {
	fmt.Println(identation + f.name)

	for _, child := range f.children {
		child.print(identation + identation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}

	var tempChildren []Inode

	for _, child := range f.children {
		cp := child.clone()

		tempChildren = append(tempChildren, cp)
	}

	cloneFolder.children = tempChildren

	return cloneFolder
}
