package lookoutequipmentinferencescheduler


type LookoutequipmentInferenceSchedulerTags struct {
	// The key for the specified tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutequipment_inference_scheduler#key LookoutequipmentInferenceScheduler#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the specified tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutequipment_inference_scheduler#value LookoutequipmentInferenceScheduler#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

