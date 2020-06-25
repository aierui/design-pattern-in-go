package templatemethod

import (
	"fmt"
	"testing"
)

func TestOpt(t *testing.T) {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)
	fmt.Println("")
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

	/*
		MS: generating random otp 1234
		SMS: saving otp: 1234 to cache
		SMS: sending sms: SMS OTP for login is 1234
		SMS: publishing metrics

		EMAIL: generating random otp 1234
		EMAIL: saving otp: 1234 to cache
		EMAIL: sending email: EMAIL OTP for login is 1234
		EMAIL: publishing metrics
	*/
}
