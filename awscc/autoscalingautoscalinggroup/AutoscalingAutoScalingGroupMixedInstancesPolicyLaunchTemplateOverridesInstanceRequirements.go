package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#memory_mi_b AutoscalingAutoScalingGroup#memory_mi_b}.
	MemoryMiB *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryMiB `field:"required" json:"memoryMiB" yaml:"memoryMiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#v_cpu_count AutoscalingAutoScalingGroup#v_cpu_count}.
	VCpuCount *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsVCpuCount `field:"required" json:"vCpuCount" yaml:"vCpuCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#accelerator_count AutoscalingAutoScalingGroup#accelerator_count}.
	AcceleratorCount *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorCount `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#accelerator_manufacturers AutoscalingAutoScalingGroup#accelerator_manufacturers}.
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#accelerator_names AutoscalingAutoScalingGroup#accelerator_names}.
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#accelerator_total_memory_mi_b AutoscalingAutoScalingGroup#accelerator_total_memory_mi_b}.
	AcceleratorTotalMemoryMiB *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorTotalMemoryMiB `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#accelerator_types AutoscalingAutoScalingGroup#accelerator_types}.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#allowed_instance_types AutoscalingAutoScalingGroup#allowed_instance_types}.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#bare_metal AutoscalingAutoScalingGroup#bare_metal}.
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#baseline_ebs_bandwidth_mbps AutoscalingAutoScalingGroup#baseline_ebs_bandwidth_mbps}.
	BaselineEbsBandwidthMbps *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselineEbsBandwidthMbps `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#burstable_performance AutoscalingAutoScalingGroup#burstable_performance}.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#cpu_manufacturers AutoscalingAutoScalingGroup#cpu_manufacturers}.
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#excluded_instance_types AutoscalingAutoScalingGroup#excluded_instance_types}.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#instance_generations AutoscalingAutoScalingGroup#instance_generations}.
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#local_storage AutoscalingAutoScalingGroup#local_storage}.
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#local_storage_types AutoscalingAutoScalingGroup#local_storage_types}.
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#max_spot_price_as_percentage_of_optimal_on_demand_price AutoscalingAutoScalingGroup#max_spot_price_as_percentage_of_optimal_on_demand_price}.
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#memory_gi_b_per_v_cpu AutoscalingAutoScalingGroup#memory_gi_b_per_v_cpu}.
	MemoryGiBPerVCpu *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryGiBPerVCpu `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#network_bandwidth_gbps AutoscalingAutoScalingGroup#network_bandwidth_gbps}.
	NetworkBandwidthGbps *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkBandwidthGbps `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#network_interface_count AutoscalingAutoScalingGroup#network_interface_count}.
	NetworkInterfaceCount *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkInterfaceCount `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#on_demand_max_price_percentage_over_lowest_price AutoscalingAutoScalingGroup#on_demand_max_price_percentage_over_lowest_price}.
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#require_hibernate_support AutoscalingAutoScalingGroup#require_hibernate_support}.
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#spot_max_price_percentage_over_lowest_price AutoscalingAutoScalingGroup#spot_max_price_percentage_over_lowest_price}.
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/autoscaling_auto_scaling_group#total_local_storage_gb AutoscalingAutoScalingGroup#total_local_storage_gb}.
	TotalLocalStorageGb *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsTotalLocalStorageGb `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
}

