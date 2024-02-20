package inspectorv2filter


type Inspectorv2FilterFilterCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#aws_account_id Inspectorv2Filter#aws_account_id}.
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#component_id Inspectorv2Filter#component_id}.
	ComponentId interface{} `field:"optional" json:"componentId" yaml:"componentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#component_type Inspectorv2Filter#component_type}.
	ComponentType interface{} `field:"optional" json:"componentType" yaml:"componentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ec_2_instance_image_id Inspectorv2Filter#ec_2_instance_image_id}.
	Ec2InstanceImageId interface{} `field:"optional" json:"ec2InstanceImageId" yaml:"ec2InstanceImageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ec_2_instance_subnet_id Inspectorv2Filter#ec_2_instance_subnet_id}.
	Ec2InstanceSubnetId interface{} `field:"optional" json:"ec2InstanceSubnetId" yaml:"ec2InstanceSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ec_2_instance_vpc_id Inspectorv2Filter#ec_2_instance_vpc_id}.
	Ec2InstanceVpcId interface{} `field:"optional" json:"ec2InstanceVpcId" yaml:"ec2InstanceVpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ecr_image_architecture Inspectorv2Filter#ecr_image_architecture}.
	EcrImageArchitecture interface{} `field:"optional" json:"ecrImageArchitecture" yaml:"ecrImageArchitecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ecr_image_hash Inspectorv2Filter#ecr_image_hash}.
	EcrImageHash interface{} `field:"optional" json:"ecrImageHash" yaml:"ecrImageHash"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ecr_image_pushed_at Inspectorv2Filter#ecr_image_pushed_at}.
	EcrImagePushedAt interface{} `field:"optional" json:"ecrImagePushedAt" yaml:"ecrImagePushedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ecr_image_registry Inspectorv2Filter#ecr_image_registry}.
	EcrImageRegistry interface{} `field:"optional" json:"ecrImageRegistry" yaml:"ecrImageRegistry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ecr_image_repository_name Inspectorv2Filter#ecr_image_repository_name}.
	EcrImageRepositoryName interface{} `field:"optional" json:"ecrImageRepositoryName" yaml:"ecrImageRepositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#ecr_image_tags Inspectorv2Filter#ecr_image_tags}.
	EcrImageTags interface{} `field:"optional" json:"ecrImageTags" yaml:"ecrImageTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#finding_arn Inspectorv2Filter#finding_arn}.
	FindingArn interface{} `field:"optional" json:"findingArn" yaml:"findingArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#finding_status Inspectorv2Filter#finding_status}.
	FindingStatus interface{} `field:"optional" json:"findingStatus" yaml:"findingStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#finding_type Inspectorv2Filter#finding_type}.
	FindingType interface{} `field:"optional" json:"findingType" yaml:"findingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#first_observed_at Inspectorv2Filter#first_observed_at}.
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#inspector_score Inspectorv2Filter#inspector_score}.
	InspectorScore interface{} `field:"optional" json:"inspectorScore" yaml:"inspectorScore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#last_observed_at Inspectorv2Filter#last_observed_at}.
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#network_protocol Inspectorv2Filter#network_protocol}.
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#port_range Inspectorv2Filter#port_range}.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#related_vulnerabilities Inspectorv2Filter#related_vulnerabilities}.
	RelatedVulnerabilities interface{} `field:"optional" json:"relatedVulnerabilities" yaml:"relatedVulnerabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#resource_id Inspectorv2Filter#resource_id}.
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#resource_tags Inspectorv2Filter#resource_tags}.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#resource_type Inspectorv2Filter#resource_type}.
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#severity Inspectorv2Filter#severity}.
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#title Inspectorv2Filter#title}.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#updated_at Inspectorv2Filter#updated_at}.
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#vendor_severity Inspectorv2Filter#vendor_severity}.
	VendorSeverity interface{} `field:"optional" json:"vendorSeverity" yaml:"vendorSeverity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#vulnerability_id Inspectorv2Filter#vulnerability_id}.
	VulnerabilityId interface{} `field:"optional" json:"vulnerabilityId" yaml:"vulnerabilityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#vulnerability_source Inspectorv2Filter#vulnerability_source}.
	VulnerabilitySource interface{} `field:"optional" json:"vulnerabilitySource" yaml:"vulnerabilitySource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#vulnerable_packages Inspectorv2Filter#vulnerable_packages}.
	VulnerablePackages interface{} `field:"optional" json:"vulnerablePackages" yaml:"vulnerablePackages"`
}

