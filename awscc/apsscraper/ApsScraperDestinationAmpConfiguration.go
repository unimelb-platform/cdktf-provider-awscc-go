package apsscraper


type ApsScraperDestinationAmpConfiguration struct {
	// ARN of an Amazon Managed Prometheus workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#workspace_arn ApsScraper#workspace_arn}
	WorkspaceArn *string `field:"optional" json:"workspaceArn" yaml:"workspaceArn"`
}

