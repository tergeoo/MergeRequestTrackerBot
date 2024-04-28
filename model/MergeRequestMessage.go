package model

type MergeRequestMessage struct {
	MessageId, MergeRequestId           int
	Title, Link, Status, AuthorUserName string
	HasComments                         bool
}

type Saved struct {
	Messages map[int]MergeRequestMessage
}
