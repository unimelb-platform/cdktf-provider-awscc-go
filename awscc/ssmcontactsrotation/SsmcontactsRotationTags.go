package ssmcontactsrotation


type SsmcontactsRotationTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_rotation#key SsmcontactsRotation#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmcontacts_rotation#value SsmcontactsRotation#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

