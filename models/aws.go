package models

type Recipient struct {
	ToEmails  []string
	CcEmails  []string
	BccEmails []string
}
