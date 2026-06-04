package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy struct {
	// Firewall deployment mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/fms_policy#firewall_deployment_model FmsPolicy#firewall_deployment_model}
	FirewallDeploymentModel *string `field:"optional" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

