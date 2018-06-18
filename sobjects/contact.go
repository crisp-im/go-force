package sobjects

type Contact struct {
	BaseSObject
	Email             string `force:",omitempty"`
	FirstName         string `force:",omitempty"`
	LastName          string `force:",omitempty"`
	Phone             string `force:",omitempty"`
	PhotoUrl          string `force:",omitempty"`
	Title             string `force:",omitempty"`
	LeadSource        string `force:",omitempty"`
	MailingStreet     string `force:",omitempty"`
	MailingCity       string `force:",omitempty"`
	MailingCountry    string `force:",omitempty"`
	MailingState      string `force:",omitempty"`
	MailingPostalCode string `force:",omitempty"`
}

func (t *Contact) ApiName() string {
	return "Contact"
}

type ContactQueryResponse struct {
	BaseQuery
	Records []Contact `json:"Records" force:"records"`
}