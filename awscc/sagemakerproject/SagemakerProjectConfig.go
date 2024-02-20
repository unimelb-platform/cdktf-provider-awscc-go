package sagemakerproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerProjectConfig struct {
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
	// The name of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#project_name SagemakerProject#project_name}
	ProjectName *string `field:"required" json:"projectName" yaml:"projectName"`
	// Input ServiceCatalog Provisioning Details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#service_catalog_provisioning_details SagemakerProject#service_catalog_provisioning_details}
	ServiceCatalogProvisioningDetails *SagemakerProjectServiceCatalogProvisioningDetails `field:"required" json:"serviceCatalogProvisioningDetails" yaml:"serviceCatalogProvisioningDetails"`
	// The description of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#project_description SagemakerProject#project_description}
	ProjectDescription *string `field:"optional" json:"projectDescription" yaml:"projectDescription"`
	// Provisioned ServiceCatalog  Details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#service_catalog_provisioned_product_details SagemakerProject#service_catalog_provisioned_product_details}
	ServiceCatalogProvisionedProductDetails *SagemakerProjectServiceCatalogProvisionedProductDetails `field:"optional" json:"serviceCatalogProvisionedProductDetails" yaml:"serviceCatalogProvisionedProductDetails"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_project#tags SagemakerProject#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

