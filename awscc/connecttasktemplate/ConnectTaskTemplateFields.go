package connecttasktemplate


type ConnectTaskTemplateFields struct {
	// the identifier (name) for the task template field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_task_template#id ConnectTaskTemplate#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *ConnectTaskTemplateFieldsId `field:"required" json:"id" yaml:"id"`
	// The type of the task template's field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_task_template#type ConnectTaskTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the task template's field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_task_template#description ConnectTaskTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// list of field options to be used with single select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_task_template#single_select_options ConnectTaskTemplate#single_select_options}
	SingleSelectOptions *[]*string `field:"optional" json:"singleSelectOptions" yaml:"singleSelectOptions"`
}

