package transferconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TransferConnectorConfig struct {
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
	// Specifies the access role for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#access_role TransferConnector#access_role}
	AccessRole *string `field:"required" json:"accessRole" yaml:"accessRole"`
	// URL for Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#url TransferConnector#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Configuration for an AS2 connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#as_2_config TransferConnector#as_2_config}
	As2Config *TransferConnectorAs2Config `field:"optional" json:"as2Config" yaml:"as2Config"`
	// Specifies the logging role for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#logging_role TransferConnector#logging_role}
	LoggingRole *string `field:"optional" json:"loggingRole" yaml:"loggingRole"`
	// Configuration for an SFTP connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#sftp_config TransferConnector#sftp_config}
	SftpConfig *TransferConnectorSftpConfig `field:"optional" json:"sftpConfig" yaml:"sftpConfig"`
	// Key-value pairs that can be used to group and search for connectors.
	//
	// Tags are metadata attached to connectors for any purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#tags TransferConnector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

