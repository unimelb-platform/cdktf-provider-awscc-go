package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kendradatasource/internal"
)

type KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference interface {
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
	CustomKnowledgeArticleTypeConfigurations() KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationCustomKnowledgeArticleTypeConfigurationsList
	CustomKnowledgeArticleTypeConfigurationsInput() interface{}
	// Experimental.
	Fqn() *string
	IncludedStates() *[]*string
	SetIncludedStates(val *[]*string)
	IncludedStatesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StandardKnowledgeArticleTypeConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationStandardKnowledgeArticleTypeConfigurationOutputReference
	StandardKnowledgeArticleTypeConfigurationInput() interface{}
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
	PutCustomKnowledgeArticleTypeConfigurations(value interface{})
	PutStandardKnowledgeArticleTypeConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationStandardKnowledgeArticleTypeConfiguration)
	ResetCustomKnowledgeArticleTypeConfigurations()
	ResetStandardKnowledgeArticleTypeConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) CustomKnowledgeArticleTypeConfigurations() KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationCustomKnowledgeArticleTypeConfigurationsList {
	var returns KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationCustomKnowledgeArticleTypeConfigurationsList
	_jsii_.Get(
		j,
		"customKnowledgeArticleTypeConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) CustomKnowledgeArticleTypeConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customKnowledgeArticleTypeConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) IncludedStates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) IncludedStatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) StandardKnowledgeArticleTypeConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationStandardKnowledgeArticleTypeConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationStandardKnowledgeArticleTypeConfigurationOutputReference
	_jsii_.Get(
		j,
		"standardKnowledgeArticleTypeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) StandardKnowledgeArticleTypeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"standardKnowledgeArticleTypeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference)SetIncludedStates(val *[]*string) {
	if err := j.validateSetIncludedStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedStates",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) PutCustomKnowledgeArticleTypeConfigurations(value interface{}) {
	if err := k.validatePutCustomKnowledgeArticleTypeConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCustomKnowledgeArticleTypeConfigurations",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) PutStandardKnowledgeArticleTypeConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationStandardKnowledgeArticleTypeConfiguration) {
	if err := k.validatePutStandardKnowledgeArticleTypeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putStandardKnowledgeArticleTypeConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) ResetCustomKnowledgeArticleTypeConfigurations() {
	_jsii_.InvokeVoid(
		k,
		"resetCustomKnowledgeArticleTypeConfigurations",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) ResetStandardKnowledgeArticleTypeConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetStandardKnowledgeArticleTypeConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

