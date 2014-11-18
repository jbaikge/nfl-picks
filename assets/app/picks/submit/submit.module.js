"use strict"

angular.module("Picks.Picks.Submit", [
	"ngRoute",
	"ui.bootstrap"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/picks/submit", {
		controller:  "Picks.Picks.SubmitController",
		templateUrl: "/app/picks/submit/submit.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// pick.service.js
angular.module("Picks.Picks.Submit").service("SubmitService", ["jsonrpc", function(jsonrpc) {
	this.currentLines = function() {
		return jsonrpc("Lines.Current")
	}

	this.submit = function(userId, picks) {
		var p = {
			UserId: userId,
			Picks:  picks
		}
		return jsonrpc("Picks.Submit", p)
	}
}])

// pick.controller.js
angular.module("Picks.Picks.Submit").controller("Picks.Picks.SubmitController", [
	"$rootScope",
	"$scope",
	"SubmitService",
	function($rootScope, $scope, SubmitService) {
		$scope.Lines = []
		$scope.Picks = {
			"BUFvMIA@20141113":"OVR"
		}

		$scope.submitPicks = function() {
			var picks = []
			var gameId
			for (gameId in $scope.Picks) {
				picks.push({
					GameId: gameId,
					Value:  $scope.Picks[gameId]
				})
			}
			SubmitService.submit($rootScope.User.Id, picks)
				.success(function(data, status) {
					console.log("SubmitService", "data", data)
				})
				.error(function(data, status) {
					console.log("SubmitService.submit", "status", status, "data", data)
				})
		}

		SubmitService.currentLines()
			.success(function(data, status) {
				$scope.Current = data.result.Current
				$scope.Lines = data.result.Lines
			})
			.error(function(data, status) {
				console.log("SubmitService.currentLines", "status", status, "data", data)
			})
	}
])
