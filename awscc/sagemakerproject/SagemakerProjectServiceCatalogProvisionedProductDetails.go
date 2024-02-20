package sagemakerproject


type SagemakerProjectServiceCatalogProvisionedProductDetails struct {
	// The identifier of the provisioning artifact (also known as a version).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#provisioned_product_id SagemakerProject#provisioned_product_id}
	ProvisionedProductId *string `field:"optional" json:"provisionedProductId" yaml:"provisionedProductId"`
	// Provisioned Product Status Message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#provisioned_product_status_message SagemakerProject#provisioned_product_status_message}
	ProvisionedProductStatusMessage *string `field:"optional" json:"provisionedProductStatusMessage" yaml:"provisionedProductStatusMessage"`
}

