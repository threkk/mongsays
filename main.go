package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const VERSION string = "0.0.1"

var DOGS = []string{

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

func printText(text string) {

}

func main() {
	// Sets the output to the stdout instead of the stderr.
	flag.CommandLine.SetOutput(os.Stdout)

	// Change the command line
	flag.Usage = func() {
		fmt.Printf("Mong says - A dog version of cowsay (v%s)\n", VERSION)
		fmt.Println()
		fmt.Println("Usage: mongsay [options] <text>")
		fmt.Println("Options:")
		fmt.Println()
		flag.PrintDefaults()
	}

	var isMute *bool = flag.Bool("mute", false, "Print the dog without the speech bubble.")
	var isVersion *bool = flag.Bool("version", false, "Print the version number.")
	var dogType *int = flag.Int("type", 0, "Dog version to use: 0 (default), 1 or 2.")
	flag.Parse()

	text := strings.Join(flag.Args(), " ")

	if *isVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if !*isMute {
		printText(text)
	}

	if *dogType >= len(DOGS) || *dogType < 0 {
		fmt.Printf("Invalid option: %d.\n\n", *dogType)
		flag.Usage()
		os.Exit(1)
	} else {
		fmt.Println(DOGS[*dogType])
	}
}
