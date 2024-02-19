package apigatewayapikey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayApiKeyConfig struct {
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
	// An MKT customer identifier, when integrating with the AWS SaaS Marketplace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#customer_id ApigatewayApiKey#customer_id}
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// The description of the ApiKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#description ApigatewayApiKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the ApiKey can be used by callers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#enabled ApigatewayApiKey#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies whether (``true``) or not (``false``) the key identifier is distinct from the created API key value.
	//
	// This parameter is deprecated and should not be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#generate_distinct_id ApigatewayApiKey#generate_distinct_id}
	GenerateDistinctId interface{} `field:"optional" json:"generateDistinctId" yaml:"generateDistinctId"`
	// A name for the API key.
	//
	// If you don't specify a name, CFN generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
	//  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#name ApigatewayApiKey#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#stage_keys ApigatewayApiKey#stage_keys}
	StageKeys interface{} `field:"optional" json:"stageKeys" yaml:"stageKeys"`
	// The key-value map of strings.
	//
	// The valid character set is [a-zA-Z+-=._:/]. The tag key can be up to 128 characters and must not start with ``aws:``. The tag value can be up to 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#tags ApigatewayApiKey#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies a value of the API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_api_key#value ApigatewayApiKey#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

