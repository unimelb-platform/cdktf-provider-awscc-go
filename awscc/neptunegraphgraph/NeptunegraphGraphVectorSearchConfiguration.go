package neptunegraphgraph


type NeptunegraphGraphVectorSearchConfiguration struct {
	// The vector search dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/neptunegraph_graph#vector_search_dimension NeptunegraphGraph#vector_search_dimension}
	VectorSearchDimension *float64 `field:"optional" json:"vectorSearchDimension" yaml:"vectorSearchDimension"`
}

