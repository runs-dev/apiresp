package apiresp

var (
	success              = "000"
	invalidSignature     = "100"
	queryError           = "101"
	unregisteredUser     = "102"
	expiredOTP           = "103"
	invalidOTP           = "104"
	emptyEmail           = "105"
	emptyParam           = "106"
	validationError      = "107"
	invalidDtTm          = "108"
	invalidRawBody       = "109"
	invalidNTP           = "110"
	internalServerError  = "111"
	masterDataDuplicated = "112"
	failedBindingParam   = "113"
	troubleRequest       = "114"
	troubleDecode        = "115"
)

var statusText = map[string]string{
	success:              "Success.",
	invalidSignature:     "Your signature is invalid.",
	queryError:           "Query error",
	unregisteredUser:     "User has not been registered. Contact your administrator.",
	expiredOTP:           "Your OTP is expired.",
	invalidOTP:           "Invalid OTP.",
	emptyEmail:           "Your account has no email yet. Contact your administrator.",
	emptyParam:           "You need to fill all parameters. In this case : ",
	validationError:      "Validation error.",
	invalidDtTm:          "End date is earlier than start date.",
	invalidRawBody:       "Invalid raw body : ",
	invalidNTP:           "Invalid NTP.",
	internalServerError:  "There is a problem with the service.",
	masterDataDuplicated: "You've entered duplicated data.",
	failedBindingParam:   "Invalid parameter binding.",
	troubleRequest:       "Failed to request data with such error : ",
	troubleDecode:        "Failed to decode data with such error : ",
}
