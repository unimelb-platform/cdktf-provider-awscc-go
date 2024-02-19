package devopsgururesourcecollection


type DevopsguruResourceCollectionResourceCollectionFilterTags struct {
	// A Tag key for DevOps Guru app boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/devopsguru_resource_collection#app_boundary_key DevopsguruResourceCollection#app_boundary_key}
	AppBoundaryKey *string `field:"optional" json:"appBoundaryKey" yaml:"appBoundaryKey"`
	// Tag values of DevOps Guru app boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/devopsguru_resource_collection#tag_values DevopsguruResourceCollection#tag_values}
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

