package pagekit

// oauthErrorCodes are the error codes an authorization response may carry:
// RFC 6749 section 4.1.2.1, OpenID Connect Core section 3.1.2.6 and RFC 8707.
var oauthErrorCodes = map[string]bool{
	"invalid_request":            true,
	"unauthorized_client":        true,
	"access_denied":              true,
	"unsupported_response_type":  true,
	"invalid_scope":              true,
	"server_error":               true,
	"temporarily_unavailable":    true,
	"interaction_required":       true,
	"login_required":             true,
	"account_selection_required": true,
	"consent_required":           true,
	"invalid_request_uri":        true,
	"invalid_request_object":     true,
	"request_not_supported":      true,
	"request_uri_not_supported":  true,
	"registration_not_supported": true,
	"invalid_target":             true,
}

// OAuthErrorCode returns code when it is an error code an authorization
// response may carry, such as access_denied or login_required, and "" for
// anything else. Anyone can send a browser to a callback with an error
// parameter of their choosing, so pass a callback's error through it before
// showing it, for example as Page.Detail.
func OAuthErrorCode(code string) string {
	if oauthErrorCodes[code] {
		return code
	}
	return ""
}
