package sesemailidentity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sesemailidentity/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity awscc_ses_email_identity}.
type SesEmailIdentity interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigurationSetAttributes() SesEmailIdentityConfigurationSetAttributesOutputReference
	ConfigurationSetAttributesInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DkimAttributes() SesEmailIdentityDkimAttributesOutputReference
	DkimAttributesInput() interface{}
	DkimDnsTokenName1() *string
	DkimDnsTokenName2() *string
	DkimDnsTokenName3() *string
	DkimDnsTokenValue1() *string
	DkimDnsTokenValue2() *string
	DkimDnsTokenValue3() *string
	DkimSigningAttributes() SesEmailIdentityDkimSigningAttributesOutputReference
	DkimSigningAttributesInput() interface{}
	EmailIdentity() *string
	SetEmailIdentity(val *string)
	EmailIdentityInput() *string
	FeedbackAttributes() SesEmailIdentityFeedbackAttributesOutputReference
	FeedbackAttributesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MailFromAttributes() SesEmailIdentityMailFromAttributesOutputReference
	MailFromAttributesInput() interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutConfigurationSetAttributes(value *SesEmailIdentityConfigurationSetAttributes)
	PutDkimAttributes(value *SesEmailIdentityDkimAttributes)
	PutDkimSigningAttributes(value *SesEmailIdentityDkimSigningAttributes)
	PutFeedbackAttributes(value *SesEmailIdentityFeedbackAttributes)
	PutMailFromAttributes(value *SesEmailIdentityMailFromAttributes)
	ResetConfigurationSetAttributes()
	ResetDkimAttributes()
	ResetDkimSigningAttributes()
	ResetFeedbackAttributes()
	ResetMailFromAttributes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SesEmailIdentity
type jsiiProxy_SesEmailIdentity struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SesEmailIdentity) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) ConfigurationSetAttributes() SesEmailIdentityConfigurationSetAttributesOutputReference {
	var returns SesEmailIdentityConfigurationSetAttributesOutputReference
	_jsii_.Get(
		j,
		"configurationSetAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) ConfigurationSetAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurationSetAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimAttributes() SesEmailIdentityDkimAttributesOutputReference {
	var returns SesEmailIdentityDkimAttributesOutputReference
	_jsii_.Get(
		j,
		"dkimAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dkimAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimDnsTokenName1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dkimDnsTokenName1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimDnsTokenName2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dkimDnsTokenName2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimDnsTokenName3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dkimDnsTokenName3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimDnsTokenValue1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dkimDnsTokenValue1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimDnsTokenValue2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dkimDnsTokenValue2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimDnsTokenValue3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dkimDnsTokenValue3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimSigningAttributes() SesEmailIdentityDkimSigningAttributesOutputReference {
	var returns SesEmailIdentityDkimSigningAttributesOutputReference
	_jsii_.Get(
		j,
		"dkimSigningAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) DkimSigningAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dkimSigningAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) EmailIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) EmailIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) FeedbackAttributes() SesEmailIdentityFeedbackAttributesOutputReference {
	var returns SesEmailIdentityFeedbackAttributesOutputReference
	_jsii_.Get(
		j,
		"feedbackAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) FeedbackAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"feedbackAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) MailFromAttributes() SesEmailIdentityMailFromAttributesOutputReference {
	var returns SesEmailIdentityMailFromAttributesOutputReference
	_jsii_.Get(
		j,
		"mailFromAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) MailFromAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mailFromAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentity) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity awscc_ses_email_identity} Resource.
func NewSesEmailIdentity(scope constructs.Construct, id *string, config *SesEmailIdentityConfig) SesEmailIdentity {
	_init_.Initialize()

	if err := validateNewSesEmailIdentityParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesEmailIdentity{}

	_jsii_.Create(
		"awscc.sesEmailIdentity.SesEmailIdentity",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_email_identity awscc_ses_email_identity} Resource.
func NewSesEmailIdentity_Override(s SesEmailIdentity, scope constructs.Construct, id *string, config *SesEmailIdentityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sesEmailIdentity.SesEmailIdentity",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetEmailIdentity(val *string) {
	if err := j.validateSetEmailIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailIdentity",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentity)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a SesEmailIdentity resource upon running "cdktf plan <stack-name>".
func SesEmailIdentity_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSesEmailIdentity_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.sesEmailIdentity.SesEmailIdentity",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SesEmailIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesEmailIdentity_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sesEmailIdentity.SesEmailIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesEmailIdentity_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesEmailIdentity_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sesEmailIdentity.SesEmailIdentity",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SesEmailIdentity_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSesEmailIdentity_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sesEmailIdentity.SesEmailIdentity",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SesEmailIdentity_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.sesEmailIdentity.SesEmailIdentity",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SesEmailIdentity) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SesEmailIdentity) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SesEmailIdentity) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SesEmailIdentity) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SesEmailIdentity) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SesEmailIdentity) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SesEmailIdentity) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SesEmailIdentity) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SesEmailIdentity) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SesEmailIdentity) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SesEmailIdentity) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SesEmailIdentity) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SesEmailIdentity) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentity) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SesEmailIdentity) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SesEmailIdentity) PutConfigurationSetAttributes(value *SesEmailIdentityConfigurationSetAttributes) {
	if err := s.validatePutConfigurationSetAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfigurationSetAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEmailIdentity) PutDkimAttributes(value *SesEmailIdentityDkimAttributes) {
	if err := s.validatePutDkimAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDkimAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEmailIdentity) PutDkimSigningAttributes(value *SesEmailIdentityDkimSigningAttributes) {
	if err := s.validatePutDkimSigningAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDkimSigningAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEmailIdentity) PutFeedbackAttributes(value *SesEmailIdentityFeedbackAttributes) {
	if err := s.validatePutFeedbackAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFeedbackAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEmailIdentity) PutMailFromAttributes(value *SesEmailIdentityMailFromAttributes) {
	if err := s.validatePutMailFromAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMailFromAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesEmailIdentity) ResetConfigurationSetAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetConfigurationSetAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentity) ResetDkimAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetDkimAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentity) ResetDkimSigningAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetDkimSigningAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentity) ResetFeedbackAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetFeedbackAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentity) ResetMailFromAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetMailFromAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentity) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentity) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentity) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentity) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

