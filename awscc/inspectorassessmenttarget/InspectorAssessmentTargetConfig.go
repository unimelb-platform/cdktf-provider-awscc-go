package inspectorassessmenttarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InspectorAssessmentTargetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspector_assessment_target#assessment_target_name InspectorAssessmentTarget#assessment_target_name}.
	AssessmentTargetName *string `field:"optional" json:"assessmentTargetName" yaml:"assessmentTargetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspector_assessment_target#resource_group_arn InspectorAssessmentTarget#resource_group_arn}.
	ResourceGroupArn *string `field:"optional" json:"resourceGroupArn" yaml:"resourceGroupArn"`
}

