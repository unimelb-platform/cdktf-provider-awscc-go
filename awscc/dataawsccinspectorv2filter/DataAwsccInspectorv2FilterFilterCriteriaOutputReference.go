package dataawsccinspectorv2filter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccinspectorv2filter/internal"
)

type DataAwsccInspectorv2FilterFilterCriteriaOutputReference interface {
	cdktf.ComplexObject
	AwsAccountId() DataAwsccInspectorv2FilterFilterCriteriaAwsAccountIdList
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
	ComponentId() DataAwsccInspectorv2FilterFilterCriteriaComponentIdList
	ComponentType() DataAwsccInspectorv2FilterFilterCriteriaComponentTypeList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Ec2InstanceImageId() DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceImageIdList
	Ec2InstanceSubnetId() DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceSubnetIdList
	Ec2InstanceVpcId() DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceVpcIdList
	EcrImageArchitecture() DataAwsccInspectorv2FilterFilterCriteriaEcrImageArchitectureList
	EcrImageHash() DataAwsccInspectorv2FilterFilterCriteriaEcrImageHashList
	EcrImagePushedAt() DataAwsccInspectorv2FilterFilterCriteriaEcrImagePushedAtList
	EcrImageRegistry() DataAwsccInspectorv2FilterFilterCriteriaEcrImageRegistryList
	EcrImageRepositoryName() DataAwsccInspectorv2FilterFilterCriteriaEcrImageRepositoryNameList
	EcrImageTags() DataAwsccInspectorv2FilterFilterCriteriaEcrImageTagsList
	FindingArn() DataAwsccInspectorv2FilterFilterCriteriaFindingArnList
	FindingStatus() DataAwsccInspectorv2FilterFilterCriteriaFindingStatusList
	FindingType() DataAwsccInspectorv2FilterFilterCriteriaFindingTypeList
	FirstObservedAt() DataAwsccInspectorv2FilterFilterCriteriaFirstObservedAtList
	// Experimental.
	Fqn() *string
	InspectorScore() DataAwsccInspectorv2FilterFilterCriteriaInspectorScoreList
	InternalValue() *DataAwsccInspectorv2FilterFilterCriteria
	SetInternalValue(val *DataAwsccInspectorv2FilterFilterCriteria)
	LastObservedAt() DataAwsccInspectorv2FilterFilterCriteriaLastObservedAtList
	NetworkProtocol() DataAwsccInspectorv2FilterFilterCriteriaNetworkProtocolList
	PortRange() DataAwsccInspectorv2FilterFilterCriteriaPortRangeList
	RelatedVulnerabilities() DataAwsccInspectorv2FilterFilterCriteriaRelatedVulnerabilitiesList
	ResourceId() DataAwsccInspectorv2FilterFilterCriteriaResourceIdList
	ResourceTags() DataAwsccInspectorv2FilterFilterCriteriaResourceTagsList
	ResourceType() DataAwsccInspectorv2FilterFilterCriteriaResourceTypeList
	Severity() DataAwsccInspectorv2FilterFilterCriteriaSeverityList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() DataAwsccInspectorv2FilterFilterCriteriaTitleList
	UpdatedAt() DataAwsccInspectorv2FilterFilterCriteriaUpdatedAtList
	VendorSeverity() DataAwsccInspectorv2FilterFilterCriteriaVendorSeverityList
	VulnerabilityId() DataAwsccInspectorv2FilterFilterCriteriaVulnerabilityIdList
	VulnerabilitySource() DataAwsccInspectorv2FilterFilterCriteriaVulnerabilitySourceList
	VulnerablePackages() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccInspectorv2FilterFilterCriteriaOutputReference
type jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) AwsAccountId() DataAwsccInspectorv2FilterFilterCriteriaAwsAccountIdList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaAwsAccountIdList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ComponentId() DataAwsccInspectorv2FilterFilterCriteriaComponentIdList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaComponentIdList
	_jsii_.Get(
		j,
		"componentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ComponentType() DataAwsccInspectorv2FilterFilterCriteriaComponentTypeList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaComponentTypeList
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceImageId() DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceImageIdList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceImageIdList
	_jsii_.Get(
		j,
		"ec2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceSubnetId() DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceSubnetIdList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceSubnetIdList
	_jsii_.Get(
		j,
		"ec2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) Ec2InstanceVpcId() DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceVpcIdList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEc2InstanceVpcIdList
	_jsii_.Get(
		j,
		"ec2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) EcrImageArchitecture() DataAwsccInspectorv2FilterFilterCriteriaEcrImageArchitectureList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEcrImageArchitectureList
	_jsii_.Get(
		j,
		"ecrImageArchitecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) EcrImageHash() DataAwsccInspectorv2FilterFilterCriteriaEcrImageHashList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEcrImageHashList
	_jsii_.Get(
		j,
		"ecrImageHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) EcrImagePushedAt() DataAwsccInspectorv2FilterFilterCriteriaEcrImagePushedAtList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEcrImagePushedAtList
	_jsii_.Get(
		j,
		"ecrImagePushedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) EcrImageRegistry() DataAwsccInspectorv2FilterFilterCriteriaEcrImageRegistryList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEcrImageRegistryList
	_jsii_.Get(
		j,
		"ecrImageRegistry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) EcrImageRepositoryName() DataAwsccInspectorv2FilterFilterCriteriaEcrImageRepositoryNameList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEcrImageRepositoryNameList
	_jsii_.Get(
		j,
		"ecrImageRepositoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) EcrImageTags() DataAwsccInspectorv2FilterFilterCriteriaEcrImageTagsList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaEcrImageTagsList
	_jsii_.Get(
		j,
		"ecrImageTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) FindingArn() DataAwsccInspectorv2FilterFilterCriteriaFindingArnList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaFindingArnList
	_jsii_.Get(
		j,
		"findingArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) FindingStatus() DataAwsccInspectorv2FilterFilterCriteriaFindingStatusList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaFindingStatusList
	_jsii_.Get(
		j,
		"findingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) FindingType() DataAwsccInspectorv2FilterFilterCriteriaFindingTypeList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaFindingTypeList
	_jsii_.Get(
		j,
		"findingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) FirstObservedAt() DataAwsccInspectorv2FilterFilterCriteriaFirstObservedAtList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaFirstObservedAtList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) InspectorScore() DataAwsccInspectorv2FilterFilterCriteriaInspectorScoreList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaInspectorScoreList
	_jsii_.Get(
		j,
		"inspectorScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) InternalValue() *DataAwsccInspectorv2FilterFilterCriteria {
	var returns *DataAwsccInspectorv2FilterFilterCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) LastObservedAt() DataAwsccInspectorv2FilterFilterCriteriaLastObservedAtList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaLastObservedAtList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) NetworkProtocol() DataAwsccInspectorv2FilterFilterCriteriaNetworkProtocolList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaNetworkProtocolList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) PortRange() DataAwsccInspectorv2FilterFilterCriteriaPortRangeList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaPortRangeList
	_jsii_.Get(
		j,
		"portRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) RelatedVulnerabilities() DataAwsccInspectorv2FilterFilterCriteriaRelatedVulnerabilitiesList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaRelatedVulnerabilitiesList
	_jsii_.Get(
		j,
		"relatedVulnerabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ResourceId() DataAwsccInspectorv2FilterFilterCriteriaResourceIdList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaResourceIdList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ResourceTags() DataAwsccInspectorv2FilterFilterCriteriaResourceTagsList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ResourceType() DataAwsccInspectorv2FilterFilterCriteriaResourceTypeList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaResourceTypeList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) Severity() DataAwsccInspectorv2FilterFilterCriteriaSeverityList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaSeverityList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) Title() DataAwsccInspectorv2FilterFilterCriteriaTitleList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaTitleList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) UpdatedAt() DataAwsccInspectorv2FilterFilterCriteriaUpdatedAtList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaUpdatedAtList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) VendorSeverity() DataAwsccInspectorv2FilterFilterCriteriaVendorSeverityList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVendorSeverityList
	_jsii_.Get(
		j,
		"vendorSeverity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) VulnerabilityId() DataAwsccInspectorv2FilterFilterCriteriaVulnerabilityIdList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerabilityIdList
	_jsii_.Get(
		j,
		"vulnerabilityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) VulnerabilitySource() DataAwsccInspectorv2FilterFilterCriteriaVulnerabilitySourceList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerabilitySourceList
	_jsii_.Get(
		j,
		"vulnerabilitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) VulnerablePackages() DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesList {
	var returns DataAwsccInspectorv2FilterFilterCriteriaVulnerablePackagesList
	_jsii_.Get(
		j,
		"vulnerablePackages",
		&returns,
	)
	return returns
}


func NewDataAwsccInspectorv2FilterFilterCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccInspectorv2FilterFilterCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccInspectorv2FilterFilterCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccInspectorv2Filter.DataAwsccInspectorv2FilterFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccInspectorv2FilterFilterCriteriaOutputReference_Override(d DataAwsccInspectorv2FilterFilterCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccInspectorv2Filter.DataAwsccInspectorv2FilterFilterCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference)SetInternalValue(val *DataAwsccInspectorv2FilterFilterCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInspectorv2FilterFilterCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

