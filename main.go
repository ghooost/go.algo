package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item interface {
	Label()
	Check(key string) bool
	Run()
}

type Title struct {
	note string
}

func (t Title) Label() {
	fmt.Printf("\033[35m%s\033[0m\n", t.note)
}
func (t Title) Check(key string) bool {
	return false
}
func (t Title) Run() {}

type SortAlgo struct {
	note string
	key string
	runner func(arr []int)
}

func (a SortAlgo) Label() {
	fmt.Printf("%s -- %s\n", a.key, a.note)
}

func (a SortAlgo) Check(key string) bool {
	return a.key == key
}

func (a SortAlgo) Run() {
	fmt.Printf("\033[32m%s\033[0m\n", a.note)
	arr := []int{2, 6, 3, 1, 5, 4, 12, 0, 1}
	fmt.Println("Before: ", arr)
	a.runner(arr[:])
	fmt.Println("After: ", arr)
}

type Op struct {
	note string
	key string
	runner func()
}

func (a Op) Label() {
	fmt.Printf("%s -- %s\n", a.key, a.note)
}

func (a Op) Check(key string) bool {
	return a.key == key
}

func (a Op) Run() {
	a.runner()
}

var ops = []Item {
	Title {
		note: "Sorts",
	},
	SortAlgo {
		note: "Bubble sort",
		key: "1",
		runner: Bubble, 
	},
	SortAlgo {
		note: "Merge sort",
		key: "2",
		runner: Merge, 
	},
	SortAlgo {
		note: "Quick sort",
		key: "3",
		runner: Quick, 
	},
	Title {
		note: "Other",
	},
	Op {
		note: "Quit",
		key: "q",
		runner: func() {
			fmt.Println("Good bye")
			os.Exit(0)
		},
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		for i:=0; i<len(ops); i++ {
			ops[i].Label()
		}

		fmt.Print("->");
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		
		found := false
		for i:=0; i<len(ops); i++ {
			if ops[i].Check(text) {
				ops[i].Run()				
			}
		}

		if found == false {
			fmt.Printf("%s not found\n", text)
		}
	}
}
