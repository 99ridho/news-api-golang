package models

type FetchNewsParam struct {
	Status   NewsStatus
	TopicIDs []int64
	*Pagination
}
