package main

import "fmt"

// Notifier defines the behavior for sending notifications.
type Notifier interface {
	Send(msg string)
}

// Implements interface
func Notify(n Notifier, msg string) {
	n.Send(msg)
}

// Implement SMS
type SMS struct{}

func (s SMS) Send(msg string) {
	fmt.Println("Sending SMS:", msg)
}

// Implement Email
type Email struct{}

func (e Email) Send(msg string) {
	fmt.Println("Sending Email", msg)
}

func main() {
	// Without Using Interface calling directly to concrete methods.
	// SMS Notification
	// sms := SMS{}
	// sms.Send("Nice to talk to you...!")

	// // Email Notification
	// email := Email{}
	// email.Send("Thank you for visiting...!")

	// implements Notifier
	// Better code Not required but good for understanding..
	var sms Notifier = SMS{}
	var email Notifier = Email{}

	fmt.Println("----------- Using Interface ------------")
	Notify(sms, "Interface... I love this topic.")
	Notify(email, "Interface... I love this topic.")

}
