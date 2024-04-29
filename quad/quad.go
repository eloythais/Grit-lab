package piscine

import "github.com/01-edu/z01"

// QuadA
func QuadA(x, y int) {
	if !(x > 0 && y > 0) {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == 0 && j == x-1) || (j == 0 && i == y-1) || (i == y-1 && j == x-1) {
				z01.PrintRune('o')
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				z01.PrintRune('-')
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				z01.PrintRune('|')
			}
			if j == x-1 && i != y-1 {
				z01.PrintRune('\n')
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

// QuadB
func QuadB(x, y int) {
	if !(x > 0 && y > 0) {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == x-1 && i != 0 && j > 0) {
				z01.PrintRune('/')
			}
			if (i == 0 && j == x-1 && j > 0) || (j == 0 && i == y-1 && i != 0) {
				z01.PrintRune('\\')
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				z01.PrintRune('*')
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				z01.PrintRune('*')
			}
			if j == x-1 && i != y-1 {
				z01.PrintRune('\n')
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

// QuadC
func QuadC(x, y int) {
	if !(x > 0 && y > 0) {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == 0 && j == x-1) {
				z01.PrintRune('A')
			}
			if (j == 0 && i == y-1 && i > 0) || (i == y-1 && j == x-1 && i > 0) {
				z01.PrintRune('C')
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				z01.PrintRune('B')
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				z01.PrintRune('B')
			}
			if j == x-1 && i != y-1 {
				z01.PrintRune('\n')
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

// QuadD
func QuadD(x, y int) {
	if !(x > 0 && y > 0) {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == 0) {
				z01.PrintRune('A')
			}
			if (i == 0 && j == x-1 && j > 0) || (i == y-1 && j == x-1 && i > 0 && j > 0) {
				z01.PrintRune('C')
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				z01.PrintRune('B')
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				z01.PrintRune('B')
			}
			if j == x-1 && i != y-1 {
				z01.PrintRune('\n')
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

// QuadE
func QuadE(x, y int) {
	if !(x > 0 && y > 0) {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 && j == 0) || (i == y-1 && j == x-1 && i > 0 && j > 0) {
				z01.PrintRune('A')
			}
			if (j == 0 && i == y-1 && i > 0) || (i == 0 && j == x-1 && j > 0) {
				z01.PrintRune('C')
			}
			if (i == 0 || i == y-1) && (j > 0 && j < x-1) {
				z01.PrintRune('B')
			}
			if i > 0 && i < y-1 && (j == 0 || j == x-1) {
				z01.PrintRune('B')
			}
			if j == x-1 && i != y-1 {
				z01.PrintRune('\n')
			}
			if i > 0 && i < y-1 && (j > 0 && j < x-1) {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
