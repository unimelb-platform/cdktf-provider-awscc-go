package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfiguration struct {
	// An array of objects representing the replication destinations and repository filters for a replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecr_replication_configuration#rules EcrReplicationConfiguration#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

