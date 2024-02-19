package resiliencehubresiliencypolicy


type ResiliencehubResiliencyPolicyPolicy struct {
	// RPO in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#rpo_in_secs ResiliencehubResiliencyPolicy#rpo_in_secs}
	RpoInSecs *float64 `field:"optional" json:"rpoInSecs" yaml:"rpoInSecs"`
	// RTO in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#rto_in_secs ResiliencehubResiliencyPolicy#rto_in_secs}
	RtoInSecs *float64 `field:"optional" json:"rtoInSecs" yaml:"rtoInSecs"`
}

