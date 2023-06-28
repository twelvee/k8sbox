// Package get is used to process get commands
package get

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/twelvee/boxie/internal/boxie"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"k8s.io/utils/strings/slices"
	"os"
	"strconv"
	"time"
)

// HandleShelfGetCommand is the boxie shelf get command handler
func HandleShelfGetCommand(context context.Context, resourceType string) {
	err := validateShelfGetRequest(resourceType)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if slices.Contains(structs.GetBoxAliaces(), resourceType) {
		err = listBoxesInShelf()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
	os.Exit(0)
}

func listBoxesInShelf() error {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Namespace", "Type", "Created")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	shelf := boxie.GetShelf(os.Getenv("BOXIE_SHELF_DRIVER"), os.Getenv("SHELF_DSN"))
	boxes, err := shelf.GetBoxes()
	if err != nil {
		return fmt.Errorf("No boxes found.")
	}
	for _, widget := range boxes {
		i, err := strconv.ParseInt(widget.Created, 10, 64)
		if err != nil {
			fmt.Println("Box " + widget.Name + " has invalid created time format.")
		}
		tm := time.Unix(i, 0)
		tbl.AddRow(widget.Name, widget.Namespace, widget.Type, tm)
	}

	tbl.Print()
	return nil
}
