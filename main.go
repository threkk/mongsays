package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Version number.
const version string = "0.0.1"

// Maximum width of the screen.
const width = 80
const padding = 4

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
func splitInLines(str string) []string {
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

func fixLength(length int, str string) string {
	verb := fmt.Sprintf("%%%ds", length)
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
	result = append(result, fmt.Sprintf(" _%20s_ ", strings.Repeat("_", lineLen)))

	switch amountOfLines {
	case 1:
		// Only one line.
		//  _______
		// < Mong! >
		//  -------
		result = append(result, fmt.Sprintf("< %20s >", lines[0]))
	case 2:
		// Two lines.
		//  ------
		// / Mong!\
		// \ Mong!/
		//  ------
		result = append(result, fmt.Sprintf("/ %20s \\", lines[0]))
		result = append(result, fmt.Sprintf("\\ %20s /", lines[1]))
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
				result = append(result, fmt.Sprintf("/ %20s \\", line))
			case amountOfLines:
				result = append(result, fmt.Sprintf("\\ %20s /", line))
			default:
				result = append(result, fmt.Sprintf("| %20s |", line))
			}
		}
	}

	// Bottom border
	result = append(result, fmt.Sprintf(" -%20s- ", strings.Repeat("-", lineLen)))
	result = append(result, fmt.Sprintf("%20s", "//  "))
	// result = append(result, "\n")

	return result
}

// Returns the version of the application and a successful status code.
func getVersion() (string, int) {
	return fmt.Sprintf("%s\n", version), 0
}

// Returns the error message for a given option and unsuccessful status code.
func getError(option int) (string, int) {
	return fmt.Sprintf("Invalid option: %d.\n", option), 1
}

func main() {
	// Error message
	flag.Usage = usage

	// Flags
	isMute := flag.Bool("mute", false, "Print the dog without the speech bubble.")
	isVersion := flag.Bool("version", false, "Print the version number.")
	dogType := flag.Int("type", 0, "Dog version to use: 0 (default), 1 or 2.")

	// Parse the CLI.
	flag.Parse()

	// Main text
	text := strings.Join(flag.Args(), " ")

	var out string
	var code int

	if *isVersion {
		out, code = getVersion()
	} else if *dogType < 0 || *dogType >= len(dogs) {
		out, code = getError(*dogType)
	} else if *isMute || len(text) == 0 {
		out = dogs[*dogType]
		code = 0
	} else {
		balloon := applyBalloon(splitInLines(text))
		out = strings.Join(append(balloon, dogs[*dogType]), "\n")
		code = 0
	}

	fmt.Printf(out)
	fmt.Println()
	os.Exit(code)
}
