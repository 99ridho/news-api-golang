package models

// NewsStatus is enumeration for representing news status
type NewsStatus byte

const (
	Published NewsStatus = iota
	Draft
	Deleted
)

func (ns NewsStatus) String() string {
	return []string{"published", "draft", "deleted"}[ns]
}
