package rumappmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RumAppMonitorConfig struct {
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
	// A name for the app monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#name RumAppMonitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// AppMonitor configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#app_monitor_configuration RumAppMonitor#app_monitor_configuration}
	AppMonitorConfiguration *RumAppMonitorAppMonitorConfiguration `field:"optional" json:"appMonitorConfiguration" yaml:"appMonitorConfiguration"`
	// AppMonitor custom events configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#custom_events RumAppMonitor#custom_events}
	CustomEvents *RumAppMonitorCustomEvents `field:"optional" json:"customEvents" yaml:"customEvents"`
	// Data collected by RUM is kept by RUM for 30 days and then deleted.
	//
	// This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#cw_log_enabled RumAppMonitor#cw_log_enabled}
	CwLogEnabled interface{} `field:"optional" json:"cwLogEnabled" yaml:"cwLogEnabled"`
	// A structure that contains the configuration for how an app monitor can deobfuscate stack traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#deobfuscation_configuration RumAppMonitor#deobfuscation_configuration}
	DeobfuscationConfiguration *RumAppMonitorDeobfuscationConfiguration `field:"optional" json:"deobfuscationConfiguration" yaml:"deobfuscationConfiguration"`
	// The top-level internet domain name for which your application has administrative authority.
	//
	// The CreateAppMonitor requires either the domain or the domain list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#domain RumAppMonitor#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The top-level internet domain names for which your application has administrative authority.
	//
	// The CreateAppMonitor requires either the domain or the domain list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#domain_list RumAppMonitor#domain_list}
	DomainList *[]*string `field:"optional" json:"domainList" yaml:"domainList"`
	// A structure that defines resource policy attached to your app monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#resource_policy RumAppMonitor#resource_policy}
	ResourcePolicy *RumAppMonitorResourcePolicy `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Assigns one or more tags (key-value pairs) to the app monitor.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/rum_app_monitor#tags RumAppMonitor#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

