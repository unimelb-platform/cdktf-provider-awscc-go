package ssmassociation


type SsmAssociationTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ssm_association#key SsmAssociation#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ssm_association#values SsmAssociation#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

