package models

type FetchNewsParam struct {
	Status       NewsStatus `query:"status"`
	TopicIDQuery string     `query:"topic_id"`
	TopicIDs     []int64
	*Pagination
}
