package main

import (
	"fmt"
	"github.com/mrbarge/aoc2022/helper"
	"math"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name   string
	files  []*File
	dirs   []*Dir
	parent *Dir
}

func problem(lines []string, partTwo bool) (int, error) {

	i := 0
	dirs := make(map[string]*Dir, 0)
	rootDir := Dir{name: "", files: make([]*File, 0), dirs: make([]*Dir, 0)}
	dirs["/"] = &rootDir
	currentDir := &rootDir
	for i < len(lines) {
		line := lines[i]
		if strings.HasPrefix(line, "$ cd") {
			dir := strings.Split(line, " ")[2]
			if dir == "/" {
				currentDir = &rootDir
				i++
				continue
			} else if dir == ".." && currentDir.name != "" {
				currentDir = currentDir.parent
				i++
				continue
			}
			dirname := currentDir.name + "/" + dir
			if _, seen := dirs[dir]; !seen {
				nd := Dir{name: dirname, files: make([]*File, 0), dirs: make([]*Dir, 0), parent: currentDir}
				dirs[dirname] = &nd
				currentDir.dirs = append(currentDir.dirs, &nd)
			}
			currentDir = dirs[dirname]
			i++
			continue
		}
		if strings.HasPrefix(line, "$ ls") {
			i++
			currentDir.files = make([]*File, 0)
			for i < len(lines) {
				lsline := lines[i]
				if strings.HasPrefix(lsline, "$ ") {
					break
				}
				if !strings.HasPrefix(lsline, "dir ") {
					fsize, _ := strconv.Atoi(strings.Split(lsline, " ")[0])
					fname := strings.Split(lsline, " ")[1]
					f := File{name: fname, size: fsize}
					currentDir.files = append(currentDir.files, &f)
				}
				i++
			}
			continue
		}
		fmt.Printf("encountered unexpected line: %v\n", line)
		i++
	}

	if partTwo {
		totalspace := rootDir.Size()
		unusedspace := 70000000 - totalspace
		requiredspace := 30000000 - unusedspace
		fmt.Printf("required: %v, unusued: %v", requiredspace, unusedspace)
		return rootDir.findSmallestDirToDelete(requiredspace), nil
	}
	return rootDir.countSmallDirs(100000), nil
}

func (d *Dir) findSmallestDirToDelete(minsize int) int {
	ds := d.Size()
	smallest := math.MaxInt
	if ds >= minsize && ds < smallest {
		smallest = ds
	}
	for _, dir := range d.dirs {
		ds := dir.findSmallestDirToDelete(minsize)
		if ds > minsize && ds < smallest {
			smallest = ds
		}
	}
	return smallest
}

func (d *Dir) countSmallDirs(limit int) int {
	r := 0
	ds := d.Size()
	if ds <= limit {
		r += ds
	}
	for _, dir := range d.dirs {
		r += dir.countSmallDirs(limit)
	}
	return r
}

func (d *Dir) Size() int {
	size := 0
	for _, file := range d.files {
		size += file.size
	}
	for _, d := range d.dirs {
		size += d.Size()
	}
	return size
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)

	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part one: %v\n", ans)

	ans, err = problem(lines, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part two: %v\n", ans)

}
