package entityresolutionidmappingworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/entityresolutionidmappingworkflow/internal"
)

type EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference interface {
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
	IntermediateS3Path() *string
	SetIntermediateS3Path(val *string)
	IntermediateS3PathInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference
type jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) IntermediateS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intermediateS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) IntermediateS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intermediateS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewEntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.entityresolutionIdMappingWorkflow.EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference_Override(e EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.entityresolutionIdMappingWorkflow.EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference)SetIntermediateS3Path(val *string) {
	if err := j.validateSetIntermediateS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intermediateS3Path",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

