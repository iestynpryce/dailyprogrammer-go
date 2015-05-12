// Laziest way to type out a sentance using hunt and peck on a qwerty keyboard
package main

import (
	"bufio"
	"fmt"
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

func key_distance(points []Point, p *Point) (distance int, q Point) {
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

func print_action(left, right Point, kb [][]rune) {
	fmt.Printf("%s: Use left hand\n", printable(kb[left.y][left.x]))
	fmt.Printf("%s: Use right hand\n", printable(kb[right.y][right.x]))
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

func laziest_route(s string, kb [][]rune, keyMap map[rune][]Point) {

	var left, right Point
	var leftSet, rightSet bool

	var totalEffort int = 0

	for _, c := range s {
		if !leftSet || !rightSet {
			r := unicode.ToLower(c)
			p := keyMap[r][0] // exploit knowledge that chars only have 1 key
			if p.x > 4 {
				left = keyMap['^'][0]
				leftSet = true
			} else {
				right = keyMap['^'][1]
				rightSet = true
			}
			if unicode.IsUpper(c) {
				if rightSet {
					left = p
					leftSet = true
				} else {
					right = p
					rightSet = true
				}
			}
			print_action(left, right, kb)
			continue
		}

		target := keyMap[unicode.ToLower(c)]
		distanceLeft, leftTemp := key_distance(target, &left)
		distanceRight, rightTemp := key_distance(target, &right)

		if distanceLeft < distanceRight {
			fmt.Printf("%s: Move left hand from %s (effort: %d)\n",
				printable(c), printable(kb[left.y][left.x]),
				distanceLeft)
			left = leftTemp
			totalEffort += distanceLeft
		} else {
			fmt.Printf("%s: Move right hand from %s (effort: %d)\n",
				printable(c), printable(kb[right.y][right.x]),
				distanceRight)
			right = rightTemp
			totalEffort += distanceRight
		}
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
