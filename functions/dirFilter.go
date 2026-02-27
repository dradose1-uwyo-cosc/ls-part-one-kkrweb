// Kane Kriz
// UWYO Linux Programming
// HW3 - ls part one
// 26 Feb 2026
// dirFilter.go

//

package functions

import (
	"os"
)

func dirFilter(entries []os.DirEntry) []os.DirEntry {
	writePos := 0

	for i := 0; i < len(entries); i++ {
		currentDirEntry := entries[i]
		name := currentDirEntry.Name()
		isHidden := len(name) > 0 && name[0] == '.'

		if !isHidden {
			entries[writePos] = currentDirEntry
			writePos++
		}
	}

	return entries[:writePos]
}

//
