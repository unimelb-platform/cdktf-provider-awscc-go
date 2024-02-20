package ssmassociation


type SsmAssociationTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_association#key SsmAssociation#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_association#values SsmAssociation#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

