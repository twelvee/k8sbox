// Package structs contain every k8sbox public structs
package structs

// Header is an http header to work with k8sbox
type Header struct {
	Name  string `toml:"name"`
	Value string `toml:"value"`
}
