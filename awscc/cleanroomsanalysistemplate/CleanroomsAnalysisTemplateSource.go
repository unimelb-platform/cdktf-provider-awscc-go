package cleanroomsanalysistemplate


type CleanroomsAnalysisTemplateSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_analysis_template#artifacts CleanroomsAnalysisTemplate#artifacts}.
	Artifacts *CleanroomsAnalysisTemplateSourceArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_analysis_template#text CleanroomsAnalysisTemplate#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

