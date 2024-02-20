package ec2instance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2InstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#additional_info Ec2Instance#additional_info}.
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#affinity Ec2Instance#affinity}.
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#availability_zone Ec2Instance#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#block_device_mappings Ec2Instance#block_device_mappings}.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#cpu_options Ec2Instance#cpu_options}.
	CpuOptions *Ec2InstanceCpuOptions `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#credit_specification Ec2Instance#credit_specification}.
	CreditSpecification *Ec2InstanceCreditSpecification `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#disable_api_termination Ec2Instance#disable_api_termination}.
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#ebs_optimized Ec2Instance#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#elastic_gpu_specifications Ec2Instance#elastic_gpu_specifications}.
	ElasticGpuSpecifications interface{} `field:"optional" json:"elasticGpuSpecifications" yaml:"elasticGpuSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#elastic_inference_accelerators Ec2Instance#elastic_inference_accelerators}.
	ElasticInferenceAccelerators interface{} `field:"optional" json:"elasticInferenceAccelerators" yaml:"elasticInferenceAccelerators"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#enclave_options Ec2Instance#enclave_options}.
	EnclaveOptions *Ec2InstanceEnclaveOptions `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#hibernation_options Ec2Instance#hibernation_options}.
	HibernationOptions *Ec2InstanceHibernationOptions `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#host_id Ec2Instance#host_id}.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#host_resource_group_arn Ec2Instance#host_resource_group_arn}.
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#iam_instance_profile Ec2Instance#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#image_id Ec2Instance#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#instance_initiated_shutdown_behavior Ec2Instance#instance_initiated_shutdown_behavior}.
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#instance_type Ec2Instance#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#ipv_6_address_count Ec2Instance#ipv_6_address_count}.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#ipv_6_addresses Ec2Instance#ipv_6_addresses}.
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#kernel_id Ec2Instance#kernel_id}.
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#key_name Ec2Instance#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#launch_template Ec2Instance#launch_template}.
	LaunchTemplate *Ec2InstanceLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#license_specifications Ec2Instance#license_specifications}.
	LicenseSpecifications interface{} `field:"optional" json:"licenseSpecifications" yaml:"licenseSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#monitoring Ec2Instance#monitoring}.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#network_interfaces Ec2Instance#network_interfaces}.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#placement_group_name Ec2Instance#placement_group_name}.
	PlacementGroupName *string `field:"optional" json:"placementGroupName" yaml:"placementGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#private_dns_name_options Ec2Instance#private_dns_name_options}.
	PrivateDnsNameOptions *Ec2InstancePrivateDnsNameOptions `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#private_ip_address Ec2Instance#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#propagate_tags_to_volume_on_creation Ec2Instance#propagate_tags_to_volume_on_creation}.
	PropagateTagsToVolumeOnCreation interface{} `field:"optional" json:"propagateTagsToVolumeOnCreation" yaml:"propagateTagsToVolumeOnCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#ramdisk_id Ec2Instance#ramdisk_id}.
	RamdiskId *string `field:"optional" json:"ramdiskId" yaml:"ramdiskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#security_group_ids Ec2Instance#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#security_groups Ec2Instance#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#source_dest_check Ec2Instance#source_dest_check}.
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#ssm_associations Ec2Instance#ssm_associations}.
	SsmAssociations interface{} `field:"optional" json:"ssmAssociations" yaml:"ssmAssociations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#subnet_id Ec2Instance#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#tags Ec2Instance#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#tenancy Ec2Instance#tenancy}.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#user_data Ec2Instance#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#volumes Ec2Instance#volumes}.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

