package main

import "fmt"

// Prototype interface
type Inode interface {
    print(string)
    clone() Inode
}

// Concrete Prototype
type File struct {
    name string
}

func (f *File) print(indentation string) {
    fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
    return &File{name: f.name + "_clone"}
}

// Concrete prototype
type Folder struct {
    children []Inode
    name     string
}

func (f *Folder) print(indentation string) {
    fmt.Println(indentation + f.name)
    for _, i := range f.children {
        i.print(indentation + indentation)
    }
}

func (f *Folder) clone() Inode {
    cloneFolder := &Folder{name: f.name + "_clone"}
    var tempChildren []Inode
    for _, i := range f.children {
        copy := i.clone()
        tempChildren = append(tempChildren, copy)
    }
    cloneFolder.children = tempChildren
    return cloneFolder
}

// client
func main() {
    file1 := &File{name: "File1"}
    file2 := &File{name: "File2"}
    file3 := &File{name: "File3"}

    folder1 := &Folder{
        children: []Inode{file1},
        name:     "Folder1",
    }

    folder2 := &Folder{
        children: []Inode{folder1, file2, file3},
        name:     "Folder2",
    }
    fmt.Println("\nPrinting hierarchy for Folder2")
    folder2.print("  ")

    cloneFolder := folder2.clone()
    fmt.Println("\nPrinting hierarchy for clone Folder")
    cloneFolder.print("  ")
}