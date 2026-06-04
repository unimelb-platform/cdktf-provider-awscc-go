package ssmguiconnectpreferences

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmguiconnectPreferencesConfig struct {
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
	// The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region.
	//
	// This includes details such as which S3 bucket recordings are stored in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ssmguiconnect_preferences#connection_recording_preferences SsmguiconnectPreferences#connection_recording_preferences}
	ConnectionRecordingPreferences *SsmguiconnectPreferencesConnectionRecordingPreferences `field:"optional" json:"connectionRecordingPreferences" yaml:"connectionRecordingPreferences"`
}

