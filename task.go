package wrikeGoAPI

import (
	"fmt"
)

// Task is a task
type Task struct {
	ID             string `json:"id"`
	AccountID      string `json:"accountId"`
	Title          string `json:"title"`
	Status         string `json:"status"`
	Importance     string `json:"importance"`
	CreatedDate    Time   `json:"createdDate"`
	UpdatedDate    Time   `json:"updatedDate"`
	Dates          Dates  `json:"dates"`
	Scope          string `json:"scope"`
	CustomStatusID string `json:"customStatusId"`
	Permalink      string `json:"permalink"`
	Priority       string `json:"priority"`
	Description    string `json:"description"`
}

// Dates are dates
type Dates struct {
	Type     string `json:"type"`
	Duration int    `json:"duration"`
	Start    string `json:"start"`
	Due      string `json:"due"`
}

// TasksResponse is Response from /tasks query
type TaskResponse struct {
	Kind string         `json:"kind"`
	Data []DetailedTask `json:"data,omitempty"`
}

// DetailedTask represents task return fron /tasks/{id} endpoint
type DetailedTask struct {
	ID               string        `json:"id"`
	AccountID        string        `json:"accountID"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	BriefDescription string        `json:"briefDescription"`
	ParentIDs        []string      `json:"parentIds"`
	SuperParentIDs   []string      `json:"superParentIds"`
	SharedIDs        []string      `json:"sharedIds"`
	ResponsibleIDs   []string      `json:"responsibleIds"`
	Status           string        `json:"status"`
	Importance       string        `json:"importance"`
	CreatedDate      Time          `json:"createdDate"`
	UpdatedDate      Time          `json:"updatedDate"`
	CompletedDate    Time          `json:"completedDate"`
	Dates            Dates         `json:"dates"`
	Scope            string        `json:"scope"`
	AuthorIds        []string      `json:"authorIds"`
	CustomStatusID   string        `json:"customStatusId"`
	HasAttachments   bool          `json:"hasAttachments"`
	Permalink        string        `json:"permalink"`
	Priority         string        `json:"priority"`
	FollowedByMe     bool          `json:"followedByMe"`
	FollowerIDs      []string      `json:"followerIds"`
	Recurrent        bool          `json:"recurrent"`
	SuperTaskIDs     []string      `json:"superTaskIds"`
	SubTaskIDs       []string      `json:"subTaskIds"`
	DependencyIDs    []string      `json:"dependencyIds"`
	Metadata         []Metadata    `json:"metadata"`
	CustomFields     []CustomField `json:"customFields"`
	BillingType      string        `json:"billingType"`
}

//UpdateTask update task by id
func (t *TaskService) UpdateTask(id string, params *UpdateTask) (*TaskResponse, *Response, error) {
	path := fmt.Sprintf("%s/%s", tasksUrl, id)
	req, err := t.client.NewRequest("PUT", path, params)
	if err != nil {
		return nil, nil, err
	}

	updateTaskResponse := new(TaskResponse)
	response, err := t.client.Do(req, updateTaskResponse)
	if err != nil {
		return nil, response, err
	}

	return updateTaskResponse, response, err
}

//CreateTask will create a task for you
func (t *TaskService) CreateTask(folderApiID string, params *CreateTask) (*TaskResponse, *Response, error) {
	path := fmt.Sprintf("%v/%v/tasks", foldersUrl, folderApiID)
	req, err := t.client.NewRequest("POST", path, params)
	if err != nil {
		return nil, nil, err
	}
	//fmt.Printf("CreateTask is: %v", req)
	taskResponse := new(TaskResponse)
	response, err := t.client.Do(req, taskResponse)
	if err != nil {
		return nil, response, err
	}

	return taskResponse, response, err
}

// CreateTask contains parameters that will be passed to GetFolderTree API.
type CreateTask struct {
	Title       string    `url:"title,omitempty"`
	Description string    `url:"description,omitempty"`
	Metadata    *Metadata `url:"metadata,omitempty"`
	Project     *bool     `url:"project,omitempty"`
	Deleted     *bool     `url:"deleted,omitempty"`
}

// UpdateTask contains parameters that will be passed to GetFolderTree API.
type UpdateTask struct {
	Title              *string        `url:"title,omitempty"`
	Description        *string        `url:"description,omitempty"`
	Importance         TaskImportance `url:"importance,omitempty"`
	AddParents         FolderIDSet    `url:"addParents,omitempty"`
	RemoveParents      FolderIDSet    `url:"removeParents,omitempty"`
	AddShareds         ContactIDSet   `url:"addShareds,omitempty"`
	RemoveShareds      ContactIDSet   `url:"removeShareds,omitempty"`
	AddResponsibles    string         `url:"addResponsibles,omitempty"`
	RemoveResponsibles ContactIDSet   `url:"removeResponsibles,omitempty"`
	AddFollowers       ContactIDSet   `url:"addFollowers,omitempty"`
	Follow             *bool          `url:"follow,omitempty"`
	PriorityBefore     TaskID         `url:"priorityBefore,omitempty"`
	PriorityAfter      TaskID         `url:"priorityAfter,omitempty"`
	CustomStatus       CustomStatusID `url:"customStatus,omitempty"`
	Restore            *bool          `url:"restore,omitempty"`
}

// ContactID is a string that represents a Wrike Contact ID
type ContactID string

type ContactIDSet []ContactID

// ContactID is a string that represents a Wrike Contact ID
type FolderID string

type FolderIDSet []FolderID

type TaskID string

type CustomStatusID string

type TaskImportance string
