// Package persona offers functionality to use Mozilla's Persona (formerly BrowserID) identification system with Go web applications.
package persona

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Verifies an assertion for a given audience. Returns the response and any errors that occured while requesting, receiving and unmarshaling the response. Verification errors have to be handled using Reponse.OK() / Reponse.Status and Response.Reason.
func VerifyAssertion(audience, assertion string) (resp Response, err error) {
	rawResp, err := http.PostForm("https://verifier.login.persona.org/verify", url.Values{"audience": {audience}, "assertion": {assertion}})
	if err != nil {
		return
	}
	defer rawResp.Body.Close()

	body, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	return
}

// Describes a verification response, i.e. the reponse you get when using the https://verifier.login.persona.org/verify service.
type Response struct {
	Status string `json:"status"` // either "okay" or "failure"

	// fields are only populated if status == "failure"
	Reason string `json:"reason,omitempty"` // reason for failure

	// fields are only populated if status == "okay"
	Email    string `json:"email,omitempty"`    // the user's email address
	Audience string `json:"audience,omitempty"` // the audience for which the assertion was verified
	Expires  int64  `json:"expires,omitempty"`  // the expiration date as a unix timestamp
	Issuer   string `json:"issuer,omitempty"`   // the hostname and port of the identity provider
}

// Returns true if the verification was successful, and false otherwise.
func (r Response) OK() bool {
	return r.Status == "okay"
}
