"use strict"

angular.module("Picks.Picks.Submitall", [
	"ngRoute",
	"ui.bootstrap"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/picks/submitall", {
		controller:  "Picks.Picks.SubmitallController",
		templateUrl: "/app/picks/submitall/submitall.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// submit.service.js
angular.module("Picks.Picks.Submitall").service("SubmitallService", ["jsonrpc", function(jsonrpc) {
	this.lines = function() {
		return jsonrpc("Lines.Current")
	}

	this.picks = function() {
		return jsonrpc("Picks.AllCurrent")
	}

	this.picksClosed = function() {
		return jsonrpc("Picks.Closed")
	}

	this.scores = function(year, week) {
		return jsonrpc("Game.Scores", { Year: year, Week: week })
	}

	this.submit = function(userId, picks) {
		var p = {
			UserId: userId,
			Picks:  picks
		}
		return jsonrpc("Picks.Submit", p)
	}

	this.tiebreaker = function(userId, week, year, value) {
		var t = {
			UserId: userId,
			TieBreaker: {
				Week: {
					Week: week,
					Year: year
				},
				Value: value
			}
		}
		return jsonrpc("TieBreaker.Submit", t)
	}

	this.tiebreakers = function() {
		return jsonrpc("TieBreaker.Current")
	}

	this.users = function() {
		return jsonrpc("User.Usernames")
	}
}])

// submit.controller.js
angular.module("Picks.Picks.Submitall").controller("Picks.Picks.SubmitallController", [
	"$log",
	"$rootScope",
	"$scope",
	"SubmitallService",
	function($log, $rootScope, $scope, SubmitallService) {
		$scope.Closed     = true
		$scope.Week       = {}
		$scope.Lines      = []
		$scope.Picks      = {}
		$scope.TieBreaker = { Value: 0, Submitting: false, Submitted: false }
		$scope.Users       = []


		var updateClosed = function() {
			SubmitallService.picksClosed()
				.success(function(data, status) {
					$scope.Closed = data.Closed
					if ($scope.Closed == false) {
						$timeout(updateClosed, 60 * 1000)
					}
				})
				.error(function(data, status) {
					$log.warn("error status: %s", status)
					$log.warn("error data: ", data)
				})
		}
		updateClosed()

		$scope.$watch("Picks", function(newValue, oldValue) {

			var changed = []
			for (var id in newValue) {
				var value = newValue[id]
				if (!oldValue.hasOwnProperty(id) || oldValue[id] != value) {
					changed.push({ GameId: id, Value:  value })
				}
			}
			// Fix up the percentages
			if (newValue.initial) {
				delete($scope.Picks.initial)
				return
			}

			SubmitallService.submit($rootScope.User.Id, changed)
				.success(function(data, status) {
					$log.log("SubmitallService", "data", data)
				})
				.error(function(data, status) {
					$log.log("SubmitallService.submit", "status", status, "data", data)
				})
		}, true)

		$scope.submitTieBreaker = function() {
			$log.log("TieBreaker", $scope.TieBreaker)
			if ($scope.Closed) {
				return
			}
			$scope.TieBreaker.Submitted = false
			$scope.TieBreaker.Submitting = true

			SubmitallService.tiebreaker($rootScope.User.Id, $scope.Week.Week, $scope.Week.Year, $scope.TieBreaker.Value)
				.success(function(data, status) {
					$scope.TieBreaker.Submitted = true
					$scope.TieBreaker.Submitting = false
					$log.log("SubmitallService.tiebreaker", "data", data)
				})
				.error(function(data, status) {
					$log.log("SubmitallService.tiebreaker", "data", data)
				})
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
			SubmitallService.submit($rootScope.User.Id, picks)
				.success(function(data, status) {
					$log.log("SubmitallService", "data", data)
				})
				.error(function(data, status) {
					$log.log("SubmitallService.submit", "status", status, "data", data)
				})
		}

		SubmitallService.users()
			.success(function(data) {
				$scope.Users = data.result.Usernames
			})

		SubmitallService.lines()
			.success(function(data, status) {
				$scope.Week = data.result.Week
				$scope.Lines = data.result.Lines
				// $scope.TieBreaker.Value = data.result.TieBreaker.Value
				// var picks = data.result.Picks
				// if (picks == null) {
				// 	return
				// }
				// var myPicks = {
				// 	// Since this is the initial load, don't let @watch re-send the existing picks
				// 	initial: true
				// }
				// for (var i = 0; i < picks.length; i++) {
				// 	myPicks[picks[i].GameId] = picks[i].Value
				// }
				// $scope.Picks = myPicks
			})
			.error(function(data, status) {
				$log.log("SubmitallService.currentLines", "status", status, "data", data)
			})
	}
])
