package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/henderiw-nephio/kform/tools/pkg/fsys"
	"github.com/henderiw-nephio/kform/tools/pkg/pkgio"
	"github.com/henderiw/logger/log"
)

const (
	hostname = "ghcr.io"
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
	pkgName := os.Getenv("INPUT_PKG_NAME")
	repository := os.Getenv("GITHUB_REPOSITORY")
	refName := os.Getenv("GITHUB_REF_NAME")

	log := log.FromContext(ctx).With("repository", repository, "refName", refName, "rootPath", rootPath, "pkgName", pkgName)
	log.Info("run kformpkg action")

	fmt.Println("repository", repository)
	fmt.Println("refName", refName)
	fmt.Println("rootPath", rootPath)
	fmt.Println("pkgName", pkgName)
	fmt.Println("github token", os.Getenv("GITHUB_TOKEN"))

	fs := fsys.NewDiskFS(".")
	f, err := fs.Stat(rootPath)
	if err != nil {
		return fmt.Errorf("cannot create a pkg, rootpath %s does not exist", rootPath)
	}
	if !f.IsDir() {
		return fmt.Errorf("cannot initialize a pkg on a file, please provide a directory instead, file: %s", rootPath)
	}

	ref := fmt.Sprintf("%s/%s/%s:%s", hostname, repository, pkgName, refName)
	fmt.Println(ref)
	pkgrw := pkgio.NewPkgPushReadWriter(rootPath, ref)
	p := pkgio.Pipeline{
		Inputs:  []pkgio.Reader{pkgrw},
		Outputs: []pkgio.Writer{pkgrw},
	}
	return p.Execute(ctx)
}
