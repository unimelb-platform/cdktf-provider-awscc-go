package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRules struct {
	// An array of objects representing the details of a replication destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_replication_configuration#destinations EcrReplicationConfiguration#destinations}
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// An array of objects representing the details of a repository filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_replication_configuration#repository_filters EcrReplicationConfiguration#repository_filters}
	RepositoryFilters interface{} `field:"optional" json:"repositoryFilters" yaml:"repositoryFilters"`
}

