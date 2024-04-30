package model

type MergeRequest struct {
	Id                                  int
	Title, Link, Status, AuthorUserName string
	HasComments                         bool
	State                               MergeRequestState
}

// ключ id mr
// значение id сообщения в тг
type Saved struct {
	MessageIds map[int]int `json:"message_ids"`
}
