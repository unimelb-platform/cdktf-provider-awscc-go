package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/appflowflow/internal"
)

type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference interface {
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

// The jsii proxy struct for AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference
type jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) FailOnFirstError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) FailOnFirstErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference{}

	_jsii_.Create(
		"awscc.appflowFlow.AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference_Override(a AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.appflowFlow.AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetBucketPrefix(val *string) {
	if err := j.validateSetBucketPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetFailOnFirstError(val interface{}) {
	if err := j.validateSetFailOnFirstErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOnFirstError",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) ResetFailOnFirstError() {
	_jsii_.InvokeVoid(
		a,
		"resetFailOnFirstError",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSapoDataErrorHandlingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
