package datasynclocationfsxontap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationFsxOntapConfig struct {
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
	// The ARNs of the security groups that are to use to configure the FSx ONTAP file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#security_group_arns DatasyncLocationFsxOntap#security_group_arns}
	SecurityGroupArns *[]*string `field:"required" json:"securityGroupArns" yaml:"securityGroupArns"`
	// The Amazon Resource Name (ARN) for the FSx ONTAP SVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#storage_virtual_machine_arn DatasyncLocationFsxOntap#storage_virtual_machine_arn}
	StorageVirtualMachineArn *string `field:"required" json:"storageVirtualMachineArn" yaml:"storageVirtualMachineArn"`
	// Configuration settings for NFS or SMB protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#protocol DatasyncLocationFsxOntap#protocol}
	Protocol *DatasyncLocationFsxOntapProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// A subdirectory in the location's path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#subdirectory DatasyncLocationFsxOntap#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_ontap#tags DatasyncLocationFsxOntap#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

