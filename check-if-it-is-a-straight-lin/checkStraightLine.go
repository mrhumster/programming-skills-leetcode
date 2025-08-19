/*
Package t1232

You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents
the coordinate of a point. Check if these points make a straight line in the XY plane.
*/
package t1232

import "math"

func arePointsCollinear(x1, y1, x2, y2, x3, y3 float64) bool {
	area := (x2-x1)*(y3-y1) - (y2-y1)*(x3-x1)
	return math.Abs(area) < 1e-10
}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}
	x1, y1 := float64(coordinates[0][0]), float64(coordinates[0][1])
	x2, y2 := float64(coordinates[1][0]), float64(coordinates[1][1])

	for i := 2; i < len(coordinates); i++ {
		xi, yi := float64(coordinates[i][0]), float64(coordinates[i][1])
		if !arePointsCollinear(x1, y1, x2, y2, xi, yi) {
			return false
		}
	}
	return true
}
