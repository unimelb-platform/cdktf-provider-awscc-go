package dataawsccec2launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccec2launchtemplate/internal"
)

type DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference interface {
	cdktf.ComplexObject
	AcceleratorCount() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorManufacturers() *[]*string
	AcceleratorNames() *[]*string
	AcceleratorTotalMemoryMiB() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	AcceleratorTypes() *[]*string
	AllowedInstanceTypes() *[]*string
	BareMetal() *string
	BaselineEbsBandwidthMbps() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	BurstablePerformance() *string
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
	CpuManufacturers() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedInstanceTypes() *[]*string
	// Experimental.
	Fqn() *string
	InstanceGenerations() *[]*string
	InternalValue() *DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirements
	SetInternalValue(val *DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirements)
	LocalStorage() *string
	LocalStorageTypes() *[]*string
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64
	MemoryGiBPerVCpu() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuOutputReference
	MemoryMiB() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiBOutputReference
	NetworkBandwidthGbps() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbpsOutputReference
	NetworkInterfaceCount() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCountOutputReference
	OnDemandMaxPricePercentageOverLowestPrice() *float64
	RequireHibernateSupport() cdktf.IResolvable
	SpotMaxPricePercentageOverLowestPrice() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalLocalStorageGb() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGbOutputReference
	VCpuCount() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCountOutputReference
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

// The jsii proxy struct for DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference
type jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) AcceleratorCount() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCountOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) AcceleratorTotalMemoryMiB() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) InternalValue() *DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirements {
	var returns *DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirements
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) MemoryGiBPerVCpu() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuOutputReference
	_jsii_.Get(
		j,
		"memoryGiBPerVCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) MemoryMiB() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiBOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"memoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) NetworkBandwidthGbps() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbpsOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbpsOutputReference
	_jsii_.Get(
		j,
		"networkBandwidthGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) NetworkInterfaceCount() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) RequireHibernateSupport() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) TotalLocalStorageGb() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) VCpuCount() DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCountOutputReference {
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCountOutputReference
	_jsii_.Get(
		j,
		"vCpuCount",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccEc2LaunchTemplate.DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference_Override(d DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccEc2LaunchTemplate.DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference)SetInternalValue(val *DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirements) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

