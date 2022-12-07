package main

import (
	"aoc2022"
	"fmt"
	"strings"
)

func main() {
	timer := aoc2022.Time()
	defer timer()

	input := aoc2022.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	rootFile := parseInput(lines)
	totalSpace := 70000000
	freeNeeded := 30000000
	unused := totalSpace - rootFile.getSize()
	toClean := freeNeeded - unused

	dirSizes := getDirSizes(rootFile, []int{rootFile.getSize()})

	closest := rootFile.getSize()
	for _, size := range dirSizes {
		if size > toClean && size < closest {
			closest = size
		}
	}

	return closest
}

func getDirSizes(file *File, sizes []int) []int {
	// Only parse directories, not files
	if file.size == 0 {
		sizes = append(sizes, file.getSize())
		for _, f := range file.files {
			sizes = getDirSizes(f, sizes)
		}
	}
	return sizes

}

func parseInput(lines []string) *File {
	root := NewFile("/", 0, nil)
	active := root

	// Skip the first cd cmd, thats always the root file
	for _, l := range lines[1:] {
		if l[0] == '$' {
			cmd := parseCmd(l)

			if cmd[0] == "cd" {
				if cmd[1] == ".." {
					active = active.parent
				} else {
					fileName := cmd[1]
					if file, ok := active.files[fileName]; ok {
						active = file
					} else {
						new := NewFile(fileName, 0, active)
						active.addFile(new)
						active = new
					}
				}

			}
		} else {
			parts := strings.Split(l, " ")
			if parts[0] == "dir" {
				new := NewFile(parts[1], 0, active)
				active.addFile(new)
			} else {
				new := NewFile(parts[1], aoc2022.Atoi(parts[0]), active)
				active.addFile(new)
			}
		}
	}

	return root
}

func parseCmd(line string) []string {
	parts := strings.Split(line, " ")
	// Don't return the `$`
	return parts[1:]
}

type File struct {
	name string
	size int

	parent *File
	files  map[string]*File
}

func NewFile(name string, size int, parent *File) *File {
	return &File{
		name:   name,
		size:   size,
		parent: parent,
		files:  map[string]*File{},
	}
}

func (f *File) getSize() int {
	size := f.size
	for _, file := range f.files {
		size += file.getSize()
	}

	return size
}

func (f *File) addFile(file *File) {
	f.files[file.name] = file
}

func (f *File) tree(prefix string) {
	fmt.Printf("%v- %v \n", prefix, f.name)
	for _, file := range f.files {
		file.tree(prefix + "  ")
	}
}
