package structs

import "math"

type Rectangle struct{
	width float64
	height float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

type Triangle struct{
	base float64
	height float64
}

func (t Triangle) Area() float64{
	return (t.base * t.height) *0.5
}

// for declaring method we write func(receiverName ReceiverType) MethodName(args)

// we can access the fields of a struct with syntax of myStruct.field (Rectangle.width in this case)

func Perimeter(rectangle Rectangle) float64{
	return 2 * (rectangle.width + rectangle.height)
}

//This is statically typed language hence In Go interface resolution is implicit. 

type Shape interface{
	Area() float64
}
