package ivsstage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IvsStageConfig struct {
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
	// Configuration object for individual participant recording, to attach to the new stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_stage#auto_participant_recording_configuration IvsStage#auto_participant_recording_configuration}
	AutoParticipantRecordingConfiguration *IvsStageAutoParticipantRecordingConfiguration `field:"optional" json:"autoParticipantRecordingConfiguration" yaml:"autoParticipantRecordingConfiguration"`
	// Stage name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_stage#name IvsStage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_stage#tags IvsStage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

