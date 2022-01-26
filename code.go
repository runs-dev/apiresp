package apiresp

//CodeSuccess : print success code
func CodeSuccess() string {
	return success
}

//CodeInvalidSignature : print invalid signature code
func CodeInvalidSignature() string {
	return invalidSignature
}

//CodeQueryError : print query error code
func CodeQueryError() string {
	return queryError
}

//CodeUnregisteredUser : print Unregistered User code
func CodeUnregisteredUser() string {
	return unregisteredUser
}

//CodeExpiredOTP : print Expired OTP code
func CodeExpiredOTP() string {
	return unregisteredUser
}

//CodeInvalidOTP : print Invalid OTP code
func CodeInvalidOTP() string {
	return invalidOTP
}

//CodeEmptyEmail : print Empty Email code
func CodeEmptyEmail() string {
	return emptyEmail
}

//CodeEmptyParam : print Empty Param code
func CodeEmptyParam() string {
	return emptyParam
}

//CodeValidationError : print Validation Error code
func CodeValidationError() string {
	return validationError
}

//CodeInvalidDtTm : print Invalid Date Time code
func CodeInvalidDtTm() string {
	return invalidDtTm
}

//CodeInvalidRawBody : print Invalid Raw Body code
func CodeInvalidRawBody() string {
	return invalidRawBody
}

//CodeInvalidNTP : print Invalid NTP code
func CodeInvalidNTP() string {
	return invalidNTP
}

//CodeInternalServerError : print Internal Server Error code
func CodeInternalServerError() string {
	return internalServerError
}

//CodeMasterDataDuplicated : print Master Data Duplicated code
func CodeMasterDataDuplicated() string {
	return masterDataDuplicated
}

//CodeFailedBindingParam : print Failed Binding Param code
func CodeFailedBindingParam() string {
	return failedBindingParam
}

//CodeTroubleRequest : print Trouble Request code
func CodeTroubleRequest() string {
	return troubleRequest
}

//CodeTroubleDecode : print Trouble Decode code
func CodeTroubleDecode() string {
	return troubleDecode
}
