package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements struct {
	// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferential chips) on an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#accelerator_count Ec2LaunchTemplate#accelerator_count}
	AcceleratorCount *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCount `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Indicates whether instance types must have accelerators by specific manufacturers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#accelerator_manufacturers Ec2LaunchTemplate#accelerator_manufacturers}
	AcceleratorManufacturers *[]*string `field:"optional" json:"acceleratorManufacturers" yaml:"acceleratorManufacturers"`
	// The accelerators that must be on the instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#accelerator_names Ec2LaunchTemplate#accelerator_names}
	AcceleratorNames *[]*string `field:"optional" json:"acceleratorNames" yaml:"acceleratorNames"`
	// The minimum and maximum amount of total accelerator memory, in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#accelerator_total_memory_mi_b Ec2LaunchTemplate#accelerator_total_memory_mi_b}
	AcceleratorTotalMemoryMiB *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiB `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// The accelerator types that must be on the instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#accelerator_types Ec2LaunchTemplate#accelerator_types}
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// The instance types to apply your specified attributes against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#allowed_instance_types Ec2LaunchTemplate#allowed_instance_types}
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Indicates whether bare metal instance types must be included, excluded, or required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#bare_metal Ec2LaunchTemplate#bare_metal}
	BareMetal *string `field:"optional" json:"bareMetal" yaml:"bareMetal"`
	// The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#baseline_ebs_bandwidth_mbps Ec2LaunchTemplate#baseline_ebs_bandwidth_mbps}
	BaselineEbsBandwidthMbps *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbps `field:"optional" json:"baselineEbsBandwidthMbps" yaml:"baselineEbsBandwidthMbps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#burstable_performance Ec2LaunchTemplate#burstable_performance}.
	BurstablePerformance *string `field:"optional" json:"burstablePerformance" yaml:"burstablePerformance"`
	// The CPU manufacturers to include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#cpu_manufacturers Ec2LaunchTemplate#cpu_manufacturers}
	CpuManufacturers *[]*string `field:"optional" json:"cpuManufacturers" yaml:"cpuManufacturers"`
	// The instance types to exclude.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#excluded_instance_types Ec2LaunchTemplate#excluded_instance_types}
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Indicates whether current or previous generation instance types are included.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#instance_generations Ec2LaunchTemplate#instance_generations}
	InstanceGenerations *[]*string `field:"optional" json:"instanceGenerations" yaml:"instanceGenerations"`
	// The user data to make available to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#local_storage Ec2LaunchTemplate#local_storage}
	LocalStorage *string `field:"optional" json:"localStorage" yaml:"localStorage"`
	// The type of local storage that is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#local_storage_types Ec2LaunchTemplate#local_storage_types}
	LocalStorageTypes *[]*string `field:"optional" json:"localStorageTypes" yaml:"localStorageTypes"`
	// The price protection threshold for Spot Instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#max_spot_price_as_percentage_of_optimal_on_demand_price Ec2LaunchTemplate#max_spot_price_as_percentage_of_optimal_on_demand_price}
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice *float64 `field:"optional" json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice"`
	// The minimum and maximum amount of memory per vCPU, in GiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#memory_gi_b_per_v_cpu Ec2LaunchTemplate#memory_gi_b_per_v_cpu}
	MemoryGiBPerVCpu *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpu `field:"optional" json:"memoryGiBPerVCpu" yaml:"memoryGiBPerVCpu"`
	// The minimum and maximum amount of memory, in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#memory_mi_b Ec2LaunchTemplate#memory_mi_b}
	MemoryMiB *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiB `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#network_bandwidth_gbps Ec2LaunchTemplate#network_bandwidth_gbps}
	NetworkBandwidthGbps *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbps `field:"optional" json:"networkBandwidthGbps" yaml:"networkBandwidthGbps"`
	// TThe minimum and maximum number of network interfaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#network_interface_count Ec2LaunchTemplate#network_interface_count}
	NetworkInterfaceCount *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCount `field:"optional" json:"networkInterfaceCount" yaml:"networkInterfaceCount"`
	// The price protection threshold for On-Demand Instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#on_demand_max_price_percentage_over_lowest_price Ec2LaunchTemplate#on_demand_max_price_percentage_over_lowest_price}
	OnDemandMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"onDemandMaxPricePercentageOverLowestPrice" yaml:"onDemandMaxPricePercentageOverLowestPrice"`
	// Indicates whether instance types must support hibernation for On-Demand Instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#require_hibernate_support Ec2LaunchTemplate#require_hibernate_support}
	RequireHibernateSupport interface{} `field:"optional" json:"requireHibernateSupport" yaml:"requireHibernateSupport"`
	// The price protection threshold for Spot Instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#spot_max_price_percentage_over_lowest_price Ec2LaunchTemplate#spot_max_price_percentage_over_lowest_price}
	SpotMaxPricePercentageOverLowestPrice *float64 `field:"optional" json:"spotMaxPricePercentageOverLowestPrice" yaml:"spotMaxPricePercentageOverLowestPrice"`
	// The minimum and maximum amount of total local storage, in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#total_local_storage_gb Ec2LaunchTemplate#total_local_storage_gb}
	TotalLocalStorageGb *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGb `field:"optional" json:"totalLocalStorageGb" yaml:"totalLocalStorageGb"`
	// The minimum and maximum number of vCPUs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#v_cpu_count Ec2LaunchTemplate#v_cpu_count}
	VCpuCount *Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCount `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}

