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
	// <p>Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.
	//
	// These subnets must support the specified instance types. </p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#ec_2_subnet_ids NimblestudioLaunchProfile#ec_2_subnet_ids}
	Ec2SubnetIds *[]*string `field:"required" json:"ec2SubnetIds" yaml:"ec2SubnetIds"`
	// <p>The version number of the protocol that is used by the launch profile.
	//
	// The only valid
	//             version is "2021-03-31".</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#launch_profile_protocol_versions NimblestudioLaunchProfile#launch_profile_protocol_versions}
	LaunchProfileProtocolVersions *[]*string `field:"required" json:"launchProfileProtocolVersions" yaml:"launchProfileProtocolVersions"`
	// <p>The name for the launch profile.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#name NimblestudioLaunchProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// <p>A configuration for a streaming session.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#stream_configuration NimblestudioLaunchProfile#stream_configuration}
	StreamConfiguration *NimblestudioLaunchProfileStreamConfiguration `field:"required" json:"streamConfiguration" yaml:"streamConfiguration"`
	// <p>Unique identifiers for a collection of studio components that can be used with this             launch profile.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#studio_component_ids NimblestudioLaunchProfile#studio_component_ids}
	StudioComponentIds *[]*string `field:"required" json:"studioComponentIds" yaml:"studioComponentIds"`
	// <p>The studio ID. </p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#studio_id NimblestudioLaunchProfile#studio_id}
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
	// <p>The description.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#description NimblestudioLaunchProfile#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_launch_profile#tags NimblestudioLaunchProfile#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

