// Package handlers is used to process Cobra commands
package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/twelvee/k8sbox/internal/k8sbox"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"k8s.io/utils/strings/slices"
)

// HandleGetCommand is the k8sbox get command handler
func HandleGetCommand(context context.Context, modelName string, namespace string) {
	if !slices.Contains(structs.GetEnvironmentAliases(), modelName) {
		fmt.Printf("An invalid argument. Available arguments: %s", strings.Join(structs.GetEnvironmentAliases(), ", "))
		os.Exit(1)
	}

	KuberExecutable(context, namespace)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Namespace", "Boxes")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	s := k8sbox.GetStorageService()
	environments, err := s.GetEnvironments(namespace)
	if err != nil {
		fmt.Println("No environments found.")
		os.Exit(1)
	}
	for _, widget := range environments {
		tbl.AddRow(widget.ID, widget.Name, widget.Namespace, formatBoxesToTable(widget.Boxes))
	}

	tbl.Print()
}

func formatBoxesToTable(boxes []structs.Box) string {
	var boxWidget []string
	for _, box := range boxes {
		boxWidget = append(boxWidget, box.Name)
	}
	return strings.Join(boxWidget, ", ")
}
