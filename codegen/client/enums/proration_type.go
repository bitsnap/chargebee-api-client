package enums

type ProrationTypeEnum string

const (
	DefaultProration     ProrationTypeEnum = "site_default"
	PartialTermProration                   = "partial_term"
	FullTermProration                      = "full_term"
)
