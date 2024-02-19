package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelpackage/internal"
)

type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference interface {
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
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataSource() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSourceOutputReference
	DataSourceInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput
	SetInternalValue(val *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput)
	SplitType() *string
	SetSplitType(val *string)
	SplitTypeInput() *string
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
	PutDataSource(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource)
	ResetCompressionType()
	ResetContentType()
	ResetSplitType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference
type jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) DataSource() SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSourceOutputReference {
	var returns SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSourceOutputReference
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) DataSourceInput() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource {
	var returns *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) InternalValue() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput {
	var returns *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) SplitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) SplitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"splitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference{}

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference_Override(s SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetInternalValue(val *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetSplitType(val *string) {
	if err := j.validateSetSplitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitType",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) PutDataSource(value *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputDataSource) {
	if err := s.validatePutDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		s,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		s,
		"resetContentType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ResetSplitType() {
	_jsii_.InvokeVoid(
		s,
		"resetSplitType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

