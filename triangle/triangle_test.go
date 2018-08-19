package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetType_TriangleInequality(t *testing.T) {
	triangleType, err := GetType(1, 2, 3)
	assert.Nil(t, triangleType, "GetType should return a nil type")
	assert.Error(t, err, "GetType should return arror")
	assert.Equal(t, err, ErrNonConformanceToTriangleInequality, "Invalid error returned")
}

func TestGetType_Equilateral(t *testing.T) {
	triangleType, err := GetType(3, 3, 3)
	assert.NotNil(t, triangleType, "GetType should return a valid type")
	assert.NoError(t, err, "GetType should not return arror")
	assert.Equal(t, triangleType, &Equilateral, "Invalid triangle type returned")
}

func TestGetType_Isosceles(t *testing.T) {
	triangleType, err := GetType(3, 3, 2)
	assert.NotNil(t, triangleType, "GetType should return a valid type")
	assert.NoError(t, err, "GetType should not return arror")
	assert.Equal(t, triangleType, &Isosceles, "Invalid triangle type returned")
}

func TestGetType_Scalene(t *testing.T) {
	triangleType, err := GetType(3, 4, 5)
	assert.NotNil(t, triangleType, "GetType should return a valid type")
	assert.NoError(t, err, "GetType should not return arror")
	assert.Equal(t, triangleType, &Scalene, "Invalid triangle type returned")
}
