package apiresp

//MsgSuccess : print success message
func MsgSuccess() string {
	return statusText[success]
}

//MsgInvalidSignature : print invalid signature message
func MsgInvalidSignature() string {
	return statusText[invalidSignature]
}

//MsgQueryError : print query error message
func MsgQueryError() string {
	return statusText[queryError]
}

//MsgUnregisteredUser : print Unregistered User message
func MsgUnregisteredUser() string {
	return statusText[unregisteredUser]
}

//MsgExpiredOTP : print Expired OTP message
func MsgExpiredOTP() string {
	return statusText[unregisteredUser]
}

//MsgInvalidOTP : print Invalid OTP message
func MsgInvalidOTP() string {
	return statusText[invalidOTP]
}

//MsgEmptyEmail : print Empty Email message
func MsgEmptyEmail() string {
	return statusText[emptyEmail]
}

//MsgEmptyParam : print Empty Param message
func MsgEmptyParam() string {
	return statusText[emptyParam]
}

//MsgValidationError : print Validation Error message
func MsgValidationError() string {
	return statusText[validationError]
}

//MsgInvalidDtTm : print Invalid Date Time message
func MsgInvalidDtTm() string {
	return statusText[invalidDtTm]
}

//MsgInvalidRawBody : print Invalid Raw Body message
func MsgInvalidRawBody() string {
	return statusText[invalidRawBody]
}

//MsgInvalidNTP : print Invalid NTP message
func MsgInvalidNTP() string {
	return statusText[invalidNTP]
}

//MsgInternalServerError : print Internal Server Error message
func MsgInternalServerError() string {
	return statusText[internalServerError]
}

//MsgMasterDataDuplicated : print Master Data Duplicated message
func MsgMasterDataDuplicated() string {
	return statusText[masterDataDuplicated]
}

//MsgFailedBindingParam : print Failed Binding Param message
func MsgFailedBindingParam() string {
	return statusText[failedBindingParam]
}

//MsgTroubleRequest : print Trouble Request message
func MsgTroubleRequest() string {
	return statusText[troubleRequest]
}

//MsgTroubleDecode : print Trouble Decode message
func MsgTroubleDecode() string {
	return statusText[troubleDecode]
}
