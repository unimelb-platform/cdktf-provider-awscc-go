package rdsdbinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsDbInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The amount of storage (in gigabytes) to be initially allocated for the database instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#allocated_storage RdsDbInstance#allocated_storage}
	AllocatedStorage *string `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// A value that indicates whether major version upgrades are allowed.
	//
	// Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#allow_major_version_upgrade RdsDbInstance#allow_major_version_upgrade}
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// The AWS Identity and Access Management (IAM) roles associated with the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#associated_roles RdsDbInstance#associated_roles}
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// Enables replication of automated backups to a different Amazon Web Services Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#automatic_backup_replication_region RdsDbInstance#automatic_backup_replication_region}
	AutomaticBackupReplicationRegion *string `field:"optional" json:"automaticBackupReplicationRegion" yaml:"automaticBackupReplicationRegion"`
	// A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#auto_minor_version_upgrade RdsDbInstance#auto_minor_version_upgrade}
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Availability Zone (AZ) where the database will be created. For information on AWS Regions and Availability Zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#availability_zone RdsDbInstance#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days for which automated backups are retained.
	//
	// Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#backup_retention_period RdsDbInstance#backup_retention_period}
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// The identifier of the CA certificate for this DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#ca_certificate_identifier RdsDbInstance#ca_certificate_identifier}
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// Returns the details of the DB instance's server certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#certificate_details RdsDbInstance#certificate_details}
	CertificateDetails *RdsDbInstanceCertificateDetails `field:"optional" json:"certificateDetails" yaml:"certificateDetails"`
	// A value that indicates whether the DB instance is restarted when you rotate your SSL/TLS certificate.
	//
	// By default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate is not updated until the DB instance is restarted.
	// If you are using SSL/TLS to connect to the DB instance, follow the appropriate instructions for your DB engine to rotate your SSL/TLS certificate
	// This setting doesn't apply to RDS Custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#certificate_rotation_restart RdsDbInstance#certificate_rotation_restart}
	CertificateRotationRestart interface{} `field:"optional" json:"certificateRotationRestart" yaml:"certificateRotationRestart"`
	// For supported engines, indicates that the DB instance should be associated with the specified character set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#character_set_name RdsDbInstance#character_set_name}
	CharacterSetName *string `field:"optional" json:"characterSetName" yaml:"characterSetName"`
	// A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance.
	//
	// By default, tags are not copied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#copy_tags_to_snapshot RdsDbInstance#copy_tags_to_snapshot}
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.
	//
	// The instance profile must meet the following requirements:
	//  * The profile must exist in your account.
	//  * The profile must have an IAM role that Amazon EC2 has permissions to assume.
	//  * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom .
	// For the list of permissions required for the IAM role, see Configure IAM and your VPC in the Amazon RDS User Guide .
	//
	// This setting is required for RDS Custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#custom_iam_instance_profile RdsDbInstance#custom_iam_instance_profile}
	CustomIamInstanceProfile *string `field:"optional" json:"customIamInstanceProfile" yaml:"customIamInstanceProfile"`
	// The identifier of the DB cluster that the instance will belong to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_cluster_identifier RdsDbInstance#db_cluster_identifier}
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The identifier for the RDS for MySQL Multi-AZ DB cluster snapshot to restore from.
	//
	// For more information on Multi-AZ DB clusters, see Multi-AZ deployments with two readable standby DB instances in the Amazon RDS User Guide .
	//
	// Constraints:
	//  * Must match the identifier of an existing Multi-AZ DB cluster snapshot.
	//  * Can't be specified when DBSnapshotIdentifier is specified.
	//  * Must be specified when DBSnapshotIdentifier isn't specified.
	//  * If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot.
	//  * Can't be the identifier of an Aurora DB cluster snapshot.
	//  * Can't be the identifier of an RDS for PostgreSQL Multi-AZ DB cluster snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_cluster_snapshot_identifier RdsDbInstance#db_cluster_snapshot_identifier}
	DbClusterSnapshotIdentifier *string `field:"optional" json:"dbClusterSnapshotIdentifier" yaml:"dbClusterSnapshotIdentifier"`
	// The compute and memory capacity of the DB instance, for example, db.m4.large. Not all DB instance classes are available in all AWS Regions, or for all database engines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_instance_class RdsDbInstance#db_instance_class}
	DbInstanceClass *string `field:"optional" json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_instance_identifier RdsDbInstance#db_instance_identifier}
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The meaning of this parameter differs according to the database engine you use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_name RdsDbInstance#db_name}
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// The name of an existing DB parameter group or a reference to an AWS::RDS::DBParameterGroup resource created in the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_parameter_group_name RdsDbInstance#db_parameter_group_name}
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// A list of the DB security groups to assign to the DB instance.
	//
	// The list can include both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup resources created in the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_security_groups RdsDbInstance#db_security_groups}
	DbSecurityGroups *[]*string `field:"optional" json:"dbSecurityGroups" yaml:"dbSecurityGroups"`
	// The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance.
	//
	// If you're restoring from a shared manual DB snapshot, you must specify the ARN of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_snapshot_identifier RdsDbInstance#db_snapshot_identifier}
	DbSnapshotIdentifier *string `field:"optional" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// A DB subnet group to associate with the DB instance.
	//
	// If you update this value, the new subnet group must be a subnet group in a new VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#db_subnet_group_name RdsDbInstance#db_subnet_group_name}
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Indicates whether the DB instance has a dedicated log volume (DLV) enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#dedicated_log_volume RdsDbInstance#dedicated_log_volume}
	DedicatedLogVolume interface{} `field:"optional" json:"dedicatedLogVolume" yaml:"dedicatedLogVolume"`
	// A value that indicates whether to remove automated backups immediately after the DB instance is deleted.
	//
	// This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#delete_automated_backups RdsDbInstance#delete_automated_backups}
	DeleteAutomatedBackups interface{} `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// A value that indicates whether the DB instance has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#deletion_protection RdsDbInstance#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	//
	// Currently, only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#domain RdsDbInstance#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The ARN for the Secrets Manager secret with the credentials for the user joining the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#domain_auth_secret_arn RdsDbInstance#domain_auth_secret_arn}
	DomainAuthSecretArn *string `field:"optional" json:"domainAuthSecretArn" yaml:"domainAuthSecretArn"`
	// The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#domain_dns_ips RdsDbInstance#domain_dns_ips}
	DomainDnsIps *[]*string `field:"optional" json:"domainDnsIps" yaml:"domainDnsIps"`
	// The fully qualified domain name (FQDN) of an Active Directory domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#domain_fqdn RdsDbInstance#domain_fqdn}
	DomainFqdn *string `field:"optional" json:"domainFqdn" yaml:"domainFqdn"`
	// Specify the name of the IAM role to be used when making API calls to the Directory Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#domain_iam_role_name RdsDbInstance#domain_iam_role_name}
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// The Active Directory organizational unit for your DB instance to join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#domain_ou RdsDbInstance#domain_ou}
	DomainOu *string `field:"optional" json:"domainOu" yaml:"domainOu"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#enable_cloudwatch_logs_exports RdsDbInstance#enable_cloudwatch_logs_exports}
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#enable_iam_database_authentication RdsDbInstance#enable_iam_database_authentication}
	EnableIamDatabaseAuthentication interface{} `field:"optional" json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// A value that indicates whether to enable Performance Insights for the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#enable_performance_insights RdsDbInstance#enable_performance_insights}
	EnablePerformanceInsights interface{} `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// Specifies the connection endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#endpoint RdsDbInstance#endpoint}
	Endpoint *RdsDbInstanceEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The name of the database engine that you want to use for this DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#engine RdsDbInstance#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The version number of the database engine to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#engine_version RdsDbInstance#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#iops RdsDbInstance#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The ARN of the AWS Key Management Service (AWS KMS) master key that's used to encrypt the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#kms_key_id RdsDbInstance#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// License model information for this DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#license_model RdsDbInstance#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// A value that indicates whether to manage the master user password with AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#manage_master_user_password RdsDbInstance#manage_master_user_password}
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// The master user name for the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#master_username RdsDbInstance#master_username}
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The password for the master user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#master_user_password RdsDbInstance#master_user_password}
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// Contains the secret managed by RDS in AWS Secrets Manager for the master user password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#master_user_secret RdsDbInstance#master_user_secret}
	MasterUserSecret *RdsDbInstanceMasterUserSecret `field:"optional" json:"masterUserSecret" yaml:"masterUserSecret"`
	// The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#max_allocated_storage RdsDbInstance#max_allocated_storage}
	MaxAllocatedStorage *float64 `field:"optional" json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
	//
	// To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#monitoring_interval RdsDbInstance#monitoring_interval}
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#monitoring_role_arn RdsDbInstance#monitoring_role_arn}
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// Specifies whether the database instance is a multiple Availability Zone deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#multi_az RdsDbInstance#multi_az}
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// The name of the NCHAR character set for the Oracle DB instance. This parameter doesn't apply to RDS Custom.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#nchar_character_set_name RdsDbInstance#nchar_character_set_name}
	NcharCharacterSetName *string `field:"optional" json:"ncharCharacterSetName" yaml:"ncharCharacterSetName"`
	// The network type of the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#network_type RdsDbInstance#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Indicates that the DB instance should be associated with the specified option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#option_group_name RdsDbInstance#option_group_name}
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The KMS key ID is the Amazon Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#performance_insights_kms_key_id RdsDbInstance#performance_insights_kms_key_id}
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// The amount of time, in days, to retain Performance Insights data. Valid values are 7 or 731 (2 years).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#performance_insights_retention_period RdsDbInstance#performance_insights_retention_period}
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// The port number on which the database accepts connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#port RdsDbInstance#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#preferred_backup_window RdsDbInstance#preferred_backup_window}
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// he weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#preferred_maintenance_window RdsDbInstance#preferred_maintenance_window}
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#processor_features RdsDbInstance#processor_features}
	ProcessorFeatures interface{} `field:"optional" json:"processorFeatures" yaml:"processorFeatures"`
	// A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#promotion_tier RdsDbInstance#promotion_tier}
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
	// Indicates whether the DB instance is an internet-facing instance.
	//
	// If you specify true, AWS CloudFormation creates an instance with a publicly resolvable DNS name, which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#publicly_accessible RdsDbInstance#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The open mode of an Oracle read replica. The default is open-read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#replica_mode RdsDbInstance#replica_mode}
	ReplicaMode *string `field:"optional" json:"replicaMode" yaml:"replicaMode"`
	// The date and time to restore from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#restore_time RdsDbInstance#restore_time}
	RestoreTime *string `field:"optional" json:"restoreTime" yaml:"restoreTime"`
	// The identifier of the Multi-AZ DB cluster that will act as the source for the read replica.
	//
	// Each DB cluster can have up to 15 read replicas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#source_db_cluster_identifier RdsDbInstance#source_db_cluster_identifier}
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The Amazon Resource Name (ARN) of the replicated automated backups from which to restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#source_db_instance_automated_backups_arn RdsDbInstance#source_db_instance_automated_backups_arn}
	SourceDbInstanceAutomatedBackupsArn *string `field:"optional" json:"sourceDbInstanceAutomatedBackupsArn" yaml:"sourceDbInstanceAutomatedBackupsArn"`
	// If you want to create a Read Replica DB instance, specify the ID of the source DB instance.
	//
	// Each DB instance can have a limited number of Read Replicas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#source_db_instance_identifier RdsDbInstance#source_db_instance_identifier}
	SourceDbInstanceIdentifier *string `field:"optional" json:"sourceDbInstanceIdentifier" yaml:"sourceDbInstanceIdentifier"`
	// The resource ID of the source DB instance from which to restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#source_dbi_resource_id RdsDbInstance#source_dbi_resource_id}
	SourceDbiResourceId *string `field:"optional" json:"sourceDbiResourceId" yaml:"sourceDbiResourceId"`
	// The ID of the region that contains the source DB instance for the Read Replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#source_region RdsDbInstance#source_region}
	SourceRegion *string `field:"optional" json:"sourceRegion" yaml:"sourceRegion"`
	// A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#storage_encrypted RdsDbInstance#storage_encrypted}
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Specifies the storage throughput for the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#storage_throughput RdsDbInstance#storage_throughput}
	StorageThroughput *float64 `field:"optional" json:"storageThroughput" yaml:"storageThroughput"`
	// Specifies the storage type to be associated with the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#storage_type RdsDbInstance#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Tags to assign to the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#tags RdsDbInstance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ARN from the key store with which to associate the instance for TDE encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#tde_credential_arn RdsDbInstance#tde_credential_arn}
	TdeCredentialArn *string `field:"optional" json:"tdeCredentialArn" yaml:"tdeCredentialArn"`
	// The password for the given ARN from the key store in order to access the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#tde_credential_password RdsDbInstance#tde_credential_password}
	TdeCredentialPassword *string `field:"optional" json:"tdeCredentialPassword" yaml:"tdeCredentialPassword"`
	// The time zone of the DB instance. The time zone parameter is currently supported only by Microsoft SQL Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#timezone RdsDbInstance#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// A value that indicates whether the DB instance class of the DB instance uses its default processor features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#use_default_processor_features RdsDbInstance#use_default_processor_features}
	UseDefaultProcessorFeatures interface{} `field:"optional" json:"useDefaultProcessorFeatures" yaml:"useDefaultProcessorFeatures"`
	// A value that indicates whether the DB instance is restored from the latest backup time.
	//
	// By default, the DB instance isn't restored from the latest backup time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#use_latest_restorable_time RdsDbInstance#use_latest_restorable_time}
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of the VPC security group IDs to assign to the DB instance.
	//
	// The list can include both the physical IDs of existing VPC security groups and references to AWS::EC2::SecurityGroup resources created in the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance#vpc_security_groups RdsDbInstance#vpc_security_groups}
	VpcSecurityGroups *[]*string `field:"optional" json:"vpcSecurityGroups" yaml:"vpcSecurityGroups"`
}

