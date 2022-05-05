// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var list = []int{1, 2, 3, 4, 5}

func main() {
	for i := range list {
		if i == 3 {
			fmt.Println(i)
			break
		}
	}
}
