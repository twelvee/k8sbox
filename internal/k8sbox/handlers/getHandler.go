// Package handlers is used to handle cobra commands
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
	"k8s.io/utils/strings/slices"
)

// HandleGetCommand is the k8sbox get command handler
func HandleGetCommand(context context.Context, modelName string, flags []string) {
	if slices.Contains(structs.GetEnvironmentAliases(), modelName) {
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
			tbl.AddRow(widget.ID, widget.Name, widget.Namespace, formatBoxesToTable(widget.Boxes))
		}

		tbl.Print()
		return
	}
	fmt.Println(fmt.Sprintf("Invalid argument. Available types: %s", strings.Join(structs.GetEnvironmentAliases(), ", ")))
	os.Exit(1)
}

func formatBoxesToTable(boxes []structs.Box) string {
	var boxWidget []string
	for _, box := range boxes {
		boxWidget = append(boxWidget, box.Name)
	}
	return strings.Join(boxWidget, ", ")
}
