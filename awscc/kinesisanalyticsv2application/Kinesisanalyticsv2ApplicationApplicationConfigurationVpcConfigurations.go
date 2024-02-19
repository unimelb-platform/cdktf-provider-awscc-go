package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationVpcConfigurations struct {
	// The array of SecurityGroup IDs used by the VPC configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#security_group_ids Kinesisanalyticsv2Application#security_group_ids}
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The array of Subnet IDs used by the VPC configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#subnet_ids Kinesisanalyticsv2Application#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

