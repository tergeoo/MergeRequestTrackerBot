package model

type Webhook struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	Project struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       any    `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         any    `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
		CiConfigPath      string `json:"ci_config_path"`
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		AssigneeID     any    `json:"assignee_id"`
		AuthorID       int    `json:"author_id"`
		Description    string `json:"description"`
		Draft          bool   `json:"draft"`
		HeadPipelineID any    `json:"head_pipeline_id"`
		ID             int    `json:"id"`
		Iid            int    `json:"iid"`
		LastEditedAt   any    `json:"last_edited_at"`
		LastEditedByID any    `json:"last_edited_by_id"`
		MergeCommitSha string `json:"merge_commit_sha"`
		MergeError     any    `json:"merge_error"`
		Note           string `json:"note"`
		MergeParams    struct {
			ForceRemoveSourceBranch  string `json:"force_remove_source_branch"`
			ShouldRemoveSourceBranch bool   `json:"should_remove_source_branch"`
		} `json:"merge_params"`
		MergeStatus               string `json:"merge_status"`
		MergeUserID               any    `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool   `json:"merge_when_pipeline_succeeds"`
		MilestoneID               any    `json:"milestone_id"`
		SourceBranch              string `json:"source_branch"`
		SourceProjectID           int    `json:"source_project_id"`
		StateID                   int    `json:"state_id"`
		TargetBranch              string `json:"target_branch"`
		TargetProjectID           int    `json:"target_project_id"`
		Title                     string `json:"title"`
		UpdatedByID               any    `json:"updated_by_id"`
		URL                       string `json:"url"`
		Source                    struct {
			ID                int    `json:"id"`
			Name              string `json:"name"`
			Description       any    `json:"description"`
			WebURL            string `json:"web_url"`
			AvatarURL         any    `json:"avatar_url"`
			GitSSHURL         string `json:"git_ssh_url"`
			GitHTTPURL        string `json:"git_http_url"`
			Namespace         string `json:"namespace"`
			VisibilityLevel   int    `json:"visibility_level"`
			PathWithNamespace string `json:"path_with_namespace"`
			DefaultBranch     string `json:"default_branch"`
			CiConfigPath      string `json:"ci_config_path"`
			Homepage          string `json:"homepage"`
			URL               string `json:"url"`
			SSHURL            string `json:"ssh_url"`
			HTTPURL           string `json:"http_url"`
		} `json:"source"`
		Target struct {
			ID                int    `json:"id"`
			Name              string `json:"name"`
			Description       any    `json:"description"`
			WebURL            string `json:"web_url"`
			AvatarURL         any    `json:"avatar_url"`
			GitSSHURL         string `json:"git_ssh_url"`
			GitHTTPURL        string `json:"git_http_url"`
			Namespace         string `json:"namespace"`
			VisibilityLevel   int    `json:"visibility_level"`
			PathWithNamespace string `json:"path_with_namespace"`
			DefaultBranch     string `json:"default_branch"`
			CiConfigPath      string `json:"ci_config_path"`
			Homepage          string `json:"homepage"`
			URL               string `json:"url"`
			SSHURL            string `json:"ssh_url"`
			HTTPURL           string `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID      string `json:"id"`
			Message string `json:"message"`
			Title   string `json:"title"`
			URL     string `json:"url"`
			Author  struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress              bool   `json:"work_in_progress"`
		AssigneeIds                 []any  `json:"assignee_ids"`
		ReviewerIds                 []any  `json:"reviewer_ids"`
		Labels                      []any  `json:"labels"`
		State                       string `json:"state"`
		BlockingDiscussionsResolved bool   `json:"blocking_discussions_resolved"`
		FirstContribution           bool   `json:"first_contribution"`
		DetailedMergeStatus         string `json:"detailed_merge_status"`
	} `json:"object_attributes"`
	Labels  []any `json:"labels"`
	Changes struct {
	} `json:"changes"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description any    `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	MergeRequest struct {
		AssigneeID     any    `json:"assignee_id"`
		AuthorID       int    `json:"author_id"`
		Description    string `json:"description"`
		Draft          bool   `json:"draft"`
		HeadPipelineID any    `json:"head_pipeline_id"`
		ID             int    `json:"id"`
		Iid            int    `json:"iid"`
		LastEditedAt   any    `json:"last_edited_at"`
		LastEditedByID any    `json:"last_edited_by_id"`
		MergeCommitSha any    `json:"merge_commit_sha"`
		MergeError     any    `json:"merge_error"`
		MergeParams    struct {
		} `json:"merge_params"`
		MergeStatus               string `json:"merge_status"`
		MergeUserID               any    `json:"merge_user_id"`
		MergeWhenPipelineSucceeds bool   `json:"merge_when_pipeline_succeeds"`
		MilestoneID               any    `json:"milestone_id"`
		SourceBranch              string `json:"source_branch"`
		SourceProjectID           int    `json:"source_project_id"`
		StateID                   int    `json:"state_id"`
		TargetBranch              string `json:"target_branch"`
		TargetProjectID           int    `json:"target_project_id"`
		Title                     string `json:"title"`
		UpdatedByID               int    `json:"updated_by_id"`
		URL                       string `json:"url"`
	} `json:"merge_request"`
}

func (receiver Webhook) CreateMergeRequestMessageFromMR() MergeRequest {
	var state = OPENED

	if receiver.ObjectAttributes.State == "approved" {
		state = APPROVED
	} else if len(receiver.ObjectAttributes.Note) > 0 {
		state = COMMENTED
	} else if receiver.ObjectAttributes.State == "merged" {
		state = MERGED
	} else if receiver.ObjectAttributes.State == "closed" {
		state = CLOSED
	}

	return MergeRequest{
		Id:             receiver.ObjectAttributes.Iid,
		Title:          receiver.ObjectAttributes.Title,
		Link:           receiver.ObjectAttributes.URL,
		Status:         receiver.ObjectAttributes.State,
		AuthorUserName: receiver.User.Username,
		HasComments:    len(receiver.ObjectAttributes.Note) > 0,
		State:          state,
	}
}
func (receiver Webhook) CreateMergeRequestMessageFromComment() MergeRequest {
	var state = OPENED

	if receiver.ObjectAttributes.DetailedMergeStatus == "approved" {
		state = APPROVED
	} else if len(receiver.ObjectAttributes.Note) > 0 {
		state = COMMENTED
	} else if receiver.ObjectAttributes.DetailedMergeStatus == "merged" {
		state = MERGED
	} else if receiver.ObjectAttributes.DetailedMergeStatus == "closed" {
		state = CLOSED
	}

	return MergeRequest{
		Id:             receiver.MergeRequest.Iid,
		Title:          receiver.MergeRequest.Title,
		Link:           receiver.MergeRequest.URL,
		Status:         receiver.MergeRequest.MergeStatus,
		AuthorUserName: receiver.User.Username,
		HasComments:    len(receiver.ObjectAttributes.Note) > 0,
		State:          state,
	}
}
