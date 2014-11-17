"use strict"

angular.module("Picks", [
	"ngRoute",
	"jsonrpc",
	"Picks.User.Login"
])

.config([ '$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
	$locationProvider.html5Mode(false)
	$locationProvider.hashPrefix("")
	$routeProvider.otherwise({ redirectTo: "/login" })
}])
