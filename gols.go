// Kane Kriz
// UWYO Linux Programming
// HW3 - ls part one
// 26 Feb 2026
// gols.go

//

package main

import (
	"bufio"
	"lspartonekkrweb/functions"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	useColor := functions.IsTerminal(os.Stdout)

	_ = functions.SimpleLS(writer, os.Args[1:], useColor)

	_ = writer.Flush()
}

//
