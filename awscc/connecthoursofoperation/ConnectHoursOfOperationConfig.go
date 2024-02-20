package connecthoursofoperation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectHoursOfOperationConfig struct {
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
	// Configuration information for the hours of operation: day, start time, and end time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#config ConnectHoursOfOperation#config}
	Config interface{} `field:"required" json:"config" yaml:"config"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#instance_arn ConnectHoursOfOperation#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the hours of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#name ConnectHoursOfOperation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The time zone of the hours of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#time_zone ConnectHoursOfOperation#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// The description of the hours of operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#description ConnectHoursOfOperation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_hours_of_operation#tags ConnectHoursOfOperation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

