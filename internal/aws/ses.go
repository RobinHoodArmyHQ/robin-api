package aws

import (
	"log"

	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// SendEmailSES sends email to specified email IDs
func SendEmailSES(messageBody string, subject string, fromEmail string, recipient models.Recipient) (err error) {
	// create new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		log.Println("Error occurred while creating aws session", err)
		return
	}

	// set to section
	var recipients []*string
	for _, r := range recipient.ToEmails {
		recipient := r
		recipients = append(recipients, &recipient)
	}

	// set cc section
	var ccRecipients []*string
	if len(recipient.CcEmails) > 0 {
		for _, r := range recipient.CcEmails {
			ccrecipient := r
			ccRecipients = append(ccRecipients, &ccrecipient)
		}
	}

	// set bcc section
	var bccRecipients []*string
	if len(recipient.BccEmails) > 0 {
		for _, r := range recipient.BccEmails {
			bccrecipient := r
			recipients = append(recipients, &bccrecipient)
		}
	}

	// create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		// Set destination emails
		Destination: &ses.Destination{
			CcAddresses:  ccRecipients,
			ToAddresses:  recipients,
			BccAddresses: bccRecipients,
		},

		// Set email message and subject
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(messageBody),
				},
			},

			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},

		// send from email
		Source: aws.String(fromEmail),
	}

	// Call AWS send email function which internally calls to SES API
	_, err = svc.SendEmail(input)
	if err != nil {
		log.Println("Error sending mail - ", err)
	}

	return
}
