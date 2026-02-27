// Kane Kriz
// UWYO Linux Programming
// HW3 - ls part one
// 26 Feb 2026
// simplels.go

//

package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func accessErrorHelper(input string, err error) {
	pathErr, ok := err.(*os.PathError)
	if ok {
		fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %s\n", input, pathErr.Err.Error())
	} else {
		fmt.Fprintf(os.Stderr, "gols: cannot access '%s': %s\n", input, err.Error())
	}
}

func SimpleLS(w io.Writer, args []string, useColor bool) error {
	argsUsed := args

	if len(argsUsed) == 0 {
		argsUsed = []string{"."}
	}

	var fileInputs []string
	var dirInputs []string

	for i := 0; i < len(argsUsed); i++ {
		input := argsUsed[i]

		infoFromLStat, err := os.Lstat(input)

		if err != nil {
			accessErrorHelper(input, err)
		} else {
			if infoFromLStat.IsDir() {
				dirInputs = append(dirInputs, input)
			} else {
				fileInputs = append(fileInputs, input)
			}
		}
	}

	sort.Strings(fileInputs)
	sort.Strings(dirInputs)

	for i := 0; i < len(fileInputs); i++ {
		input := fileInputs[i]

		infoFromLStat, err := os.Lstat(input)

		if err != nil {
			accessErrorHelper(input, err)
		} else {
			printName := filepath.Base(input)
			colorToUse := defColor

			if useColor {
				fileMode := infoFromLStat.Mode()
				isRegularFile := fileMode.IsRegular()
				hasFileModeBitmask := (fileMode & 0111) != 0

				switch {
				case infoFromLStat.IsDir():
					colorToUse = blue
				case isRegularFile && hasFileModeBitmask:
					colorToUse = green
				default:
					colorToUse = defColor
				}
			}

			colorToUse.ColorPrint(w, printName)
			_, writeErr := io.WriteString(w, "\n")

			if writeErr != nil {
				return writeErr
			}
		}
	}

	multipleDirsBool := len(dirInputs) > 1

	for i := 0; i < len(dirInputs); i++ {
		dirName := dirInputs[i]

		if multipleDirsBool {
			if i > 0 {
				_, writeErr := io.WriteString(w, "\n")
				if writeErr != nil {
					return writeErr
				}
			}

			_, writeErr := io.WriteString(w, dirName)
			if writeErr != nil {
				return writeErr
			}
			_, writeErr = io.WriteString(w, ":\n")
			if writeErr != nil {
				return writeErr
			}
		}

		dirEntries, err := os.ReadDir(dirName)

		if err != nil {
			accessErrorHelper(dirName, err)
		} else {
			dirEntries = dirFilter(dirEntries)
			var entryNames []string

			for i := 0; i < len(dirEntries); i++ {
				entryNames = append(entryNames, dirEntries[i].Name())
			}

			sort.Strings(entryNames)

			for i := 0; i < len(entryNames); i++ {
				itemName := entryNames[i]
				fullPath := filepath.Join(dirName, itemName)

				infoFromLStat, err := os.Lstat(fullPath)

				if err != nil {
					accessErrorHelper(fullPath, err)
				} else {
					colorToUse := defColor
					if useColor {
						fileMode := infoFromLStat.Mode()
						isRegularFile := fileMode.IsRegular()
						hasFileModeBitmask := (fileMode & 0111) != 0

						switch {
						case infoFromLStat.IsDir():
							colorToUse = blue
						case isRegularFile && hasFileModeBitmask:
							colorToUse = green
						default:
							colorToUse = defColor
						}
					}

					colorToUse.ColorPrint(w, itemName)
					_, writeErr := io.WriteString(w, "\n")

					if writeErr != nil {
						return writeErr
					}
				}
			}
		}
	}

	return nil
}

//
