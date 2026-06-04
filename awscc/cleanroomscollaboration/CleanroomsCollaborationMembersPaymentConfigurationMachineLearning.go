package cleanroomscollaboration


type CleanroomsCollaborationMembersPaymentConfigurationMachineLearning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#model_inference CleanroomsCollaboration#model_inference}.
	ModelInference *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelInference `field:"optional" json:"modelInference" yaml:"modelInference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_collaboration#model_training CleanroomsCollaboration#model_training}.
	ModelTraining *CleanroomsCollaborationMembersPaymentConfigurationMachineLearningModelTraining `field:"optional" json:"modelTraining" yaml:"modelTraining"`
}

