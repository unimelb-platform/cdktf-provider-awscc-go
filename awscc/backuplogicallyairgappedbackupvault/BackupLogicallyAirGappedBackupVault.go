package backuplogicallyairgappedbackupvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/backuplogicallyairgappedbackupvault/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault awscc_backup_logically_air_gapped_backup_vault}.
type BackupLogicallyAirGappedBackupVault interface {
	cdktf.TerraformResource
	AccessPolicy() *string
	SetAccessPolicy(val *string)
	AccessPolicyInput() *string
	BackupVaultArn() *string
	BackupVaultName() *string
	SetBackupVaultName(val *string)
	BackupVaultNameInput() *string
	BackupVaultTags() *map[string]*string
	SetBackupVaultTags(val *map[string]*string)
	BackupVaultTagsInput() *map[string]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	EncryptionKeyArn() *string
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
	MaxRetentionDays() *float64
	SetMaxRetentionDays(val *float64)
	MaxRetentionDaysInput() *float64
	MinRetentionDays() *float64
	SetMinRetentionDays(val *float64)
	MinRetentionDaysInput() *float64
	// The tree node.
	Node() constructs.Node
	Notifications() BackupLogicallyAirGappedBackupVaultNotificationsOutputReference
	NotificationsInput() interface{}
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
	VaultState() *string
	VaultType() *string
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
	PutNotifications(value *BackupLogicallyAirGappedBackupVaultNotifications)
	ResetAccessPolicy()
	ResetBackupVaultTags()
	ResetNotifications()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for BackupLogicallyAirGappedBackupVault
type jsiiProxy_BackupLogicallyAirGappedBackupVault struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) AccessPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) AccessPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) BackupVaultArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) BackupVaultName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) BackupVaultNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupVaultNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) BackupVaultTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"backupVaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) BackupVaultTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"backupVaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) EncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) MaxRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) MaxRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) MinRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) MinRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Notifications() BackupLogicallyAirGappedBackupVaultNotificationsOutputReference {
	var returns BackupLogicallyAirGappedBackupVaultNotificationsOutputReference
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) NotificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) VaultState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault) VaultType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault awscc_backup_logically_air_gapped_backup_vault} Resource.
func NewBackupLogicallyAirGappedBackupVault(scope constructs.Construct, id *string, config *BackupLogicallyAirGappedBackupVaultConfig) BackupLogicallyAirGappedBackupVault {
	_init_.Initialize()

	if err := validateNewBackupLogicallyAirGappedBackupVaultParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupLogicallyAirGappedBackupVault{}

	_jsii_.Create(
		"awscc.backupLogicallyAirGappedBackupVault.BackupLogicallyAirGappedBackupVault",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault awscc_backup_logically_air_gapped_backup_vault} Resource.
func NewBackupLogicallyAirGappedBackupVault_Override(b BackupLogicallyAirGappedBackupVault, scope constructs.Construct, id *string, config *BackupLogicallyAirGappedBackupVaultConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.backupLogicallyAirGappedBackupVault.BackupLogicallyAirGappedBackupVault",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetAccessPolicy(val *string) {
	if err := j.validateSetAccessPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPolicy",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetBackupVaultName(val *string) {
	if err := j.validateSetBackupVaultNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupVaultName",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetBackupVaultTags(val *map[string]*string) {
	if err := j.validateSetBackupVaultTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupVaultTags",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetMaxRetentionDays(val *float64) {
	if err := j.validateSetMaxRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetentionDays",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetMinRetentionDays(val *float64) {
	if err := j.validateSetMinRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetentionDays",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BackupLogicallyAirGappedBackupVault)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a BackupLogicallyAirGappedBackupVault resource upon running "cdktf plan <stack-name>".
func BackupLogicallyAirGappedBackupVault_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBackupLogicallyAirGappedBackupVault_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.backupLogicallyAirGappedBackupVault.BackupLogicallyAirGappedBackupVault",
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
func BackupLogicallyAirGappedBackupVault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupLogicallyAirGappedBackupVault_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.backupLogicallyAirGappedBackupVault.BackupLogicallyAirGappedBackupVault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupLogicallyAirGappedBackupVault_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupLogicallyAirGappedBackupVault_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.backupLogicallyAirGappedBackupVault.BackupLogicallyAirGappedBackupVault",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BackupLogicallyAirGappedBackupVault_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBackupLogicallyAirGappedBackupVault_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.backupLogicallyAirGappedBackupVault.BackupLogicallyAirGappedBackupVault",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BackupLogicallyAirGappedBackupVault_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.backupLogicallyAirGappedBackupVault.BackupLogicallyAirGappedBackupVault",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) PutNotifications(value *BackupLogicallyAirGappedBackupVaultNotifications) {
	if err := b.validatePutNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putNotifications",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ResetAccessPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetAccessPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ResetBackupVaultTags() {
	_jsii_.InvokeVoid(
		b,
		"resetBackupVaultTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ResetNotifications() {
	_jsii_.InvokeVoid(
		b,
		"resetNotifications",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupLogicallyAirGappedBackupVault) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

