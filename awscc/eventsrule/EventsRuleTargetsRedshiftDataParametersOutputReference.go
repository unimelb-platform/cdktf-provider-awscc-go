package eventsrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/eventsrule/internal"
)

type EventsRuleTargetsRedshiftDataParametersOutputReference interface {
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
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	DbUser() *string
	SetDbUser(val *string)
	DbUserInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecretManagerArn() *string
	SetSecretManagerArn(val *string)
	SecretManagerArnInput() *string
	Sql() *string
	SetSql(val *string)
	SqlInput() *string
	Sqls() *[]*string
	SetSqls(val *[]*string)
	SqlsInput() *[]*string
	StatementName() *string
	SetStatementName(val *string)
	StatementNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WithEvent() interface{}
	SetWithEvent(val interface{})
	WithEventInput() interface{}
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
	ResetDbUser()
	ResetSecretManagerArn()
	ResetSql()
	ResetSqls()
	ResetStatementName()
	ResetWithEvent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventsRuleTargetsRedshiftDataParametersOutputReference
type jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) DbUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) DbUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) SecretManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) SecretManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) Sql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) SqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) Sqls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sqls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) SqlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sqlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) StatementName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) StatementNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) WithEvent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withEvent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) WithEventInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withEventInput",
		&returns,
	)
	return returns
}


func NewEventsRuleTargetsRedshiftDataParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventsRuleTargetsRedshiftDataParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEventsRuleTargetsRedshiftDataParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference{}

	_jsii_.Create(
		"awscc.eventsRule.EventsRuleTargetsRedshiftDataParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsRuleTargetsRedshiftDataParametersOutputReference_Override(e EventsRuleTargetsRedshiftDataParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.eventsRule.EventsRuleTargetsRedshiftDataParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetDbUser(val *string) {
	if err := j.validateSetDbUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbUser",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetSecretManagerArn(val *string) {
	if err := j.validateSetSecretManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretManagerArn",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetSql(val *string) {
	if err := j.validateSetSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sql",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetSqls(val *[]*string) {
	if err := j.validateSetSqlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqls",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetStatementName(val *string) {
	if err := j.validateSetStatementNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementName",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference)SetWithEvent(val interface{}) {
	if err := j.validateSetWithEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"withEvent",
		val,
	)
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ResetDbUser() {
	_jsii_.InvokeVoid(
		e,
		"resetDbUser",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ResetSecretManagerArn() {
	_jsii_.InvokeVoid(
		e,
		"resetSecretManagerArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ResetSql() {
	_jsii_.InvokeVoid(
		e,
		"resetSql",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ResetSqls() {
	_jsii_.InvokeVoid(
		e,
		"resetSqls",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ResetStatementName() {
	_jsii_.InvokeVoid(
		e,
		"resetStatementName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ResetWithEvent() {
	_jsii_.InvokeVoid(
		e,
		"resetWithEvent",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventsRuleTargetsRedshiftDataParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

