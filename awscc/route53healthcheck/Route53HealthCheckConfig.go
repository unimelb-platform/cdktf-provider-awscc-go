package route53healthcheck

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53HealthCheckConfig struct {
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
	// A complex type that contains information about the health check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_health_check#health_check_config Route53HealthCheck#health_check_config}
	HealthCheckConfig *Route53HealthCheckHealthCheckConfig `field:"required" json:"healthCheckConfig" yaml:"healthCheckConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53_health_check#health_check_tags Route53HealthCheck#health_check_tags}
	HealthCheckTags interface{} `field:"optional" json:"healthCheckTags" yaml:"healthCheckTags"`
}

