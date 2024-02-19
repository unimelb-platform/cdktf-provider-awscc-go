package evidentlyfeature


type EvidentlyFeatureEntityOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_feature#entity_id EvidentlyFeature#entity_id}.
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_feature#variation EvidentlyFeature#variation}.
	Variation *string `field:"optional" json:"variation" yaml:"variation"`
}

