package applicationinsightsapplication


type ApplicationinsightsApplicationCustomComponents struct {
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#component_name ApplicationinsightsApplication#component_name}
	ComponentName *string `field:"required" json:"componentName" yaml:"componentName"`
	// The list of resource ARNs that belong to the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#resource_list ApplicationinsightsApplication#resource_list}
	ResourceList *[]*string `field:"required" json:"resourceList" yaml:"resourceList"`
}

