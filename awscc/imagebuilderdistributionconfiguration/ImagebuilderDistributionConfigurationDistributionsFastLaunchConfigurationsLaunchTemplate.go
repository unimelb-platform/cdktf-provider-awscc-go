package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsFastLaunchConfigurationsLaunchTemplate struct {
	// The ID of the launch template to use for faster launching for a Windows AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#launch_template_id ImagebuilderDistributionConfiguration#launch_template_id}
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template to use for faster launching for a Windows AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#launch_template_name ImagebuilderDistributionConfiguration#launch_template_name}
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version of the launch template to use for faster launching for a Windows AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#launch_template_version ImagebuilderDistributionConfiguration#launch_template_version}
	LaunchTemplateVersion *string `field:"optional" json:"launchTemplateVersion" yaml:"launchTemplateVersion"`
}

