"use strict"

angular.module("Picks.User.Login").service("LoginService", ["jsonrpc", function(jsonrpc) {
	this.auth = function(username, pin) {
		var data = {
			Username: username,
			PIN: pin
		}
		return jsonrpc("User.Auth", data)
	}

	this.usernames = function() {
		return jsonrpc("User.Usernames")
	}
}])
