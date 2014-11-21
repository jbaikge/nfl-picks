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
}])

// submit.controller.js
angular.module("Picks.Picks.Submit").controller("Picks.Picks.SubmitController", [
	"$rootScope",
	"$scope",
	"SubmitService",
	function($rootScope, $scope, SubmitService) {
		$scope.Lines = []
		$scope.Picks = {}
		$scope.Progress = {}

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
			console.log("counts:", o, "changed:", changed)

			if (newValue.initial) {
				delete($scope.Picks.initial)
				return
			}

			SubmitService.submit($rootScope.User.Id, changed)
				.success(function(data, status) {
					console.log("SubmitService", "data", data)
				})
				.error(function(data, status) {
					console.log("SubmitService.submit", "status", status, "data", data)
				})
			$scope.Progress = o
		}, true)

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

		SubmitService.currentLines($rootScope.User.Id)
			.success(function(data, status) {
				$scope.Week = data.result.Week
				$scope.Lines = data.result.Lines
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
				console.log("SubmitService.currentLines", "status", status, "data", data)
			})
	}
])
