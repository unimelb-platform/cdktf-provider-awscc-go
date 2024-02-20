package cleanroomsanalysistemplate


type CleanroomsAnalysisTemplateAnalysisParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_analysis_template#name CleanroomsAnalysisTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_analysis_template#type CleanroomsAnalysisTemplate#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_analysis_template#default_value CleanroomsAnalysisTemplate#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
}

