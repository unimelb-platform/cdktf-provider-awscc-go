package osispipeline


type OsisPipelineBufferOptions struct {
	// Whether persistent buffering should be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/osis_pipeline#persistent_buffer_enabled OsisPipeline#persistent_buffer_enabled}
	PersistentBufferEnabled interface{} `field:"required" json:"persistentBufferEnabled" yaml:"persistentBufferEnabled"`
}

