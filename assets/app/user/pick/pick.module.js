"use strict"

angular.module("Picks.User.Pick", [
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/user/pick", {
		controller:  "Picks.User.PickController",
		templateUrl: "/app/user/pick/pick.partial.html",
		auth:        function(user) {
			return angular.isDefined(user)
		}
	})
}])

