package constant

import "errors"

var (
	ErrBodyInvalid = errors.New("request body is invalid")

	ErrMissingBodyURLTag = "ErrMissingBodyURL"
	ErrMissingBodyURL = errors.New("The url is not present.")

	ErrShortcodeNotFoundTag = "ErrShortcodeNotFound"
	ErrShortcodeNotFound = errors.New("The shortcode cannot be found in the system.") // 404

	ErrShortcodeAlreadyInUseTag = "ErrShortcodeAlreadyInUse"
	ErrShortcodeAlreadyInUse = errors.New("The desired shortcode is already in use.") // 409

	ErrShortcodeNotAlphaNumericTag = "ErrShortcodeNotAlphaNumeric"
	ErrShortcodeNotAlphaNumeric = errors.New("The shortcode must be 6 characters consists only letter and number.") // 422
)
