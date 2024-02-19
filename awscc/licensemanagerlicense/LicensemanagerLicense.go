package licensemanagerlicense

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/licensemanagerlicense/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license awscc_licensemanager_license}.
type LicensemanagerLicense interface {
	cdktf.TerraformResource
	Beneficiary() *string
	SetBeneficiary(val *string)
	BeneficiaryInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumptionConfiguration() LicensemanagerLicenseConsumptionConfigurationOutputReference
	ConsumptionConfigurationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Entitlements() LicensemanagerLicenseEntitlementsList
	EntitlementsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HomeRegion() *string
	SetHomeRegion(val *string)
	HomeRegionInput() *string
	Id() *string
	Issuer() LicensemanagerLicenseIssuerOutputReference
	IssuerInput() interface{}
	LicenseArn() *string
	LicenseMetadata() LicensemanagerLicenseLicenseMetadataList
	LicenseMetadataInput() interface{}
	LicenseName() *string
	SetLicenseName(val *string)
	LicenseNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ProductName() *string
	SetProductName(val *string)
	ProductNameInput() *string
	ProductSku() *string
	SetProductSku(val *string)
	ProductSkuInput() *string
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
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Validity() LicensemanagerLicenseValidityOutputReference
	ValidityInput() interface{}
	Version() *string
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutConsumptionConfiguration(value *LicensemanagerLicenseConsumptionConfiguration)
	PutEntitlements(value interface{})
	PutIssuer(value *LicensemanagerLicenseIssuer)
	PutLicenseMetadata(value interface{})
	PutValidity(value *LicensemanagerLicenseValidity)
	ResetBeneficiary()
	ResetLicenseMetadata()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProductSku()
	ResetStatus()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LicensemanagerLicense
type jsiiProxy_LicensemanagerLicense struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LicensemanagerLicense) Beneficiary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beneficiary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) BeneficiaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beneficiaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ConsumptionConfiguration() LicensemanagerLicenseConsumptionConfigurationOutputReference {
	var returns LicensemanagerLicenseConsumptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"consumptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ConsumptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Entitlements() LicensemanagerLicenseEntitlementsList {
	var returns LicensemanagerLicenseEntitlementsList
	_jsii_.Get(
		j,
		"entitlements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) EntitlementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entitlementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) HomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) HomeRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Issuer() LicensemanagerLicenseIssuerOutputReference {
	var returns LicensemanagerLicenseIssuerOutputReference
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) IssuerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) LicenseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) LicenseMetadata() LicensemanagerLicenseLicenseMetadataList {
	var returns LicensemanagerLicenseLicenseMetadataList
	_jsii_.Get(
		j,
		"licenseMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) LicenseMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) LicenseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) LicenseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ProductNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ProductSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ProductSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Validity() LicensemanagerLicenseValidityOutputReference {
	var returns LicensemanagerLicenseValidityOutputReference
	_jsii_.Get(
		j,
		"validity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) ValidityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicense) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license awscc_licensemanager_license} Resource.
func NewLicensemanagerLicense(scope constructs.Construct, id *string, config *LicensemanagerLicenseConfig) LicensemanagerLicense {
	_init_.Initialize()

	if err := validateNewLicensemanagerLicenseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicensemanagerLicense{}

	_jsii_.Create(
		"awscc.licensemanagerLicense.LicensemanagerLicense",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/licensemanager_license awscc_licensemanager_license} Resource.
func NewLicensemanagerLicense_Override(l LicensemanagerLicense, scope constructs.Construct, id *string, config *LicensemanagerLicenseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.licensemanagerLicense.LicensemanagerLicense",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetBeneficiary(val *string) {
	if err := j.validateSetBeneficiaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"beneficiary",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetHomeRegion(val *string) {
	if err := j.validateSetHomeRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homeRegion",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetLicenseName(val *string) {
	if err := j.validateSetLicenseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseName",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetProductName(val *string) {
	if err := j.validateSetProductNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetProductSku(val *string) {
	if err := j.validateSetProductSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productSku",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicense)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Generates CDKTF code for importing a LicensemanagerLicense resource upon running "cdktf plan <stack-name>".
func LicensemanagerLicense_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLicensemanagerLicense_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.licensemanagerLicense.LicensemanagerLicense",
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
func LicensemanagerLicense_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLicensemanagerLicense_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.licensemanagerLicense.LicensemanagerLicense",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LicensemanagerLicense_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLicensemanagerLicense_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.licensemanagerLicense.LicensemanagerLicense",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LicensemanagerLicense_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLicensemanagerLicense_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.licensemanagerLicense.LicensemanagerLicense",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LicensemanagerLicense_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.licensemanagerLicense.LicensemanagerLicense",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) PutConsumptionConfiguration(value *LicensemanagerLicenseConsumptionConfiguration) {
	if err := l.validatePutConsumptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConsumptionConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) PutEntitlements(value interface{}) {
	if err := l.validatePutEntitlementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEntitlements",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) PutIssuer(value *LicensemanagerLicenseIssuer) {
	if err := l.validatePutIssuerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIssuer",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) PutLicenseMetadata(value interface{}) {
	if err := l.validatePutLicenseMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLicenseMetadata",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) PutValidity(value *LicensemanagerLicenseValidity) {
	if err := l.validatePutValidityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putValidity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicense) ResetBeneficiary() {
	_jsii_.InvokeVoid(
		l,
		"resetBeneficiary",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicense) ResetLicenseMetadata() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseMetadata",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicense) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicense) ResetProductSku() {
	_jsii_.InvokeVoid(
		l,
		"resetProductSku",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicense) ResetStatus() {
	_jsii_.InvokeVoid(
		l,
		"resetStatus",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicense) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicense) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

