package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kendradatasource/internal"
)

type KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference interface {
	cdktf.ComplexObject
	ChangeDetectingColumns() *[]*string
	SetChangeDetectingColumns(val *[]*string)
	ChangeDetectingColumnsInput() *[]*string
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
	DocumentDataColumnName() *string
	SetDocumentDataColumnName(val *string)
	DocumentDataColumnNameInput() *string
	DocumentIdColumnName() *string
	SetDocumentIdColumnName(val *string)
	DocumentIdColumnNameInput() *string
	DocumentTitleColumnName() *string
	SetDocumentTitleColumnName(val *string)
	DocumentTitleColumnNameInput() *string
	FieldMappings() KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationFieldMappingsList
	FieldMappingsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration
	SetInternalValue(val *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration)
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
	PutFieldMappings(value interface{})
	ResetDocumentTitleColumnName()
	ResetFieldMappings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ChangeDetectingColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"changeDetectingColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ChangeDetectingColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"changeDetectingColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) DocumentDataColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentDataColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) DocumentDataColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentDataColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) DocumentIdColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentIdColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) DocumentIdColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentIdColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) DocumentTitleColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTitleColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) DocumentTitleColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTitleColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) FieldMappings() KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationFieldMappingsList {
	var returns KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationFieldMappingsList
	_jsii_.Get(
		j,
		"fieldMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) FieldMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) InternalValue() *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration {
	var returns *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetChangeDetectingColumns(val *[]*string) {
	if err := j.validateSetChangeDetectingColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeDetectingColumns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetDocumentDataColumnName(val *string) {
	if err := j.validateSetDocumentDataColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentDataColumnName",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetDocumentIdColumnName(val *string) {
	if err := j.validateSetDocumentIdColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentIdColumnName",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetDocumentTitleColumnName(val *string) {
	if err := j.validateSetDocumentTitleColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentTitleColumnName",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetInternalValue(val *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) PutFieldMappings(value interface{}) {
	if err := k.validatePutFieldMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putFieldMappings",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ResetDocumentTitleColumnName() {
	_jsii_.InvokeVoid(
		k,
		"resetDocumentTitleColumnName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ResetFieldMappings() {
	_jsii_.InvokeVoid(
		k,
		"resetFieldMappings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

