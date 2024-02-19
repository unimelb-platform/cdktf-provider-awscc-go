package shieldprotection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ShieldProtectionConfig struct {
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
	// Friendly name for the Protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_protection#name ShieldProtection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_protection#resource_arn ShieldProtection#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The automatic application layer DDoS mitigation settings for a Protection.
	//
	// This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_protection#application_layer_automatic_response_configuration ShieldProtection#application_layer_automatic_response_configuration}
	ApplicationLayerAutomaticResponseConfiguration *ShieldProtectionApplicationLayerAutomaticResponseConfiguration `field:"optional" json:"applicationLayerAutomaticResponseConfiguration" yaml:"applicationLayerAutomaticResponseConfiguration"`
	// The Amazon Resource Names (ARNs) of the health check to associate with the protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_protection#health_check_arns ShieldProtection#health_check_arns}
	HealthCheckArns *[]*string `field:"optional" json:"healthCheckArns" yaml:"healthCheckArns"`
	// One or more tag key-value pairs for the Protection object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/shield_protection#tags ShieldProtection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

