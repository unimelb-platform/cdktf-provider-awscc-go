package emrserverlessapplication


type EmrserverlessApplicationNetworkConfiguration struct {
	// The ID of the security groups in the VPC to which you want to connect your job or application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#security_group_ids EmrserverlessApplication#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect your job or application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#subnet_ids EmrserverlessApplication#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

