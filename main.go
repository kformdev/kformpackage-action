package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/henderiw-nephio/kform/tools/pkg/fsys"
	"github.com/henderiw/logger/log"
)

func main() {
	// init logging
	l := log.NewLogger(&log.HandlerOptions{Name: "kformpackage-action-logger", AddSource: false})
	slog.SetDefault(l)

	// init context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	ctx = log.IntoContext(ctx, l)

	if err := runMain(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s \n", err.Error())
		cancel()
		os.Exit(1)
	}
	os.Exit(0)
}

func runMain(ctx context.Context) error {
	rootPath := os.Getenv("INPUT_PKG_DIR")

	fmt.Println("reposiotry", os.Getenv("GITHUB_REPOSITORY"))
	fmt.Println("refName", os.Getenv("GITHUB_REF_NAME"))
	fmt.Println("github token", os.Getenv("GITHUB_TOKEN"))

	fs := fsys.NewDiskFS(".")
	f, err := fs.Stat(rootPath)
	if err != nil {
		return fmt.Errorf("cannot create a pkg, rootpath %s does not exist", rootPath)
	}
	if !f.IsDir() {
		return fmt.Errorf("cannot initialize a pkg on a file, please provide a directory instead, file: %s", rootPath)
	}
	return nil
}
