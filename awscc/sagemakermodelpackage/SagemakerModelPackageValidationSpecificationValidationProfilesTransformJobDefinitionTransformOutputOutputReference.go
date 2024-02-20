package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelpackage/internal"
)

type SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference interface {
	cdktf.ComplexObject
	Accept() *string
	SetAccept(val *string)
	AcceptInput() *string
	AssembleWith() *string
	SetAssembleWith(val *string)
	AssembleWithInput() *string
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
	InternalValue() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput
	SetInternalValue(val *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	S3OutputPath() *string
	SetS3OutputPath(val *string)
	S3OutputPathInput() *string
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
	ResetAccept()
	ResetAssembleWith()
	ResetKmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference
type jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) Accept() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accept",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) AcceptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) AssembleWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assembleWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) AssembleWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assembleWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) InternalValue() *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput {
	var returns *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) S3OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3OutputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) S3OutputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3OutputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference{}

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference_Override(s SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetAccept(val *string) {
	if err := j.validateSetAcceptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accept",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetAssembleWith(val *string) {
	if err := j.validateSetAssembleWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assembleWith",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetInternalValue(val *SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetS3OutputPath(val *string) {
	if err := j.validateSetS3OutputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3OutputPath",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) ResetAccept() {
	_jsii_.InvokeVoid(
		s,
		"resetAccept",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) ResetAssembleWith() {
	_jsii_.InvokeVoid(
		s,
		"resetAssembleWith",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerModelPackageValidationSpecificationValidationProfilesTransformJobDefinitionTransformOutputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

