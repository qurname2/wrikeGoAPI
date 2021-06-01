package wrikeGoAPI

type Folder struct {
	Kind string `json:"kind"`
	Data []struct {
		ID             string        `json:"id"`
		AccountID      string        `json:"accountId"`
		Title          string        `json:"title"`
		CreatedDate    string        `json:"createdAt"`
		UpdatedDate    string        `json:"updatedAt"`
		Description    string        `json:"description"`
		SharedIds      []string      `json:"sharedIds,omitempty"`
		ParentIds      []string      `json:"parentIds,omitempty"`
		ChildIds       []string      `json:"childIds,omitempty"`
		SuperParentIds []string      `json:"superParentIds,omitempty"`
		Scope          string        `json:"scope"`
		HasAttachments bool          `json:"hasAttachments"`
		Permalink      string        `json:"permalink"`
		WorkflowID     string        `json:"workflowId"`
		Metadata       []Metadata    `json:"metadata,omitempty"`
		CustomFields   []CustomField `json:"customFields"`
		Project        Project       `json:"project"`
	} `json:"data,omitempty"`
}

// TasksResponse is Response from /tasks query
type GetFolderApiIDByPermalinkParams struct {
	permalink string
}

// GetFolderApiIDByPermalink from id, see Wrike API: https://developers.wrike.com/documentation/api/methods/get-folder
func (s *FolderService) GetFolderApiIDByPermalink(perm string) (*Folder, *Response, error) {

	var params = GetFolderApiIDByPermalinkParams{perm}

	req, err := s.client.NewRequest("GET", foldersUrl, params)
	if err != nil {
		return nil, nil, err
	}

	newFolderObject := new(Folder)
	resp, err := s.client.Do(req, newFolderObject)

	if err != nil {
		return nil, resp, err
	}
	return newFolderObject, resp, err
}

func (s *FolderService) GetFolderApiID(params *GetFolderTree) (*Folder, *Response, error) {

	req, err := s.client.NewRequest("GET", foldersUrl, params)
	if err != nil {
		return nil, nil, err
	}

	newFolderObject := new(Folder)
	resp, err := s.client.Do(req, newFolderObject)

	if err != nil {
		return nil, resp, err
	}
	return newFolderObject, resp, err
}

// GetFolderTree contains parameters that will be passed to GetFolderTree API.
type GetFolderTree struct {
	Permalink   string            `url:"permalink,omitempty"`
	Descendants *bool              `url:"descendants,omitempty"`
	Metadata    *Metadata          `url:"metadata,omitempty"`
	//CustomField *CustomFieldFilter `url:"customField,omitempty"`
	//UpdatedDate *DateRange         `url:"updatedDate,omitempty"`
	Project     *bool              `url:"project,omitempty"`
	Deleted     *bool              `url:"deleted,omitempty"`
	//Fields      *FieldSet          `url:"fields,omitempty"`
}