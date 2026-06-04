package gluecrawler


type GlueCrawlerTargetsDynamoDbTargets struct {
	// The name of the DynamoDB table to crawl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_crawler#path GlueCrawler#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

