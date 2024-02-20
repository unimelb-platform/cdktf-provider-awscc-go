package codeguruprofilerprofilinggroup


type CodeguruprofilerProfilingGroupTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	Value *string `field:"required" json:"value" yaml:"value"`
}

