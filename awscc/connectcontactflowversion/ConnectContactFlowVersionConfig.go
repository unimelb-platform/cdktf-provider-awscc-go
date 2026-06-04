package connectcontactflowversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectContactFlowVersionConfig struct {
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
	// The ARN of the contact flow this version is tied to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_contact_flow_version#contact_flow_id ConnectContactFlowVersion#contact_flow_id}
	ContactFlowId *string `field:"required" json:"contactFlowId" yaml:"contactFlowId"`
	// The description of the version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connect_contact_flow_version#description ConnectContactFlowVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

