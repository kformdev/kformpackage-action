package main

import (
	"fmt"
	"os"
)

func main() {
	pkgDir := os.Getenv("INPUT_PKG_DIR")
	versionTag := os.Getenv("INPUT_VERSION_TAG")

	fmt.Println("pkgDir", pkgDir)
	fmt.Println("versionTag", versionTag)
}
