package configorganizationconformancepack


type ConfigOrganizationConformancePackConformancePackInputParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_organization_conformance_pack#parameter_name ConfigOrganizationConformancePack#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_organization_conformance_pack#parameter_value ConfigOrganizationConformancePack#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

