package imagebuilderworkflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderWorkflowConfig struct {
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
	// The name of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#name ImagebuilderWorkflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the workflow denotes whether the workflow is used to build, test, or distribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#type ImagebuilderWorkflow#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The version of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#version ImagebuilderWorkflow#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The change description of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#change_description ImagebuilderWorkflow#change_description}
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// The data of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#data ImagebuilderWorkflow#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#description ImagebuilderWorkflow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#kms_key_id ImagebuilderWorkflow#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The tags associated with the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#tags ImagebuilderWorkflow#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The uri of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_workflow#uri ImagebuilderWorkflow#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

