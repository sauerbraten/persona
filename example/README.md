# Files and Folders

- *css/*, *html/*, *js/* are self explanatory
- *cookie_handling.go* initializes a securecookie instance and contains methods to set and revoke session cookies, using [gorilla/securecookie](http://www.gorillatoolkit.org/pkg/securecookie). It also reads email addresses out of those cookies.
- *user_management.go* sets up the user "database" and contains the sign-in and sign-out handlers (`signIn()` and `signOut()` respectively), plus dummy methods to simulate user management (adding users, checking if a user already existst).
- *server.go* contains `main()` and a simple landing page handler. `main()` sets up the HTTP routes and fires up the server.

# How it works:

The page uses a [template](https://github.com/sauerbraten/persona/blob/master/example/html/landing.html) that is filled with the email address read from the clients session cookie. If there is no such cookie, the page offers to sign in with persona, if there is a cookie (and thus an email address), the page displays that address and offers to sign out.

The page includes [*js/persona.js*](https://github.com/sauerbraten/persona/blob/master/example/js/persona.js), where the javascript methods for sing-in and sign-out are written.

Sign-in works like this: it calls `navigator.id.request()` (as specified [here](https://developer.mozilla.org/en-US/docs/Web/API/navigator.id)). `navigator.id.request()` then calls `signIn()` (in [*js/persona.js*](https://github.com/sauerbraten/persona/blob/master/example/js/persona.js)), which POSTs the assertion to `localhost:8080/signin` and reacts depending on the response:

- if the server returns a `200 OK`, the page is reloaded. Since the server sent a cookie along with the `200 OK`, the client is recognized this time, and the email address is put into the template.
- if the server returns anything else, `signIn()` calls `navigator.id.logout()` (as specified [here](https://developer.mozilla.org/en-US/docs/Web/API/navigator.id)). This again calls `signOut()`, which cleans up all session cookies there might be by POSTing to `localhost:8080/signout`, and reloads the page. This way everything should be cleaned up and sign-in should work the next time you try.
