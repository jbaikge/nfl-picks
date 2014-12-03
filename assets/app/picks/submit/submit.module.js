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

// submit.service.js
angular.module("Picks.Picks.Submit").service("SubmitService", ["jsonrpc", function(jsonrpc) {
	this.currentLines = function(userId) {
		return jsonrpc("Lines.Current", {UserId: userId})
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
}])

// submit.controller.js
angular.module("Picks.Picks.Submit").controller("Picks.Picks.SubmitController", [
	"$log",
	"$rootScope",
	"$scope",
	"SubmitService",
	function($log, $rootScope, $scope, SubmitService) {
		$scope.Week       = {}
		$scope.Lines      = []
		$scope.Picks      = {}
		$scope.Progress   = {}
		$scope.TieBreaker = { Value: 0, Submitting: false, Submitted: false }

		// TODO - Turn this into an API call
		var now = new Date
		$scope.Closed = !(now.getDay() == 3 && now.getHours() >= 17 || now.getDay() == 4 && now.getHours() <= 12)

		$scope.$watch("Picks", function(newValue, oldValue) {

			var o = { Home: 0, Away: 0, Over: 0, Under: 0 }
			var changed = []
			for (var id in newValue) {
				var value = newValue[id]
				if (!oldValue.hasOwnProperty(id) || oldValue[id] != value) {
					changed.push({ GameId: id, Value:  value })
				}
				switch (true) {
				case value == "OVER":        o.Over++;  break
				case value == "UNDER":       o.Under++; break
				case id.indexOf(value) == 0: o.Away++;  break
				case id.indexOf(value) > 0:  o.Home++;  break
				}
			}
			// Fix up the percentages
			var total = 0
			for (var k in o) {
				o[k] = Math.round(o[k] / $scope.Lines.length * 100)
				total += o[k]
			}
			if (total > 100) {
				o.Under -= total - 100
			}
			$log.log("counts:", o, "changed:", changed)

			if (newValue.initial) {
				delete($scope.Picks.initial)
				return
			}
			$scope.Progress = o

			// Do not communicate with server if picking is closed
			if ($scope.Closed) {
				return
			}
			SubmitService.submit($rootScope.User.Id, changed)
				.success(function(data, status) {
					$log.log("SubmitService", "data", data)
				})
				.error(function(data, status) {
					$log.log("SubmitService.submit", "status", status, "data", data)
				})
		}, true)

		$scope.submitTieBreaker = function() {
			$log.log("TieBreaker", $scope.TieBreaker)
			if ($scope.Closed) {
				return
			}
			$scope.TieBreaker.Submitted = false
			$scope.TieBreaker.Submitting = true

			SubmitService.tiebreaker($rootScope.User.Id, $scope.Week.Week, $scope.Week.Year, $scope.TieBreaker.Value)
				.success(function(data, status) {
					$scope.TieBreaker.Submitted = true
					$scope.TieBreaker.Submitting = false
					$log.log("SubmitService.tiebreaker", "data", data)
				})
				.error(function(data, status) {
					$log.log("SubmitService.tiebreaker", "data", data)
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
			SubmitService.submit($rootScope.User.Id, picks)
				.success(function(data, status) {
					$log.log("SubmitService", "data", data)
				})
				.error(function(data, status) {
					$log.log("SubmitService.submit", "status", status, "data", data)
				})
		}

		SubmitService.currentLines($rootScope.User.Id)
			.success(function(data, status) {
				$scope.Week = data.result.Week
				$scope.Lines = data.result.Lines
				$scope.TieBreaker.Value = data.result.TieBreaker.Value
				var picks = data.result.Picks
				if (picks == null) {
					return
				}
				var myPicks = {
					// Since this is the initial load, don't let @watch re-send the existing picks
					initial: true
				}
				for (var i = 0; i < picks.length; i++) {
					myPicks[picks[i].GameId] = picks[i].Value
				}
				$scope.Picks = myPicks
			})
			.error(function(data, status) {
				$log.log("SubmitService.currentLines", "status", status, "data", data)
			})
	}
])
