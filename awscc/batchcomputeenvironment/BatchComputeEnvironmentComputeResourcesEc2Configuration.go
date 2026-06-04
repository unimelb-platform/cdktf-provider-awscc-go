package batchcomputeenvironment


type BatchComputeEnvironmentComputeResourcesEc2Configuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_compute_environment#image_id_override BatchComputeEnvironment#image_id_override}.
	ImageIdOverride *string `field:"optional" json:"imageIdOverride" yaml:"imageIdOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_compute_environment#image_kubernetes_version BatchComputeEnvironment#image_kubernetes_version}.
	ImageKubernetesVersion *string `field:"optional" json:"imageKubernetesVersion" yaml:"imageKubernetesVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_compute_environment#image_type BatchComputeEnvironment#image_type}.
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
}

