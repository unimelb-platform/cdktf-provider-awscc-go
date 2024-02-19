package codedeploydeploymentconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/codedeploydeploymentconfig/internal"
)

type CodedeployDeploymentConfigZonalConfigOutputReference interface {
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
	FirstZoneMonitorDurationInSeconds() *float64
	SetFirstZoneMonitorDurationInSeconds(val *float64)
	FirstZoneMonitorDurationInSecondsInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumHealthyHostsPerZone() CodedeployDeploymentConfigZonalConfigMinimumHealthyHostsPerZoneOutputReference
	MinimumHealthyHostsPerZoneInput() interface{}
	MonitorDurationInSeconds() *float64
	SetMonitorDurationInSeconds(val *float64)
	MonitorDurationInSecondsInput() *float64
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
	PutMinimumHealthyHostsPerZone(value *CodedeployDeploymentConfigZonalConfigMinimumHealthyHostsPerZone)
	ResetFirstZoneMonitorDurationInSeconds()
	ResetMinimumHealthyHostsPerZone()
	ResetMonitorDurationInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodedeployDeploymentConfigZonalConfigOutputReference
type jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) FirstZoneMonitorDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstZoneMonitorDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) FirstZoneMonitorDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstZoneMonitorDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) MinimumHealthyHostsPerZone() CodedeployDeploymentConfigZonalConfigMinimumHealthyHostsPerZoneOutputReference {
	var returns CodedeployDeploymentConfigZonalConfigMinimumHealthyHostsPerZoneOutputReference
	_jsii_.Get(
		j,
		"minimumHealthyHostsPerZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) MinimumHealthyHostsPerZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"minimumHealthyHostsPerZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) MonitorDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) MonitorDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCodedeployDeploymentConfigZonalConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodedeployDeploymentConfigZonalConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCodedeployDeploymentConfigZonalConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference{}

	_jsii_.Create(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfigZonalConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodedeployDeploymentConfigZonalConfigOutputReference_Override(c CodedeployDeploymentConfigZonalConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.codedeployDeploymentConfig.CodedeployDeploymentConfigZonalConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference)SetFirstZoneMonitorDurationInSeconds(val *float64) {
	if err := j.validateSetFirstZoneMonitorDurationInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstZoneMonitorDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference)SetMonitorDurationInSeconds(val *float64) {
	if err := j.validateSetMonitorDurationInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) PutMinimumHealthyHostsPerZone(value *CodedeployDeploymentConfigZonalConfigMinimumHealthyHostsPerZone) {
	if err := c.validatePutMinimumHealthyHostsPerZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMinimumHealthyHostsPerZone",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) ResetFirstZoneMonitorDurationInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetFirstZoneMonitorDurationInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) ResetMinimumHealthyHostsPerZone() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumHealthyHostsPerZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) ResetMonitorDurationInSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetMonitorDurationInSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodedeployDeploymentConfigZonalConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

