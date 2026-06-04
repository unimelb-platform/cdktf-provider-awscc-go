package sesvdmattributes


type SesVdmAttributesDashboardAttributes struct {
	// Whether emails sent from this account have engagement tracking enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_vdm_attributes#engagement_metrics SesVdmAttributes#engagement_metrics}
	EngagementMetrics *string `field:"optional" json:"engagementMetrics" yaml:"engagementMetrics"`
}

