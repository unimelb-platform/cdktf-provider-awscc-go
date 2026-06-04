package securityhubstandard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubStandardConfig struct {
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
	// The ARN of the standard that you want to enable.
	//
	// To view a list of available ASH standards and their ARNs, use the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_standard#standards_arn SecurityhubStandard#standards_arn}
	StandardsArn *string `field:"required" json:"standardsArn" yaml:"standardsArn"`
	// Specifies which controls are to be disabled in a standard.   *Maximum*: ``100``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_standard#disabled_standards_controls SecurityhubStandard#disabled_standards_controls}
	DisabledStandardsControls interface{} `field:"optional" json:"disabledStandardsControls" yaml:"disabledStandardsControls"`
}

