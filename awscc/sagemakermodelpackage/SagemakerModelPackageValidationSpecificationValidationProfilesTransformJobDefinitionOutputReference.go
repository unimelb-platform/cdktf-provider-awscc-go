package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelpackage/internal"
)

type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference interface {
	cdktf.ComplexObject
	BatchStrategy() *string
	SetBatchStrategy(val *string)
	BatchStrategyInput() *string
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
	Environment() *map[string]*string
	SetEnvironment(val *map[string]*string)
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinition
	SetInternalValue(val *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinition)
	MaxConcurrentTransforms() *float64
	SetMaxConcurrentTransforms(val *float64)
	MaxConcurrentTransformsInput() *float64
	MaxPayloadInMb() *float64
	SetMaxPayloadInMb(val *float64)
	MaxPayloadInMbInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformInput() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference
	TransformInputInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput
	TransformOutput() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference
	TransformOutputInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput
	TransformResources() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResourcesOutputReference
	TransformResourcesInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResources
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
	PutTransformInput(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput)
	PutTransformOutput(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput)
	PutTransformResources(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResources)
	ResetBatchStrategy()
	ResetEnvironment()
	ResetMaxConcurrentTransforms()
	ResetMaxPayloadInMb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference
type jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) BatchStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) BatchStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) InternalValue() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinition {
	var returns *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxConcurrentTransforms() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTransforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxConcurrentTransformsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTransformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxPayloadInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) MaxPayloadInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPayloadInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformInput() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference {
	var returns SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference
	_jsii_.Get(
		j,
		"transformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformInputInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput {
	var returns *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput
	_jsii_.Get(
		j,
		"transformInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformOutput() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference {
	var returns SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference
	_jsii_.Get(
		j,
		"transformOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformOutputInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput {
	var returns *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput
	_jsii_.Get(
		j,
		"transformOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformResources() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResourcesOutputReference {
	var returns SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResourcesOutputReference
	_jsii_.Get(
		j,
		"transformResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) TransformResourcesInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResources {
	var returns *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResources
	_jsii_.Get(
		j,
		"transformResourcesInput",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference{}

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference_Override(s SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetBatchStrategy(val *string) {
	if err := j.validateSetBatchStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchStrategy",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetInternalValue(val *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetMaxConcurrentTransforms(val *float64) {
	if err := j.validateSetMaxConcurrentTransformsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentTransforms",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetMaxPayloadInMb(val *float64) {
	if err := j.validateSetMaxPayloadInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPayloadInMb",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) PutTransformInput(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput) {
	if err := s.validatePutTransformInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformInput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) PutTransformOutput(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput) {
	if err := s.validatePutTransformOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformOutput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) PutTransformResources(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformResources) {
	if err := s.validatePutTransformResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformResources",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetBatchStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetBatchStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		s,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetMaxConcurrentTransforms() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxConcurrentTransforms",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ResetMaxPayloadInMb() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxPayloadInMb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

