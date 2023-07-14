// Package structs contain every boxie public structs
package structs

// Cluster is your cluster in a struct
type Cluster struct {
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

// ClusterService is a public ClusterService
type ClusterService struct {
	TestConnection func(cluster Cluster) (bool, error)
}
