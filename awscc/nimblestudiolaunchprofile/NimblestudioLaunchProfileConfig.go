package nimblestudiolaunchprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NimblestudioLaunchProfileConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#ec_2_subnet_ids NimblestudioLaunchProfile#ec_2_subnet_ids}.
	Ec2SubnetIds *[]*string `field:"required" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#launch_profile_protocol_versions NimblestudioLaunchProfile#launch_profile_protocol_versions}.
	LaunchProfileProtocolVersions *[]*string `field:"required" json:"launchProfileProtocolVersions" yaml:"launchProfileProtocolVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#name NimblestudioLaunchProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#stream_configuration NimblestudioLaunchProfile#stream_configuration}.
	StreamConfiguration *NimblestudioLaunchProfileStreamConfiguration `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#studio_component_ids NimblestudioLaunchProfile#studio_component_ids}.
	StudioComponentIds *[]*string `field:"required" json:"studioComponentIds" yaml:"studioComponentIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#studio_id NimblestudioLaunchProfile#studio_id}.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#description NimblestudioLaunchProfile#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/nimblestudio_launch_profile#tags NimblestudioLaunchProfile#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

