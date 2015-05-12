// Laziest way to type out a sentance using hunt and peck on a qwerty keyboard
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"unicode"
)

type Point struct {
	x int
	y int
}

func (p *Point) distance(q *Point) int {
	xdiff := p.x - q.x
	if xdiff < 0 {
		xdiff *= -1
	}
	ydiff := p.y - q.y
	if ydiff < 0 {
		ydiff *= -1
	}
	return xdiff + ydiff
}

func min_key_distance(points []Point, p *Point) (distance int, q Point) {
	distance = 100 // pick a very high value
	for _, point := range points {
		point_dist := point.distance(p)
		if point_dist < distance {
			distance = point_dist
			q = point
		}
	}
	return
}

func max_key_distance(points []Point, p *Point) (distance int, q Point) {
	distance = -1 // pick a low value
	for _, point := range points {
		point_dist := point.distance(p)
		if point_dist >= distance {
			distance = point_dist
			q = point
		}
	}
	return
}

/* Create a 4x10 querty keybard:
 * . = key which does nothing
 * ^ = shift key
 * # = space key */
func create_keyboard() (keyboard [][]rune) {
	keyboard = make([][]rune, 4)
	keyboard[0] = []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'}
	keyboard[1] = []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', '.'}
	keyboard[2] = []rune{'^', 'z', 'x', 'c', 'v', 'b', 'n', 'm', '.', '^'}
	keyboard[3] = []rune{'.', '.', '.', ' ', ' ', ' ', ' ', ' ', '.', '.'}
	return
}

// Map each key to a point in the 2D array
func key_to_position(k [][]rune) (mapping map[rune][]Point) {

	mapping = make(map[rune][]Point)

	for y := range k {
		for x := range k[y] {
			c := k[y][x]
			mapping[c] = append(mapping[c], Point{x, y})
		}
	}
	return
}

func print_action(c rune, p *Point, direction string, distance int, kb [][]rune) {
	if distance > 0 {
		fmt.Printf("%s: Move %s hand from %s (effort: %d)\n",
			printable(c), direction, printable(kb[p.y][p.x]),
			distance)
	} else {
		fmt.Printf("%s: Use %s hand again (effort: %d)\n",
			printable(c), direction, distance)
	}
}

func print_first(p *Point, direction string, kb [][]rune) {
	fmt.Printf("%s: Use %s hand\n", printable(kb[p.y][p.x]), direction)
}

/* turn a rune into a printable string i.e. uppercase and replacing special
 * chars with their names i.e. ' ' == "Space"
 */
func printable(r rune) string {
	var s string
	if r == '^' {
		s = "Shift"
	} else if r == ' ' {
		s = "Space"
	} else {
		s = string(unicode.ToUpper(r))
	}
	return s
}

func next_key(c rune, left *Point, right *Point, kb [][]rune,
	keyMap map[rune][]Point) int {
	target := keyMap[unicode.ToLower(c)]
	distanceLeft, leftTemp := min_key_distance(target, left)
	distanceRight, rightTemp := min_key_distance(target, right)

	var distance int

	if distanceLeft < distanceRight {
		if unicode.IsUpper(c) {
			distanceRight, rightTemp := min_key_distance(keyMap['^'], right)
			print_action('^', right, "right", distanceRight, kb)
			*right = rightTemp
			distance += distanceRight
		}
		print_action(c, left, "left", distanceLeft, kb)
		*left = leftTemp
		distance += distanceLeft
	} else {
		if unicode.IsUpper(c) {
			distanceLeft, leftTemp := min_key_distance(keyMap['^'], left)
			print_action('^', left, "left", distanceLeft, kb)
			*left = leftTemp
			distance += distanceLeft
		}
		print_action(c, right, "right", distanceRight, kb)
		*right = rightTemp
		distance += distanceRight
	}
	return distance
}

func laziest_route(s string, kb [][]rune, keyMap map[rune][]Point) {

	var left, right Point
	var leftSet, rightSet bool

	var totalEffort int = 0

	for _, c := range s {

		target := keyMap[unicode.ToLower(c)]

		if !leftSet || !rightSet {
			// Randomly pick a start key if there is a choice
			p := target[rand.Intn(len(target))]

			if !leftSet && p.x > 4 {
				if unicode.IsUpper(c) {
					_, right = max_key_distance(keyMap['^'], &left)
					rightSet = true
					print_first(&right, "right", kb)
				}
				left = p
				leftSet = true
				print_first(&left, "left", kb)
			}
			if !rightSet && p.x <= 4 {
				if unicode.IsUpper(c) {
					_, left = max_key_distance(keyMap['^'], &right)
					leftSet = true
					print_first(&left, "left", kb)
				}
				right = p
				rightSet = true
				print_first(&right, "right", kb)
			}
			continue
		}

		totalEffort += next_key(c, &left, &right, kb, keyMap)

	}
	fmt.Printf("Total effort: %d\n", totalEffort)
}

func main() {
	kb := create_keyboard()
	keyMap := key_to_position(kb)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		laziest_route(scanner.Text(), kb, keyMap)
	}
}
