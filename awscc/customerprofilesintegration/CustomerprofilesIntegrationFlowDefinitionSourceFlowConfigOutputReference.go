package customerprofilesintegration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/customerprofilesintegration/internal"
)

type CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference interface {
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
	ConnectorProfileName() *string
	SetConnectorProfileName(val *string)
	ConnectorProfileNameInput() *string
	ConnectorType() *string
	SetConnectorType(val *string)
	ConnectorTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IncrementalPullConfig() CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigIncrementalPullConfigOutputReference
	IncrementalPullConfigInput() interface{}
	InternalValue() *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig
	SetInternalValue(val *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig)
	SourceConnectorProperties() CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesOutputReference
	SourceConnectorPropertiesInput() *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties
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
	PutIncrementalPullConfig(value *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigIncrementalPullConfig)
	PutSourceConnectorProperties(value *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties)
	ResetConnectorProfileName()
	ResetIncrementalPullConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference
type jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ConnectorProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ConnectorProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorProfileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ConnectorTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) IncrementalPullConfig() CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigIncrementalPullConfigOutputReference {
	var returns CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigIncrementalPullConfigOutputReference
	_jsii_.Get(
		j,
		"incrementalPullConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) IncrementalPullConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incrementalPullConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) InternalValue() *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig {
	var returns *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) SourceConnectorProperties() CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesOutputReference {
	var returns CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesOutputReference
	_jsii_.Get(
		j,
		"sourceConnectorProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) SourceConnectorPropertiesInput() *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties {
	var returns *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties
	_jsii_.Get(
		j,
		"sourceConnectorPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference{}

	_jsii_.Create(
		"awscc.customerprofilesIntegration.CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference_Override(c CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.customerprofilesIntegration.CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference)SetConnectorProfileName(val *string) {
	if err := j.validateSetConnectorProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorProfileName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference)SetConnectorType(val *string) {
	if err := j.validateSetConnectorTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorType",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference)SetInternalValue(val *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) PutIncrementalPullConfig(value *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigIncrementalPullConfig) {
	if err := c.validatePutIncrementalPullConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIncrementalPullConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) PutSourceConnectorProperties(value *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties) {
	if err := c.validatePutSourceConnectorPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceConnectorProperties",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ResetConnectorProfileName() {
	_jsii_.InvokeVoid(
		c,
		"resetConnectorProfileName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ResetIncrementalPullConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetIncrementalPullConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

