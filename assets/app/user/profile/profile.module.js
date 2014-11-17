"use strict"

angular.module("Picks.User.Profile", [
	"ngCookies",
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/user/profile", {
		controller:  "Picks.User.ProfileController",
		templateUrl: "/app/user/profile/profile.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// profile.service.js
angular.module("Picks.User.Profile").service("ProfileService", ["jsonrpc", function(jsonrpc) {

}])

// profile.controller.js
angular.module("Picks.User.Profile").controller("Picks.User.ProfileController", [
	"$cookieStore",
	"$rootScope",
	"$scope",
	"ProfileService",
	function($cookieStore, $rootScope, $scope, ProfileService) {
	
	}
])
