// Package structs contain every boxie public structs
package structs

// Header is an http header to work with boxie
type Header struct {
	Name  string `toml:"name"`
	Value string `toml:"value"`
}
