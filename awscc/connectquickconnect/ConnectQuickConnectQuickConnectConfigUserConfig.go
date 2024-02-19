package connectquickconnect


type ConnectQuickConnectQuickConnectConfigUserConfig struct {
	// The identifier of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#contact_flow_arn ConnectQuickConnect#contact_flow_arn}
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The identifier of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#user_arn ConnectQuickConnect#user_arn}
	UserArn *string `field:"required" json:"userArn" yaml:"userArn"`
}

