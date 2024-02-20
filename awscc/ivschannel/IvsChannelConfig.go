package ivschannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvsChannelConfig struct {
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
	// Whether the channel is authorized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#authorized IvsChannel#authorized}
	Authorized interface{} `field:"optional" json:"authorized" yaml:"authorized"`
	// Whether the channel allows insecure ingest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#insecure_ingest IvsChannel#insecure_ingest}
	InsecureIngest interface{} `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// Channel latency mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#latency_mode IvsChannel#latency_mode}
	LatencyMode *string `field:"optional" json:"latencyMode" yaml:"latencyMode"`
	// Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#name IvsChannel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional transcode preset for the channel.
	//
	// This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string ("").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#preset IvsChannel#preset}
	Preset *string `field:"optional" json:"preset" yaml:"preset"`
	// Recording Configuration ARN.
	//
	// A value other than an empty string indicates that recording is enabled. Default: "" (recording is disabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#recording_configuration_arn IvsChannel#recording_configuration_arn}
	RecordingConfigurationArn *string `field:"optional" json:"recordingConfigurationArn" yaml:"recordingConfigurationArn"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#tags IvsChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_channel#type IvsChannel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

