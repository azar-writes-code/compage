package core

import (
	"github.com/kube-tarian/compage-core/internal/core/edge"
	"github.com/kube-tarian/compage-core/internal/core/node"
	"time"
)

// ModificationDetails has creation and modification details.
type ModificationDetails struct {
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Project is top level struct depicting the monorepo and resembling to github repository.
// It has a single compage.json and can have multiple nodes and edges (projects and connections) internally.
type Project struct {
	Name           string                 `json:"name"`
	RepositoryName string                 `json:"repositoryName"`
	UserName       string                 `json:"userName"`
	CompageYaml    *CompageYaml           `json:"compageYaml"`
	Metadata       map[string]interface{} `json:"metadata"`
	ModificationDetails
}

// CompageYaml has all edges and nodes
type CompageYaml struct {
	// Linkages (connection details) between Nodes
	Edges []*edge.Edge `json:"edges"`
	// Nodes represent components and has details of the component.
	Nodes []*node.Node `json:"nodes"`
}

// Rest Protocol
const Rest = "REST"

// Grpc Protocol
const Grpc = "GRPC"

// Ws Protocol
const Ws = "WS"
