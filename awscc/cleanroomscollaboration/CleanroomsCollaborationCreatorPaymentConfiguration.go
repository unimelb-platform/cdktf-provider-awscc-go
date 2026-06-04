package cleanroomscollaboration


type CleanroomsCollaborationCreatorPaymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#job_compute CleanroomsCollaboration#job_compute}.
	JobCompute *CleanroomsCollaborationCreatorPaymentConfigurationJobCompute `field:"optional" json:"jobCompute" yaml:"jobCompute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#machine_learning CleanroomsCollaboration#machine_learning}.
	MachineLearning *CleanroomsCollaborationCreatorPaymentConfigurationMachineLearning `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#query_compute CleanroomsCollaboration#query_compute}.
	QueryCompute *CleanroomsCollaborationCreatorPaymentConfigurationQueryCompute `field:"optional" json:"queryCompute" yaml:"queryCompute"`
}

