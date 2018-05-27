package main

import (
	"fmt"
)

var solapp int = 0
var solapps string = ""
var nbs int = 0

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func search(c []int, s int, cs string) (end bool) {
	found := false

	for i, a := range c {
		if a == -1 {
			continue
		}
		for j, b := range c {
			if i >= j {
				continue
			}
			for k := 0; k < 4; k++ {
				var op string
				var cal int
				switch k {
				case 0:
					{
						cal = a + b
						op = fmt.Sprintf("%d + %d = %d", a, b, cal)
					}
				case 1:
					{
						cal = a * b
						op = fmt.Sprintf("%d * %d = %d", a, b, cal)
					}
				case 2:
					{
						var e, f int
						if a == b {
							continue
						}

						if a > b {
							e = a
							f = b
						}
						if a < b {
							e = b
							f = a
						}
						cal = e - f
						op = fmt.Sprintf("%d - %d = %d", e, f, cal)
					}
				case 3:
					{
						if a == 1 || b == 1 {
							continue
						}
						if a%b == 0 {
							cal = a / b
							op = fmt.Sprintf("%d / %d = %d", a, b, cal)
						} else if b%a == 0 {
							cal = b / a
							op = fmt.Sprintf("%d / %d = %d", b, a, cal)
						} else {
							continue
						}

					}
				}

				cs1 := cs + "    " + op

				if abs(cal-s) < abs(solapp-s) {
					solapp = cal
					solapps = cs1
				}

				if cal == s {
					nbs++
					fmt.Println(nbs, cs1)
					found = true
					//return true
				}

				c1 := []int{cal}

				for l, d := range c {
					if (l == i) || (l == j) {
						continue
					}
					c1 = append(c1, d)
				}

				if search(c1, s, cs1) {
					return true
				}

			}

		}

	}
	return found
}

func main() {

	cartes := []int{1, 4, 5, 25, 2, 7}
	solution := 768
	var calcul string

	if !search(cartes, solution, calcul) {
		fmt.Println(solapps)
	}

}
