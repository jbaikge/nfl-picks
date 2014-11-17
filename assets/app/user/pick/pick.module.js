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
}])

// pick.controller.js
angular.module("Picks.User.Pick").controller("Picks.User.PickController", [
	"$scope",
	"PickService",
	function($log, $scope, PickService) {
	
	}
])
