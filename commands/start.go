package commands

import (
	"errors"
	"fmt"
	"github.com/ewave-com/devbox/utils"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

// It's an entry point for project start flow. It allows:
// - Start project if it selected by name
// - Start all projects if --all flag used
// - Start prompt select if no project provided with input
func StartAction(c *cli.Context) error {

	var error error = nil
	projectName := c.Args().Get(0)

	if projectName != "" {
		if !IsValidProjectName(projectName) {
			return errors.New(
				fmt.Sprintf("Project with name '%s' is not exists.", projectName),
			)
		}
		StartProject(projectName)
	} else if utils.InArray("all", c.FlagNames()) {
		var projects []utils.Project
		projects, error = utils.GetProjectList()
		if error != nil {
			return error
		}

		for _, project := range projects {
			StartProject(project.Name)
		}
	} else {
		projectName, error = GetProjectNameFromPrompt()
		if error != nil {
			return error
		}

		StartProject(projectName)
	}

	return nil
}

// Check if entered project name exists in the projects list.
func IsValidProjectName(name string) bool {
	projects, error := utils.GetProjectList()
	if error != nil {
		return false
	}

	for _, project := range projects {
		if name == project.Name {
			return true
		}
	}

	return false
}

// It starts a prompt with a list of projects.
// After the user selects one, it retrieves its name.
func GetProjectNameFromPrompt() (string, error) {
	var _ int
	var projectName string
	var error error
	var projects []utils.Project
	projects, error = utils.GetProjectList()
	if error != nil {
		return "", error
	}

	var projectNames []string
	for _, project := range projects {
		projectNames = append(projectNames, project.Name)
	}

	prompt := promptui.Select{
		Label: "Select a project from the following list",
		Items: projectNames,
	}

	_, projectName, error = prompt.Run()

	return projectName, error
}

// Start project with the certain name selected
func StartProject(projectName string) error {
	fmt.Println()
	color.Yellow("================================================")
	color.Yellow(
		fmt.Sprintf("[START PROJECT] : %s", projectName),
	)
	color.Yellow("================================================")

	return nil
}
