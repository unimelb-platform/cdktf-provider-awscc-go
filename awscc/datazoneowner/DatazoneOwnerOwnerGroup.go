package datazoneowner


type DatazoneOwnerOwnerGroup struct {
	// The ID of the domain unit owners group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_owner#group_identifier DatazoneOwner#group_identifier}
	GroupIdentifier *string `field:"optional" json:"groupIdentifier" yaml:"groupIdentifier"`
}

