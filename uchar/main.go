// Bash completion: ../cmdctl/completions/uchar.bash

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"crdx.org/col"
	"github.com/samber/lo"
	"golang.org/x/text/unicode/runenames"
)

func getInput() []string {
	if len(os.Args) > 1 {
		return os.Args[1:]
	} else {
		bytes := lo.Must(io.ReadAll(os.Stdin))
		return strings.Split(strings.TrimRight(string(bytes), "\n"), " ")
	}
}

func main() {
	col.Init()

	for i, arg := range getInput() {
		if i != 0 {
			fmt.Println()
		}
		for _, r := range arg {
			fmt.Println(describe(r))
		}
	}
}

func describe(char rune) string {
	name := runenames.Name(char)

	var types []string
	if unicode.IsControl(char) {
		types = append(types, "control")
	}

	if unicode.IsDigit(char) {
		types = append(types, "digit")
	}

	if unicode.IsGraphic(char) {
		types = append(types, "graphic")
	}

	if unicode.IsLetter(char) {
		types = append(types, "letter")
	}

	if unicode.IsLower(char) {
		types = append(types, "lower")
	}

	if unicode.IsMark(char) {
		types = append(types, "mark")
	}

	if unicode.IsNumber(char) {
		types = append(types, "number")
	}

	if unicode.IsPrint(char) {
		types = append(types, "printable")
	}

	if unicode.IsPunct(char) {
		types = append(types, "punct")
	}

	if unicode.IsSpace(char) {
		types = append(types, "space")
	}

	if unicode.IsSymbol(char) {
		types = append(types, "symbol")
	}

	if unicode.IsTitle(char) {
		types = append(types, "title")
	}

	if unicode.IsUpper(char) {
		types = append(types, "upper")
	}

	return fmt.Sprintf(
		"%q @ %d — %s — %s",
		char,
		char,
		col.Yellow(name),
		col.Cyan(strings.Join(types, ", ")),
	)
}
