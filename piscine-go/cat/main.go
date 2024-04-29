package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		str, err0 := io.ReadAll(os.Stdin)
		if err0 != nil {
			for _, c := range "ERROR: " + err0.Error() {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
		res := string(str)
		for _, c := range res {
			z01.PrintRune(rune(c))
		}
	} else {
		for _, c := range args {
			file, err := os.Open(c)
			fileState, err2 := file.Stat()
			if err != nil {
				for _, c := range "ERROR: " + err.Error() {
					z01.PrintRune(c)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			} else if err2 != nil {
				for _, c := range "ERROR: " + err2.Error() {
					z01.PrintRune(c)
				}
			} else {
				arr := make([]byte, fileState.Size())
				file.Read(arr)
				file.Close()
				for _, c := range string(arr) {
					z01.PrintRune(c)
				}
			}
		}
	}
}
