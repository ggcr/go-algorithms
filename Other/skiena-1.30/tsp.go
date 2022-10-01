// The Algorithm Design Manual -- Steven S. Skiena
// Chapter 1
// Problem 1-30
// Implement the two TSP heuristics of Section 1.1 (page 5). Which of them gives better solutions in practice? Can you devise a heuristic that works better than both of them?

// TO-DO: EN VEZ DE PASAR EL INDICE ETC, PODEMOS USAR UNA REFERENCIA DE ESTA MANERA NO HACE FALTA TENER EN QUENTA LA MASCARA
// TO-DO: currentChain should be array of indices ([]int) or points ([]Points) ?

package tsp

import (
	"math"
)

// PROBLEM SPECIFICATION

type Point struct {
	x float64
	y float64
}

type SetPoints []Point

var Circle SetPoints
var Line_Horizontal SetPoints
var Line_Veritcal SetPoints

func init() {
	Circle = SetPoints{
		Point{0, 5}, Point{4, 3}, Point{5, 0}, Point{4, -3}, Point{0, -5}, Point{-4, -3}, Point{-5, 0}, Point{-4, 3},
	}

	Line_Horizontal = SetPoints{
		Point{-11, 0}, Point{-6, 0}, Point{-2, 0}, Point{0, 0}, Point{2, 0}, Point{5, 0}, Point{10, 0},
	}

	Line_Veritcal = SetPoints{}
	for _, p := range Line_Horizontal {
		Line_Veritcal = append(Line_Veritcal, Point{p.y, p.x})
	}
}

// NEAREST NEIGHBORUR

// Euclidean distance
func dist(a, b Point) (d float64) {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

// Calculate the shortest neighbour of a given point
func calculateMinDist(point_idx int, points SetPoints, mask []bool) (idx int) {
	p_t := points[point_idx]
	min_dist := math.MaxFloat64
	i_t := -1
	for id, p := range points {
		if !mask[id] {
			d := dist(p_t, p)
			if d < min_dist {
				min_dist = d
				i_t = id
			}
		}
	}
	return i_t
}

// Main Nearest Neighbour heuristic loop
func NearestNeigh(points SetPoints, initial int) []int {
	mask := make([]bool, len(points))
	mask[initial] = !mask[initial]
	id := initial
	camino := []int{initial}
	for { // O(n^2)
		next := calculateMinDist(id, points, mask) // O(n)
		if next == -1 {
			return camino
		}
		camino = append(camino, next)
		mask[next] = !mask[next]
		id = next
	}
}

// CLOSEST PAIR HEURISTIC
// Instead of writting a function such as `assertIdNotInChain` we simply use a mask
func findClosestPair(currentChain []int, points SetPoints, mask []bool) int {
	min_dist := math.MaxFloat64
	i_t := -1
	for _, v := range currentChain {
		for id, p := range points {
			if !mask[id] {
				d := dist(points[v], p)
				if d < min_dist {
					min_dist = d
					i_t = id
				}
			}
		}
	}
	return i_t
}

func CompareDistanceXY(a, b Point) bool {
	return a.x >= b.x && a.y >= b.y
}

// Main Closest Pair heuristic loop
func ClosestPair(points SetPoints, initial int) []int {
	currentChain := []int{initial}
	mask := make([]bool, len(points))
	mask[initial] = !mask[initial]
	for {
		next := findClosestPair(currentChain, points, mask) // O(n^2)
		if next == -1 {
			return currentChain
		}
		if CompareDistanceXY(points[next], points[currentChain[len(currentChain)-1]]) {
			currentChain = append(currentChain, next)
		} else {
			currentChain = append([]int{next}, currentChain...)
		}
		mask[next] = !mask[next]
	}
}

// func main() {
// 	res := NearestNeigh(Circle, 0)
// 	fmt.Println(res)
// 	res = NearestNeigh(Line_Horizontal, 3)
// 	fmt.Println(res)
// 	res = NearestNeigh(Line_Veritcal, 3)
// 	fmt.Println(res)
// }
