package apprunnerservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApprunnerServiceConfig struct {
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
	// Source Code configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#source_configuration ApprunnerService#source_configuration}
	SourceConfiguration *ApprunnerServiceSourceConfiguration `field:"required" json:"sourceConfiguration" yaml:"sourceConfiguration"`
	// Autoscaling configuration ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#auto_scaling_configuration_arn ApprunnerService#auto_scaling_configuration_arn}
	AutoScalingConfigurationArn *string `field:"optional" json:"autoScalingConfigurationArn" yaml:"autoScalingConfigurationArn"`
	// Encryption configuration (KMS key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#encryption_configuration ApprunnerService#encryption_configuration}
	EncryptionConfiguration *ApprunnerServiceEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Health check configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#health_check_configuration ApprunnerService#health_check_configuration}
	HealthCheckConfiguration *ApprunnerServiceHealthCheckConfiguration `field:"optional" json:"healthCheckConfiguration" yaml:"healthCheckConfiguration"`
	// Instance Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#instance_configuration ApprunnerService#instance_configuration}
	InstanceConfiguration *ApprunnerServiceInstanceConfiguration `field:"optional" json:"instanceConfiguration" yaml:"instanceConfiguration"`
	// Network configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#network_configuration ApprunnerService#network_configuration}
	NetworkConfiguration *ApprunnerServiceNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Service observability configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#observability_configuration ApprunnerService#observability_configuration}
	ObservabilityConfiguration *ApprunnerServiceObservabilityConfiguration `field:"optional" json:"observabilityConfiguration" yaml:"observabilityConfiguration"`
	// The AppRunner Service Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#service_name ApprunnerService#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#tags ApprunnerService#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

