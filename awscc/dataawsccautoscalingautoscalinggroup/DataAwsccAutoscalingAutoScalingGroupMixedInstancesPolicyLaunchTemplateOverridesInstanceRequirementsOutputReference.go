package dataawsccautoscalingautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccautoscalingautoscalinggroup/internal"
)

type DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference interface {
	cdktf.ComplexObject
	AcceleratorCount() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorCountOutputReference
	AcceleratorManufacturers() *[]*string
	AcceleratorNames() *[]*string
	AcceleratorTotalMemoryMiB() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	AcceleratorTypes() *[]*string
	AllowedInstanceTypes() *[]*string
	BareMetal() *string
	BaselineEbsBandwidthMbps() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
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
	InternalValue() *DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements
	SetInternalValue(val *DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements)
	LocalStorage() *string
	LocalStorageTypes() *[]*string
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64
	MemoryGiBPerVCpu() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryGiBPerVCpuOutputReference
	MemoryMiB() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryMiBOutputReference
	NetworkBandwidthGbps() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkBandwidthGbpsOutputReference
	NetworkInterfaceCount() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkInterfaceCountOutputReference
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
	TotalLocalStorageGb() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsTotalLocalStorageGbOutputReference
	VCpuCount() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsVCpuCountOutputReference
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

// The jsii proxy struct for DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference
type jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) AcceleratorCount() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorCountOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorCountOutputReference
	_jsii_.Get(
		j,
		"acceleratorCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) AcceleratorManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) AcceleratorNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) AcceleratorTotalMemoryMiB() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"acceleratorTotalMemoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) AcceleratorTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"acceleratorTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) AllowedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) BareMetal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) BaselineEbsBandwidthMbps() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference
	_jsii_.Get(
		j,
		"baselineEbsBandwidthMbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) BurstablePerformance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"burstablePerformance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) CpuManufacturers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cpuManufacturers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) ExcludedInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) InstanceGenerations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGenerations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) InternalValue() *DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements {
	var returns *DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) LocalStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) LocalStorageTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localStorageTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) MaxSpotPriceAsPercentageOfOptimalOnDemandPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSpotPriceAsPercentageOfOptimalOnDemandPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) MemoryGiBPerVCpu() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryGiBPerVCpuOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryGiBPerVCpuOutputReference
	_jsii_.Get(
		j,
		"memoryGiBPerVCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) MemoryMiB() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryMiBOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsMemoryMiBOutputReference
	_jsii_.Get(
		j,
		"memoryMiB",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) NetworkBandwidthGbps() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkBandwidthGbpsOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkBandwidthGbpsOutputReference
	_jsii_.Get(
		j,
		"networkBandwidthGbps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) NetworkInterfaceCount() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkInterfaceCountOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsNetworkInterfaceCountOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) OnDemandMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) RequireHibernateSupport() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"requireHibernateSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) SpotMaxPricePercentageOverLowestPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotMaxPricePercentageOverLowestPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) TotalLocalStorageGb() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsTotalLocalStorageGbOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsTotalLocalStorageGbOutputReference
	_jsii_.Get(
		j,
		"totalLocalStorageGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) VCpuCount() DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsVCpuCountOutputReference {
	var returns DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsVCpuCountOutputReference
	_jsii_.Get(
		j,
		"vCpuCount",
		&returns,
	)
	return returns
}


func NewDataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference_Override(d DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccAutoscalingAutoScalingGroup.DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference)SetInternalValue(val *DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccAutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

