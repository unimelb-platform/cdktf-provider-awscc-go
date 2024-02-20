package xraygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type XrayGroupConfig struct {
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
	// The case-sensitive name of the new group. Names must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#group_name XrayGroup#group_name}
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The filter expression defining criteria by which to group traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#filter_expression XrayGroup#filter_expression}
	FilterExpression *string `field:"optional" json:"filterExpression" yaml:"filterExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#insights_configuration XrayGroup#insights_configuration}.
	InsightsConfiguration *XrayGroupInsightsConfiguration `field:"optional" json:"insightsConfiguration" yaml:"insightsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#tags XrayGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

