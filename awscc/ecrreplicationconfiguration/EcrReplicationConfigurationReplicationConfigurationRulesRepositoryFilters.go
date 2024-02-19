package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRulesRepositoryFilters struct {
	// The repository filter to be applied for replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_replication_configuration#filter EcrReplicationConfiguration#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Type of repository filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_replication_configuration#filter_type EcrReplicationConfiguration#filter_type}
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
}

