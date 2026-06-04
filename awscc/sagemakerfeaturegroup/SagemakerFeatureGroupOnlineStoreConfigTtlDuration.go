package sagemakerfeaturegroup


type SagemakerFeatureGroupOnlineStoreConfigTtlDuration struct {
	// Unit of ttl configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_feature_group#unit SagemakerFeatureGroup#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// Value of ttl configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_feature_group#value SagemakerFeatureGroup#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

