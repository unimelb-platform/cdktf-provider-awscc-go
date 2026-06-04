package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentS3AccessPoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#policy FsxS3AccessPointAttachment#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fsx_s3_access_point_attachment#vpc_configuration FsxS3AccessPointAttachment#vpc_configuration}.
	VpcConfiguration *FsxS3AccessPointAttachmentS3AccessPointVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

