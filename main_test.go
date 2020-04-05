package main_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAssert_AssertToBeTrue(t *testing.T) {
	hello := "Hello"
	fmt.Println(hello)
}

// call function //
func add(x int, y int) int {
	return x + y
}

func TestSimpleAssert_CallFunctionAdd(t *testing.T) {
	assert.Equal(t, add(2, 3), 5)
}

///////////////////

// call function multiple return //
func swap(x, y int) (int, int) {
	return y, x
}

func TestSimpleAssert_CallFunctionSwapMultipleReturn(t *testing.T) {
	a, b := swap(1, 3)
	assert.Equal(t, 3, a)
	assert.Equal(t, 1, b)
}

///////////////////

// call function named return //
func declare() (xx int, yy int) {
	xx = 100
	yy = 200
	return
}

func TestSimpleAssert_CallFunctionNamedReturn(t *testing.T) {
	xx, yy := declare()
	assert.Equal(t, xx, 100)
	assert.Equal(t, yy, 200)
}

///////////////////

// simple variable declaration //
func TestSimpleAssert_DeclareVariable(t *testing.T) {
	var golang string
	fmt.Println("golang", golang)
	var c int
	fmt.Println("c", c)
	var html float64
	fmt.Println("html", html)
	var javascript bool
	fmt.Println("javascript", javascript)

	// shorthand variable declaration
	php := 123
	fmt.Println("php", php)
	python := "test"
	fmt.Println("python", python)
}

///////////////////

// type conversion //
func TestSimpleAssert_TypeConversion(t *testing.T) {
	targetAsInteger := 123
	fmt.Println(reflect.TypeOf(targetAsInteger))
	targetAsFloat64 := float64(targetAsInteger)
	fmt.Println(reflect.TypeOf(targetAsFloat64))

	aToI, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Println("error while convert type")
		t.Fail()
		return
	}
	fmt.Println("Atoi: ", aToI, reflect.TypeOf(aToI))

	iToA := strconv.Itoa(123)
	fmt.Println("iToA: ", iToA, reflect.TypeOf(iToA))
}

///////////////////

// type inference //
func TestSimpleAssert_TypeInference(t *testing.T) {
	targetAsInteger := 123
	fmt.Println("targetAsInteger", reflect.TypeOf(targetAsInteger))

	targetAsFloat := 123.0021
	fmt.Println("targetAsFloat", reflect.TypeOf(targetAsFloat))

	var str string = "test"
	fmt.Println("str", str)
}

///////////////////

// constant //
func TestSimpleAssert_Constant(t *testing.T) {
	const pi = 3.14
	fmt.Println(pi)
}

///////////////////

// simple for loop //
func TestSimpleAssert_SimpleForLoop(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	sum2 := 0
	for sum2 < 1000 {
		sum2 = sum2 + 5
	}
	fmt.Println(sum2)

	sum3 := 0
	for {
		fmt.Println(sum3)
		if sum3 > 10 {
			break
		}
		sum3 = sum3 + 1
	}

	langArr := []string{"test1", "test2", "test3"}
	for key, value := range langArr {
		fmt.Println(key, value)
	}
}

///////////////////

// switch case //
func simpleCheckBill(x int) {
	switch x {
	case 10:
		fmt.Println("too bad")
	case 20:
		fmt.Println("normal")
	default:
		fmt.Println("default")
	}
}

func TestSimpleAssert_SwitchCase(t *testing.T) {
	simpleCheckBill(30)
}

///////////////////

// if case //
func simpleIfCheckBill(x int) {
	if x < 10 {
		fmt.Println("less than 10")
	} else if x < 20 {
		fmt.Println("less than 20")
	} else {
		fmt.Println("else")
	}
}

func TestSimpleAssert_IfElse(t *testing.T) {
	simpleIfCheckBill(30)
}

///////////////////

// defer //
func deferCase() {
	fmt.Println("Open connection")
	defer fmt.Println("<Defer here> Close connection1")

	defer fmt.Println("<Defer here> Close connection2")
	fmt.Println("Connect to DB ......")
	fmt.Println("<Defer here> Close connection1")

}

