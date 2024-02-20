package sagemakerfeaturegroup


type SagemakerFeatureGroupOnlineStoreConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#enable_online_store SagemakerFeatureGroup#enable_online_store}.
	EnableOnlineStore interface{} `field:"optional" json:"enableOnlineStore" yaml:"enableOnlineStore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#security_config SagemakerFeatureGroup#security_config}.
	SecurityConfig *SagemakerFeatureGroupOnlineStoreConfigSecurityConfig `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#storage_type SagemakerFeatureGroup#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
}

