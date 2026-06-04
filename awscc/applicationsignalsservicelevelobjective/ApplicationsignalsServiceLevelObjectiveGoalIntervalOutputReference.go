package applicationsignalsservicelevelobjective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/applicationsignalsservicelevelobjective/internal"
)

type ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference interface {
	cdktf.ComplexObject
	CalendarInterval() ApplicationsignalsServiceLevelObjectiveGoalIntervalCalendarIntervalOutputReference
	CalendarIntervalInput() interface{}
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
	RollingInterval() ApplicationsignalsServiceLevelObjectiveGoalIntervalRollingIntervalOutputReference
	RollingIntervalInput() interface{}
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
	PutCalendarInterval(value *ApplicationsignalsServiceLevelObjectiveGoalIntervalCalendarInterval)
	PutRollingInterval(value *ApplicationsignalsServiceLevelObjectiveGoalIntervalRollingInterval)
	ResetCalendarInterval()
	ResetRollingInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference
type jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) CalendarInterval() ApplicationsignalsServiceLevelObjectiveGoalIntervalCalendarIntervalOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveGoalIntervalCalendarIntervalOutputReference
	_jsii_.Get(
		j,
		"calendarInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) CalendarIntervalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"calendarIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) RollingInterval() ApplicationsignalsServiceLevelObjectiveGoalIntervalRollingIntervalOutputReference {
	var returns ApplicationsignalsServiceLevelObjectiveGoalIntervalRollingIntervalOutputReference
	_jsii_.Get(
		j,
		"rollingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) RollingIntervalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference{}

	_jsii_.Create(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference_Override(a ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.applicationsignalsServiceLevelObjective.ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) PutCalendarInterval(value *ApplicationsignalsServiceLevelObjectiveGoalIntervalCalendarInterval) {
	if err := a.validatePutCalendarIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCalendarInterval",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) PutRollingInterval(value *ApplicationsignalsServiceLevelObjectiveGoalIntervalRollingInterval) {
	if err := a.validatePutRollingIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRollingInterval",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) ResetCalendarInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetCalendarInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) ResetRollingInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetRollingInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsServiceLevelObjectiveGoalIntervalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

