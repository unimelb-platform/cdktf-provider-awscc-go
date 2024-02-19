package datasynclocationazureblob


type DatasyncLocationAzureBlobAzureBlobSasConfiguration struct {
	// Specifies the shared access signature (SAS) token, which indicates the permissions DataSync needs to access your Azure Blob Storage container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_azure_blob#azure_blob_sas_token DatasyncLocationAzureBlob#azure_blob_sas_token}
	AzureBlobSasToken *string `field:"required" json:"azureBlobSasToken" yaml:"azureBlobSasToken"`
}

