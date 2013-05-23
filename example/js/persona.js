var $ = function (id) {
	return document.getElementById(id);
}

function signIn(assertion) {
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "/signin", true);

	xhr.onreadystatechange = function() {
		if (xhr.readyState == 4) {
			if (xhr.status == 200) {
				console.log("signin worked, reloading page");
				// reload page to reflect new login state
				window.location.reload();
			} else {
				console.log("signin didn't work, calling logout()");
				navigator.id.logout();
			}
		}
	};
	
	var fd = new FormData();
	fd.append("assertion", assertion)

	xhr.send(fd);
}

function signOut() {
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "/signout", true);

	xhr.onreadystatechange = function() {
		if (xhr.readyState == 4) {
			if (xhr.status == 200) {
				console.log("signout worked, reloading page");
				// reload page to reflect new login state
				window.location.reload();
			} else {
				console.log("signout didn't work, calling logout()");
				navigator.id.logout();
			}
		}
	};

	xhr.send();
}
