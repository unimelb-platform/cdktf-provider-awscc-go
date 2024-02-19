package datazonedatasource


type DatazoneDataSourceAssetFormsInput struct {
	// The name of the metadata form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#form_name DatazoneDataSource#form_name}
	FormName *string `field:"required" json:"formName" yaml:"formName"`
	// The content of the metadata form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#content DatazoneDataSource#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The ID of the metadata form type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#type_identifier DatazoneDataSource#type_identifier}
	TypeIdentifier *string `field:"optional" json:"typeIdentifier" yaml:"typeIdentifier"`
	// The revision of the metadata form type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#type_revision DatazoneDataSource#type_revision}
	TypeRevision *string `field:"optional" json:"typeRevision" yaml:"typeRevision"`
}

