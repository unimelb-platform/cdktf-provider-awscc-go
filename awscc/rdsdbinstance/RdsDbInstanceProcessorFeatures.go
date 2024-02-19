package rdsdbinstance


type RdsDbInstanceProcessorFeatures struct {
	// The name of the processor feature. Valid names are coreCount and threadsPerCore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#name RdsDbInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of a processor feature name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#value RdsDbInstance#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

