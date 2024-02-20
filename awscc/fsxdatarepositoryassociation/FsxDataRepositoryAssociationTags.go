package fsxdatarepositoryassociation


type FsxDataRepositoryAssociationTags struct {
	// A value that specifies the ``TagKey``, the name of the tag.
	//
	// Tag keys must be unique for the resource to which they are attached.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#key FsxDataRepositoryAssociation#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A value that specifies the ``TagValue``, the value assigned to the corresponding tag key.
	//
	// Tag values can be null and don't have to be unique in a tag set. For example, you can have a key-value pair in a tag set of ``finances : April`` and also of ``payroll : April``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fsx_data_repository_association#value FsxDataRepositoryAssociation#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

