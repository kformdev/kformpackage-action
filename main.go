package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("args", os.Args)
	fmt.Println("pkgDir", os.Getenv("INPUT_PKG_DIR"))
	fmt.Println("versionTag", os.Getenv("INPUT_VERSION_TAG"))
	fmt.Println("github token", os.Getenv("GITHUB_TOKEN"))
}
