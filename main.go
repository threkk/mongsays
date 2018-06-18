package main

import (
	"flag"
	"fmt"
	"github.com/threkk/mongsays/terminal"
	"os"
	"strings"
)

const (
	version string = "1.0.0"
	padding        = 4
)

var (
	isQuiet   bool
	isVersion bool
	dogType   int
)

// Array containing three different types of ASCII art for the dogs.
var dogs = []string{

	0: ` \ ______/ U'-,
  }        /~~
 /_)^ --,r'
|b      |b`,
	1: `         __
        /  \
       / ..|\
      (_\  |_)
      /  \@'
     /     \
_   /      |
\\/  \  | _\
 \   /_ || \\_
  \____)|_) \_) `,
	2: `      /)---(\
\\   (/ . . \)
 \\__)-\(*)/
 \_       (_
 (___/-(____)`,
}

// Change the command line usage text
var usage = func() {
	fmt.Printf("Mong says - A dog version of cowsay (v%s)\n", version)
	fmt.Printf("\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("  mongsays [-type number] <text>...\n")
	fmt.Printf("  mongsays -version\n")
	fmt.Printf("  mongsays -h | --help\n")
	fmt.Printf("\n")
	fmt.Printf("Options:\n")
	fmt.Printf("\n")

	flag.PrintDefaults()
}

// Splits a "s" string in lines of length "width" with a total "padding" per line.
// It returns an array with a line item.
func splitInLines(str string, width int) []string {
	length := width - padding
	line := ""
	lines := make([]string, 0)

	chars := []rune(str)
	end := len(chars)
	for i, char := range chars {
		line = line + string(char)
		if (i+1)%length == 0 {
			lines = append(lines, line)
			line = ""
		} else if (i + 1) == end {
			lines = append(lines, line)
		}
	}

	return lines
}

func alignRight(str string, length int) string {
	verb := fmt.Sprintf("%%%ds", length)
	return fmt.Sprintf(verb, str)
}

func alignLeft(str string, length int) string {
	verb := fmt.Sprintf("%%-%ds", length)
	return fmt.Sprintf(verb, str)
}

// Apply a balloon to the given array of strings. It will wrap all the strings
// with an ASCII balloon. The ascii balloon shape depends on the amount of lines
// that it has.
func applyBalloon(lines []string) []string {
	result := make([]string, 0)
	amountOfLines := len(lines)
	lineLen := len(lines[0])

	// Top border
	result = append(result, fmt.Sprintf(" _%s_ ", strings.Repeat("_", lineLen)))

	switch amountOfLines {
	case 1:
		// Only one line.
		//  _______
		// < Mong! >
		//  -------
		result = append(result, fmt.Sprintf("< %s >", lines[0]))
	case 2:
		// Two lines.
		//  ------
		// / Mong!\
		// \ Mong!/
		//  ------
		result = append(result, fmt.Sprintf("/ %s \\", lines[0]))
		result = append(result, fmt.Sprintf("\\ %s /", alignLeft(lines[1], lineLen)))
	default:
		// More than two.
		//  -------
		// / Mong! \
		// | Mong! |
		// \ Mong! /
		//  -------
		for i, line := range lines {
			switch i {
			case 0:
				result = append(result, fmt.Sprintf("/ %s \\", line))
			default:
				result = append(result, fmt.Sprintf("| %s |", line))
			case amountOfLines - 1:
				result = append(result, fmt.Sprintf("\\ %s /", alignLeft(line, lineLen)))
			}
		}
	}

	// Bottom border
	result = append(result, fmt.Sprintf(" -%s- ", strings.Repeat("-", lineLen)))
	result = append(result, fmt.Sprintf("%s", "//  "))

	for i, r := range result {
		result[i] = alignRight(r, 20)
	}

	return result
}

// Returns the error message for a given option and unsuccessful status code.
func getError(option int) (string, int) {
	return fmt.Sprintf("Invalid option: %d.\n", option), 1
}

func init() {

	// Error message
	flag.Usage = usage

	// Flags
	flag.BoolVar(&isQuiet, "quiet", false, "Print the dog without the speech bubble.")
	flag.BoolVar(&isVersion, "version", false, "Print the version number.")
	flag.IntVar(&dogType, "type", 0, "Dog version to use: 0 (default), 1 or 2.")
}

func main() {
	// Parse the CLI.
	flag.Parse()

	// Main text
	text := strings.Join(flag.Args(), " ")

	if isVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if dogType < 0 || dogType >= len(dogs) {
		fmt.Printf("Invalid option: %d.\n", dogType)
		os.Exit(1)
	}

	balloon := make([]string, 0)
	if !isQuiet && len(text) > 0 {
		width := terminal.GetColumns()
		balloon = applyBalloon(splitInLines(text, int(width)))
	}

	fmt.Printf(strings.Join(append(balloon, dogs[dogType]), "\n"))
	os.Exit(0)
}
