package rdsdbcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsDbClusterConfig struct {
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
	// The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#allocated_storage RdsDbCluster#allocated_storage}
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#associated_roles RdsDbCluster#associated_roles}
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// A value that indicates whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.
	//
	// By default, minor engine upgrades are applied automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#auto_minor_version_upgrade RdsDbCluster#auto_minor_version_upgrade}
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// A list of Availability Zones (AZs) where instances in the DB cluster can be created.
	//
	// For information on AWS Regions and Availability Zones, see Choosing the Regions and Availability Zones in the Amazon Aurora User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#availability_zones RdsDbCluster#availability_zones}
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The target backtrack window, in seconds. To disable backtracking, set this value to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#backtrack_window RdsDbCluster#backtrack_window}
	BacktrackWindow *float64 `field:"optional" json:"backtrackWindow" yaml:"backtrackWindow"`
	// The number of days for which automated backups are retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#backup_retention_period RdsDbCluster#backup_retention_period}
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster.
	//
	// The default is not to copy them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#copy_tags_to_snapshot RdsDbCluster#copy_tags_to_snapshot}
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The name of your database.
	//
	// If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see Naming Constraints in the Amazon RDS User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#database_name RdsDbCluster#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The DB cluster identifier. This parameter is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#db_cluster_identifier RdsDbCluster#db_cluster_identifier}
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example db.m6g.xlarge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#db_cluster_instance_class RdsDbCluster#db_cluster_instance_class}
	DbClusterInstanceClass *string `field:"optional" json:"dbClusterInstanceClass" yaml:"dbClusterInstanceClass"`
	// The name of the DB cluster parameter group to associate with this DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#db_cluster_parameter_group_name RdsDbCluster#db_cluster_parameter_group_name}
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// The name of the DB parameter group to apply to all instances of the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#db_instance_parameter_group_name RdsDbCluster#db_instance_parameter_group_name}
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// A DB subnet group that you want to associate with this DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#db_subnet_group_name RdsDbCluster#db_subnet_group_name}
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Reserved for future use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#db_system_id RdsDbCluster#db_system_id}
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// A value that indicates whether the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#deletion_protection RdsDbCluster#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB cluster in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#domain RdsDbCluster#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Specify the name of the IAM role to be used when making API calls to the Directory Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#domain_iam_role_name RdsDbCluster#domain_iam_role_name}
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	//
	// The values in the list depend on the DB engine being used. For more information, see Publishing Database Logs to Amazon CloudWatch Logs in the Amazon Aurora User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#enable_cloudwatch_logs_exports RdsDbCluster#enable_cloudwatch_logs_exports}
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database).
	//
	// By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#enable_global_write_forwarding RdsDbCluster#enable_global_write_forwarding}
	EnableGlobalWriteForwarding interface{} `field:"optional" json:"enableGlobalWriteForwarding" yaml:"enableGlobalWriteForwarding"`
	// A value that indicates whether to enable the HTTP endpoint for DB cluster.
	//
	// By default, the HTTP endpoint is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#enable_http_endpoint RdsDbCluster#enable_http_endpoint}
	EnableHttpEndpoint interface{} `field:"optional" json:"enableHttpEndpoint" yaml:"enableHttpEndpoint"`
	// A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	//
	// By default, mapping is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#enable_iam_database_authentication RdsDbCluster#enable_iam_database_authentication}
	EnableIamDatabaseAuthentication interface{} `field:"optional" json:"enableIamDatabaseAuthentication" yaml:"enableIamDatabaseAuthentication"`
	// The name of the database engine to be used for this DB cluster.
	//
	// Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora), and aurora-postgresql
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#engine RdsDbCluster#engine}
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The DB engine mode of the DB cluster, either provisioned, serverless, parallelquery, global, or multimaster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#engine_mode RdsDbCluster#engine_mode}
	EngineMode *string `field:"optional" json:"engineMode" yaml:"engineMode"`
	// The version number of the database engine to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#engine_version RdsDbCluster#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster.
	//
	// To define the primary database cluster of the global cluster, use the AWS::RDS::GlobalCluster resource.
	//
	// If you aren't configuring a global database cluster, don't specify this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#global_cluster_identifier RdsDbCluster#global_cluster_identifier}
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#iops RdsDbCluster#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The Amazon Resource Name (ARN) of the AWS Key Management Service master key that is used to encrypt the database instances in the DB cluster, such as arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	// If you enable the StorageEncrypted property but don't specify this property, the default master key is used. If you specify this property, you must set the StorageEncrypted property to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#kms_key_id RdsDbCluster#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A value that indicates whether to manage the master user password with AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#manage_master_user_password RdsDbCluster#manage_master_user_password}
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// The name of the master user for the DB cluster.
	//
	// You must specify MasterUsername, unless you specify SnapshotIdentifier. In that case, don't specify MasterUsername.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#master_username RdsDbCluster#master_username}
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The master password for the DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#master_user_password RdsDbCluster#master_user_password}
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// Contains the secret managed by RDS in AWS Secrets Manager for the master user password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#master_user_secret RdsDbCluster#master_user_secret}
	MasterUserSecret *RdsDbClusterMasterUserSecret `field:"optional" json:"masterUserSecret" yaml:"masterUserSecret"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.
	//
	// To turn off collecting Enhanced Monitoring metrics, specify 0. The default is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#monitoring_interval RdsDbCluster#monitoring_interval}
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#monitoring_role_arn RdsDbCluster#monitoring_role_arn}
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// The network type of the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#network_type RdsDbCluster#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// A value that indicates whether to turn on Performance Insights for the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#performance_insights_enabled RdsDbCluster#performance_insights_enabled}
	PerformanceInsightsEnabled interface{} `field:"optional" json:"performanceInsightsEnabled" yaml:"performanceInsightsEnabled"`
	// The Amazon Web Services KMS key identifier for encryption of Performance Insights data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#performance_insights_kms_key_id RdsDbCluster#performance_insights_kms_key_id}
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// The amount of time, in days, to retain Performance Insights data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#performance_insights_retention_period RdsDbCluster#performance_insights_retention_period}
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// The port number on which the instances in the DB cluster accept connections.
	//
	// Default: 3306 if engine is set as aurora or 5432 if set to aurora-postgresql.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#port RdsDbCluster#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region. To see the time blocks available, see Adjusting the Preferred DB Cluster Maintenance Window in the Amazon Aurora User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#preferred_backup_window RdsDbCluster#preferred_backup_window}
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see Adjusting the Preferred DB Cluster Maintenance Window in the Amazon Aurora User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#preferred_maintenance_window RdsDbCluster#preferred_maintenance_window}
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A value that indicates whether the DB cluster is publicly accessible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#publicly_accessible RdsDbCluster#publicly_accessible}
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#read_endpoint RdsDbCluster#read_endpoint}.
	ReadEndpoint *RdsDbClusterReadEndpoint `field:"optional" json:"readEndpoint" yaml:"readEndpoint"`
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a Read Replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#replication_source_identifier RdsDbCluster#replication_source_identifier}
	ReplicationSourceIdentifier *string `field:"optional" json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// The date and time to restore the DB cluster to.
	//
	// Value must be a time in Universal Coordinated Time (UTC) format. An example: 2015-03-07T23:45:00Z
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#restore_to_time RdsDbCluster#restore_to_time}
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// The type of restore to be performed.
	//
	// You can specify one of the following values:
	// full-copy - The new DB cluster is restored as a full copy of the source DB cluster.
	// copy-on-write - The new DB cluster is restored as a clone of the source DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#restore_type RdsDbCluster#restore_type}
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// The ScalingConfiguration property type specifies the scaling configuration of an Aurora Serverless DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#scaling_configuration RdsDbCluster#scaling_configuration}
	ScalingConfiguration *RdsDbClusterScalingConfiguration `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// Contains the scaling configuration of an Aurora Serverless v2 DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#serverless_v2_scaling_configuration RdsDbCluster#serverless_v2_scaling_configuration}
	ServerlessV2ScalingConfiguration *RdsDbClusterServerlessV2ScalingConfiguration `field:"optional" json:"serverlessV2ScalingConfiguration" yaml:"serverlessV2ScalingConfiguration"`
	// The identifier for the DB snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
	// After you restore a DB cluster with a SnapshotIdentifier property, you must specify the same SnapshotIdentifier property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed. However, if you don't specify the SnapshotIdentifier property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the specified SnapshotIdentifier property, and the original DB cluster is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#snapshot_identifier RdsDbCluster#snapshot_identifier}
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The identifier of the source DB cluster from which to restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#source_db_cluster_identifier RdsDbCluster#source_db_cluster_identifier}
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, us-east-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#source_region RdsDbCluster#source_region}
	SourceRegion *string `field:"optional" json:"sourceRegion" yaml:"sourceRegion"`
	// Indicates whether the DB instance is encrypted.
	//
	// If you specify the DBClusterIdentifier, SnapshotIdentifier, or SourceDBInstanceIdentifier property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#storage_encrypted RdsDbCluster#storage_encrypted}
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Specifies the storage type to be associated with the DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#storage_type RdsDbCluster#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#tags RdsDbCluster#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A value that indicates whether to restore the DB cluster to the latest restorable backup time.
	//
	// By default, the DB cluster is not restored to the latest restorable backup time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#use_latest_restorable_time RdsDbCluster#use_latest_restorable_time}
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of EC2 VPC security groups to associate with this DB cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster#vpc_security_group_ids RdsDbCluster#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

