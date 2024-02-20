package mskreplicator


type MskReplicatorReplicationInfoListTopicReplication struct {
	// List of regular expression patterns indicating the topics to copy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#topics_to_replicate MskReplicator#topics_to_replicate}
	TopicsToReplicate *[]*string `field:"required" json:"topicsToReplicate" yaml:"topicsToReplicate"`
	// Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#copy_access_control_lists_for_topics MskReplicator#copy_access_control_lists_for_topics}
	CopyAccessControlListsForTopics interface{} `field:"optional" json:"copyAccessControlListsForTopics" yaml:"copyAccessControlListsForTopics"`
	// Whether to periodically configure remote topics to match their corresponding upstream topics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#copy_topic_configurations MskReplicator#copy_topic_configurations}
	CopyTopicConfigurations interface{} `field:"optional" json:"copyTopicConfigurations" yaml:"copyTopicConfigurations"`
	// Whether to periodically check for new topics and partitions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#detect_and_copy_new_topics MskReplicator#detect_and_copy_new_topics}
	DetectAndCopyNewTopics interface{} `field:"optional" json:"detectAndCopyNewTopics" yaml:"detectAndCopyNewTopics"`
	// List of regular expression patterns indicating the topics that should not be replicated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#topics_to_exclude MskReplicator#topics_to_exclude}
	TopicsToExclude *[]*string `field:"optional" json:"topicsToExclude" yaml:"topicsToExclude"`
}

