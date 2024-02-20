package emrserverlessapplication


type EmrserverlessApplicationInitialCapacity struct {
	// Worker type for an analytics framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#key EmrserverlessApplication#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#value EmrserverlessApplication#value}.
	Value *EmrserverlessApplicationInitialCapacityValue `field:"required" json:"value" yaml:"value"`
}

