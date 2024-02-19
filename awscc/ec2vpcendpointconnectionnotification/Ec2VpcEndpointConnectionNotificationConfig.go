package ec2vpcendpointconnectionnotification

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpcEndpointConnectionNotificationConfig struct {
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
	// The endpoint events for which to receive notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_endpoint_connection_notification#connection_events Ec2VpcEndpointConnectionNotification#connection_events}
	ConnectionEvents *[]*string `field:"required" json:"connectionEvents" yaml:"connectionEvents"`
	// The ARN of the SNS topic for the notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_endpoint_connection_notification#connection_notification_arn Ec2VpcEndpointConnectionNotification#connection_notification_arn}
	ConnectionNotificationArn *string `field:"required" json:"connectionNotificationArn" yaml:"connectionNotificationArn"`
	// The ID of the endpoint service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_endpoint_connection_notification#service_id Ec2VpcEndpointConnectionNotification#service_id}
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
	// The ID of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_endpoint_connection_notification#vpc_endpoint_id Ec2VpcEndpointConnectionNotification#vpc_endpoint_id}
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

