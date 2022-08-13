# Nearest Neighbour vs Closest Pair heuristic in the TSP problem
This is the first introductory algorithmic thinking exercise we encounter for the book "The Algorithm Design Manual" by Steven S. Skiena, Chapter 1, problem 1.30.

[Check Out the blog post](https://ggcr.vercel.app/blog/algorithms-with-go-(1):-tsp-with-heuristics)

## Table of Contents
-  How this works
-  Travelling Salesman Problem
-  First Problem Specification: Circle
   -  Nearest Neighbour Heuristic
-  First Problem Specification: Nearest Neightbour Counterexample
   -  Closest Pair heuristic
-  Conclusions: The tests

*\* Table of contents generated with my custom tool [go-auto-toc](https://github.com/ggcr/go-auto-toc)*

## How this works
I work with a TDD methodology. Each folder has its [name].go file and its [name]_test.go file. In order to run the go code just go to the desired path and run the following command:
```bash
go test [-v] // The -v flag is optional but is handy
```

## Travelling Salesman Problem
TSP stands for the Travelling Salesman Problem, it is a very common problem in the algorithm world of Computer Science and it is used as a problem in every Computer Science Degree.

This problem is about finding the shortest path for a given Salesman that has to go through a given set of locations.

## First Problem Specification: Circle
We have the following abstraction of the problem.
![circle](https://i.imgur.com/5mpvqO4.jpeg)
The dots are locations and we have to find the shortest sequence of edges that minimize the distance between each other.
In the code I have declared the Circle with the SetPoints type, a set of Points{x,y}.

### Nearest Neighbour Heuristic
The first approach the book intend us to use is the **Nearest Neighbour** heuristic.

It is a very simple heuristic that stands that we have to choose, every time, the shortest possible location we have from the current location.

The function I've made is the following:
```go
// Main Nearest Neighbour heuristic loop
func NearestNeigh(points SetPoints, initial int) []int {
	mask := make([]bool, len(points))
	mask[initial] = !mask[initial]
	id := initial
	camino := []int{initial}
	for {
		next := calculateMinDist(id, points, mask) // O(n)
		if next == -1 {
			return camino
		}
		camino = append(camino, next)
		mask[next] = !mask[next]
		id = next
	}
}
```

Where SetPoints is the type of the Circle. It's just a static array of Points, each point with its X and Y coordinate.

The code is pretty self-explanatory; we start with an initial point index. Then we will calculate the distance from this current point to ALL the other points available and choose the closest one.
In order to discard the aready visited points we will use a bit mask, so whenever we have the next point index, we can turn its bit to 0 (or False).

In order to calculate the closest neighbour I used this function:

```go
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
```

It's brute force, we loop for all unvisited nodes (that's why we have the condition of !mask[id]), and accumulate the minimum distance possible. Finally we return the closest neighbour index.

The overall algorithm will return the path we have to make in the indices.
```bash
NEAREST NEIGHBOUT RETURN FOR CIRCLE => {0, 1, 2, 3, 4, 5, 6, 7}
```
![neigh_circle](https://i.imgur.com/2JAzkfo.jpeg)

## Second Problem Specification: Nearest Neightbour Counterexample
The following next scenario it is made to demonstrate that the Nearest Neighbout heuristic algorithm is simply **incorrect**.

In the words of the book author, Steven S. Skiana:
```
An algorithm is a procedure that takes any of the possible input instances and transforms it to the desired output.
We seek algorithms that are correct and efficient, while being easy to implement.
```

![line_spec](https://i.imgur.com/2wau5GM.png)
If we run the Nearest Neighbour algorithm we see that it does not return the correct shortest path:
```bash
$ go test
--- FAIL: Test_NearestNeigh (0.00s)
    --- FAIL: Test_NearestNeigh/horizontal_line (0.00s)
        tsp_test.go:45: got [3 2 1 0 4 5 6] want [0 1 2 3 4 5 6]
    --- FAIL: Test_NearestNeigh/vertical_line (0.00s)
        tsp_test.go:51: got [3 2 1 0 4 5 6] want [0 1 2 3 4 5 6]
```

And this is exactly correct, it does not return the correct path. Because how the algorithm's heuristic works, it will start to jump to the nearest neighbour without having the global context of all the visited nodes, the workflow for Nearest Neighour in the second scenario will be something like this:
![incorrect_line_neigh](https://i.imgur.com/dHxAt45.jpeg)

### Closest Pair heuristic
And for this second scenario we are presented with a new heuristic, the Closest pair heuristic. This heuristic will group vertices into a sole one group, what we will call **chain** or ```currentChain```.

And it will always choose the shortest edge from the current set of vertices or vertex chain.

```go
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
```

The workflow is very similar but now we have a differents function for finding the shortest neighbour from A SET OF POINTS, that's why is it a new function, previously we were seeking for a neighbur from a point.

Once we have our closest neighbour we will append it to the end of the path array or to the start, depending on the distance, depending if it's more to the "left" or to the "right", it is really depending if the Point is "bigger" or "smaller".

I've made a compare function for this. It is not robust and does not pretend to be. I am adapting the code to our current scenarios. We are not meant to generalize this code for any kind of problems as this will come later on!

```go
func CompareDistanceXY(a, b Point) bool {
	return a.x >= b.x && a.y >= b.y
}
```

This new heuristic will return the correct path for the first and the second scenario!!!

```bash
CLOSEST PAIR RETURN FOR CIRCLE => {7, 6, 5, 4, 3, 2, 1, 0} (Correct)
CLOSEST PAIR RETURN FOR HORIZONTAL LINE => {0, 1, 2, 3, 4, 5, 6} (Correct)
CLOSEST PAIR RETURN FOR VERTICAL LINE => {0, 1, 2, 3, 4, 5, 6} (Correct)
```

## Conclusions: The tests
I encourage you to explore the tests!
If you run the tests this is what it will be outputed:
```bash
=== RUN   Test_Utils
=== RUN   Test_Utils/euclidean_distance
--- PASS: Test_Utils (0.00s)
    --- PASS: Test_Utils/euclidean_distance (0.00s)
=== RUN   Test_NearestNeigh
=== RUN   Test_NearestNeigh/circle
=== RUN   Test_NearestNeigh/horizontal_line
    tsp_test.go:45: got [3 2 1 0 4 5 6] want [0 1 2 3 4 5 6]
=== RUN   Test_NearestNeigh/vertical_line
    tsp_test.go:51: got [3 2 1 0 4 5 6] want [0 1 2 3 4 5 6]
--- FAIL: Test_NearestNeigh (0.00s)
    --- PASS: Test_NearestNeigh/circle (0.00s)
    --- FAIL: Test_NearestNeigh/horizontal_line (0.00s)
    --- FAIL: Test_NearestNeigh/vertical_line (0.00s)
=== RUN   Test_ClosestPair
=== RUN   Test_ClosestPair/circle
=== RUN   Test_ClosestPair/horizontal_line
=== RUN   Test_ClosestPair/vertical_line
--- PASS: Test_ClosestPair (0.00s)
    --- PASS: Test_ClosestPair/circle (0.00s)
    --- PASS: Test_ClosestPair/horizontal_line (0.00s)
    --- PASS: Test_ClosestPair/vertical_line (0.00s)
```

As you can see, the Nearest Neighbour algorithm for our second scenario is failing! And it is intended to fail.
