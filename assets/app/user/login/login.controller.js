"use strict"

angular.module("Picks.User.Login").controller("Picks.User.LoginController", [
	"$cookieStore",
	"$location",
	"$log",
	"$rootScope",
	"$scope",
	"LoginService",
	function($cookieStore, $location, $log, $rootScope, $scope, LoginService) {
		$scope.Form = {}

		$scope.Usernames = []
		LoginService.usernames()
			.success(function(data, status) {
				$scope.Usernames = data.result.Usernames
				var defaultOption = "Chose your name"
				$scope.Usernames.unshift(defaultOption)
				$scope.Form.Username = defaultOption
			})
			.error(function(data, status) {
				$log.warn("error status: %s", status)
				$log.warn("error data: ", data)
			})

		$scope.doLogin = function() {
			if (!angular.isDefined($scope.Form.Username)) {
				return
			}

			if (!angular.isDefined($scope.Form.PIN)) {
				return
			}

			LoginService.auth($scope.Form.Username, $scope.Form.PIN)
				.success(function(data, status) {
					var user = {
						Id:       data.result.Id,
						IsAdmin:  data.result.IsAdmin,
						Username: data.result.Username
					}
					$rootScope.User = user
					$cookieStore.put("User", user)
					$location.path("/user/pick")
				}).error(function(data, status) {
					$log.warn("error status: %s", status)
					$log.warn("error data: ", data)
				})
		}
	}
])
