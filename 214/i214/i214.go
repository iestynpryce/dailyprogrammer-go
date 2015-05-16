package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Rect struct {
	c int
	x int
	y int
	w int
	h int
}

type Point struct {
	x int
	y int
}

type Count struct {
	value int
	count int
}

type Counts []*Count

func (s Counts) Len() int      { return len(s) }
func (s Counts) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByCount struct{ Counts }

func (s ByCount) Less(i, j int) bool {
	return s.Counts[i].count < s.Counts[j].count
}

func calculate_colors(rects []Rect) {
	nrects := len(rects)
	canvas := make(map[Point]int)

	for i := nrects - 1; i >= 0; i-- {
		r := rects[i]
		for j := r.x; j < (r.w + r.x); j++ {
			for k := r.y; k < (r.h + r.y); k++ {
				if canvas[Point{j, k}] == 0 {
					canvas[Point{j, k}] = r.c
				}
			}
		}
	}

	colour_map := make(map[int]int)
	for _, colour := range canvas {
		colour_map[colour]++
	}
	colour_array := make(Counts, len(colour_map))

	i := 0
	for k, v := range colour_map {
		colour_array[i] = &Count{k, v}
		i++
	}
	sort.Sort(ByCount{colour_array})
	for i := range colour_array {
		fmt.Println(colour_array[i].value, colour_array[i].count)
	}
}

func main() {
	rects := make([]Rect, 1)
	rects[0] = Rect{0, 0, 0, 0, 0}

	var firstLine bool = true

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var r Rect
		if firstLine {
			line := scanner.Text()
			for i, c := range strings.Fields(line) {
				num, err := strconv.Atoi(c)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(2)
				}
				switch i {
				case 0:
					rects[0].w = num
				case 1:
					rects[0].h = num
					firstLine = false
				}
			}
		}
		for i, c := range strings.Split(scanner.Text(), " ") {
			num, err := strconv.Atoi(c)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
			switch i {
			case 0:
				r.c = num
			case 1:
				r.x = num
			case 2:
				r.y = num
			case 3:
				r.w = num
			case 4:
				r.h = num
			}
		}
		rects = append(rects, r)
	}

	calculate_colors(rects)
}
