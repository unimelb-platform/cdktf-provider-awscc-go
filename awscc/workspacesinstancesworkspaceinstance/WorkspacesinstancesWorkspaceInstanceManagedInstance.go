package workspacesinstancesworkspaceinstance


type WorkspacesinstancesWorkspaceInstanceManagedInstance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#block_device_mappings WorkspacesinstancesWorkspaceInstance#block_device_mappings}.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#cpu_options WorkspacesinstancesWorkspaceInstance#cpu_options}.
	CpuOptions *WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptions `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#credit_specification WorkspacesinstancesWorkspaceInstance#credit_specification}.
	CreditSpecification *WorkspacesinstancesWorkspaceInstanceManagedInstanceCreditSpecification `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#disable_api_stop WorkspacesinstancesWorkspaceInstance#disable_api_stop}.
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#ebs_optimized WorkspacesinstancesWorkspaceInstance#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#enclave_options WorkspacesinstancesWorkspaceInstance#enclave_options}.
	EnclaveOptions *WorkspacesinstancesWorkspaceInstanceManagedInstanceEnclaveOptions `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#hibernation_options WorkspacesinstancesWorkspaceInstance#hibernation_options}.
	HibernationOptions *WorkspacesinstancesWorkspaceInstanceManagedInstanceHibernationOptions `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#iam_instance_profile WorkspacesinstancesWorkspaceInstance#iam_instance_profile}.
	IamInstanceProfile *WorkspacesinstancesWorkspaceInstanceManagedInstanceIamInstanceProfile `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#image_id WorkspacesinstancesWorkspaceInstance#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#instance_type WorkspacesinstancesWorkspaceInstance#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#key_name WorkspacesinstancesWorkspaceInstance#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#maintenance_options WorkspacesinstancesWorkspaceInstance#maintenance_options}.
	MaintenanceOptions *WorkspacesinstancesWorkspaceInstanceManagedInstanceMaintenanceOptions `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#metadata_options WorkspacesinstancesWorkspaceInstance#metadata_options}.
	MetadataOptions *WorkspacesinstancesWorkspaceInstanceManagedInstanceMetadataOptions `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#monitoring WorkspacesinstancesWorkspaceInstance#monitoring}.
	Monitoring *WorkspacesinstancesWorkspaceInstanceManagedInstanceMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#network_interfaces WorkspacesinstancesWorkspaceInstance#network_interfaces}.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#network_performance_options WorkspacesinstancesWorkspaceInstance#network_performance_options}.
	NetworkPerformanceOptions *WorkspacesinstancesWorkspaceInstanceManagedInstanceNetworkPerformanceOptions `field:"optional" json:"networkPerformanceOptions" yaml:"networkPerformanceOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#placement WorkspacesinstancesWorkspaceInstance#placement}.
	Placement *WorkspacesinstancesWorkspaceInstanceManagedInstancePlacement `field:"optional" json:"placement" yaml:"placement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#private_dns_name_options WorkspacesinstancesWorkspaceInstance#private_dns_name_options}.
	PrivateDnsNameOptions *WorkspacesinstancesWorkspaceInstanceManagedInstancePrivateDnsNameOptions `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#tag_specifications WorkspacesinstancesWorkspaceInstance#tag_specifications}.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#user_data WorkspacesinstancesWorkspaceInstance#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

