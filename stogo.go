package main

import (
	"fmt"
	"log"
	"os"
	fp "path/filepath"
)

// IDEA:
// stogo sourcedir targetdir
//   this will make symlinks in targetdir to the files in sourcedir

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("usage: stogo sourcedir targetdir")
		return
	}

	source_path := args[0]
	target_path := args[1]

	source_dir, err := os.Open(source_path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := source_dir.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	source_stat, err := source_dir.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if !source_stat.IsDir() {
		fmt.Println("error: sourcedir has to be a directory")
		return
	}

	target_stat, err := os.Stat(target_path)
	if err != nil {
		log.Fatal(err)
	}
	if !target_stat.IsDir() {
		fmt.Println("error: targetdir has to be a directory")
		return
	}

	source_dents, err := source_dir.ReadDir(0)
	if err != nil {
		log.Fatal(err)
	}
	rel_source_path, err := fp.Rel(target_path, source_path)
	if err != nil {
		log.Fatal(err)
	}
	for _, dent := range source_dents {
		link_target := fp.Join(rel_source_path, dent.Name())
		link_path := fp.Join(target_path, dent.Name())
		if err := os.Symlink(link_target, link_path); err != nil {
			log.Fatal(err)
		}
	}
}
