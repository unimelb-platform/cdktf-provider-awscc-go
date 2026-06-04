package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowResolutionTechniquesRuleBasedProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#attribute_matching_model EntityresolutionMatchingWorkflow#attribute_matching_model}.
	AttributeMatchingModel *string `field:"optional" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#match_purpose EntityresolutionMatchingWorkflow#match_purpose}.
	MatchPurpose *string `field:"optional" json:"matchPurpose" yaml:"matchPurpose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#rules EntityresolutionMatchingWorkflow#rules}.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

