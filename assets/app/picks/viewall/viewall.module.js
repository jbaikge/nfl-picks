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
	"$rootScope",
	"$scope",
	"ViewallService",
	function($rootScope, $scope, ViewallService) {
		$scope.Week   = {}
		$scope.Users  = []
		$scope.Games  = []
		$scope.Picks  = {}
		$scope.Scores = []

		ViewallService.users()
			.then(function(response) {
				$scope.Users = response.data.result.Usernames
			})

		ViewallService.lines()
			.then(function(response) {
				$scope.Week = response.data.result.Week
				$scope.Games = response.data.result.Lines
				return ViewallService.scores($scope.Week.Year, $scope.Week.Week)
			})
			.then(function(response) {
				$scope.Scores = response.data.result.Scores
				console.log($scope.Scores)
			})

		ViewallService.picks()
			.then(function(response) {
				$scope.Picks = response.data.result.Picks
			})
	}
])
