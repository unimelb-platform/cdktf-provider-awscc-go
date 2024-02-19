package appflowconnectorprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowConnectorProfileConfig struct {
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
	// Mode in which data transfer should be enabled.
	//
	// Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connection_mode AppflowConnectorProfile#connection_mode}
	ConnectionMode *string `field:"required" json:"connectionMode" yaml:"connectionMode"`
	// The maximum number of items to retrieve in a single batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connector_profile_name AppflowConnectorProfile#connector_profile_name}
	ConnectorProfileName *string `field:"required" json:"connectorProfileName" yaml:"connectorProfileName"`
	// List of Saas providers that need connector profile to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connector_type AppflowConnectorProfile#connector_type}
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// The label of the connector.
	//
	// The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connector_label AppflowConnectorProfile#connector_label}
	ConnectorLabel *string `field:"optional" json:"connectorLabel" yaml:"connectorLabel"`
	// Connector specific configurations needed to create connector profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connector_profile_config AppflowConnectorProfile#connector_profile_config}
	ConnectorProfileConfig *AppflowConnectorProfileConnectorProfileConfig `field:"optional" json:"connectorProfileConfig" yaml:"connectorProfileConfig"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables.
	//
	// If it's not provided, AWS Lambda uses a default service key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#kms_arn AppflowConnectorProfile#kms_arn}
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}

