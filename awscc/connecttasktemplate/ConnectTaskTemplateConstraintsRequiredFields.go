package connecttasktemplate


type ConnectTaskTemplateConstraintsRequiredFields struct {
	// the identifier (name) for the task template field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_task_template#id ConnectTaskTemplate#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *ConnectTaskTemplateConstraintsRequiredFieldsId `field:"required" json:"id" yaml:"id"`
}

