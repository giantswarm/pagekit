package pagekit

// OAuthErrorCode returns code when it is an error code an authorization
// response may carry, such as access_denied or login_required, and "" for
// anything else. Anyone can send a browser to a callback with an error
// parameter of their choosing, so pass a callback's error through it before
// showing it, for example as Page.Detail, and log the raw value instead.
//
// The known codes are those the IANA OAuth Extensions Error Registry lists
// for the authorization endpoint, plus invalid_target, which RFC 8707
// registers for the implicit grant's error response.
func OAuthErrorCode(code string) string {
	switch code {
	case
		// RFC 6749
		"invalid_request", "unauthorized_client", "access_denied", "unsupported_response_type",
		"invalid_scope", "server_error", "temporarily_unavailable", "invalid_client",
		// OpenID Connect Core
		"interaction_required", "login_required", "account_selection_required", "consent_required",
		"invalid_request_uri", "invalid_request_object", "request_not_supported",
		"request_uri_not_supported", "registration_not_supported",
		// RFC 9396
		"invalid_authorization_details",
		// OpenID Federation
		"invalid_issuer", "invalid_subject", "invalid_trust_anchor", "invalid_trust_chain",
		"invalid_metadata", "not_found", "unsupported_parameter",
		// OpenID for Verifiable Presentations
		"vp_formats_not_supported", "invalid_request_uri_method", "wallet_unavailable",
		// RFC 8707
		"invalid_target":
		return code
	}
	return ""
}
