package filestructure

import (
	"fmt"
	"os"
	"path/filepath"
)

var output []string

// traverse recursively traverses the directory tree.
func traverse(path string, prefix string) {
	files, err := os.ReadDir(path)
	if err != nil {
		output = append(output, prefix+"[Error reading directory: "+err.Error()+"]")
		return
	}

	for i, file := range files {
		connector := "├── "
		if i == len(files)-1 {
			connector = "└── "
		}

		line := prefix + connector + file.Name()
		output = append(output, line)

		if file.IsDir() {
			newPrefix := prefix
			if i == len(files)-1 {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			traverse(filepath.Join(path, file.Name()), newPrefix)
		}
	}
}

func File_structure() {
	root := "." // current directory
	output = append(output, root)
	traverse(root, "")

	// Write output to file
	f, err := os.Create("file_Structure.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	for _, line := range output {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Directory structure written to file_Structure.txt")
}
