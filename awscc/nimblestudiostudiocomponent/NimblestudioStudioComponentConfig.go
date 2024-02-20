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
	// <p>The name for the studio component.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#name NimblestudioStudioComponent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>The studio ID. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#studio_id NimblestudioStudioComponent#studio_id}
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#type NimblestudioStudioComponent#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// <p>The configuration of the studio component, based on component type.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#configuration NimblestudioStudioComponent#configuration}
	Configuration *NimblestudioStudioComponentConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// <p>The description.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#description NimblestudioStudioComponent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// <p>The EC2 security groups that control access to the studio component.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#ec_2_security_group_ids NimblestudioStudioComponent#ec_2_security_group_ids}
	Ec2SecurityGroupIds *[]*string `field:"optional" json:"ec2SecurityGroupIds" yaml:"ec2SecurityGroupIds"`
	// <p>Initialization scripts for studio components.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#initialization_scripts NimblestudioStudioComponent#initialization_scripts}
	InitializationScripts interface{} `field:"optional" json:"initializationScripts" yaml:"initializationScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#runtime_role_arn NimblestudioStudioComponent#runtime_role_arn}.
	RuntimeRoleArn *string `field:"optional" json:"runtimeRoleArn" yaml:"runtimeRoleArn"`
	// <p>Parameters for the studio component scripts.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#script_parameters NimblestudioStudioComponent#script_parameters}
	ScriptParameters interface{} `field:"optional" json:"scriptParameters" yaml:"scriptParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#secure_initialization_role_arn NimblestudioStudioComponent#secure_initialization_role_arn}.
	SecureInitializationRoleArn *string `field:"optional" json:"secureInitializationRoleArn" yaml:"secureInitializationRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#subtype NimblestudioStudioComponent#subtype}.
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#tags NimblestudioStudioComponent#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

