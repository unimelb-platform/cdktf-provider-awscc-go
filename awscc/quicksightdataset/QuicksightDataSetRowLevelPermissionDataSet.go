package quicksightdataset


type QuicksightDataSetRowLevelPermissionDataSet struct {
	// <p>The Amazon Resource Name (ARN) of the permission dataset.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#arn QuicksightDataSet#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#permission_policy QuicksightDataSet#permission_policy}.
	PermissionPolicy *string `field:"required" json:"permissionPolicy" yaml:"permissionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#format_version QuicksightDataSet#format_version}.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// <p>The namespace associated with the row-level permissions dataset.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#namespace QuicksightDataSet#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

