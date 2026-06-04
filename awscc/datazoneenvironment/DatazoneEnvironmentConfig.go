package datazoneenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneEnvironmentConfig struct {
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
	// The identifier of the Amazon DataZone domain in which the environment would be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#domain_identifier DatazoneEnvironment#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#name DatazoneEnvironment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the Amazon DataZone project in which the environment would be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#project_identifier DatazoneEnvironment#project_identifier}
	ProjectIdentifier *string `field:"required" json:"projectIdentifier" yaml:"projectIdentifier"`
	// The description of the Amazon DataZone environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#description DatazoneEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS account in which the Amazon DataZone environment is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#environment_account_identifier DatazoneEnvironment#environment_account_identifier}
	EnvironmentAccountIdentifier *string `field:"optional" json:"environmentAccountIdentifier" yaml:"environmentAccountIdentifier"`
	// The AWS region in which the Amazon DataZone environment is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#environment_account_region DatazoneEnvironment#environment_account_region}
	EnvironmentAccountRegion *string `field:"optional" json:"environmentAccountRegion" yaml:"environmentAccountRegion"`
	// The ID of the environment profile with which the Amazon DataZone environment would be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#environment_profile_identifier DatazoneEnvironment#environment_profile_identifier}
	EnvironmentProfileIdentifier *string `field:"optional" json:"environmentProfileIdentifier" yaml:"environmentProfileIdentifier"`
	// Environment role arn for custom aws environment permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#environment_role_arn DatazoneEnvironment#environment_role_arn}
	EnvironmentRoleArn *string `field:"optional" json:"environmentRoleArn" yaml:"environmentRoleArn"`
	// The glossary terms that can be used in the Amazon DataZone environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#glossary_terms DatazoneEnvironment#glossary_terms}
	GlossaryTerms *[]*string `field:"optional" json:"glossaryTerms" yaml:"glossaryTerms"`
	// The user parameters of the Amazon DataZone environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_environment#user_parameters DatazoneEnvironment#user_parameters}
	UserParameters interface{} `field:"optional" json:"userParameters" yaml:"userParameters"`
}

