package ec2networkinsightsaccessscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2NetworkInsightsAccessScopeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_access_scope#exclude_paths Ec2NetworkInsightsAccessScope#exclude_paths}.
	ExcludePaths interface{} `field:"optional" json:"excludePaths" yaml:"excludePaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_access_scope#match_paths Ec2NetworkInsightsAccessScope#match_paths}.
	MatchPaths interface{} `field:"optional" json:"matchPaths" yaml:"matchPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_insights_access_scope#tags Ec2NetworkInsightsAccessScope#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

