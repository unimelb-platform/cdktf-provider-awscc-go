package pcaconnectorscepconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PcaconnectorscepConnectorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcaconnectorscep_connector#certificate_authority_arn PcaconnectorscepConnector#certificate_authority_arn}.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcaconnectorscep_connector#mobile_device_management PcaconnectorscepConnector#mobile_device_management}.
	MobileDeviceManagement *PcaconnectorscepConnectorMobileDeviceManagement `field:"optional" json:"mobileDeviceManagement" yaml:"mobileDeviceManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcaconnectorscep_connector#tags PcaconnectorscepConnector#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

