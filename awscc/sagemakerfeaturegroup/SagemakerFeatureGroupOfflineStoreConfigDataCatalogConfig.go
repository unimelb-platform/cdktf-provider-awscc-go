package sagemakerfeaturegroup


type SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#catalog SagemakerFeatureGroup#catalog}.
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#database SagemakerFeatureGroup#database}.
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#table_name SagemakerFeatureGroup#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

