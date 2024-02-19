package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataLicenseSpecifications struct {
	// The Amazon Resource Name (ARN) of the license configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#license_configuration_arn Ec2LaunchTemplate#license_configuration_arn}
	LicenseConfigurationArn *string `field:"optional" json:"licenseConfigurationArn" yaml:"licenseConfigurationArn"`
}

