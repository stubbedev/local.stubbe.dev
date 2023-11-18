package server

import (
	"log"
	"os"
	"strings"
)

func getStaticPaths(root string) []string {
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
			subdirs := getStaticPaths(subpath)
			folder_paths = append(folder_paths, subdirs...)
		}
	}
	return folder_paths
}

func getStaticRoutes(root_path string) []string {
	file_path_routes := append(getStaticPaths(root_path), "/")
	clean_paths := []string{}
	for _, p := range file_path_routes {
		clean_paths = append(clean_paths, strings.ReplaceAll(p, root_path, ""))
	}
	return clean_paths
}
