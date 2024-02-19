package cloudfrontfunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontFunctionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#function_code CloudfrontFunction#function_code}.
	FunctionCode *string `field:"required" json:"functionCode" yaml:"functionCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#function_config CloudfrontFunction#function_config}.
	FunctionConfig *CloudfrontFunctionFunctionConfig `field:"required" json:"functionConfig" yaml:"functionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#name CloudfrontFunction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#auto_publish CloudfrontFunction#auto_publish}.
	AutoPublish interface{} `field:"optional" json:"autoPublish" yaml:"autoPublish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#function_metadata CloudfrontFunction#function_metadata}.
	FunctionMetadata *CloudfrontFunctionFunctionMetadata `field:"optional" json:"functionMetadata" yaml:"functionMetadata"`
}

