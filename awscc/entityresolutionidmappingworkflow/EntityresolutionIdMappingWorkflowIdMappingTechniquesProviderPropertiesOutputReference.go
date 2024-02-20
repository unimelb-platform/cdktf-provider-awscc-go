package entityresolutionidmappingworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/entityresolutionidmappingworkflow/internal"
)

type EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference interface {
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
	IntermediateSourceConfiguration() EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference
	IntermediateSourceConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProviderConfiguration() *map[string]*string
	SetProviderConfiguration(val *map[string]*string)
	ProviderConfigurationInput() *map[string]*string
	ProviderServiceArn() *string
	SetProviderServiceArn(val *string)
	ProviderServiceArnInput() *string
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
	PutIntermediateSourceConfiguration(value *EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfiguration)
	ResetIntermediateSourceConfiguration()
	ResetProviderConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference
type jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) IntermediateSourceConfiguration() EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference {
	var returns EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"intermediateSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) IntermediateSourceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intermediateSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ProviderConfiguration() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"providerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ProviderConfigurationInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"providerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ProviderServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerServiceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ProviderServiceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerServiceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewEntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference{}

	_jsii_.Create(
		"awscc.entityresolutionIdMappingWorkflow.EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference_Override(e EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.entityresolutionIdMappingWorkflow.EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference)SetProviderConfiguration(val *map[string]*string) {
	if err := j.validateSetProviderConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerConfiguration",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference)SetProviderServiceArn(val *string) {
	if err := j.validateSetProviderServiceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerServiceArn",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) PutIntermediateSourceConfiguration(value *EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfiguration) {
	if err := e.validatePutIntermediateSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIntermediateSourceConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ResetIntermediateSourceConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetIntermediateSourceConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ResetProviderConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetProviderConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

