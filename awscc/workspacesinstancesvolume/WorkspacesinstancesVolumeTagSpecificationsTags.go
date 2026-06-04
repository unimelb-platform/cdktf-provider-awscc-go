package workspacesinstancesvolume


type WorkspacesinstancesVolumeTagSpecificationsTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume#key WorkspacesinstancesVolume#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume#value WorkspacesinstancesVolume#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

