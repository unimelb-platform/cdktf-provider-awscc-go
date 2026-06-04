package workspacesinstancesvolume


type WorkspacesinstancesVolumeTagSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume#resource_type WorkspacesinstancesVolume#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The tags to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume#tags WorkspacesinstancesVolume#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

