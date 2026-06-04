package eventsconnection


type EventsConnectionAuthParametersInvocationHttpParametersBodyParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/events_connection#is_value_secret EventsConnection#is_value_secret}.
	IsValueSecret interface{} `field:"optional" json:"isValueSecret" yaml:"isValueSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/events_connection#key EventsConnection#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/events_connection#value EventsConnection#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

