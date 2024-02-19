package cleanroomsmembership


type CleanroomsMembershipPaymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_membership#query_compute CleanroomsMembership#query_compute}.
	QueryCompute *CleanroomsMembershipPaymentConfigurationQueryCompute `field:"required" json:"queryCompute" yaml:"queryCompute"`
}

