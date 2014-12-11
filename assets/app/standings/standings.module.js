"use strict"

angular.module("Picks.Standings", [
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/standings", {
		controller:  "Picks.StandingsController",
		templateUrl: "/app/standings/standings.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// profile.service.js
angular.module("Picks.Standings").service("StandingsService", ["jsonrpc", function(jsonrpc) {
	this.standings = function() {
		return jsonrpc("Standings.Overall")
	}

	this.users = function() {
		return jsonrpc("User.Usernames")
	}
}])

// profile.controller.js
angular.module("Picks.Standings").controller("Picks.StandingsController", [
	"$rootScope",
	"$scope",
	"StandingsService",
	function($rootScope, $scope, StandingsService) {
		$scope.Users      = []
		$scope.Standings  = []
		$scope.UserTotals = {}
		$scope.TotalPicks = 0

		StandingsService.users()
			.success(function(data) {
				$scope.Users = data.result.Usernames
			})

		StandingsService.standings()
			.success(function(data) {
				$scope.Standings = data.result.Standings
				for (var i = 0; i < $scope.Standings.length; i++) {
					var s = $scope.Standings[i]
					$scope.TotalPicks += s.TotalGames
					for (var u in s.UserWins) {
						if (isNaN($scope.UserTotals[u])) {
							$scope.UserTotals[u] = 0
						}
						$scope.UserTotals[u] += s.UserWins[u]
					}
				}
			})
	}
])
