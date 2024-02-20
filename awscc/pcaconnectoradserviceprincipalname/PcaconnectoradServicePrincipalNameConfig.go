package pcaconnectoradserviceprincipalname

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PcaconnectoradServicePrincipalNameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_service_principal_name#connector_arn PcaconnectoradServicePrincipalName#connector_arn}.
	ConnectorArn *string `field:"optional" json:"connectorArn" yaml:"connectorArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_service_principal_name#directory_registration_arn PcaconnectoradServicePrincipalName#directory_registration_arn}.
	DirectoryRegistrationArn *string `field:"optional" json:"directoryRegistrationArn" yaml:"directoryRegistrationArn"`
}

