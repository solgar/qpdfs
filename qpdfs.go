package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

var verbose = false

var children []int
var n int
var solutionsTotal int
var currRow int

func main() {
	argsLen := len(os.Args)
	if argsLen < 2 || argsLen > 3 ||
		(argsLen > 1 && (os.Args[1] == "h" || os.Args[1] == "-h" || os.Args[1] == "help" || os.Args[1] == "-help")) {
		fmt.Println("usage: qpdfs n [verbose=0]")
		fmt.Println("\tn - queens and puzzle size")
		fmt.Println("\tverbose - print solutions")
		return
	}

	for idx, arg := range os.Args {
		switch idx {
		case 1:
			val, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Cannot read \"n\".")
				return
			}
			n = val
		case 2:
			val, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Cannot read \"verbose\".")
				return
			}
			verbose = val != 0
		}
	}

	children = make([]int, n)
	children[0] = 0
	var col int = 0
	currRow = 0
	solutionsTotal = 0

	start := time.Now()
mainLoop:
	for {
		childAdded := false

		for ; col < n; col += 1 {
			if doCollideWithRest(currRow, col) == false {
				children[currRow] = col
				childAdded = true
				currRow += 1
				col = 0
				break
			}
		}

		if childAdded && currRow == n {
			solutionsTotal += 1
			if verbose {
				fmt.Println("Solution number ", solutionsTotal, ":")
				printField()
			}
		} else if !childAdded {
			for {
				currRow -= 1
				if currRow < 0 {
					break mainLoop
				}
				col = children[currRow] + 1
				if col < n {
					break
				}
			}
		}
	}

	end := time.Now()
	difference := end.Sub(start)

	fmt.Println("solutionsTotal:", solutionsTotal)
	fmt.Println("took:", difference)
	fmt.Println("solutions per second:", float64(solutionsTotal)/difference.Seconds())
}

func doCollideWithRest(row, col int) bool {

	for r := 0; r < currRow; r += 1 {
		if r == row || children[r] == col || (math.Abs(float64(r-row)) == math.Abs(float64(children[r]-col))) {
			return true
		}
	}
	return false
}

func printField() {
	fmt.Println("Solution number ", solutionsTotal, ":")
	fmt.Print("Fields:")
	for r := 0; r < currRow; r += 1 {
		fmt.Print(" [", r, ", ", children[r], "]")
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			founded := false
			for r := 0; r < currRow; r += 1 {
				if i == r && j == children[r] {
					founded = true
					break
				}
			}
			if founded {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
