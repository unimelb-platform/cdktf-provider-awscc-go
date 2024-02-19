package wisdomassistantassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WisdomAssistantAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_assistant_association#assistant_id WisdomAssistantAssociation#assistant_id}.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_assistant_association#association WisdomAssistantAssociation#association}.
	Association *WisdomAssistantAssociationAssociation `field:"required" json:"association" yaml:"association"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_assistant_association#association_type WisdomAssistantAssociation#association_type}.
	AssociationType *string `field:"required" json:"associationType" yaml:"associationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_assistant_association#tags WisdomAssistantAssociation#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

