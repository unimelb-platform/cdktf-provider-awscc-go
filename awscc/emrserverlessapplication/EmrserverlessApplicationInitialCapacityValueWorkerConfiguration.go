package emrserverlessapplication


type EmrserverlessApplicationInitialCapacityValueWorkerConfiguration struct {
	// Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#cpu EmrserverlessApplication#cpu}
	Cpu *string `field:"required" json:"cpu" yaml:"cpu"`
	// Per worker memory resource. GB is the only supported unit and specifying GB is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#memory EmrserverlessApplication#memory}
	Memory *string `field:"required" json:"memory" yaml:"memory"`
	// Per worker Disk resource. GB is the only supported unit and specifying GB is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#disk EmrserverlessApplication#disk}
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
}

