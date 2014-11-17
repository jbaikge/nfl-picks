"use strict"

angular.module("Picks.User.Login", [
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/login", {
		controller:  "Picks.User.LoginController",
		templateUrl: "/app/user/login/login.partial.html"
	})
}])
