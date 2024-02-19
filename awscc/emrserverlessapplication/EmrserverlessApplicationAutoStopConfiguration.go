package emrserverlessapplication


type EmrserverlessApplicationAutoStopConfiguration struct {
	// If set to true, the Application will automatically stop after being idle. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#enabled EmrserverlessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#idle_timeout_minutes EmrserverlessApplication#idle_timeout_minutes}
	IdleTimeoutMinutes *float64 `field:"optional" json:"idleTimeoutMinutes" yaml:"idleTimeoutMinutes"`
}

