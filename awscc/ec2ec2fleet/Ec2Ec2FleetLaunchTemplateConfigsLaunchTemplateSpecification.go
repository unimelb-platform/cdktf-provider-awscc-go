package ec2ec2fleet


type Ec2Ec2FleetLaunchTemplateConfigsLaunchTemplateSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#version Ec2Ec2Fleet#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#launch_template_id Ec2Ec2Fleet#launch_template_id}.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#launch_template_name Ec2Ec2Fleet#launch_template_name}.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
}

