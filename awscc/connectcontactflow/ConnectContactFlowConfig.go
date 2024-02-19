package connectcontactflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectContactFlowConfig struct {
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
	// The content of the contact flow in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow#content ConnectContactFlow#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the Amazon Connect instance (ARN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow#instance_arn ConnectContactFlow#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow#name ConnectContactFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow#type ConnectContactFlow#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow#description ConnectContactFlow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow#state ConnectContactFlow#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_contact_flow#tags ConnectContactFlow#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

