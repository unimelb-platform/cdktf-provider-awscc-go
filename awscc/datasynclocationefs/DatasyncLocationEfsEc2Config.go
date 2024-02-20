package datasynclocationefs


type DatasyncLocationEfsEc2Config struct {
	// The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_efs#security_group_arns DatasyncLocationEfs#security_group_arns}
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The ARN of the subnet that DataSync uses to access the target EFS file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_efs#subnet_arn DatasyncLocationEfs#subnet_arn}
	SubnetArn *string `field:"required" json:"subnetArn" yaml:"subnetArn"`
}

