package inspectorv2cisscanconfiguration


type Inspectorv2CisScanConfigurationTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#account_ids Inspectorv2CisScanConfiguration#account_ids}.
	AccountIds *[]*string `field:"required" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#target_resource_tags Inspectorv2CisScanConfiguration#target_resource_tags}.
	TargetResourceTags interface{} `field:"required" json:"targetResourceTags" yaml:"targetResourceTags"`
}

