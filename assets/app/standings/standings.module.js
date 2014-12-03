"use strict"

angular.module("Picks.Standings", [
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/standings", {
		controller:  "Picks.StandingsController",
		templateUrl: "/app/standings.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// profile.service.js
angular.module("Picks.Standings").service("StandingsService", ["jsonrpc", function(jsonrpc) {
	this.get = function() {
		return jsonrpc("Picks.Standings")
	}
}])

// profile.controller.js
angular.module("Picks.User.Profile").controller("Picks.User.ProfileController", [
	"$rootScope",
	"$scope",
	"StandingsService",
	function($rootScope, $scope, StandingsService) {
		$scope.Usernames = []
		$scope.Standings = {}
	}
])
