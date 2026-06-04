package sagemakerendpoint


type SagemakerEndpointExcludeRetainedVariantProperties struct {
	// The type of variant property (e.g., 'DesiredInstanceCount', 'DesiredWeight', 'DataCaptureConfig').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_endpoint#variant_property_type SagemakerEndpoint#variant_property_type}
	VariantPropertyType *string `field:"optional" json:"variantPropertyType" yaml:"variantPropertyType"`
}

