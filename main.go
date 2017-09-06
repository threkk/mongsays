package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const version string = "0.0.1"
const width = 80

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

func splitInLines(s string) []string {
	lines := make([]string, 0)
	words := strings.Split(s, " ")

	acc := 0
	index := 0
	for i, word := range words {
		acc = acc + len(word) + 1 // Separation between words needs to be added.
		if acc > (width - 4) {
			line := strings.Join(words[index:i-1], " ")
			lines = append(lines, line)
			index = i
			acc = 0
		}
	}
	return lines
}

func showBalloon(lines []string) {
	if len(lines) == 1 && len(lines[0]) < 80 {
		// Line is smaller tha 80, we need to wrap it.
		//  _______
		// < Mong! >
		//  -------
	}

}

func showDog(dogType int) {
	fmt.Println(dogs[dogType])
}

func showVersion() {
	fmt.Println(version)
	os.Exit(0)
}

func showError(option int) {
	fmt.Printf("Invalid option: %d.\n\n", option)
	flag.Usage()
	os.Exit(1)
}

func main() {
	// Sets the output to the stdout instead of the stderr.
	flag.CommandLine.SetOutput(os.Stdout)

	// Change the command line
	flag.Usage = func() {
		fmt.Printf("Mong says - A dog version of cowsay (v%s)\n", version)
		fmt.Println()
		fmt.Println("Usage: mongsay [options] <text>")
		fmt.Println("Options:")
		fmt.Println()
		flag.PrintDefaults()
	}

	isMute := flag.Bool("mute", false, "Print the dog without the speech bubble.")
	isVersion := flag.Bool("version", false, "Print the version number.")
	dogType := flag.Int("type", 0, "Dog version to use: 0 (default), 1 or 2.")
	flag.Parse()

	text := strings.Join(flag.Args(), " ")

	if *isVersion {
		showVersion()
	}

	if *dogType >= len(dogs) || *dogType < 0 {
		showError(*dogType)
	}

	if !*isMute {
		showBalloon(splitInLines(text))
	}
	showDog(*dogType)
}
