package grade_test

import (
	"ssm/grade"
	"testing"

	"gotest.tools/assert"
)

func Test_input_80_should_return_A(t *testing.T) {
	score := 80.0

	actual := grade.Calculate(score)

	expected := "A"
	assert.Equal(t, expected, actual)
}

func Test_input_79_should_return_B_plus(t *testing.T) {
	score := 79.0

	actual := grade.Calculate(score)

	expected := "B+"
	assert.Equal(t, expected, actual)
}

func Test_input_70_should_return_B(t *testing.T) {
	score := 70.0

	actual := grade.Calculate(score)

	expected := "B"
	assert.Equal(t, expected, actual)
}

func Test_input_65_should_return_C_plus(t *testing.T) {
	score := 65.0

	actual := grade.Calculate(score)

	expected := "C+"
	assert.Equal(t, expected, actual)
}

func Test_input_60_should_return_C(t *testing.T) {
	score := 60.0

	actual := grade.Calculate(score)

	expected := "C"
	assert.Equal(t, expected, actual)
}

func Test_input_59_should_return_D_plus(t *testing.T) {
	score := 59.0

	actual := grade.Calculate(score)

	expected := "D+"
	assert.Equal(t, expected, actual)
}

func Test_input_99_should_return_A(t *testing.T) {
	score := 99.0

	actual := grade.Calculate(score)

	expected := "A"
	assert.Equal(t, expected, actual)
}

func Test_input_69_should_return_C_plus(t *testing.T) {
	score := 69.0

	actual := grade.Calculate(score)

	expected := "C+"
	assert.Equal(t, expected, actual)
}

func Test_input_75_should_return_B_plus(t *testing.T) {
	score := 75.0

	actual := grade.Calculate(score)

	expected := "B+"
	assert.Equal(t, expected, actual)
}

func Test_input_50_should_return_D(t *testing.T) {
	score := 50.0

	actual := grade.Calculate(score)

	expected := "D"
	assert.Equal(t, expected, actual)
}

func Test_input_40_should_return_F(t *testing.T) {
	score := 40.0

	actual := grade.Calculate(score)

	expected := "F"
	assert.Equal(t, expected, actual)
}

func Test_input_0_should_return_F(t *testing.T) {
	score := 0.0

	actual := grade.Calculate(score)

	expected := "F"
	assert.Equal(t, expected, actual)
}
