package ecrreplicationconfiguration


type EcrReplicationConfigurationReplicationConfigurationRulesDestinations struct {
	// A Region to replicate to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_replication_configuration#region EcrReplicationConfiguration#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The account ID of the destination registry to replicate to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_replication_configuration#registry_id EcrReplicationConfiguration#registry_id}
	RegistryId *string `field:"required" json:"registryId" yaml:"registryId"`
}

