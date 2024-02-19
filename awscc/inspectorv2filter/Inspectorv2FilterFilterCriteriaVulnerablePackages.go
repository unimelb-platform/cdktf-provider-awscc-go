package inspectorv2filter


type Inspectorv2FilterFilterCriteriaVulnerablePackages struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#architecture Inspectorv2Filter#architecture}.
	Architecture *Inspectorv2FilterFilterCriteriaVulnerablePackagesArchitecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#epoch Inspectorv2Filter#epoch}.
	Epoch *Inspectorv2FilterFilterCriteriaVulnerablePackagesEpoch `field:"optional" json:"epoch" yaml:"epoch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#name Inspectorv2Filter#name}.
	Name *Inspectorv2FilterFilterCriteriaVulnerablePackagesName `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#release Inspectorv2Filter#release}.
	Release *Inspectorv2FilterFilterCriteriaVulnerablePackagesRelease `field:"optional" json:"release" yaml:"release"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#source_layer_hash Inspectorv2Filter#source_layer_hash}.
	SourceLayerHash *Inspectorv2FilterFilterCriteriaVulnerablePackagesSourceLayerHash `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#version Inspectorv2Filter#version}.
	Version *Inspectorv2FilterFilterCriteriaVulnerablePackagesVersion `field:"optional" json:"version" yaml:"version"`
}

