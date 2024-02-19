package appstreamappblock

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppstreamAppBlockConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#name AppstreamAppBlock#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#source_s3_location AppstreamAppBlock#source_s3_location}.
	SourceS3Location *AppstreamAppBlockSourceS3Location `field:"required" json:"sourceS3Location" yaml:"sourceS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#description AppstreamAppBlock#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#display_name AppstreamAppBlock#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#packaging_type AppstreamAppBlock#packaging_type}.
	PackagingType *string `field:"optional" json:"packagingType" yaml:"packagingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#post_setup_script_details AppstreamAppBlock#post_setup_script_details}.
	PostSetupScriptDetails *AppstreamAppBlockPostSetupScriptDetails `field:"optional" json:"postSetupScriptDetails" yaml:"postSetupScriptDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#setup_script_details AppstreamAppBlock#setup_script_details}.
	SetupScriptDetails *AppstreamAppBlockSetupScriptDetails `field:"optional" json:"setupScriptDetails" yaml:"setupScriptDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#tags AppstreamAppBlock#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

