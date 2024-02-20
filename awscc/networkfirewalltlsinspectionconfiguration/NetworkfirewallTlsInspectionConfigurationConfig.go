package networkfirewalltlsinspectionconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkfirewallTlsInspectionConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_tls_inspection_configuration#tls_inspection_configuration NetworkfirewallTlsInspectionConfiguration#tls_inspection_configuration}.
	TlsInspectionConfiguration *NetworkfirewallTlsInspectionConfigurationTlsInspectionConfiguration `field:"required" json:"tlsInspectionConfiguration" yaml:"tlsInspectionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_tls_inspection_configuration#tls_inspection_configuration_name NetworkfirewallTlsInspectionConfiguration#tls_inspection_configuration_name}.
	TlsInspectionConfigurationName *string `field:"required" json:"tlsInspectionConfigurationName" yaml:"tlsInspectionConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_tls_inspection_configuration#description NetworkfirewallTlsInspectionConfiguration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_tls_inspection_configuration#tags NetworkfirewallTlsInspectionConfiguration#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

