package apigatewaydeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/apigatewaydeployment/internal"
)

type ApigatewayDeploymentStageDescriptionCanarySettingOutputReference interface {
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
	PercentTraffic() *float64
	SetPercentTraffic(val *float64)
	PercentTrafficInput() *float64
	StageVariableOverrides() *map[string]*string
	SetStageVariableOverrides(val *map[string]*string)
	StageVariableOverridesInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseStageCache() interface{}
	SetUseStageCache(val interface{})
	UseStageCacheInput() interface{}
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
	ResetPercentTraffic()
	ResetStageVariableOverrides()
	ResetUseStageCache()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigatewayDeploymentStageDescriptionCanarySettingOutputReference
type jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) PercentTraffic() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) PercentTrafficInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) StageVariableOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariableOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) StageVariableOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariableOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) UseStageCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStageCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) UseStageCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStageCacheInput",
		&returns,
	)
	return returns
}


func NewApigatewayDeploymentStageDescriptionCanarySettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigatewayDeploymentStageDescriptionCanarySettingOutputReference {
	_init_.Initialize()

	if err := validateNewApigatewayDeploymentStageDescriptionCanarySettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference{}

	_jsii_.Create(
		"awscc.apigatewayDeployment.ApigatewayDeploymentStageDescriptionCanarySettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigatewayDeploymentStageDescriptionCanarySettingOutputReference_Override(a ApigatewayDeploymentStageDescriptionCanarySettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.apigatewayDeployment.ApigatewayDeploymentStageDescriptionCanarySettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetPercentTraffic(val *float64) {
	if err := j.validateSetPercentTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentTraffic",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetStageVariableOverrides(val *map[string]*string) {
	if err := j.validateSetStageVariableOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stageVariableOverrides",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference)SetUseStageCache(val interface{}) {
	if err := j.validateSetUseStageCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useStageCache",
		val,
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) ResetPercentTraffic() {
	_jsii_.InvokeVoid(
		a,
		"resetPercentTraffic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) ResetStageVariableOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetStageVariableOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) ResetUseStageCache() {
	_jsii_.InvokeVoid(
		a,
		"resetUseStageCache",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigatewayDeploymentStageDescriptionCanarySettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

