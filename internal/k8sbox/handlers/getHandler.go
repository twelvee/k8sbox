package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
)

func HandleGetCommand(getType string, flags []string, context context.Context) {
	if getType == "environment" {
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("ID", "Name", "Namespace", "Boxes")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		environments, err := utils.GetEnvironments()
		if err != nil {
			fmt.Println("No environments found")
			os.Exit(1)
		}
		for _, widget := range environments {
			tbl.AddRow(widget.Id, widget.Name, widget.Namespace, formatBoxesToTable(widget.Boxes))
		}

		tbl.Print()
		return
	}
	fmt.Println("Invalid argument. Available types: environment")
	os.Exit(1)
}

func formatBoxesToTable(boxes []structs.Box) string {
	var boxWidget []string
	for _, box := range boxes {
		boxWidget = append(boxWidget, box.Name)
	}
	return strings.Join(boxWidget, ", ")
}
