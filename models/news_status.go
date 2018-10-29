package models

// NewsStatus is enumeration for representing news status
type NewsStatus byte

const (
	NewsStatusPublished NewsStatus = iota
	NewsStatusDraft
	NewsStatusDeleted
)

func (ns NewsStatus) String() string {
	return []string{"published", "draft", "deleted"}[ns]
}
