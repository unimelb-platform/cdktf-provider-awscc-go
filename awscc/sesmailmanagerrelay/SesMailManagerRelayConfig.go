package sesmailmanagerrelay

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SesMailManagerRelayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_relay#authentication SesMailManagerRelay#authentication}.
	Authentication *SesMailManagerRelayAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_relay#server_name SesMailManagerRelay#server_name}.
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_relay#server_port SesMailManagerRelay#server_port}.
	ServerPort *float64 `field:"required" json:"serverPort" yaml:"serverPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_relay#relay_name SesMailManagerRelay#relay_name}.
	RelayName *string `field:"optional" json:"relayName" yaml:"relayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_relay#tags SesMailManagerRelay#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

