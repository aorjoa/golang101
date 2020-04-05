package average_test

import (
	"ssm/average"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage_1_2_3_5_should_return_2_75(t *testing.T) {
	// arrange
	input := []float64{1, 2, 3, 5}

	// action
	actual := average.Average(input)

	// assert
	expected := 2.75
	assert.Equal(t, expected, actual)

}

func TestAverage_3_3_3_3_should_return_3(t *testing.T) {
	input := []float64{3, 3, 3, 3}

	actual := average.Average(input)

	expected := 3.0
	assert.Equal(t, expected, actual)
}
