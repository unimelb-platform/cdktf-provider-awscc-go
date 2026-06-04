package logsintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsIntegrationConfig struct {
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
	// User provided identifier for integration, unique to the user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_integration#integration_name LogsIntegration#integration_name}
	IntegrationName *string `field:"required" json:"integrationName" yaml:"integrationName"`
	// The type of the Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_integration#integration_type LogsIntegration#integration_type}
	IntegrationType *string `field:"required" json:"integrationType" yaml:"integrationType"`
	// OpenSearchResourceConfig for the given Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_integration#resource_config LogsIntegration#resource_config}
	ResourceConfig *LogsIntegrationResourceConfig `field:"required" json:"resourceConfig" yaml:"resourceConfig"`
}

