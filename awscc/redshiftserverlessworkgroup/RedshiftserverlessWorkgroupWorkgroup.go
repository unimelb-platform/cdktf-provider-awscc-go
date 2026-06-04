package redshiftserverlessworkgroup


type RedshiftserverlessWorkgroupWorkgroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/redshiftserverless_workgroup#config_parameters RedshiftserverlessWorkgroup#config_parameters}.
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/redshiftserverless_workgroup#endpoint RedshiftserverlessWorkgroup#endpoint}.
	Endpoint *RedshiftserverlessWorkgroupWorkgroupEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/redshiftserverless_workgroup#price_performance_target RedshiftserverlessWorkgroup#price_performance_target}.
	PricePerformanceTarget *RedshiftserverlessWorkgroupWorkgroupPricePerformanceTarget `field:"optional" json:"pricePerformanceTarget" yaml:"pricePerformanceTarget"`
}

