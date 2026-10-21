package pagekit_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/giantswarm/pagekit"
)

func TestOAuthErrorCode_KeepsOnlyKnownCodes(t *testing.T) {
	testCases := []struct {
		code string
		want string
	}{
		// RFC 6749 section 4.1.2.1
		{code: "access_denied", want: "access_denied"},
		{code: "temporarily_unavailable", want: "temporarily_unavailable"},
		// OpenID Connect Core section 3.1.2.6
		{code: "login_required", want: "login_required"},
		{code: "registration_not_supported", want: "registration_not_supported"},
		// RFC 8707
		{code: "invalid_target", want: "invalid_target"},
		// Anything else, which anyone can put into a callback URL
		{code: "", want: ""},
		{code: "ACCESS_DENIED", want: ""},
		{code: " access_denied", want: ""},
		{code: "your_account_is_locked_call_555", want: ""},
		{code: "Call support at 555", want: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			require.Equal(t, tc.want, pagekit.OAuthErrorCode(tc.code))
		})
	}
}
