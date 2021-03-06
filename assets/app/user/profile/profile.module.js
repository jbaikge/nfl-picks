"use strict"

angular.module("Picks.User.Profile", [
	"ngCookies",
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/user/profile", {
		controller:  "Picks.User.ProfileController",
		templateUrl: "/app/user/profile/profile.partial.html",
		auth:        function(user) { return angular.isDefined(user) }
	})
}])

// profile.service.js
angular.module("Picks.User.Profile").service("ProfileService", ["jsonrpc", function(jsonrpc) {
	this.update = function(id, newUsername, beer, oldPIN, newPIN, theme) {
		return jsonrpc("User.Update", {
			Id:          id,
			NewUsername: newUsername,
			Beer:        beer,
			OldPIN:      oldPIN,
			NewPIN:      newPIN,
			Theme:       theme
		})
	}
}])

// profile.controller.js
angular.module("Picks.User.Profile").controller("Picks.User.ProfileController", [
	"$cookieStore",
	"$rootScope",
	"$scope",
	"ProfileService",
	function($cookieStore, $rootScope, $scope, ProfileService) {
		$scope.Form = {
			Username: $rootScope.User.Username,
			Beer:     $rootScope.User.Beer,
			Theme:    $rootScope.Theme
		}

		$scope.doUpdate = function() {
			var f = $scope.Form
			ProfileService.update($rootScope.User.Id, f.Username, f.Beer, f.OldPIN, f.NewPIN, f.Theme)
				.success(function(data, status) {
					$scope.Alert = {}
					if (data.error != null) {
						$scope.Alert.Danger = data.error
						return
					}
					$rootScope.User.Username = f.Username
					$rootScope.User.Theme = f.Theme
					$rootScope.User.Beer = f.Beer
					$rootScope.Theme = f.Theme
					$cookieStore.put("User", $rootScope.User)
					$scope.Alert.Success = "Information updated"
				})
				.error(function(data, status) {
					console.log("Error", "status", status, "data", data)
				})
		}
	}
])
