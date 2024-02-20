package auditmanagerassessment


type AuditmanagerAssessmentScope struct {
	// The AWS accounts included in scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#aws_accounts AuditmanagerAssessment#aws_accounts}
	AwsAccounts interface{} `field:"optional" json:"awsAccounts" yaml:"awsAccounts"`
	// The AWS services included in scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/auditmanager_assessment#aws_services AuditmanagerAssessment#aws_services}
	AwsServices interface{} `field:"optional" json:"awsServices" yaml:"awsServices"`
}

