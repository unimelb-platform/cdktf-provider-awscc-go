package eventsapidestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventsApiDestinationConfig struct {
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
	// The arn of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_api_destination#connection_arn EventsApiDestination#connection_arn}
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_api_destination#http_method EventsApiDestination#http_method}.
	HttpMethod *string `field:"required" json:"httpMethod" yaml:"httpMethod"`
	// Url endpoint to invoke.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_api_destination#invocation_endpoint EventsApiDestination#invocation_endpoint}
	InvocationEndpoint *string `field:"required" json:"invocationEndpoint" yaml:"invocationEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_api_destination#description EventsApiDestination#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_api_destination#invocation_rate_limit_per_second EventsApiDestination#invocation_rate_limit_per_second}.
	InvocationRateLimitPerSecond *float64 `field:"optional" json:"invocationRateLimitPerSecond" yaml:"invocationRateLimitPerSecond"`
	// Name of the apiDestination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/events_api_destination#name EventsApiDestination#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

