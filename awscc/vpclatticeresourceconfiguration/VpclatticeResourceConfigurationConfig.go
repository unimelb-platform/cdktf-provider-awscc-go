package vpclatticeresourceconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpclatticeResourceConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#name VpclatticeResourceConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#resource_configuration_type VpclatticeResourceConfiguration#resource_configuration_type}.
	ResourceConfigurationType *string `field:"required" json:"resourceConfigurationType" yaml:"resourceConfigurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#allow_association_to_sharable_service_network VpclatticeResourceConfiguration#allow_association_to_sharable_service_network}.
	AllowAssociationToSharableServiceNetwork interface{} `field:"optional" json:"allowAssociationToSharableServiceNetwork" yaml:"allowAssociationToSharableServiceNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#port_ranges VpclatticeResourceConfiguration#port_ranges}.
	PortRanges *[]*string `field:"optional" json:"portRanges" yaml:"portRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#protocol_type VpclatticeResourceConfiguration#protocol_type}.
	ProtocolType *string `field:"optional" json:"protocolType" yaml:"protocolType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#resource_configuration_auth_type VpclatticeResourceConfiguration#resource_configuration_auth_type}.
	ResourceConfigurationAuthType *string `field:"optional" json:"resourceConfigurationAuthType" yaml:"resourceConfigurationAuthType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#resource_configuration_definition VpclatticeResourceConfiguration#resource_configuration_definition}.
	ResourceConfigurationDefinition *VpclatticeResourceConfigurationResourceConfigurationDefinition `field:"optional" json:"resourceConfigurationDefinition" yaml:"resourceConfigurationDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#resource_configuration_group_id VpclatticeResourceConfiguration#resource_configuration_group_id}.
	ResourceConfigurationGroupId *string `field:"optional" json:"resourceConfigurationGroupId" yaml:"resourceConfigurationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#resource_gateway_id VpclatticeResourceConfiguration#resource_gateway_id}.
	ResourceGatewayId *string `field:"optional" json:"resourceGatewayId" yaml:"resourceGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_resource_configuration#tags VpclatticeResourceConfiguration#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

