package appintegrationsapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppintegrationsApplicationConfig struct {
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
	// Application source config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#application_source_config AppintegrationsApplication#application_source_config}
	ApplicationSourceConfig *AppintegrationsApplicationApplicationSourceConfig `field:"required" json:"applicationSourceConfig" yaml:"applicationSourceConfig"`
	// The application description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#description AppintegrationsApplication#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#name AppintegrationsApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#namespace AppintegrationsApplication#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The configuration of events or requests that the application has access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#permissions AppintegrationsApplication#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// The tags (keys and values) associated with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appintegrations_application#tags AppintegrationsApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

