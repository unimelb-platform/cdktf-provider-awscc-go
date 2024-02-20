package dmsdataprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dmsdataprovider/internal"
)

type DmsDataProviderSettingsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MicrosoftSqlServerSettings() DmsDataProviderSettingsMicrosoftSqlServerSettingsOutputReference
	MicrosoftSqlServerSettingsInput() interface{}
	MySqlSettings() DmsDataProviderSettingsMySqlSettingsOutputReference
	MySqlSettingsInput() interface{}
	OracleSettings() DmsDataProviderSettingsOracleSettingsOutputReference
	OracleSettingsInput() interface{}
	PostgreSqlSettings() DmsDataProviderSettingsPostgreSqlSettingsOutputReference
	PostgreSqlSettingsInput() interface{}
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
	PutMicrosoftSqlServerSettings(value *DmsDataProviderSettingsMicrosoftSqlServerSettings)
	PutMySqlSettings(value *DmsDataProviderSettingsMySqlSettings)
	PutOracleSettings(value *DmsDataProviderSettingsOracleSettings)
	PutPostgreSqlSettings(value *DmsDataProviderSettingsPostgreSqlSettings)
	ResetMicrosoftSqlServerSettings()
	ResetMySqlSettings()
	ResetOracleSettings()
	ResetPostgreSqlSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsDataProviderSettingsOutputReference
type jsiiProxy_DmsDataProviderSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MicrosoftSqlServerSettings() DmsDataProviderSettingsMicrosoftSqlServerSettingsOutputReference {
	var returns DmsDataProviderSettingsMicrosoftSqlServerSettingsOutputReference
	_jsii_.Get(
		j,
		"microsoftSqlServerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MicrosoftSqlServerSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSqlServerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MySqlSettings() DmsDataProviderSettingsMySqlSettingsOutputReference {
	var returns DmsDataProviderSettingsMySqlSettingsOutputReference
	_jsii_.Get(
		j,
		"mySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) MySqlSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mySqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) OracleSettings() DmsDataProviderSettingsOracleSettingsOutputReference {
	var returns DmsDataProviderSettingsOracleSettingsOutputReference
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) OracleSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) PostgreSqlSettings() DmsDataProviderSettingsPostgreSqlSettingsOutputReference {
	var returns DmsDataProviderSettingsPostgreSqlSettingsOutputReference
	_jsii_.Get(
		j,
		"postgreSqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) PostgreSqlSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgreSqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsDataProviderSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DmsDataProviderSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsDataProviderSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsDataProviderSettingsOutputReference{}

	_jsii_.Create(
		"awscc.dmsDataProvider.DmsDataProviderSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDmsDataProviderSettingsOutputReference_Override(d DmsDataProviderSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dmsDataProvider.DmsDataProviderSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsDataProviderSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutMicrosoftSqlServerSettings(value *DmsDataProviderSettingsMicrosoftSqlServerSettings) {
	if err := d.validatePutMicrosoftSqlServerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMicrosoftSqlServerSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutMySqlSettings(value *DmsDataProviderSettingsMySqlSettings) {
	if err := d.validatePutMySqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMySqlSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutOracleSettings(value *DmsDataProviderSettingsOracleSettings) {
	if err := d.validatePutOracleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracleSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) PutPostgreSqlSettings(value *DmsDataProviderSettingsPostgreSqlSettings) {
	if err := d.validatePutPostgreSqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgreSqlSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetMicrosoftSqlServerSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftSqlServerSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetMySqlSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMySqlSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetOracleSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetOracleSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ResetPostgreSqlSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgreSqlSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataProviderSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

