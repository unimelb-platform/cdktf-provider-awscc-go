package datasynclocationazureblob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationAzureBlobConfig struct {
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
	// The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#agent_arns DatasyncLocationAzureBlob#agent_arns}
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#azure_access_tier DatasyncLocationAzureBlob#azure_access_tier}
	AzureAccessTier *string `field:"optional" json:"azureAccessTier" yaml:"azureAccessTier"`
	// The specific authentication type that you want DataSync to use to access your Azure Blob Container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#azure_blob_authentication_type DatasyncLocationAzureBlob#azure_blob_authentication_type}
	AzureBlobAuthenticationType *string `field:"optional" json:"azureBlobAuthenticationType" yaml:"azureBlobAuthenticationType"`
	// The URL of the Azure Blob container that was described.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#azure_blob_container_url DatasyncLocationAzureBlob#azure_blob_container_url}
	AzureBlobContainerUrl *string `field:"optional" json:"azureBlobContainerUrl" yaml:"azureBlobContainerUrl"`
	// Specifies the shared access signature (SAS) that DataSync uses to access your Azure Blob Storage container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#azure_blob_sas_configuration DatasyncLocationAzureBlob#azure_blob_sas_configuration}
	AzureBlobSasConfiguration *DatasyncLocationAzureBlobAzureBlobSasConfiguration `field:"optional" json:"azureBlobSasConfiguration" yaml:"azureBlobSasConfiguration"`
	// Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#azure_blob_type DatasyncLocationAzureBlob#azure_blob_type}
	AzureBlobType *string `field:"optional" json:"azureBlobType" yaml:"azureBlobType"`
	// The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#subdirectory DatasyncLocationAzureBlob#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#tags DatasyncLocationAzureBlob#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

