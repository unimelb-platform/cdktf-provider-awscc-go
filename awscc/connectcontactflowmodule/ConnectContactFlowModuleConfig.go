package connectcontactflowmodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectContactFlowModuleConfig struct {
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
	// The content of the contact flow module in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow_module#content ConnectContactFlowModule#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the Amazon Connect instance (ARN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow_module#instance_arn ConnectContactFlowModule#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the contact flow module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow_module#name ConnectContactFlowModule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the contact flow module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow_module#description ConnectContactFlowModule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state of the contact flow module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow_module#state ConnectContactFlowModule#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow_module#tags ConnectContactFlowModule#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

