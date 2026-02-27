// Kane Kriz
// UWYO Linux Programming
// HW3 - ls part one
// 26 Feb 2026
// color.go

//

package functions

import (
	"io"
)

type color string

const (
	blue          color = "\x1b[34m"
	green         color = "\x1b[32m"
	defColor      color = ""
	endFormatting color = "\x1b[0m"
)

func (c color) ColorPrint(w io.Writer, s string) {
	switch c {
	case defColor:
		_, _ = io.WriteString(w, s)
	default:
		_, _ = io.WriteString(w, string(c))
		_, _ = io.WriteString(w, s)
		_, _ = io.WriteString(w, string(endFormatting))
	}
}

//
