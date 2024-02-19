package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateData struct {
	// The block device mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#block_device_mappings Ec2LaunchTemplate#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Specifies an instance's Capacity Reservation targeting option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#capacity_reservation_specification Ec2LaunchTemplate#capacity_reservation_specification}
	CapacityReservationSpecification *Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// specifies the CPU options for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#cpu_options Ec2LaunchTemplate#cpu_options}
	CpuOptions *Ec2LaunchTemplateLaunchTemplateDataCpuOptions `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// The user data to make available to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#credit_specification Ec2LaunchTemplate#credit_specification}
	CreditSpecification *Ec2LaunchTemplateLaunchTemplateDataCreditSpecification `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Indicates whether to enable the instance for stop protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#disable_api_stop Ec2LaunchTemplate#disable_api_stop}
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#disable_api_termination Ec2LaunchTemplate#disable_api_termination}
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Indicates whether the instance is optimized for Amazon EBS I/O.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ebs_optimized Ec2LaunchTemplate#ebs_optimized}
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// An elastic GPU to associate with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#elastic_gpu_specifications Ec2LaunchTemplate#elastic_gpu_specifications}
	ElasticGpuSpecifications interface{} `field:"optional" json:"elasticGpuSpecifications" yaml:"elasticGpuSpecifications"`
	// The elastic inference accelerator for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#elastic_inference_accelerators Ec2LaunchTemplate#elastic_inference_accelerators}
	ElasticInferenceAccelerators interface{} `field:"optional" json:"elasticInferenceAccelerators" yaml:"elasticInferenceAccelerators"`
	// Indicates whether the instance is enabled for AWS Nitro Enclaves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#enclave_options Ec2LaunchTemplate#enclave_options}
	EnclaveOptions *Ec2LaunchTemplateLaunchTemplateDataEnclaveOptions `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Specifies whether your instance is configured for hibernation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#hibernation_options Ec2LaunchTemplate#hibernation_options}
	HibernationOptions *Ec2LaunchTemplateLaunchTemplateDataHibernationOptions `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// Specifies an IAM instance profile, which is a container for an IAM role for your instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#iam_instance_profile Ec2LaunchTemplate#iam_instance_profile}
	IamInstanceProfile *Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfile `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// The ID of the AMI.
	//
	// Alternatively, you can specify a Systems Manager parameter, which will resolve to an AMI ID on launch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#image_id Ec2LaunchTemplate#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#instance_initiated_shutdown_behavior Ec2LaunchTemplate#instance_initiated_shutdown_behavior}
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// The market (purchasing) option for the instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#instance_market_options Ec2LaunchTemplate#instance_market_options}
	InstanceMarketOptions *Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// The attributes for the instance types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#instance_requirements Ec2LaunchTemplate#instance_requirements}
	InstanceRequirements *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#instance_type Ec2LaunchTemplate#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The ID of the kernel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#kernel_id Ec2LaunchTemplate#kernel_id}
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// The name of the EC2 key pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#key_name Ec2LaunchTemplate#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// The license configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#license_specifications Ec2LaunchTemplate#license_specifications}
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// The maintenance options of your instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#maintenance_options Ec2LaunchTemplate#maintenance_options}
	MaintenanceOptions *Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptions `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// The metadata options for the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#metadata_options Ec2LaunchTemplate#metadata_options}
	MetadataOptions *Ec2LaunchTemplateLaunchTemplateDataMetadataOptions `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Specifies whether detailed monitoring is enabled for an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#monitoring Ec2LaunchTemplate#monitoring}
	Monitoring *Ec2LaunchTemplateLaunchTemplateDataMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// If you specify a network interface, you must specify any security groups and subnets as part of the network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#network_interfaces Ec2LaunchTemplate#network_interfaces}
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// Specifies the placement of an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#placement Ec2LaunchTemplate#placement}
	Placement *Ec2LaunchTemplateLaunchTemplateDataPlacement `field:"optional" json:"placement" yaml:"placement"`
	// Describes the options for instance hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#private_dns_name_options Ec2LaunchTemplate#private_dns_name_options}
	PrivateDnsNameOptions *Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptions `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ram_disk_id Ec2LaunchTemplate#ram_disk_id}.
	RamDiskId *string `field:"optional" json:"ramDiskId" yaml:"ramDiskId"`
	// One or more security group IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#security_group_ids Ec2LaunchTemplate#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// One or more security group names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#security_groups Ec2LaunchTemplate#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The tags to apply to the resources that are created during instance launch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#tag_specifications Ec2LaunchTemplate#tag_specifications}
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The user data to make available to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#user_data Ec2LaunchTemplate#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

