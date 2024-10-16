package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/xlab/treeprint"
)


func main() {
	var dir string
	fmt.Println("Enter the directory path: ")
	fmt.Scanf("%s", &dir)
	Generator(dir)
	
}
func Generator(rootDir string) {
	tree := treeprint.New()

	var root treeprint.Tree
	dirs := make(map[string]treeprint.Tree) 
	

    err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		
        if err != nil {
            fmt.Println(err)
            return err
        }
		if root == nil {
			root = tree.AddBranch(info.Name())
		} else if info.IsDir() {
			dirs[path] = root.AddBranch(info.Name())
		} else {
			
			parentDir := filepath.Dir(path)
			if parentBranch, ok := dirs[parentDir]; ok {
				parentBranch.AddNode(info.Name())
			}else{
				root.AddNode(info.Name())
			}
			
		}
        return nil
    })
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(tree.String())
}