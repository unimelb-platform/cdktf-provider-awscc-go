package inspectorv2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/inspectorv2filter/internal"
)

type Inspectorv2FilterFilterCriteriaOutputReference interface {
	cdktf.ComplexObject
	AwsAccountId() Inspectorv2FilterFilterCriteriaAwsAccountIdList
	AwsAccountIdInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComponentId() Inspectorv2FilterFilterCriteriaComponentIdList
	ComponentIdInput() interface{}
	ComponentType() Inspectorv2FilterFilterCriteriaComponentTypeList
	ComponentTypeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Ec2InstanceImageId() Inspectorv2FilterFilterCriteriaEc2InstanceImageIdList
	Ec2InstanceImageIdInput() interface{}
	Ec2InstanceSubnetId() Inspectorv2FilterFilterCriteriaEc2InstanceSubnetIdList
	Ec2InstanceSubnetIdInput() interface{}
	Ec2InstanceVpcId() Inspectorv2FilterFilterCriteriaEc2InstanceVpcIdList
	Ec2InstanceVpcIdInput() interface{}
	EcrImageArchitecture() Inspectorv2FilterFilterCriteriaEcrImageArchitectureList
	EcrImageArchitectureInput() interface{}
	EcrImageHash() Inspectorv2FilterFilterCriteriaEcrImageHashList
	EcrImageHashInput() interface{}
	EcrImagePushedAt() Inspectorv2FilterFilterCriteriaEcrImagePushedAtList
	EcrImagePushedAtInput() interface{}
	EcrImageRegistry() Inspectorv2FilterFilterCriteriaEcrImageRegistryList
	EcrImageRegistryInput() interface{}
	EcrImageRepositoryName() Inspectorv2FilterFilterCriteriaEcrImageRepositoryNameList
	EcrImageRepositoryNameInput() interface{}
	EcrImageTags() Inspectorv2FilterFilterCriteriaEcrImageTagsList
	EcrImageTagsInput() interface{}
	FindingArn() Inspectorv2FilterFilterCriteriaFindingArnList
	FindingArnInput() interface{}
	FindingStatus() Inspectorv2FilterFilterCriteriaFindingStatusList
	FindingStatusInput() interface{}
	FindingType() Inspectorv2FilterFilterCriteriaFindingTypeList
	FindingTypeInput() interface{}
	FirstObservedAt() Inspectorv2FilterFilterCriteriaFirstObservedAtList
	FirstObservedAtInput() interface{}
	// Experimental.
	Fqn() *string
	InspectorScore() Inspectorv2FilterFilterCriteriaInspectorScoreList
	InspectorScoreInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastObservedAt() Inspectorv2FilterFilterCriteriaLastObservedAtList
	LastObservedAtInput() interface{}
	NetworkProtocol() Inspectorv2FilterFilterCriteriaNetworkProtocolList
	NetworkProtocolInput() interface{}
	PortRange() Inspectorv2FilterFilterCriteriaPortRangeList
	PortRangeInput() interface{}
	RelatedVulnerabilities() Inspectorv2FilterFilterCriteriaRelatedVulnerabilitiesList
	RelatedVulnerabilitiesInput() interface{}
	ResourceId() Inspectorv2FilterFilterCriteriaResourceIdList
	ResourceIdInput() interface{}
	ResourceTags() Inspectorv2FilterFilterCriteriaResourceTagsList
	ResourceTagsInput() interface{}
	ResourceType() Inspectorv2FilterFilterCriteriaResourceTypeList
	ResourceTypeInput() interface{}
	Severity() Inspectorv2FilterFilterCriteriaSeverityList
	SeverityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() Inspectorv2FilterFilterCriteriaTitleList
	TitleInput() interface{}
	UpdatedAt() Inspectorv2FilterFilterCriteriaUpdatedAtList
	UpdatedAtInput() interface{}
	VendorSeverity() Inspectorv2FilterFilterCriteriaVendorSeverityList
	VendorSeverityInput() interface{}
	VulnerabilityId() Inspectorv2FilterFilterCriteriaVulnerabilityIdList
	VulnerabilityIdInput() interface{}
	VulnerabilitySource() Inspectorv2FilterFilterCriteriaVulnerabilitySourceList
	VulnerabilitySourceInput() interface{}
	VulnerablePackages() Inspectorv2FilterFilterCriteriaVulnerablePackagesList
	VulnerablePackagesInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAwsAccountId(value interface{})
	PutComponentId(value interface{})
	PutComponentType(value interface{})
	PutEc2InstanceImageId(value interface{})
	PutEc2InstanceSubnetId(value interface{})
	PutEc2InstanceVpcId(value interface{})
	PutEcrImageArchitecture(value interface{})
	PutEcrImageHash(value interface{})
	PutEcrImagePushedAt(value interface{})
	PutEcrImageRegistry(value interface{})
	PutEcrImageRepositoryName(value interface{})
	PutEcrImageTags(value interface{})
	PutFindingArn(value interface{})
	PutFindingStatus(value interface{})
	PutFindingType(value interface{})
	PutFirstObservedAt(value interface{})
	PutInspectorScore(value interface{})
	PutLastObservedAt(value interface{})
	PutNetworkProtocol(value interface{})
	PutPortRange(value interface{})
	PutRelatedVulnerabilities(value interface{})
	PutResourceId(value interface{})
	PutResourceTags(value interface{})
	PutResourceType(value interface{})
	PutSeverity(value interface{})
	PutTitle(value interface{})
	PutUpdatedAt(value interface{})
	PutVendorSeverity(value interface{})
	PutVulnerabilityId(value interface{})
	PutVulnerabilitySource(value interface{})
	PutVulnerablePackages(value interface{})
	ResetAwsAccountId()
	ResetComponentId()
	ResetComponentType()
	ResetEc2InstanceImageId()
	ResetEc2InstanceSubnetId()
	ResetEc2InstanceVpcId()
	ResetEcrImageArchitecture()
	ResetEcrImageHash()
	ResetEcrImagePushedAt()
	ResetEcrImageRegistry()
	ResetEcrImageRepositoryName()
	ResetEcrImageTags()
	ResetFindingArn()
	ResetFindingStatus()
	ResetFindingType()
	ResetFirstObservedAt()
	ResetInspectorScore()
	ResetLastObservedAt()
	ResetNetworkProtocol()
	ResetPortRange()
	ResetRelatedVulnerabilities()
	ResetResourceId()
	ResetResourceTags()
	ResetResourceType()
	ResetSeverity()
	ResetTitle()
	ResetUpdatedAt()
	ResetVendorSeverity()
	ResetVulnerabilityId()
	ResetVulnerabilitySource()
	ResetVulnerablePackages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Inspectorv2FilterFilterCriteriaOutputReference
type jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) AwsAccountId() Inspectorv2FilterFilterCriteriaAwsAccountIdList {
	var returns Inspectorv2FilterFilterCriteriaAwsAccountIdList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) AwsAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ComponentId() Inspectorv2FilterFilterCriteriaComponentIdList {
	var returns Inspectorv2FilterFilterCriteriaComponentIdList
	_jsii_.Get(
		j,
		"componentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ComponentIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ComponentType() Inspectorv2FilterFilterCriteriaComponentTypeList {
	var returns Inspectorv2FilterFilterCriteriaComponentTypeList
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ComponentTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"componentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceImageId() Inspectorv2FilterFilterCriteriaEc2InstanceImageIdList {
	var returns Inspectorv2FilterFilterCriteriaEc2InstanceImageIdList
	_jsii_.Get(
		j,
		"ec2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceImageIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceSubnetId() Inspectorv2FilterFilterCriteriaEc2InstanceSubnetIdList {
	var returns Inspectorv2FilterFilterCriteriaEc2InstanceSubnetIdList
	_jsii_.Get(
		j,
		"ec2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceSubnetIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceVpcId() Inspectorv2FilterFilterCriteriaEc2InstanceVpcIdList {
	var returns Inspectorv2FilterFilterCriteriaEc2InstanceVpcIdList
	_jsii_.Get(
		j,
		"ec2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceVpcIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2InstanceVpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageArchitecture() Inspectorv2FilterFilterCriteriaEcrImageArchitectureList {
	var returns Inspectorv2FilterFilterCriteriaEcrImageArchitectureList
	_jsii_.Get(
		j,
		"ecrImageArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageArchitectureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageArchitectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageHash() Inspectorv2FilterFilterCriteriaEcrImageHashList {
	var returns Inspectorv2FilterFilterCriteriaEcrImageHashList
	_jsii_.Get(
		j,
		"ecrImageHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageHashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImagePushedAt() Inspectorv2FilterFilterCriteriaEcrImagePushedAtList {
	var returns Inspectorv2FilterFilterCriteriaEcrImagePushedAtList
	_jsii_.Get(
		j,
		"ecrImagePushedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImagePushedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImagePushedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageRegistry() Inspectorv2FilterFilterCriteriaEcrImageRegistryList {
	var returns Inspectorv2FilterFilterCriteriaEcrImageRegistryList
	_jsii_.Get(
		j,
		"ecrImageRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageRegistryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRegistryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageRepositoryName() Inspectorv2FilterFilterCriteriaEcrImageRepositoryNameList {
	var returns Inspectorv2FilterFilterCriteriaEcrImageRepositoryNameList
	_jsii_.Get(
		j,
		"ecrImageRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageRepositoryNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageRepositoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageTags() Inspectorv2FilterFilterCriteriaEcrImageTagsList {
	var returns Inspectorv2FilterFilterCriteriaEcrImageTagsList
	_jsii_.Get(
		j,
		"ecrImageTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) EcrImageTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrImageTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FindingArn() Inspectorv2FilterFilterCriteriaFindingArnList {
	var returns Inspectorv2FilterFilterCriteriaFindingArnList
	_jsii_.Get(
		j,
		"findingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FindingArnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FindingStatus() Inspectorv2FilterFilterCriteriaFindingStatusList {
	var returns Inspectorv2FilterFilterCriteriaFindingStatusList
	_jsii_.Get(
		j,
		"findingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FindingStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FindingType() Inspectorv2FilterFilterCriteriaFindingTypeList {
	var returns Inspectorv2FilterFilterCriteriaFindingTypeList
	_jsii_.Get(
		j,
		"findingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FindingTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"findingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FirstObservedAt() Inspectorv2FilterFilterCriteriaFirstObservedAtList {
	var returns Inspectorv2FilterFilterCriteriaFirstObservedAtList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) FirstObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) InspectorScore() Inspectorv2FilterFilterCriteriaInspectorScoreList {
	var returns Inspectorv2FilterFilterCriteriaInspectorScoreList
	_jsii_.Get(
		j,
		"inspectorScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) InspectorScoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectorScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) LastObservedAt() Inspectorv2FilterFilterCriteriaLastObservedAtList {
	var returns Inspectorv2FilterFilterCriteriaLastObservedAtList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) LastObservedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastObservedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) NetworkProtocol() Inspectorv2FilterFilterCriteriaNetworkProtocolList {
	var returns Inspectorv2FilterFilterCriteriaNetworkProtocolList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) NetworkProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PortRange() Inspectorv2FilterFilterCriteriaPortRangeList {
	var returns Inspectorv2FilterFilterCriteriaPortRangeList
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PortRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"portRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) RelatedVulnerabilities() Inspectorv2FilterFilterCriteriaRelatedVulnerabilitiesList {
	var returns Inspectorv2FilterFilterCriteriaRelatedVulnerabilitiesList
	_jsii_.Get(
		j,
		"relatedVulnerabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) RelatedVulnerabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedVulnerabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResourceId() Inspectorv2FilterFilterCriteriaResourceIdList {
	var returns Inspectorv2FilterFilterCriteriaResourceIdList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResourceIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResourceTags() Inspectorv2FilterFilterCriteriaResourceTagsList {
	var returns Inspectorv2FilterFilterCriteriaResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResourceType() Inspectorv2FilterFilterCriteriaResourceTypeList {
	var returns Inspectorv2FilterFilterCriteriaResourceTypeList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Severity() Inspectorv2FilterFilterCriteriaSeverityList {
	var returns Inspectorv2FilterFilterCriteriaSeverityList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Title() Inspectorv2FilterFilterCriteriaTitleList {
	var returns Inspectorv2FilterFilterCriteriaTitleList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) TitleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) UpdatedAt() Inspectorv2FilterFilterCriteriaUpdatedAtList {
	var returns Inspectorv2FilterFilterCriteriaUpdatedAtList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) UpdatedAtInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updatedAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VendorSeverity() Inspectorv2FilterFilterCriteriaVendorSeverityList {
	var returns Inspectorv2FilterFilterCriteriaVendorSeverityList
	_jsii_.Get(
		j,
		"vendorSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VendorSeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vendorSeverityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VulnerabilityId() Inspectorv2FilterFilterCriteriaVulnerabilityIdList {
	var returns Inspectorv2FilterFilterCriteriaVulnerabilityIdList
	_jsii_.Get(
		j,
		"vulnerabilityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VulnerabilityIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VulnerabilitySource() Inspectorv2FilterFilterCriteriaVulnerabilitySourceList {
	var returns Inspectorv2FilterFilterCriteriaVulnerabilitySourceList
	_jsii_.Get(
		j,
		"vulnerabilitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VulnerabilitySourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilitySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VulnerablePackages() Inspectorv2FilterFilterCriteriaVulnerablePackagesList {
	var returns Inspectorv2FilterFilterCriteriaVulnerablePackagesList
	_jsii_.Get(
		j,
		"vulnerablePackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) VulnerablePackagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerablePackagesInput",
		&returns,
	)
	return returns
}


func NewInspectorv2FilterFilterCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Inspectorv2FilterFilterCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewInspectorv2FilterFilterCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference{}

	_jsii_.Create(
		"awscc.inspectorv2Filter.Inspectorv2FilterFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInspectorv2FilterFilterCriteriaOutputReference_Override(i Inspectorv2FilterFilterCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.inspectorv2Filter.Inspectorv2FilterFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutAwsAccountId(value interface{}) {
	if err := i.validatePutAwsAccountIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAwsAccountId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutComponentId(value interface{}) {
	if err := i.validatePutComponentIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putComponentId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutComponentType(value interface{}) {
	if err := i.validatePutComponentTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putComponentType",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEc2InstanceImageId(value interface{}) {
	if err := i.validatePutEc2InstanceImageIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEc2InstanceImageId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEc2InstanceSubnetId(value interface{}) {
	if err := i.validatePutEc2InstanceSubnetIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEc2InstanceSubnetId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEc2InstanceVpcId(value interface{}) {
	if err := i.validatePutEc2InstanceVpcIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEc2InstanceVpcId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEcrImageArchitecture(value interface{}) {
	if err := i.validatePutEcrImageArchitectureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageArchitecture",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEcrImageHash(value interface{}) {
	if err := i.validatePutEcrImageHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageHash",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEcrImagePushedAt(value interface{}) {
	if err := i.validatePutEcrImagePushedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImagePushedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEcrImageRegistry(value interface{}) {
	if err := i.validatePutEcrImageRegistryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageRegistry",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEcrImageRepositoryName(value interface{}) {
	if err := i.validatePutEcrImageRepositoryNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageRepositoryName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutEcrImageTags(value interface{}) {
	if err := i.validatePutEcrImageTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEcrImageTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutFindingArn(value interface{}) {
	if err := i.validatePutFindingArnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFindingArn",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutFindingStatus(value interface{}) {
	if err := i.validatePutFindingStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFindingStatus",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutFindingType(value interface{}) {
	if err := i.validatePutFindingTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFindingType",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutFirstObservedAt(value interface{}) {
	if err := i.validatePutFirstObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFirstObservedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutInspectorScore(value interface{}) {
	if err := i.validatePutInspectorScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putInspectorScore",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutLastObservedAt(value interface{}) {
	if err := i.validatePutLastObservedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLastObservedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutNetworkProtocol(value interface{}) {
	if err := i.validatePutNetworkProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNetworkProtocol",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutPortRange(value interface{}) {
	if err := i.validatePutPortRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPortRange",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutRelatedVulnerabilities(value interface{}) {
	if err := i.validatePutRelatedVulnerabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRelatedVulnerabilities",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutResourceId(value interface{}) {
	if err := i.validatePutResourceIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourceId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutResourceTags(value interface{}) {
	if err := i.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutResourceType(value interface{}) {
	if err := i.validatePutResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResourceType",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutSeverity(value interface{}) {
	if err := i.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSeverity",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutTitle(value interface{}) {
	if err := i.validatePutTitleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTitle",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutUpdatedAt(value interface{}) {
	if err := i.validatePutUpdatedAtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUpdatedAt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutVendorSeverity(value interface{}) {
	if err := i.validatePutVendorSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVendorSeverity",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutVulnerabilityId(value interface{}) {
	if err := i.validatePutVulnerabilityIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVulnerabilityId",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutVulnerabilitySource(value interface{}) {
	if err := i.validatePutVulnerabilitySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVulnerabilitySource",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) PutVulnerablePackages(value interface{}) {
	if err := i.validatePutVulnerablePackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putVulnerablePackages",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		i,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetComponentId() {
	_jsii_.InvokeVoid(
		i,
		"resetComponentId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetComponentType() {
	_jsii_.InvokeVoid(
		i,
		"resetComponentType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEc2InstanceImageId() {
	_jsii_.InvokeVoid(
		i,
		"resetEc2InstanceImageId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEc2InstanceSubnetId() {
	_jsii_.InvokeVoid(
		i,
		"resetEc2InstanceSubnetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEc2InstanceVpcId() {
	_jsii_.InvokeVoid(
		i,
		"resetEc2InstanceVpcId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEcrImageArchitecture() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageArchitecture",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEcrImageHash() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageHash",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEcrImagePushedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImagePushedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEcrImageRegistry() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageRegistry",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEcrImageRepositoryName() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageRepositoryName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetEcrImageTags() {
	_jsii_.InvokeVoid(
		i,
		"resetEcrImageTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetFindingArn() {
	_jsii_.InvokeVoid(
		i,
		"resetFindingArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetFindingStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetFindingStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetFindingType() {
	_jsii_.InvokeVoid(
		i,
		"resetFindingType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetFirstObservedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetFirstObservedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetInspectorScore() {
	_jsii_.InvokeVoid(
		i,
		"resetInspectorScore",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetLastObservedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetLastObservedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetNetworkProtocol() {
	_jsii_.InvokeVoid(
		i,
		"resetNetworkProtocol",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetPortRange() {
	_jsii_.InvokeVoid(
		i,
		"resetPortRange",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetRelatedVulnerabilities() {
	_jsii_.InvokeVoid(
		i,
		"resetRelatedVulnerabilities",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		i,
		"resetSeverity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		i,
		"resetTitle",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetUpdatedAt() {
	_jsii_.InvokeVoid(
		i,
		"resetUpdatedAt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetVendorSeverity() {
	_jsii_.InvokeVoid(
		i,
		"resetVendorSeverity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetVulnerabilityId() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerabilityId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetVulnerabilitySource() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerabilitySource",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ResetVulnerablePackages() {
	_jsii_.InvokeVoid(
		i,
		"resetVulnerablePackages",
		nil, // no parameters
	)
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

