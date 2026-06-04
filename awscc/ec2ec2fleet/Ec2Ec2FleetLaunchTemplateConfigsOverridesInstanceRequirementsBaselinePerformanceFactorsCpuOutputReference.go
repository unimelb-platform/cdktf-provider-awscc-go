package ec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2ec2fleet/internal"
)

type Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference interface {
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
	References() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList
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

// The jsii proxy struct for Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference
type jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) References() Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList {
	var returns Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList
	_jsii_.Get(
		j,
		"references",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ReferencesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference {
	_init_.Initialize()

	if err := validateNewEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference{}

	_jsii_.Create(
		"awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference_Override(e Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) PutReferences(value interface{}) {
	if err := e.validatePutReferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putReferences",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ResetReferences() {
	_jsii_.InvokeVoid(
		e,
		"resetReferences",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

