package entityresolutionidmappingworkflow

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EntityresolutionIdMappingWorkflowConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#id_mapping_techniques EntityresolutionIdMappingWorkflow#id_mapping_techniques}.
	IdMappingTechniques *EntityresolutionIdMappingWorkflowIdMappingTechniques `field:"required" json:"idMappingTechniques" yaml:"idMappingTechniques"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#input_source_config EntityresolutionIdMappingWorkflow#input_source_config}.
	InputSourceConfig interface{} `field:"required" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#output_source_config EntityresolutionIdMappingWorkflow#output_source_config}.
	OutputSourceConfig interface{} `field:"required" json:"outputSourceConfig" yaml:"outputSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#role_arn EntityresolutionIdMappingWorkflow#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the IdMappingWorkflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#workflow_name EntityresolutionIdMappingWorkflow#workflow_name}
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// The description of the IdMappingWorkflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#description EntityresolutionIdMappingWorkflow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#tags EntityresolutionIdMappingWorkflow#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

