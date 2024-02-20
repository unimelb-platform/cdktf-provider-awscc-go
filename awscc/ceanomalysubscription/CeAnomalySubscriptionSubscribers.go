package ceanomalysubscription


type CeAnomalySubscriptionSubscribers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_anomaly_subscription#address CeAnomalySubscription#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_anomaly_subscription#type CeAnomalySubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ce_anomaly_subscription#status CeAnomalySubscription#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

