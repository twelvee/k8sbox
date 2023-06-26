// Package get is used to process get commands
package get

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/internal/boxie/handlers"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
)

// HandleGetCommand is the boxie get command handler
func HandleGetCommand(context context.Context, modelName string, namespace string) {
	if !slices.Contains(structs.GetEnvironmentAliases(), modelName) {
		fmt.Printf("An invalid argument. Available arguments: %s", strings.Join(structs.GetEnvironmentAliases(), ", "))
		os.Exit(1)
	}

	handlers.KuberExecutable(context, namespace)

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Namespace", "Boxes")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	s := boxie.GetStorageService()
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
