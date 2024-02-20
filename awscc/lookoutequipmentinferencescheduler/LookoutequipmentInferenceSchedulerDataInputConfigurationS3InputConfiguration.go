package lookoutequipmentinferencescheduler


type LookoutequipmentInferenceSchedulerDataInputConfigurationS3InputConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutequipment_inference_scheduler#bucket LookoutequipmentInferenceScheduler#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutequipment_inference_scheduler#prefix LookoutequipmentInferenceScheduler#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

