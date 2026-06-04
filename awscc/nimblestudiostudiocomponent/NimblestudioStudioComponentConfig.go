package nimblestudiostudiocomponent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NimblestudioStudioComponentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#name NimblestudioStudioComponent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#studio_id NimblestudioStudioComponent#studio_id}.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#type NimblestudioStudioComponent#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#configuration NimblestudioStudioComponent#configuration}.
	Configuration *NimblestudioStudioComponentConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#description NimblestudioStudioComponent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#ec_2_security_group_ids NimblestudioStudioComponent#ec_2_security_group_ids}.
	Ec2SecurityGroupIds *[]*string `field:"optional" json:"ec2SecurityGroupIds" yaml:"ec2SecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#initialization_scripts NimblestudioStudioComponent#initialization_scripts}.
	InitializationScripts interface{} `field:"optional" json:"initializationScripts" yaml:"initializationScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#script_parameters NimblestudioStudioComponent#script_parameters}.
	ScriptParameters interface{} `field:"optional" json:"scriptParameters" yaml:"scriptParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#subtype NimblestudioStudioComponent#subtype}.
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_studio_component#tags NimblestudioStudioComponent#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

