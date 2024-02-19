package sagemakerfeaturegroup


type SagemakerFeatureGroupFeatureDefinitions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#feature_name SagemakerFeatureGroup#feature_name}.
	FeatureName *string `field:"required" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_feature_group#feature_type SagemakerFeatureGroup#feature_type}.
	FeatureType *string `field:"required" json:"featureType" yaml:"featureType"`
}

