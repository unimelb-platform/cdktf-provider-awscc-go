package ec2spotfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2spotfleet/internal"
)

type Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	References() Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList
	ReferencesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutReferences(value interface{})
	ResetReferences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference
type jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) References() Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList {
	var returns Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ReferencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference {
	_init_.Initialize()

	if err := validateNewEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference{}

	_jsii_.Create(
		"awscc.ec2SpotFleet.Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference_Override(e Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2SpotFleet.Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) PutReferences(value interface{}) {
	if err := e.validatePutReferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putReferences",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ResetReferences() {
	_jsii_.InvokeVoid(
		e,
		"resetReferences",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

