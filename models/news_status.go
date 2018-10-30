package models

// NewsStatus is enumeration for representing news status
type NewsStatus string

const (
	NewsStatusPublished NewsStatus = "published"
	NewsStatusDraft     NewsStatus = "draft"
	NewsStatusDeleted   NewsStatus = "deleted"
)