func TestSimpleAssert_Defer(t *testing.T) {
	deferCase()
}

///////////////////

// pointer //

func mutate(i *int) {
	fmt.Printf("in mutate function : %p\n", i)
	*i = 200
	mutate2(&i)
}

func mutate2(i **int) {
	fmt.Printf("in mutate2 function : %p\n", i)
	**i = 300
}

func TestSimpleAssert_Pointer(t *testing.T) {
	i := 100
	fmt.Printf("in main test function : %p\n", &i)
	mutate(&i)

	fmt.Println("i: ", i)
}

///////////////////

// struct //
type Coordinate struct {
	X int
	Y int
}

func TestSimpleAssert_Struct(t *testing.T) {
	schoolMap := Coordinate{X: 10, Y: 20}
	fmt.Println(schoolMap)
	schoolMap2 := Coordinate{10, 20}
	fmt.Println(schoolMap2)
	schoolMap3 := Coordinate{Y: 20}
	fmt.Println(schoolMap3)
}

///////////////////

// array //
func arrReceiver(arr [3]int) {
	fmt.Printf("in arrReceiver >>>> %p\n", &arr)
}

func TestSimpleAssert_ArrayReceiver(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Printf("in main function >>>> %p\n", &arr)
	arrReceiver(arr)
}

func TestSimpleAssert_ArrayLooping(t *testing.T) {
	arr := [3]int{1, 2, 3}
	for _, value := range arr {
		fmt.Println(value)
	}
}

///////////////////

// slice //

func TestSimpleAssert_Slice(t *testing.T) {
	sl1 := []int{1, 2, 3, 4}
	fmt.Println(sl1[:2])
	fmt.Println(sl1[3:4])
	fmt.Println(sl1[2:])
}

///////////////////

// map //
func getPriceTag(tagName string) float64 {
	var priceTag = make(map[string]float64)
	priceTag["banana"] = 160.5
	element, ok := priceTag[tagName]
	if !ok {
		fmt.Println("error: not found in map")
		return -1
	}
	return element
}

func TestSimpleAssert_Map(t *testing.T) {
	fmt.Println(getPriceTag("banana"))
	fmt.Println(getPriceTag("apple"))
}

///////////////////

// function value //
type godFunction func(int, int) int

func doSum(fn godFunction, x int, y int) int {
	return fn(x, y)
}

func TestSimpleAssert_FunctionValue(t *testing.T) {
	sumFn := func(x, y int) int {
		return x + y
	}
	fmt.Println(doSum(sumFn, 1, 5))
}

///////////////////

// Method //
type CoordinateMethod struct {
	X int
	Y int
}

func (c CoordinateMethod) getSum() int {
	return c.X + c.Y
}

func TestSimpleAssert_Method(t *testing.T) {
	schoolMap := CoordinateMethod{X: 10, Y: 20}
	fmt.Println(schoolMap.getSum())
}

///////////////////

// Interface //
type flyable interface {
	Fly()
	Buy()
}

type BuySt struct {
	Name string
}

func (b BuySt) Buy() {
	fmt.Println(b.Name + " is buying")
}

type Bird struct {
	BuySt
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name, " is flying")
}

type Airplane struct {
	BuySt
}

func (a Airplane) Fly() {
	fmt.Println("Airplane is onboarding")
}

func doFly(fni flyable) {
	fni.Fly()
	fni.Buy()
}

func TestSimpleAssert_Interface(t *testing.T) {
	bird := Bird{BuySt: BuySt{"Test"}, Name: "JibJib"}
	doFly(bird)
	airplane := Airplane{}
	doFly(airplane)
}

///////////////////

// Type assertion //
func doPrint(value interface{}) {
	fmt.Println(reflect.TypeOf(value), reflect.TypeOf(value).Kind(), value)
}

func TestSimpleAssert_TypeAssertion(t *testing.T) {
	doPrint(Coordinate{})
}

///////////////////
