package ssmguiconnectpreferences


type SsmguiconnectPreferencesIdleConnection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmguiconnect_preferences#alert SsmguiconnectPreferences#alert}.
	Alert *SsmguiconnectPreferencesIdleConnectionAlert `field:"optional" json:"alert" yaml:"alert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmguiconnect_preferences#timeout SsmguiconnectPreferences#timeout}.
	Timeout *SsmguiconnectPreferencesIdleConnectionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

