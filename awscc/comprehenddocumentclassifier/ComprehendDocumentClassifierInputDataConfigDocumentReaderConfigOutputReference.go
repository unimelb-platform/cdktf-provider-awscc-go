package comprehenddocumentclassifier

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/comprehenddocumentclassifier/internal"
)

type ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference interface {
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
	DocumentReadAction() *string
	SetDocumentReadAction(val *string)
	DocumentReadActionInput() *string
	DocumentReadMode() *string
	SetDocumentReadMode(val *string)
	DocumentReadModeInput() *string
	FeatureTypes() *[]*string
	SetFeatureTypes(val *[]*string)
	FeatureTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetDocumentReadMode()
	ResetFeatureTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference
type jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) DocumentReadAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentReadAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) DocumentReadActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentReadActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) DocumentReadMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentReadMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) DocumentReadModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentReadModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) FeatureTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"featureTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) FeatureTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"featureTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference{}

	_jsii_.Create(
		"awscc.comprehendDocumentClassifier.ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference_Override(c ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.comprehendDocumentClassifier.ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetDocumentReadAction(val *string) {
	if err := j.validateSetDocumentReadActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentReadAction",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetDocumentReadMode(val *string) {
	if err := j.validateSetDocumentReadModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentReadMode",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetFeatureTypes(val *[]*string) {
	if err := j.validateSetFeatureTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureTypes",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) ResetDocumentReadMode() {
	_jsii_.InvokeVoid(
		c,
		"resetDocumentReadMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) ResetFeatureTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetFeatureTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComprehendDocumentClassifierInputDataConfigDocumentReaderConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

