package ssmresourcedatasync


type SsmResourceDataSyncSyncSourceAwsOrganizationsSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_resource_data_sync#organization_source_type SsmResourceDataSync#organization_source_type}.
	OrganizationSourceType *string `field:"required" json:"organizationSourceType" yaml:"organizationSourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssm_resource_data_sync#organizational_units SsmResourceDataSync#organizational_units}.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
}

