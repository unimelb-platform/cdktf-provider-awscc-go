package iotsitewisedashboard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotsitewiseDashboardConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The dashboard definition specified in a JSON literal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_dashboard#dashboard_definition IotsitewiseDashboard#dashboard_definition}
	DashboardDefinition *string `field:"required" json:"dashboardDefinition" yaml:"dashboardDefinition"`
	// A description for the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_dashboard#dashboard_description IotsitewiseDashboard#dashboard_description}
	DashboardDescription *string `field:"required" json:"dashboardDescription" yaml:"dashboardDescription"`
	// A friendly name for the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_dashboard#dashboard_name IotsitewiseDashboard#dashboard_name}
	DashboardName *string `field:"required" json:"dashboardName" yaml:"dashboardName"`
	// The ID of the project in which to create the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_dashboard#project_id IotsitewiseDashboard#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// A list of key-value pairs that contain metadata for the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_dashboard#tags IotsitewiseDashboard#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

