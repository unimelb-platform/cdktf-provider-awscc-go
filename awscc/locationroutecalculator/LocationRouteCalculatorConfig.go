package locationroutecalculator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LocationRouteCalculatorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_route_calculator#calculator_name LocationRouteCalculator#calculator_name}.
	CalculatorName *string `field:"required" json:"calculatorName" yaml:"calculatorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_route_calculator#data_source LocationRouteCalculator#data_source}.
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_route_calculator#description LocationRouteCalculator#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_route_calculator#pricing_plan LocationRouteCalculator#pricing_plan}.
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_route_calculator#tags LocationRouteCalculator#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

