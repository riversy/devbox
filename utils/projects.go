package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Project struct {
	Name     string
	FilePath string
}

func GetProjectList() ([]Project, error) {
	var projects []Project
	projectsPath := "./projects"

	currentPath, error := os.Getwd()
	if error != nil {
		return nil, error
	}

	files, error := ioutil.ReadDir(projectsPath)
	if error != nil {
		return nil, error
	}

	for _, file := range files {

		if !file.IsDir() {
			continue
		}

		project := Project{
			Name:     file.Name(),
			FilePath: filepath.Join(currentPath, file.Name()),
		}

		projects = append(projects, project)
	}

	return projects, nil
}
