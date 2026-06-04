package gluecrawler


type GlueCrawlerRecrawlPolicy struct {
	// Specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run.
	//
	// A value of CRAWL_EVERYTHING specifies crawling the entire dataset again. A value of CRAWL_NEW_FOLDERS_ONLY specifies crawling only folders that were added since the last crawler run. A value of CRAWL_EVENT_MODE specifies crawling only the changes identified by Amazon S3 events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_crawler#recrawl_behavior GlueCrawler#recrawl_behavior}
	RecrawlBehavior *string `field:"optional" json:"recrawlBehavior" yaml:"recrawlBehavior"`
}

