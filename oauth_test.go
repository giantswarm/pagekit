package pagekit_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/giantswarm/pagekit"
)

// knownCodes are the error codes the IANA OAuth Extensions Error Registry
// lists for the authorization endpoint, plus invalid_target, which RFC 8707
// registers for the implicit grant's error response.
var knownCodes = []string{
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
	"invalid_target",
}

func TestOAuthErrorCode_KeepsKnownCodes(t *testing.T) {
	for _, code := range knownCodes {
		t.Run(code, func(t *testing.T) {
			require.Equal(t, code, pagekit.OAuthErrorCode(code))
		})
	}
}

func TestOAuthErrorCode_DropsEverythingElse(t *testing.T) {
	for _, code := range []string{
		"",
		"ACCESS_DENIED",
		" access_denied",
		"access_denied ",
		"redirect_uri_mismatch",
		"your_account_is_locked_call_555",
		"Call support at 555",
	} {
		t.Run(code, func(t *testing.T) {
			require.Empty(t, pagekit.OAuthErrorCode(code))
		})
	}
}
