//interfaces
package main
import (
	"fmt"
	"math"
)
type shape interface{           //polymorphism -->>> many different shapes but one behviour is common i.e area
	area() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.height * r.height             // h * w
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius   // pi * r2
}

func getArea(s shape) float64{
	return s.area()
}

func main() {
	r1 := rect{2, 2}
	r2 := rect{3, 4}
	c1 := circle{4.2}

	shapes := []shape{&r1, &r2, c1}     //interface could be of any type
    
	for _, shape := range shapes{
		fmt.Println(getArea(shape)) 
	} 

}

//https://zetcode.com/golang/interface/
