"use strict"

angular.module("Picks.User.Pick", [
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/user/pick", {
		controller:  "Picks.User.PickController",
		templateUrl: "/app/user/pick/pick.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// pick.service.js
angular.module("Picks.User.Pick").service("PickService", ["jsonrpc", function(jsonrpc) {
	this.currentGames = function() {
		return jsonrpc("Game.CurrentGames")
	}
}])

// pick.controller.js
angular.module("Picks.User.Pick").controller("Picks.User.PickController", [
	"$scope",
	"PickService",
	function($scope, PickService) {
		$scope.Games = []

		PickService.currentGames()
			.success(function(data, status) {
				$scope.Games = data.result.Games
			})
			.error(function(data, status) {
				console.log("PickService", "status", status, "data", data)
			})
	}
])
