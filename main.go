package main

import (
	"fmt"
	"os"
)

func main() {
	pkgDir := os.Getenv("INPUT_PKGDIR")
	versionTag := os.Getenv("INPUT_VERSIONTAG")

	fmt.Println("pkgDir", pkgDir)
	fmt.Println("versionTag", versionTag)
}
