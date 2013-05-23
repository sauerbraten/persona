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

1. you write your own sign in handler
2. your sign in handler calls `persona.VerifyAssertion()` and passes your site's audience and the assertion you received with the sign in request from your site's javascript
3. `persona.VerifyAssertion()` returns a persona.Response
4. based on this response and on `persona.Response.OK()`, you either set a session cookie for access to your site or return an error, which should make your site's javascript call `navigator.id.logout()`

## Example

There is an example implementation in [*example/*](https://github.com/sauerbraten/persona/blob/master/example). Run `go run *.go` in that directory to get a demo site running at [localhost:8080](http://localhost:8080/).

## Documentation

Full documentation is at http://godoc.org/github.com/sauerbraten/persona.

## License

This code is licensed under a BSD License:

Copyright (c) 2013 Alexander Willing. All rights reserved.

- Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
- Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.