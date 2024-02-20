package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesSalesforce struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#object CustomerprofilesIntegration#object}.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#enable_dynamic_field_update CustomerprofilesIntegration#enable_dynamic_field_update}.
	EnableDynamicFieldUpdate interface{} `field:"optional" json:"enableDynamicFieldUpdate" yaml:"enableDynamicFieldUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#include_deleted_records CustomerprofilesIntegration#include_deleted_records}.
	IncludeDeletedRecords interface{} `field:"optional" json:"includeDeletedRecords" yaml:"includeDeletedRecords"`
}

