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
}])

// viewall.controller.js
angular.module("Picks.Picks.Submit").controller("Picks.Picks.ViewallController", [
	"$rootScope",
	"$scope",
	"ViewallService",
	function($rootScope, $scope, ViewallService) {
		$scope.Users = []
		$scope.Games = []
		$scope.Picks = {}

		ViewallService.users()
			.success(function(data, status) {
				$scope.Users = data.result.Usernames
			})

		ViewallService.lines()
			.success(function(data, status) {
				$scope.Games = data.result.Lines
			})

		ViewallService.picks()
			.success(function(data, status) {
				$scope.Picks = data.result.Picks
			})
	}
])
