package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	downloadArg, err := getArgsFromTerminal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "args error: %v\n", err)
		os.Exit(1)
	}

	if err := downloadRepoToDir(downloadArg); err != nil {
		fmt.Fprintf(os.Stderr, "\nexec error: %v\ncheck your args\n", err)
		os.Exit(1)
	}
}

func getArgsFromTerminal() (*downloadArgs, error) {
	outDirPath := flag.String("out", "", "output dir")
	flag.Parse()

	if flag.NArg() != 1 {
		return nil, errors.New("args error, used by: godld [repo url]")
	}
	rawRepoURL := flag.Arg(0)

	return newDownloadArgs(rawRepoURL, *outDirPath)
}

func downloadRepoToDir(args *downloadArgs) error {
	cmd := exec.Command("git", "clone", args.repoURL, args.outDirPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Fprintln(os.Stdout, strings.Join([]string{"git", "clone", args.repoURL, args.outDirPath}, " "))

	return cmd.Run()
}
