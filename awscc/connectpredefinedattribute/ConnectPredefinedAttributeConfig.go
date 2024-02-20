package connectpredefinedattribute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectPredefinedAttributeConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_predefined_attribute#instance_arn ConnectPredefinedAttribute#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the predefined attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_predefined_attribute#name ConnectPredefinedAttribute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values of a predefined attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_predefined_attribute#values ConnectPredefinedAttribute#values}
	Values *ConnectPredefinedAttributeValues `field:"required" json:"values" yaml:"values"`
}

