package cleanroomsmembership


type CleanroomsMembershipPaymentConfigurationMachineLearning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_membership#model_inference CleanroomsMembership#model_inference}.
	ModelInference *CleanroomsMembershipPaymentConfigurationMachineLearningModelInference `field:"optional" json:"modelInference" yaml:"modelInference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_membership#model_training CleanroomsMembership#model_training}.
	ModelTraining *CleanroomsMembershipPaymentConfigurationMachineLearningModelTraining `field:"optional" json:"modelTraining" yaml:"modelTraining"`
}

