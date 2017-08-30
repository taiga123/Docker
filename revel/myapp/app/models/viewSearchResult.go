package models

type ViewSearchResult struct {
	Number      string
	Name        string
	Affiliation string
	Position    string
	MailAddress string
	PhoneNumber string
	EntryMonth  string
	Birthday    string
}

func (view *ViewSearchResult) SetView(number string, name string, affiliation string, position string, mailAddress string, phoneNumber string, entryMonth string, birthday string) {
	view.Number = number
	view.Name = name
	view.Affiliation = affiliation
	view.Position = position
	view.MailAddress = mailAddress
	view.PhoneNumber = phoneNumber
	view.EntryMonth = entryMonth
	view.Birthday = birthday
}
