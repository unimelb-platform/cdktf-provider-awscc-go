package ivsrecordingconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvsRecordingConfigurationConfig struct {
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
	// Recording Destination Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#destination_configuration IvsRecordingConfiguration#destination_configuration}
	DestinationConfiguration *IvsRecordingConfigurationDestinationConfiguration `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Recording Configuration Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#name IvsRecordingConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Recording Reconnect Window Seconds. (0 means disabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#recording_reconnect_window_seconds IvsRecordingConfiguration#recording_reconnect_window_seconds}
	RecordingReconnectWindowSeconds *float64 `field:"optional" json:"recordingReconnectWindowSeconds" yaml:"recordingReconnectWindowSeconds"`
	// Rendition Configuration describes which renditions should be recorded for a stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#rendition_configuration IvsRecordingConfiguration#rendition_configuration}
	RenditionConfiguration *IvsRecordingConfigurationRenditionConfiguration `field:"optional" json:"renditionConfiguration" yaml:"renditionConfiguration"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#tags IvsRecordingConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Recording Thumbnail Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#thumbnail_configuration IvsRecordingConfiguration#thumbnail_configuration}
	ThumbnailConfiguration *IvsRecordingConfigurationThumbnailConfiguration `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
}

