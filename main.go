package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("pkgDir", os.Getenv("INPUT_PKG_DIR"))
	fmt.Println("versionTag", os.Getenv("INPUT_VERSION_TAG"))
}
