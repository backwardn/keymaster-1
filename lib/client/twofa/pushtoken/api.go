// Package vip does two factor authentication with Symantec VIP or Okta
package pushtoken

import (
	"net/http"

	"github.com/Cloud-Foundations/Dominator/lib/log"
)

// DoVIPAuthenticate performs two factor authentication with Symantec VIP
func DoVIPAuthenticate(
	client *http.Client,
	baseURL string,
	userAgentString string,
	logger log.DebugLogger) error {
	return doGenericTokenPushAuthenticate(client, baseURL, "vip", userAgentString, logger)
}

func DoOktaAuthenticate(
	client *http.Client,
	baseURL string,
	userAgentString string,
	logger log.DebugLogger) error {
	return doGenericTokenPushAuthenticate(client, baseURL, "okta", userAgentString, logger)
}
