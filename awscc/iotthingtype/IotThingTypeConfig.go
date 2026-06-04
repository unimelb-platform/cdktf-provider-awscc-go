package iotthingtype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotThingTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_type#deprecate_thing_type IotThingType#deprecate_thing_type}.
	DeprecateThingType interface{} `field:"optional" json:"deprecateThingType" yaml:"deprecateThingType"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_type#tags IotThingType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_type#thing_type_name IotThingType#thing_type_name}.
	ThingTypeName *string `field:"optional" json:"thingTypeName" yaml:"thingTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_type#thing_type_properties IotThingType#thing_type_properties}.
	ThingTypeProperties *IotThingTypeThingTypeProperties `field:"optional" json:"thingTypeProperties" yaml:"thingTypeProperties"`
}

