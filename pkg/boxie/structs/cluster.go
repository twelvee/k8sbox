// Package structs contain every boxie public structs
package structs

// Cluster is your cluster in a struct
type Cluster struct {
	ID         int32
	Name       string
	Kubeconfig string
	IsActive   bool
	CreatedAt  string
}

// DeleteClusterRequest is rest api request for delete cluster method
type DeleteClusterRequest struct {
	Name string `json:"name"`
}

// GetClusterRequest is rest api request for get cluster method
type GetClusterRequest struct {
	Name string `json:"name"`
}
