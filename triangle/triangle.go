package triangle

import (
	"errors"
)

// Type of the triangle
type Type string

var (
	// Equilateral all sides of triangle are equal
	Equilateral = Type("Equilateral")
	// Isosceles any two sides of triangle are equal
	Isosceles = Type("Isosceles")
	// Scalene None of the sides of triangle are equal
	Scalene = Type("Scalene")
)

var (
	// ErrNonConformanceToTriangleInequality - Error when the input side lengths are invalid according to triangle inequality axiom
	ErrNonConformanceToTriangleInequality = errors.New("The given input is invalid as it does not conform to the triangle inequality (sum of any 2 sides of a triangle must be greater than the measure of the third side)")
)

// GetType - returns the type of triangle based on the side lengths, if
// input side lengths do not form a valid triangle, returns an error
func GetType(side1 float64, side2 float64, side3 float64) (*Type, error) {
	if !checkTriangleInequality(side1, side2, side3) {
		return nil, ErrNonConformanceToTriangleInequality
	}
	if side1 == side2 && side2 == side3 {
		return &Equilateral, nil
	}
	if side1 == side2 || side2 == side3 || side3 == side1 {
		return &Isosceles, nil
	}
	return &Scalene, nil
}

// checkTriangleInequality - sum of any 2 sides of a triangle must be greater than the measure of the third side
func checkTriangleInequality(side1 float64, side2 float64, side3 float64) bool {
	return side1+side2 > side3 && side2+side3 > side1 && side3+side1 > side2
}
