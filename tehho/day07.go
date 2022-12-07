package main

import (
	"strconv"
	"strings"
)

type Node07 struct {
	parent *Node07
	name   string
	size   int
	list   []*Node07
}

func (n *Node07) GetChild(s string) *Node07 {
	if s == ".." {
		return n.parent
	} else {
		for i := 0; i < len(n.list); i++ {
			if n.list[i].name == s {
				return n.list[i]
			}
		}
	}
	return n
}

func (n *Node07) GetSize() int {
	if n.size != 0 {
		return n.size
	}

	for i := 0; i < len(n.list); i++ {
		n.size += n.list[i].GetSize()
	}
	return n.size
}

func (n *Node07) GetSizeOfAllDirs(cap int) int {
	val := 0
	for i := 0; i < len(n.list); i++ {
		if n.list[i].IsDir() {
			val += n.list[i].GetSizeOfAllDirs(cap)
			temp := n.list[i].GetSize()
			if temp < cap {
				val += temp
			}
		}
	}
	return val
}

func (n *Node07) AddNode(n1 *Node07) {
	n.list = append(n.list, n1)
}

func (n *Node07) IsDir() bool {
	return len(n.list) != 0
}

func Day07part01(input string) int {
	root := Node07{
		list:   make([]*Node07, 0),
		parent: nil,
		name:   "/",
		size:   0,
	}

	rows := strings.Split(input, "\n")
	current := &root

	for i := 0; i < len(rows); i++ {
		temp := strings.Split(rows[i], " ")
		if temp[0] == "" {

		} else if temp[0] == "$" {
			cmd := temp[1]
			if cmd == "cd" {
				current = current.GetChild(temp[2])
			}
		} else {
			if temp[0] == "dir" {
				current.AddNode(&Node07{
					list:   make([]*Node07, 0),
					name:   temp[1],
					parent: current,
					size:   0,
				})
			} else {
				val, err := strconv.Atoi(temp[0])
				if err != nil {
					panic(err)
				}

				current.AddNode(&Node07{
					list:   make([]*Node07, 0),
					name:   temp[1],
					parent: current,
					size:   val,
				})
			}
		}
	}

	cap := 100000
	root.GetSize()

	return root.GetSizeOfAllDirs(cap)
}

func (n *Node07) GetSmallestGreaterThan(cap int) int {
	val := n.GetSize()

	for i := 0; i < len(n.list); i++ {
		temp := n.list[i].GetSmallestGreaterThan(cap)
		if n.list[i].IsDir() && temp < val && cap < temp {
			val = temp
		}
	}

	return val
}

func Day07part02(input string) int {
	root := Node07{
		list:   make([]*Node07, 0),
		parent: nil,
		name:   "/",
		size:   0,
	}

	rows := strings.Split(input, "\n")
	current := &root

	for i := 0; i < len(rows); i++ {
		temp := strings.Split(rows[i], " ")
		if temp[0] == "" {

		} else if temp[0] == "$" {
			cmd := temp[1]
			if cmd == "cd" {
				current = current.GetChild(temp[2])
			}
		} else {
			if temp[0] == "dir" {
				current.AddNode(&Node07{
					list:   make([]*Node07, 0),
					name:   temp[1],
					parent: current,
					size:   0,
				})
			} else {
				val, err := strconv.Atoi(temp[0])
				if err != nil {
					panic(err)
				}

				current.AddNode(&Node07{
					list:   make([]*Node07, 0),
					name:   temp[1],
					parent: current,
					size:   val,
				})
			}
		}
	}

	max := 70000000
	needed := 30000000
	used := root.GetSize()

	freeNeeded := needed - (max - used)

	return root.GetSmallestGreaterThan(freeNeeded)

}
