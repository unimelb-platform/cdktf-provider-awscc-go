package inspectorassessmenttemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InspectorAssessmentTemplateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspector_assessment_template#assessment_target_arn InspectorAssessmentTemplate#assessment_target_arn}.
	AssessmentTargetArn *string `field:"required" json:"assessmentTargetArn" yaml:"assessmentTargetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspector_assessment_template#duration_in_seconds InspectorAssessmentTemplate#duration_in_seconds}.
	DurationInSeconds *float64 `field:"required" json:"durationInSeconds" yaml:"durationInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspector_assessment_template#rules_package_arns InspectorAssessmentTemplate#rules_package_arns}.
	RulesPackageArns *[]*string `field:"required" json:"rulesPackageArns" yaml:"rulesPackageArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspector_assessment_template#assessment_template_name InspectorAssessmentTemplate#assessment_template_name}.
	AssessmentTemplateName *string `field:"optional" json:"assessmentTemplateName" yaml:"assessmentTemplateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspector_assessment_template#user_attributes_for_findings InspectorAssessmentTemplate#user_attributes_for_findings}.
	UserAttributesForFindings interface{} `field:"optional" json:"userAttributesForFindings" yaml:"userAttributesForFindings"`
}

