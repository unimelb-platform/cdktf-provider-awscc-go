package rdsoptiongroup


type RdsOptionGroupOptionConfigurationsOptionSettings struct {
	// The name of the option that has settings that you can set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#name RdsOptionGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The current value of the option setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#value RdsOptionGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

