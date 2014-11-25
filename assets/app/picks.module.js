"use strict"

angular.module("Picks", [
	"ngCookies",
	"ngRoute",
	"jsonrpc",
	"quarter",
	"Picks.Picks.Viewall",
	"Picks.Picks.Submit",
	"Picks.User.Login",
	"Picks.User.Profile",
	"ui.bootstrap"
])

.config(["$routeProvider", "$locationProvider", function($routeProvider, $locationProvider) {
	$locationProvider.html5Mode(false)
	$locationProvider.hashPrefix("")
	$routeProvider.otherwise({ redirectTo: "/picks/viewall" })
}])

.run(["$cookieStore", "$location", "$rootScope", function($cookieStore, $location, $rootScope) {
	$rootScope.$on("$routeChangeStart", function(event, next, current) {
		if (next.$$route) {
			if (!angular.isDefined($rootScope.User)) {
				$rootScope.User = $cookieStore.get("User")
			}
			var user = $rootScope.User
			var auth = next.$$route.auth
			if (auth && !auth(user)) {
				$location.path("/")
			}
			if (angular.isDefined(user)) {
				$rootScope.Theme = user.Theme
			}
		}
	})
}])
