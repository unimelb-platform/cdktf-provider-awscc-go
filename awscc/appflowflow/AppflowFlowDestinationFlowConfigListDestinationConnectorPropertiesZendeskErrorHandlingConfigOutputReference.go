package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/appflowflow/internal"
)

type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference interface {
	cdktf.ComplexObject
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	BucketPrefix() *string
	SetBucketPrefix(val *string)
	BucketPrefixInput() *string
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
	FailOnFirstError() interface{}
	SetFailOnFirstError(val interface{})
	FailOnFirstErrorInput() interface{}
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
	ResetBucketName()
	ResetBucketPrefix()
	ResetFailOnFirstError()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference
type jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) FailOnFirstError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) FailOnFirstErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference{}

	_jsii_.Create(
		"awscc.appflowFlow.AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference_Override(a AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.appflowFlow.AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetBucketPrefix(val *string) {
	if err := j.validateSetBucketPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetFailOnFirstError(val interface{}) {
	if err := j.validateSetFailOnFirstErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOnFirstError",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) ResetFailOnFirstError() {
	_jsii_.InvokeVoid(
		a,
		"resetFailOnFirstError",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesZendeskErrorHandlingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

