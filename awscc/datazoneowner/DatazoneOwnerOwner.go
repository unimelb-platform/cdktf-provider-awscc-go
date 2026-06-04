package datazoneowner


type DatazoneOwnerOwner struct {
	// The properties of the domain unit owners group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_owner#group DatazoneOwner#group}
	Group *DatazoneOwnerOwnerGroup `field:"optional" json:"group" yaml:"group"`
	// The properties of the owner user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_owner#user DatazoneOwner#user}
	User *DatazoneOwnerOwnerUser `field:"optional" json:"user" yaml:"user"`
}

