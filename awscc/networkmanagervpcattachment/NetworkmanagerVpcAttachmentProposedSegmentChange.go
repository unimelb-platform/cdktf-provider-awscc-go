package networkmanagervpcattachment


type NetworkmanagerVpcAttachmentProposedSegmentChange struct {
	// The rule number in the policy document that applies to this change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#attachment_policy_rule_number NetworkmanagerVpcAttachment#attachment_policy_rule_number}
	AttachmentPolicyRuleNumber *float64 `field:"optional" json:"attachmentPolicyRuleNumber" yaml:"attachmentPolicyRuleNumber"`
	// The name of the segment to change.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#segment_name NetworkmanagerVpcAttachment#segment_name}
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// The key-value tags that changed for the segment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_vpc_attachment#tags NetworkmanagerVpcAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

