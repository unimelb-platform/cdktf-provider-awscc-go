package devopsgururesourcecollection


type DevopsguruResourceCollectionResourceCollectionFilterCloudformation struct {
	// An array of CloudFormation stack names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/devopsguru_resource_collection#stack_names DevopsguruResourceCollection#stack_names}
	StackNames *[]*string `field:"optional" json:"stackNames" yaml:"stackNames"`
}

