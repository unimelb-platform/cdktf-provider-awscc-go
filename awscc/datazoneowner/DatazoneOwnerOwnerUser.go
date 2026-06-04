package datazoneowner


type DatazoneOwnerOwnerUser struct {
	// The ID of the owner user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_owner#user_identifier DatazoneOwner#user_identifier}
	UserIdentifier *string `field:"optional" json:"userIdentifier" yaml:"userIdentifier"`
}

