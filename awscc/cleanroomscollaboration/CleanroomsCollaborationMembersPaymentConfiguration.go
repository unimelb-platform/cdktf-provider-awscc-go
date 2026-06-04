package cleanroomscollaboration


type CleanroomsCollaborationMembersPaymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#job_compute CleanroomsCollaboration#job_compute}.
	JobCompute *CleanroomsCollaborationMembersPaymentConfigurationJobCompute `field:"optional" json:"jobCompute" yaml:"jobCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#machine_learning CleanroomsCollaboration#machine_learning}.
	MachineLearning *CleanroomsCollaborationMembersPaymentConfigurationMachineLearning `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#query_compute CleanroomsCollaboration#query_compute}.
	QueryCompute *CleanroomsCollaborationMembersPaymentConfigurationQueryCompute `field:"optional" json:"queryCompute" yaml:"queryCompute"`
}

