package handlers

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
)

func HandleGetCommand(getType string, flags []string, context context.Context) {
	if getType == "environment" {
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("ID", "Name", "Namespace")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		environments, err := utils.GetEnvironments()
		if err != nil {
			fmt.Println("No environments found")
			os.Exit(1)
		}
		for _, widget := range environments {
			tbl.AddRow(widget.Id, widget.Name, widget.Namespace)
		}

		tbl.Print()
		return
	}
	fmt.Println("Invalid argument. Available types: environment")
	os.Exit(1)
}
