package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy struct {
	// Firewall deployment mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy#firewall_deployment_model FmsPolicy#firewall_deployment_model}
	FirewallDeploymentModel *string `field:"required" json:"firewallDeploymentModel" yaml:"firewallDeploymentModel"`
}

