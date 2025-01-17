package main

import "fmt"

// Template method
type IOtp interface {
    genRandomOTP(int) string
    saveOTPCache(string)
    getMessage(string) string
    sendNotification(string) error
}

// type otp struct {
// }

// func (o *otp) genAndSendOTP(iOtp iOtp, otpLength int) error {
//  otp := iOtp.genRandomOTP(otpLength)
//  iOtp.saveOTPCache(otp)
//  message := iOtp.getMessage(otp)
//  err := iOtp.sendNotification(message)
//  if err != nil {
//      return err
//  }
//  return nil
// }

type Otp struct {
    iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
    otp := o.iOtp.genRandomOTP(otpLength)
    o.iOtp.saveOTPCache(otp)
    message := o.iOtp.getMessage(otp)
    err := o.iOtp.sendNotification(message)
    if err != nil {
        return err
    }
    return nil
}

// Concrete implementation
type Sms struct {
    Otp
}

func (s *Sms) genRandomOTP(len int) string {
    randomOTP := "1234"
    fmt.Printf("SMS: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
    fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
    return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
    fmt.Printf("SMS: sending sms: %s\n", message)
    return nil
}

// Concrete implementation
type Email struct {
    Otp
}

func (s *Email) genRandomOTP(len int) string {
    randomOTP := "1234"
    fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
    return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
    fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
    return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
    fmt.Printf("EMAIL: sending email: %s\n", message)
    return nil
}

// Client code
func main() {
    // otp := otp{}

    // smsOTP := &sms{
    //  otp: otp,
    // }

    // smsOTP.genAndSendOTP(smsOTP, 4)

    // emailOTP := &email{
    //  otp: otp,
    // }
    // emailOTP.genAndSendOTP(emailOTP, 4)
    // fmt.Scanln()
    smsOTP := &Sms{}
    o := Otp{
        iOtp: smsOTP,
    }
    o.genAndSendOTP(4)

    fmt.Println("")
    emailOTP := &Email{}
    o = Otp{
        iOtp: emailOTP,
    }
    o.genAndSendOTP(4)

}