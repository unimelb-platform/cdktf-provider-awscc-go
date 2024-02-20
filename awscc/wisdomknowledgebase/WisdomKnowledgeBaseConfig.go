package wisdomknowledgebase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WisdomKnowledgeBaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#knowledge_base_type WisdomKnowledgeBase#knowledge_base_type}.
	KnowledgeBaseType *string `field:"required" json:"knowledgeBaseType" yaml:"knowledgeBaseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#name WisdomKnowledgeBase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#description WisdomKnowledgeBase#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#rendering_configuration WisdomKnowledgeBase#rendering_configuration}.
	RenderingConfiguration *WisdomKnowledgeBaseRenderingConfiguration `field:"optional" json:"renderingConfiguration" yaml:"renderingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#server_side_encryption_configuration WisdomKnowledgeBase#server_side_encryption_configuration}.
	ServerSideEncryptionConfiguration *WisdomKnowledgeBaseServerSideEncryptionConfiguration `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#source_configuration WisdomKnowledgeBase#source_configuration}.
	SourceConfiguration *WisdomKnowledgeBaseSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#tags WisdomKnowledgeBase#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

