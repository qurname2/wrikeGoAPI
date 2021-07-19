package wrikeGoAPI

type Contact struct {
	Kind string `json:"kind"`
	Data []struct {
		ID          string     `json:"id"`
		FirstName   string     `json:"firstName"`
		LastName    string     `json:"lastName"`
		Type        string     `json:"type"`
		UpdatedDate string     `json:"updatedAt"`
		Description string     `json:"description"`
		Profiles    []Profiles `json:"profiles"`
		AvatarUrl   string     `json:"AvatarUrl"`
		Timezone    string     `json:"timezone"`
		Locale      string     `json:"locale"`
		Deleted     bool       `json:"deleted"`
		Uid         string     `json:"uid"`
	} `json:"data,omitempty"`
}

// Profiles struct
type Profiles struct {
	Email string `json:"email"`
}

//GetContacts get list of all Users
func (cl *ContactService) GetContacts() (*Contact, *Response, error) {

	req, err := cl.client.NewRequest("GET", contactsUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	newContactObject := new(Contact)
	resp, err := cl.client.Do(req, newContactObject)
	if err != nil {
		return nil, resp, err
	}

	return newContactObject, resp, err
}
