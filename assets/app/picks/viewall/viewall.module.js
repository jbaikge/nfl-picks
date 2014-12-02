"use strict"

angular.module("Picks.Picks.Viewall", [
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/picks/viewall", {
		controller:  "Picks.Picks.ViewallController",
		templateUrl: "/app/picks/viewall/viewall.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// viewall.service.js
angular.module("Picks.Picks.Viewall").service("ViewallService", ["jsonrpc", function(jsonrpc) {
	this.picks = function() {
		return jsonrpc("Picks.AllCurrent")
	}

	this.lines = function() {
		return jsonrpc("Lines.Current")
	}

	this.users = function() {
		return jsonrpc("User.Usernames")
	}

	this.scores = function(year, week) {
		return jsonrpc("Game.Scores", { Year: year, Week: week })
	}
}])

// viewall.controller.js
angular.module("Picks.Picks.Submit").controller("Picks.Picks.ViewallController", [
	"$log",
	"$rootScope",
	"$scope",
	"$timeout",
	"ViewallService",
	function($log, $rootScope, $scope, $timeout, ViewallService) {
		$scope.Week        = {}
		$scope.Users       = []
		$scope.Games       = []
		$scope.Picks       = {}
		$scope.Scores      = {}
		$scope.Totals      = {}
		$scope.TieBreakers = {}

		var updateScores = function(delay) {
			$log.log("updateScores %d", delay)
			var doUpdate = function() {
				ViewallService.scores($scope.Week.Year, $scope.Week.Week)
					.success(function(data) {
						var scores = data.result.Scores
						for (var i = 0; i < scores.length; i++) {
							$scope.Scores[scores[i].Id] = scores[i]
						}
						var nextUpdate = data.result.NextUpdate / 1e6
						updateScores(nextUpdate)
					})

				if (!$rootScope.User.IsAdmin && new Date($scope.Games[0].Start) > new Date) {
					$log.log("Skipping:", new Date($scope.Games[0].Start), ">", new Date)
					return
				}

				ViewallService.picks()
					.success(function(data) {
						$scope.Picks = data.result.Picks
						$scope.Totals = data.result.Totals
					})
			}
			return $timeout(doUpdate, delay)
		}

		ViewallService.users()
			.success(function(data) {
				$scope.Users = data.result.Usernames
			})

		ViewallService.lines()
			.success(function(data) {
				$scope.Week = data.result.Week
				$scope.Games = data.result.Lines
				updateScores(0)
			})
	}
])
