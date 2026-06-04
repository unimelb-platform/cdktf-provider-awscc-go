package securitylakedatalake


type SecuritylakeDataLakeLifecycleConfigurationExpiration struct {
	// Number of days before data expires in the Amazon Security Lake object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securitylake_data_lake#days SecuritylakeDataLake#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

