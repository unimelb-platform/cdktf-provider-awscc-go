package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailsActionIncludeResources struct {
	// Use to configure lifecycle actions on AMIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#amis ImagebuilderLifecyclePolicy#amis}
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Use to configure lifecycle actions on containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#containers ImagebuilderLifecyclePolicy#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// Use to configure lifecycle actions on snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#snapshots ImagebuilderLifecyclePolicy#snapshots}
	Snapshots interface{} `field:"optional" json:"snapshots" yaml:"snapshots"`
}

