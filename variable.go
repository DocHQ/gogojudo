package gogojudo

import (
	"errors"
)

// Card ID's
var CardIDs = map[int]string{
	0:  "Unknown",
	1:  "Visa",
	2:  "Mastercard",
	3:  "Visa Electron",
	4:  "Switch",
	5:  "Solo",
	6:  "Laser",
	7:  "China Union Pay",
	8:  "Amex",
	9:  "JCB",
	10: "Maestro",
	11: "Visa Debit",
	12: "Mastercard Debit",
	13: "Visa Purchasing",
	14: "Discover",
	15: "Carnet",
	16: "Carte Bancaire",
	17: "Diners Club",
	18: "ELO",
	19: "Farmers cards",
	20: "Soriana",
	21: "Private Label Card",
	22: "Q Card",
	23: "Style",
	24: "True Rewards",
	25: "UATP",
	26: "Bankard",
	27: "Banamex costco",
}

var ErrorMap = map[int]error{
	0:      errors.New("Sorry, an error has occurred. Please try again later."),
	1:      errors.New("Sorry, we're unable to process your request. Please check your details and try again."),
	7:      errors.New("Sorry, we were unable to authorize this request. Please check your details and permissions before trying again."),
	9:      errors.New("Sorry, we were unable to process your payment. Please try again later."),
	11:     errors.New("Sorry, we were unable to process your payment at this time. Your card has not been charged."),
	12:     errors.New("Sorry, we were unable to process your payment. Please check your details and try again."),
	19:     errors.New("Sorry, but the transaction specified was not found."),
	20:     errors.New("Validation passed successfully."),
	21:     errors.New("Sorry, an error has occurred. Please try again later."),
	22:     errors.New("Sorry, an error has occurred. Please try again later."),
	23:     errors.New("Please ensure your 'From' date is in dd/mm/yyyy format."),
	24:     errors.New("Please ensure your 'To' date is in dd/mm/yyyy format."),
	25:     errors.New("We were unable to find the webpayment identified by this reference"),
	26:     errors.New("Sorry, an error has occurred. Please try again later."),
	27:     errors.New("This endpoint is not available as the Api-Version requested is too low."),
	28:     errors.New("This endpoint is not available as it has been deprecated for the Api-Version requested."),
	31:     errors.New("This endpoint is only available if the account is enabled to use the Judo Card Vault."),
	39:     errors.New("API-Version not supported"),
	40:     errors.New("Sorry, but the API version you are targeting is invalid. Please review this and try again."),
	41:     errors.New("Sorry, but the API-Version is missing from your header. Please check your details and try again."),
	42:     errors.New("Sorry, it looks like the PreAuth you are referencing has expired."),
	43:     errors.New("Sorry, it looks like you're trying to make a collection on an invalid transaction type. Collections can only be performed on PreAuths."),
	44:     errors.New("Sorry, it looks like the currency you're specifying doesn't match your original transaction."),
	45:     errors.New("Sorry, this transaction has been voided. You cannot perform a collection on a voided transaction."),
	46:     errors.New("Sorry, but the amount you're trying to collect is greater than the pre-auth"),
	47:     errors.New("Sorry, but it looks like the transaction you are trying to refund is invalid. Refunds can only be performed on Sales and Collections."),
	48:     errors.New("Sorry, this transaction has been voided. You cannot perform a refund on a voided transaction."),
	49:     errors.New("Sorry, but the amount you're trying to refund is greater than the original transaction."),
	50:     errors.New("Sorry, but it looks like the original transaction you're specifying is an invalid type for the request you're trying to perform."),
	51:     errors.New("Sorry, but it looks like the transaction you are trying to void has already been voided."),
	52:     errors.New("Sorry, but it looks like the transaction you are trying to void has already been collected."),
	53:     errors.New("Sorry, but it looks like the void you are trying to process is for a different amount than the original preauth."),
	54:     errors.New("Sorry, but we are currently unable to accept payments to this account. Please contact customer services."),
	55:     errors.New("Sorry, we're unable to find the judo ID specified. Please confirm your judo ID and try again."),
	56:     errors.New("Sorry, but the transaction specified was not found. Please check your details and try again."),
	57:     errors.New("Sorry, but there was no consumer found for the transaction specified. Please check your details and try again."),
	58:     errors.New("Sorry, but this transaction does not require 3D-Secure authorization. Please check you're sending authorization to the correct transaction."),
	59:     errors.New("Sorry, but it looks like this transaction has already been authorised via 3D-Secure."),
	60:     errors.New("3D-Secure was not successful for this transaction."),
	61:     errors.New("We've been unable to decrypt the supplied Apple Pay token. Please check your API client configuration in the dashboard."),
	62:     errors.New("Sorry, we were unable to find the transaction you have referenced in your request."),
	63:     errors.New("Sorry, but the transaction you are referencing was not successful."),
	64:     errors.New("Sorry, but it looks like you're using a test card. Only real cards are valid in this environment."),
	65:     errors.New("Sorry, but the collection request you've specified is not valid. Please check your details and try again."),
	66:     errors.New("Sorry, but we were unable to find your original transaction to refund. Please check your details and try again."),
	67:     errors.New("Sorry, but your refund request was not valid. Please check that the original transaction was a Sale or Collection, has not previously been refunded, was not for a lesser amount than your refund request."),
	68:     errors.New("Sorry, but your void request is not valid. Please check your details and try again."),
	69:     errors.New("Our server encountered a problem when processing your transaction."),
	70:     errors.New("Sorry, but it looks like the card token specified is not valid. Please check your details and try again."),
	71:     errors.New("Sorry, an error has occurred. Please try again later."),
	72:     errors.New("Sorry, we're currently unable to route this transaction. Please check your account details and try again. If this issue persists, please contact customer services."),
	73:     errors.New("Sorry, but this card type is not currently supported on this account."),
	74:     errors.New("The CV2 entered is invalid."),
	75:     errors.New("Sorry, but the Card Token specified does not match the Consumer it was originally generated against. Please check your records and try again."),
	76:     errors.New("Sorry, but it appears the web payment reference you've specified is not valid."),
	77:     errors.New("Judo id not found, please check the judo id."),
	78:     errors.New("Sorry, but you are attempting to register a card with an incorrect transaction type. Please check your details and try again."),
	79:     errors.New("Sorry, but it looks you're trying to register this card with an invalid amount. Please check the amount of your request and try again."),
	80:     errors.New("Sorry, but the content-type was not specified or is currently unsupported. Currently only application/json is supported."),
	81:     errors.New("Sorry, we encountered an error while authenticating your request. Please check your authentication details and try again. If this issue continues, please contact customer services."),
	82:     errors.New("Sorry, the transaction you've specified has not been found. Please check your details and try again."),
	83:     errors.New("Sorry, the resource you are targeting has not been found. Please review your endpoints and try again."),
	84:     errors.New("Sorry, it looks like permissions are not enabled to perform this request. Application permissions can be configured in the judo Dashboard."),
	85:     errors.New("The content-type was not specified or is unsupported for the request made to the Judopay API. Currently supported content-type is limited to application/json."),
	86:     errors.New("Sorry, this payment has been stopped as it is a duplicate transaction."),
	130:    errors.New("We've been unable to decrypt the supplied Visa Checkout payload. Please confirm the public key in your app and check your API client configuration in the dashboard."),
	140:    errors.New("Something went wrong encryting card details. We apologise for any inconvenience caused. Please try again later"),
	150:    errors.New("We've been unable to decrypt the supplied Encrypted payload. Please ensure the message has not been modified."),
	151:    errors.New("The encrypted blob is for another account."),
	152:    errors.New("Something went wrong retrieving card details. We apologise for any inconvenience caused. Please try again later"),
	153:    errors.New("The one time token is not valid. It could have expired. Please try again"),
	154:    errors.New("Your provided consumer reference is incorrect, please check your details and try again."),
	400:    errors.New("We've been unable to decrypt the supplied Android Pay wallet. Please confirm the public key in your app and check your API client configuration in the dashboard."),
	403:    errors.New("Sorry, but we were unable to authenticate your request. Please check your details and try again."),
	404:    errors.New("Sorry, but the requested resource was not found. Please check your details and try again."),
	601:    errors.New("Unable to decrypt the Google Pay payload."),
	602:    errors.New("Unable to access the Google Pay decryption service, please try again later."),
	603:    errors.New("Non-tokenized Google Pay cards not supported. Please verify the card can be tokenized for contactless Google Pay payments."),
	4002:   errors.New("Sorry, you cannot process marketplace preauths without using the Access Token. For more information please contact customer services."),
	20000:  errors.New("Sorry, but the ApplicationModel must not be null. Please check your details and try again."),
	20001:  errors.New("Sorry, but the Application Reference must not be null. Please check your details and try again."),
	20002:  errors.New("Sorry, but this application has already gone live and cannot be updated via this endpoint."),
	20003:  errors.New("Sorry, but the product selection is missing. Please check your details and try again."),
	20004:  errors.New("Sorry, but this application has already gone live and cannot be updated via this endpoint."),
	20005:  errors.New("Sorry, but you have not set your Api Application Rec ID."),
	20006:  errors.New("Sorry, but this request is not formatted correctly. Please review your request and try again."),
	20007:  errors.New("Sorry, but no application was found for the reference you supplied. Please check your details and try again."),
	20008:  errors.New("Sorry, but the file type specified is not currently supported."),
	20009:  errors.New("Sorry, there was an error with the file upload. Please try again later."),
	20010:  errors.New("Sorry, but the Application Reference must not be empty. Please supply a reference and try again."),
	20011:  errors.New("Sorry, but an application with the Application Reference supplied was not found. Please check your details and try again."),
	20013:  errors.New("Sorry, but you have not specified a sort type. Valid sort types include time-descending and time-ascending."),
	20014:  errors.New("Sorry, but the page size needs to include at least one record per page."),
	20015:  errors.New("Sorry, but the page size is currently limited to 500 or less."),
	20016:  errors.New("Sorry, but the offset supplied needs to be greater than zero."),
	20017:  errors.New("Sorry, but the Merchant ID supplied was not found. Please check your details and try again."),
	20018:  errors.New("Sorry, but the Merchant ID supplied was not found. Please check your details and try again."),
	20019:  errors.New("Sorry, no Products were found. Please check your details and try again."),
	20020:  errors.New("Sorry, but only a judo partner can submit simple applications. Please contact customer service."),
	20021:  errors.New("We're unable to parse this document, please ensure it conforms to our API requirements and then try again."),
	20022:  errors.New("Sorry, something went wrong, please check your authentication details and judo, then try again."),
	20023:  errors.New("To create a web payment, you'll need to perform a POST to the URL."),
	20025:  errors.New("Sorry, but MD parameter specified is invalid. Please check your details and try again."),
	20026:  errors.New("The receipt ID specified is not valid for the request you are attempting."),
	20027:  errors.New("Email address or password incorrect"),
	20028:  errors.New("Permanent Lockout - We detected suspicious activity on this account and have locked it for your protection. Please contact our customer service team on 0203 503 600 to unlock your account."),
	20029:  errors.New("Temporary Lockout - We detected suspicious activity on this account and have locked it for your protection. Your account has been locked for {0} minutes."),
	20030:  errors.New("Sorry, this account does not currently support the requested currency. Please contact customer services for more information."),
	20032:  errors.New("Sorry, this email address is already in use."),
	667003: errors.New("WebPayment reference is invalid"),
	667004: errors.New("JudoId not found, please check the judo id."),
	667005: errors.New("This account is currently unable to accept payments. Please contact customer services."),
}