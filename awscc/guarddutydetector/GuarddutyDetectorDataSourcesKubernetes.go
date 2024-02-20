package guarddutydetector


type GuarddutyDetectorDataSourcesKubernetes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_detector#audit_logs GuarddutyDetector#audit_logs}.
	AuditLogs *GuarddutyDetectorDataSourcesKubernetesAuditLogs `field:"required" json:"auditLogs" yaml:"auditLogs"`
}

