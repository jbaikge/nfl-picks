"use strict"

angular.module("Picks.Picks.Submit", [
	"ngRoute"
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

	this.submit = function(id, value) {
		return jsonrpc("Picks.Submit", {Id: id, Value: value})
	}
}])

// pick.controller.js
angular.module("Picks.Picks.Submit").controller("Picks.Picks.SubmitController", [
	"$scope",
	"SubmitService",
	function($scope, SubmitService) {
		$scope.Lines = []

		SubmitService.currentLines()
			.success(function(data, status) {
				$scope.Current = data.result.Current
				$scope.Lines = data.result.Lines
			})
			.error(function(data, status) {
				console.log("SubmitService", "status", status, "data", data)
			})
	}
])
