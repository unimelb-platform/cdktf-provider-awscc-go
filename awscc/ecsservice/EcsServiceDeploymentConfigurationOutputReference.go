package ecsservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ecsservice/internal"
)

type EcsServiceDeploymentConfigurationOutputReference interface {
	cdktf.ComplexObject
	Alarms() EcsServiceDeploymentConfigurationAlarmsOutputReference
	AlarmsInput() interface{}
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
	DeploymentCircuitBreaker() EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference
	DeploymentCircuitBreakerInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaximumPercent() *float64
	SetMaximumPercent(val *float64)
	MaximumPercentInput() *float64
	MinimumHealthyPercent() *float64
	SetMinimumHealthyPercent(val *float64)
	MinimumHealthyPercentInput() *float64
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
	PutAlarms(value *EcsServiceDeploymentConfigurationAlarms)
	PutDeploymentCircuitBreaker(value *EcsServiceDeploymentConfigurationDeploymentCircuitBreaker)
	ResetAlarms()
	ResetDeploymentCircuitBreaker()
	ResetMaximumPercent()
	ResetMinimumHealthyPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsServiceDeploymentConfigurationOutputReference
type jsiiProxy_EcsServiceDeploymentConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) Alarms() EcsServiceDeploymentConfigurationAlarmsOutputReference {
	var returns EcsServiceDeploymentConfigurationAlarmsOutputReference
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) AlarmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) DeploymentCircuitBreaker() EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference {
	var returns EcsServiceDeploymentConfigurationDeploymentCircuitBreakerOutputReference
	_jsii_.Get(
		j,
		"deploymentCircuitBreaker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) DeploymentCircuitBreakerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentCircuitBreakerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MaximumPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MaximumPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MinimumHealthyPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) MinimumHealthyPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumHealthyPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsServiceDeploymentConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcsServiceDeploymentConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEcsServiceDeploymentConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsServiceDeploymentConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.ecsService.EcsServiceDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsServiceDeploymentConfigurationOutputReference_Override(e EcsServiceDeploymentConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ecsService.EcsServiceDeploymentConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetMaximumPercent(val *float64) {
	if err := j.validateSetMaximumPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumPercent",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetMinimumHealthyPercent(val *float64) {
	if err := j.validateSetMinimumHealthyPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumHealthyPercent",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) PutAlarms(value *EcsServiceDeploymentConfigurationAlarms) {
	if err := e.validatePutAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAlarms",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) PutDeploymentCircuitBreaker(value *EcsServiceDeploymentConfigurationDeploymentCircuitBreaker) {
	if err := e.validatePutDeploymentCircuitBreakerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeploymentCircuitBreaker",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetAlarms() {
	_jsii_.InvokeVoid(
		e,
		"resetAlarms",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetDeploymentCircuitBreaker() {
	_jsii_.InvokeVoid(
		e,
		"resetDeploymentCircuitBreaker",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetMaximumPercent() {
	_jsii_.InvokeVoid(
		e,
		"resetMaximumPercent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ResetMinimumHealthyPercent() {
	_jsii_.InvokeVoid(
		e,
		"resetMinimumHealthyPercent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsServiceDeploymentConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

