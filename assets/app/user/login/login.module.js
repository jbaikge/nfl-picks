"use strict"

angular.module("Picks.User.Login", [
	"ngCookies",
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/user/login", {
		controller:  "Picks.User.LoginController",
		templateUrl: "/app/user/login/login.partial.html"
	})
}])
