package iotbillinggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotBillingGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_billing_group#billing_group_name IotBillingGroup#billing_group_name}.
	BillingGroupName *string `field:"optional" json:"billingGroupName" yaml:"billingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_billing_group#billing_group_properties IotBillingGroup#billing_group_properties}.
	BillingGroupProperties *IotBillingGroupBillingGroupProperties `field:"optional" json:"billingGroupProperties" yaml:"billingGroupProperties"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_billing_group#tags IotBillingGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

