package auditmanagerassessment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerAssessmentConfig struct {
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
	// The destination in which evidence reports are stored for the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#assessment_reports_destination AuditmanagerAssessment#assessment_reports_destination}
	AssessmentReportsDestination *AuditmanagerAssessmentAssessmentReportsDestination `field:"optional" json:"assessmentReportsDestination" yaml:"assessmentReportsDestination"`
	// The AWS account associated with the assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#aws_account AuditmanagerAssessment#aws_account}
	AwsAccount *AuditmanagerAssessmentAwsAccount `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// The list of delegations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#delegations AuditmanagerAssessment#delegations}
	Delegations interface{} `field:"optional" json:"delegations" yaml:"delegations"`
	// The description of the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#description AuditmanagerAssessment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier for the specified framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#framework_id AuditmanagerAssessment#framework_id}
	FrameworkId *string `field:"optional" json:"frameworkId" yaml:"frameworkId"`
	// The name of the related assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#name AuditmanagerAssessment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of roles for the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#roles AuditmanagerAssessment#roles}
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// The wrapper that contains the AWS accounts and AWS services in scope for the assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#scope AuditmanagerAssessment#scope}
	Scope *AuditmanagerAssessmentScope `field:"optional" json:"scope" yaml:"scope"`
	// The status of the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#status AuditmanagerAssessment#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags associated with the assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#tags AuditmanagerAssessment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

