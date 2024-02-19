package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/appflowflow/internal"
)

type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference interface {
	cdktf.ComplexObject
	AggregationConfig() AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigOutputReference
	AggregationConfigInput() interface{}
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
	FileType() *string
	SetFileType(val *string)
	FileTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrefixConfig() AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigOutputReference
	PrefixConfigInput() interface{}
	PreserveSourceDataTyping() interface{}
	SetPreserveSourceDataTyping(val interface{})
	PreserveSourceDataTypingInput() interface{}
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
	PutAggregationConfig(value *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig)
	PutPrefixConfig(value *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfig)
	ResetAggregationConfig()
	ResetFileType()
	ResetPrefixConfig()
	ResetPreserveSourceDataTyping()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference
type jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) AggregationConfig() AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigOutputReference {
	var returns AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigOutputReference
	_jsii_.Get(
		j,
		"aggregationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) AggregationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) FileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) FileTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) PrefixConfig() AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigOutputReference {
	var returns AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigOutputReference
	_jsii_.Get(
		j,
		"prefixConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) PrefixConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prefixConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) PreserveSourceDataTyping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveSourceDataTyping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) PreserveSourceDataTypingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveSourceDataTypingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference{}

	_jsii_.Create(
		"awscc.appflowFlow.AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference_Override(a AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.appflowFlow.AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)SetFileType(val *string) {
	if err := j.validateSetFileTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileType",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)SetPreserveSourceDataTyping(val interface{}) {
	if err := j.validateSetPreserveSourceDataTypingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveSourceDataTyping",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) PutAggregationConfig(value *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfig) {
	if err := a.validatePutAggregationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAggregationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) PutPrefixConfig(value *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfig) {
	if err := a.validatePutPrefixConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrefixConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ResetAggregationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAggregationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ResetFileType() {
	_jsii_.InvokeVoid(
		a,
		"resetFileType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ResetPrefixConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefixConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ResetPreserveSourceDataTyping() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveSourceDataTyping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesS3S3OutputFormatConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

