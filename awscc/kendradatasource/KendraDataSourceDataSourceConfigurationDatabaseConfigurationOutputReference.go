package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kendradatasource/internal"
)

type KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference interface {
	cdktf.ComplexObject
	AclConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationAclConfigurationOutputReference
	AclConfigurationInput() interface{}
	ColumnConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference
	ColumnConfigurationInput() *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration
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
	ConnectionConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfigurationOutputReference
	ConnectionConfigurationInput() *KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseEngineType() *string
	SetDatabaseEngineType(val *string)
	DatabaseEngineTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SqlConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationSqlConfigurationOutputReference
	SqlConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationVpcConfigurationOutputReference
	VpcConfigurationInput() interface{}
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
	PutAclConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationAclConfiguration)
	PutColumnConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration)
	PutConnectionConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration)
	PutSqlConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationSqlConfiguration)
	PutVpcConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationVpcConfiguration)
	ResetAclConfiguration()
	ResetSqlConfiguration()
	ResetVpcConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) AclConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationAclConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationDatabaseConfigurationAclConfigurationOutputReference
	_jsii_.Get(
		j,
		"aclConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) AclConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aclConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ColumnConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfigurationOutputReference
	_jsii_.Get(
		j,
		"columnConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ColumnConfigurationInput() *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration {
	var returns *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration
	_jsii_.Get(
		j,
		"columnConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ConnectionConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfigurationOutputReference
	_jsii_.Get(
		j,
		"connectionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ConnectionConfigurationInput() *KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration {
	var returns *KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration
	_jsii_.Get(
		j,
		"connectionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) DatabaseEngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEngineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) DatabaseEngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseEngineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) SqlConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationSqlConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationDatabaseConfigurationSqlConfigurationOutputReference
	_jsii_.Get(
		j,
		"sqlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) SqlConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) VpcConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationVpcConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationDatabaseConfigurationVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) VpcConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference)SetDatabaseEngineType(val *string) {
	if err := j.validateSetDatabaseEngineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseEngineType",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) PutAclConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationAclConfiguration) {
	if err := k.validatePutAclConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAclConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) PutColumnConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration) {
	if err := k.validatePutColumnConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putColumnConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) PutConnectionConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration) {
	if err := k.validatePutConnectionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putConnectionConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) PutSqlConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationSqlConfiguration) {
	if err := k.validatePutSqlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSqlConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) PutVpcConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfigurationVpcConfiguration) {
	if err := k.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ResetAclConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetAclConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ResetSqlConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSqlConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

