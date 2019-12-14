package commands

import (
	"github.com/ewave-com/devbox/utils"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func ListAction(c *cli.Context) error {

	projects, error := utils.GetProjectList()
	if error != nil {
		return error
	}

	color.Yellow("List of available projects:")

	for _, project := range projects {
		color.Green(project.Name)
	}

	return nil
}
