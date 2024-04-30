package model

type MergeRequestState int

const (
	OPENED MergeRequestState = iota + 1
	COMMENTED
	APPROVED
	MERGED
	CLOSED
)

// String - Creating common behavior - give the type a String function
func (w MergeRequestState) String() string {
	return [...]string{"opened", "commented", "approved", "merged", "closed"}[w-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w MergeRequestState) EnumIndex() int {
	return int(w)
}
