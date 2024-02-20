package greengrassv2deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/greengrassv2deployment/internal"
)

type Greengrassv2DeploymentComponentsRunWithOutputReference interface {
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
	PosixUser() *string
	SetPosixUser(val *string)
	PosixUserInput() *string
	SystemResourceLimits() Greengrassv2DeploymentComponentsRunWithSystemResourceLimitsOutputReference
	SystemResourceLimitsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsUser() *string
	SetWindowsUser(val *string)
	WindowsUserInput() *string
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
	PutSystemResourceLimits(value *Greengrassv2DeploymentComponentsRunWithSystemResourceLimits)
	ResetPosixUser()
	ResetSystemResourceLimits()
	ResetWindowsUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Greengrassv2DeploymentComponentsRunWithOutputReference
type jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) PosixUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) PosixUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"posixUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) SystemResourceLimits() Greengrassv2DeploymentComponentsRunWithSystemResourceLimitsOutputReference {
	var returns Greengrassv2DeploymentComponentsRunWithSystemResourceLimitsOutputReference
	_jsii_.Get(
		j,
		"systemResourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) SystemResourceLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemResourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) WindowsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) WindowsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsUserInput",
		&returns,
	)
	return returns
}


func NewGreengrassv2DeploymentComponentsRunWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Greengrassv2DeploymentComponentsRunWithOutputReference {
	_init_.Initialize()

	if err := validateNewGreengrassv2DeploymentComponentsRunWithOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference{}

	_jsii_.Create(
		"awscc.greengrassv2Deployment.Greengrassv2DeploymentComponentsRunWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGreengrassv2DeploymentComponentsRunWithOutputReference_Override(g Greengrassv2DeploymentComponentsRunWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.greengrassv2Deployment.Greengrassv2DeploymentComponentsRunWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference)SetPosixUser(val *string) {
	if err := j.validateSetPosixUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"posixUser",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference)SetWindowsUser(val *string) {
	if err := j.validateSetWindowsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsUser",
		val,
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) PutSystemResourceLimits(value *Greengrassv2DeploymentComponentsRunWithSystemResourceLimits) {
	if err := g.validatePutSystemResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSystemResourceLimits",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) ResetPosixUser() {
	_jsii_.InvokeVoid(
		g,
		"resetPosixUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) ResetSystemResourceLimits() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemResourceLimits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) ResetWindowsUser() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsUser",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentComponentsRunWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

