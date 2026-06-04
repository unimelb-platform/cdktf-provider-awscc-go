package odbodbnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OdbOdbNetworkConfig struct {
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
	// The AWS Availability Zone (AZ) where the ODB network is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#availability_zone OdbOdbNetwork#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID of the AZ where the ODB network is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#availability_zone_id OdbOdbNetwork#availability_zone_id}
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The CIDR range of the backup subnet in the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#backup_subnet_cidr OdbOdbNetwork#backup_subnet_cidr}
	BackupSubnetCidr *string `field:"optional" json:"backupSubnetCidr" yaml:"backupSubnetCidr"`
	// The CIDR range of the client subnet in the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#client_subnet_cidr OdbOdbNetwork#client_subnet_cidr}
	ClientSubnetCidr *string `field:"optional" json:"clientSubnetCidr" yaml:"clientSubnetCidr"`
	// The DNS prefix to the default DNS domain name. The default DNS domain name is oraclevcn.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#default_dns_prefix OdbOdbNetwork#default_dns_prefix}
	DefaultDnsPrefix *string `field:"optional" json:"defaultDnsPrefix" yaml:"defaultDnsPrefix"`
	// Specifies whether to delete associated OCI networking resources along with the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#delete_associated_resources OdbOdbNetwork#delete_associated_resources}
	DeleteAssociatedResources interface{} `field:"optional" json:"deleteAssociatedResources" yaml:"deleteAssociatedResources"`
	// The user-friendly name of the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#display_name OdbOdbNetwork#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Tags to assign to the Odb Network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/odb_odb_network#tags OdbOdbNetwork#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

