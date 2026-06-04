package rumappmonitor


type RumAppMonitorDeobfuscationConfiguration struct {
	// A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#java_script_source_maps RumAppMonitor#java_script_source_maps}
	JavaScriptSourceMaps *RumAppMonitorDeobfuscationConfigurationJavaScriptSourceMaps `field:"optional" json:"javaScriptSourceMaps" yaml:"javaScriptSourceMaps"`
}

