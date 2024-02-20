package datazoneenvironmentblueprintconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneEnvironmentBlueprintConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#domain_identifier DatazoneEnvironmentBlueprintConfiguration#domain_identifier}.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#enabled_regions DatazoneEnvironmentBlueprintConfiguration#enabled_regions}.
	EnabledRegions *[]*string `field:"required" json:"enabledRegions" yaml:"enabledRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#environment_blueprint_identifier DatazoneEnvironmentBlueprintConfiguration#environment_blueprint_identifier}.
	EnvironmentBlueprintIdentifier *string `field:"required" json:"environmentBlueprintIdentifier" yaml:"environmentBlueprintIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#manage_access_role_arn DatazoneEnvironmentBlueprintConfiguration#manage_access_role_arn}.
	ManageAccessRoleArn *string `field:"optional" json:"manageAccessRoleArn" yaml:"manageAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#provisioning_role_arn DatazoneEnvironmentBlueprintConfiguration#provisioning_role_arn}.
	ProvisioningRoleArn *string `field:"optional" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_environment_blueprint_configuration#regional_parameters DatazoneEnvironmentBlueprintConfiguration#regional_parameters}.
	RegionalParameters interface{} `field:"optional" json:"regionalParameters" yaml:"regionalParameters"`
}

