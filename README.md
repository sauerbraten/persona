# Persona

This is a Go package providing easy access to the Persona verification API.

## Usage

Get the package:

	$ go get github.com/sauerbraten/persona

Import the package:

	import (
		"github.com/sauerbraten/persona"
	)

Using persona works like this:

1. you write yout own sign in handler
2. your sign in handler calls persona.VerifyAssertion() and passes your site's audience and the assertion you received with the sign in request from your site's javascript
3. persona.VerifyAssertion() returns a persona.Response
4. based on this response and on persona.Response.OK(), you either set a session cookie for access to your site or return an error, which should make your site's javascript call navigator.id.logout()

## Example

There is an example implementation in [`example/`](https://github.com/sauerbraten/persona/blob/master/example). Run `go run *.go` in that directory to get a demo site running at [localhost:8080](http://localhost:8080/).

Here is how it works:

1. The page uses a template that is filled with the email address read from the clients session cookie. If there is no such cookie, it offers to sign in with persona, if there is a cookie (and thus an email address), it displays that and offers to sign out.
2. The page includes [`js/persona.js`](https://github.com/sauerbraten/persona/blob/master/example/js/persona.js), where the javascript methods for sing-in and sign-out are written.
3. Sign-in works like this: it calls `navigator.id.request()` (as specified [here](https://developer.mozilla.org/en-US/docs/Web/API/navigator.id)). `navigator.id.request()` then calls `signIn()` (in [`js/persona.js`](https://github.com/sauerbraten/persona/blob/master/example/js/persona.js)), which POSTs the assertion to localhost:8080/signIn and reacts depending on the response.
4. If the server returns a `200 OK`, the page is reloaded. Since the server sent a cookie along with the `200 OK`, the client is recognized this time, and the email address is put into the template.
5. If the server returns anything else, `signIn()` calls `navigator.id.logout()` (as specified [here](https://developer.mozilla.org/en-US/docs/Web/API/navigator.id)). This again calls `signOut()`, which cleans up all session cookies there might be by POSTing to `localhost:8080/signOut`, and reloads the page. This way everything should be cleaned up and sign-in should work the next time you try.

## Documentation

Full documentation is at http://godoc.org/github.com/sauerbraten/persona.

## License

This code is licensed under a BSD License:

Copyright (c) 2013 Alexander Willing. All rights reserved.

- Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
- Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.