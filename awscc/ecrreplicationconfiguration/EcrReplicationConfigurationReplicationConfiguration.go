package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfiguration struct {
	// An array of objects representing the replication rules for a replication configuration.
	//
	// A replication configuration may contain a maximum of 10 rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_replication_configuration#rules EcrReplicationConfiguration#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

