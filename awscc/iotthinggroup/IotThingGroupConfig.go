package iotthinggroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotThingGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_group#parent_group_name IotThingGroup#parent_group_name}.
	ParentGroupName *string `field:"optional" json:"parentGroupName" yaml:"parentGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_group#query_string IotThingGroup#query_string}.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_group#tags IotThingGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_group#thing_group_name IotThingGroup#thing_group_name}.
	ThingGroupName *string `field:"optional" json:"thingGroupName" yaml:"thingGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_group#thing_group_properties IotThingGroup#thing_group_properties}.
	ThingGroupProperties *IotThingGroupThingGroupProperties `field:"optional" json:"thingGroupProperties" yaml:"thingGroupProperties"`
}

