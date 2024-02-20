package wafv2regexpatternset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2RegexPatternSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_regex_pattern_set#regular_expression_list Wafv2RegexPatternSet#regular_expression_list}.
	RegularExpressionList *[]*string `field:"required" json:"regularExpressionList" yaml:"regularExpressionList"`
	// Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_regex_pattern_set#scope Wafv2RegexPatternSet#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Description of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_regex_pattern_set#description Wafv2RegexPatternSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the RegexPatternSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_regex_pattern_set#name Wafv2RegexPatternSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_regex_pattern_set#tags Wafv2RegexPatternSet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

