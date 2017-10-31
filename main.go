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

// Splits a "s" string in lines of length "width" with a total "padding" per line.
func splitInLines(str string, width int, padding int) []string {
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

func showBalloon(lines []string) {
	amountOfLines := len(lines)
	firstLineLen := len(lines[0])
	topBorder := strings.Repeat("_", firstLineLen)
	bottomBorder := strings.Repeat("-", firstLineLen)

	// Top border
	fmt.Printf(" _%s_ \n", topBorder)
	switch amountOfLines {
	case 1:
		// Only one line.
		//  _______
		// < Mong! >
		//  -------
		fmt.Printf("< %s >\n", lines[0])
	case 2:
		// Two lines.
		//  ------
		// / Mong!\
		// \ Mong!/
		//  ------
		fmt.Printf("/ %s \\\n", lines[0])
		fmt.Printf("\\ %s /\n", lines[1])
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
				fmt.Printf("/ %s \\\n", line)
			case amountOfLines:
				fmt.Printf("\\ %s /\n", line)
			default:
				fmt.Printf("| %s |\n", line)
			}
		}
	}

	// Bottom border
	fmt.Printf(" -%s- \n", bottomBorder)
	fmt.Printf("   \\ \n")
	fmt.Printf("\n")
}

// Displays the version of the application.
func showVersion() {
	fmt.Printf("%s\n", version)
	os.Exit(0)
}

// Displays the error message in case the option is incorrect.
func showError(option int) {
	fmt.Printf("Invalid option: %d.\n\n", option)
	flag.Usage()
	os.Exit(1)
}

func main() {
	// Gets the command name and the arguments.
	cmdName, cmdArgs := os.Args[0], os.Args[1:]

	// Creates a new CLI parser.
	cmd := flag.NewFlagSet(cmdName, flag.ExitOnError)
	cmd.SetOutput(os.Stdout)

	isMute := cmd.Bool("mute", false, "Print the dog without the speech bubble.")
	isVersion := cmd.Bool("version", false, "Print the version number.")
	dogType := cmd.Int("type", 0, "Dog version to use: 0 (default), 1 or 2.")

	// Change the command line usage text.
	cmd.Usage = func() {
		fmt.Printf("Mong says - A dog version of cowsay (v%s)\n", version)
		fmt.Printf("\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("  mongsays [-type number] <text>...\n")
		fmt.Printf("  mongsays -version\n")
		fmt.Printf("  mongsays -h | --help\n")
		fmt.Printf("\n")
		fmt.Printf("Options:\n")
		fmt.Printf("\n")

		cmd.PrintDefaults()
	}

	cmd.Parse(cmdArgs)
	text := strings.Join(cmd.Args(), " ")

	if *isVersion {
		showVersion()
	}

	if *dogType < 0 || *dogType >= len(dogs) {
		showError(*dogType)
	}

	if !*isMute && len(text) > 0 {
		showBalloon(splitInLines(text, width, 4))
	}

	fmt.Printf("%s\n", dogs[*dogType])
	os.Exit(0)
}
