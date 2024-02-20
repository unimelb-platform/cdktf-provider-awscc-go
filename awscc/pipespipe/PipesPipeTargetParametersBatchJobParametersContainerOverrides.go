package pipespipe


type PipesPipeTargetParametersBatchJobParametersContainerOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#command PipesPipe#command}.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#environment PipesPipe#environment}.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#instance_type PipesPipe#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#resource_requirements PipesPipe#resource_requirements}.
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
}

