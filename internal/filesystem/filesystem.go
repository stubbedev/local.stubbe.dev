package filesystem

import (
	"log"
	"os"
	"strings"
)

func GetFolderPaths(root string) []string {
	folder_paths := []string{}
	entries, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}

	for _, e := range entries {
		if e.IsDir() {
			subpath := root + e.Name()
			if !strings.HasSuffix(subpath, "/") {
				subpath = subpath + "/"
			}
			folder_paths = append(folder_paths, subpath)
			subdirs := GetFolderPaths(subpath)
			folder_paths = append(folder_paths, subdirs...)
		}
	}
	return folder_paths
}

func RemovePathPrefix(path string, prefix string) string {
	return strings.Replace(path, prefix, "/", 1)
}
