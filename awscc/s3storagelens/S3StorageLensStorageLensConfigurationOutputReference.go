package s3storagelens

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3storagelens/internal"
)

type S3StorageLensStorageLensConfigurationOutputReference interface {
	cdktf.ComplexObject
	AccountLevel() S3StorageLensStorageLensConfigurationAccountLevelOutputReference
	AccountLevelInput() interface{}
	AwsOrg() S3StorageLensStorageLensConfigurationAwsOrgOutputReference
	AwsOrgInput() interface{}
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
	DataExport() S3StorageLensStorageLensConfigurationDataExportOutputReference
	DataExportInput() interface{}
	Exclude() S3StorageLensStorageLensConfigurationExcludeOutputReference
	ExcludeInput() interface{}
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Include() S3StorageLensStorageLensConfigurationIncludeOutputReference
	IncludeInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	StorageLensArn() *string
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
	PutAccountLevel(value *S3StorageLensStorageLensConfigurationAccountLevel)
	PutAwsOrg(value *S3StorageLensStorageLensConfigurationAwsOrg)
	PutDataExport(value *S3StorageLensStorageLensConfigurationDataExport)
	PutExclude(value *S3StorageLensStorageLensConfigurationExclude)
	PutInclude(value *S3StorageLensStorageLensConfigurationInclude)
	ResetAwsOrg()
	ResetDataExport()
	ResetExclude()
	ResetInclude()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3StorageLensStorageLensConfigurationOutputReference
type jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) AccountLevel() S3StorageLensStorageLensConfigurationAccountLevelOutputReference {
	var returns S3StorageLensStorageLensConfigurationAccountLevelOutputReference
	_jsii_.Get(
		j,
		"accountLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) AccountLevelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) AwsOrg() S3StorageLensStorageLensConfigurationAwsOrgOutputReference {
	var returns S3StorageLensStorageLensConfigurationAwsOrgOutputReference
	_jsii_.Get(
		j,
		"awsOrg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) AwsOrgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsOrgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) DataExport() S3StorageLensStorageLensConfigurationDataExportOutputReference {
	var returns S3StorageLensStorageLensConfigurationDataExportOutputReference
	_jsii_.Get(
		j,
		"dataExport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) DataExportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataExportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) Exclude() S3StorageLensStorageLensConfigurationExcludeOutputReference {
	var returns S3StorageLensStorageLensConfigurationExcludeOutputReference
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ExcludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) Include() S3StorageLensStorageLensConfigurationIncludeOutputReference {
	var returns S3StorageLensStorageLensConfigurationIncludeOutputReference
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) IncludeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) StorageLensArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageLensArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3StorageLensStorageLensConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3StorageLensStorageLensConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewS3StorageLensStorageLensConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.s3StorageLens.S3StorageLensStorageLensConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3StorageLensStorageLensConfigurationOutputReference_Override(s S3StorageLensStorageLensConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3StorageLens.S3StorageLensStorageLensConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference)SetIsEnabled(val interface{}) {
	if err := j.validateSetIsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) PutAccountLevel(value *S3StorageLensStorageLensConfigurationAccountLevel) {
	if err := s.validatePutAccountLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAccountLevel",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) PutAwsOrg(value *S3StorageLensStorageLensConfigurationAwsOrg) {
	if err := s.validatePutAwsOrgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsOrg",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) PutDataExport(value *S3StorageLensStorageLensConfigurationDataExport) {
	if err := s.validatePutDataExportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDataExport",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) PutExclude(value *S3StorageLensStorageLensConfigurationExclude) {
	if err := s.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExclude",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) PutInclude(value *S3StorageLensStorageLensConfigurationInclude) {
	if err := s.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInclude",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ResetAwsOrg() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsOrg",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ResetDataExport() {
	_jsii_.InvokeVoid(
		s,
		"resetDataExport",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ResetExclude() {
	_jsii_.InvokeVoid(
		s,
		"resetExclude",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ResetInclude() {
	_jsii_.InvokeVoid(
		s,
		"resetInclude",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3StorageLensStorageLensConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

