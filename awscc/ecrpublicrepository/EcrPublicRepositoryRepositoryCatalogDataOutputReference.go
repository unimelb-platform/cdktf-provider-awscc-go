package ecrpublicrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ecrpublicrepository/internal"
)

type EcrPublicRepositoryRepositoryCatalogDataOutputReference interface {
	cdktf.ComplexObject
	AboutText() *string
	SetAboutText(val *string)
	AboutTextInput() *string
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OperatingSystems() *[]*string
	SetOperatingSystems(val *[]*string)
	OperatingSystemsInput() *[]*string
	RepositoryDescription() *string
	SetRepositoryDescription(val *string)
	RepositoryDescriptionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageText() *string
	SetUsageText(val *string)
	UsageTextInput() *string
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
	ResetAboutText()
	ResetArchitectures()
	ResetOperatingSystems()
	ResetRepositoryDescription()
	ResetUsageText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcrPublicRepositoryRepositoryCatalogDataOutputReference
type jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) AboutText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboutText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) AboutTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboutTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) OperatingSystems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatingSystems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) OperatingSystemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatingSystemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) RepositoryDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) RepositoryDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) UsageText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) UsageTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageTextInput",
		&returns,
	)
	return returns
}


func NewEcrPublicRepositoryRepositoryCatalogDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcrPublicRepositoryRepositoryCatalogDataOutputReference {
	_init_.Initialize()

	if err := validateNewEcrPublicRepositoryRepositoryCatalogDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference{}

	_jsii_.Create(
		"awscc.ecrPublicRepository.EcrPublicRepositoryRepositoryCatalogDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcrPublicRepositoryRepositoryCatalogDataOutputReference_Override(e EcrPublicRepositoryRepositoryCatalogDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ecrPublicRepository.EcrPublicRepositoryRepositoryCatalogDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetAboutText(val *string) {
	if err := j.validateSetAboutTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aboutText",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetOperatingSystems(val *[]*string) {
	if err := j.validateSetOperatingSystemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystems",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetRepositoryDescription(val *string) {
	if err := j.validateSetRepositoryDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryDescription",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference)SetUsageText(val *string) {
	if err := j.validateSetUsageTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageText",
		val,
	)
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ResetAboutText() {
	_jsii_.InvokeVoid(
		e,
		"resetAboutText",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ResetArchitectures() {
	_jsii_.InvokeVoid(
		e,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ResetOperatingSystems() {
	_jsii_.InvokeVoid(
		e,
		"resetOperatingSystems",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ResetRepositoryDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetRepositoryDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ResetUsageText() {
	_jsii_.InvokeVoid(
		e,
		"resetUsageText",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcrPublicRepositoryRepositoryCatalogDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

