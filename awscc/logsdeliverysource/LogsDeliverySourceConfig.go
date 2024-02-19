package logsdeliverysource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsDeliverySourceConfig struct {
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
	// The unique name of the Log source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery_source#name LogsDeliverySource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of logs being delivered.
	//
	// Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery_source#log_type LogsDeliverySource#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// The ARN of the resource that will be sending the logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery_source#resource_arn LogsDeliverySource#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The tags that have been assigned to this delivery source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_delivery_source#tags LogsDeliverySource#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

