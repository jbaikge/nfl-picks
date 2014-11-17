"use strict"

angular.module("Picks.User.Login").controller("Picks.User.LoginController", [
	"$log",
	"$scope",
	"$rootScope",
	"LoginService",
	function($log, $scope, $rootScope, LoginService) {
		$log.log("Hello from LoginController")
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
					$rootScope.User = {
						Id:       data.result.Id,
						IsAdmin:  data.result.IsAdmin,
						Username: data.result.Username
					}
				}).error(function(data, status) {
					$log.warn("error status: %s", status)
					$log.warn("error data: ", data)
				})
		}
	}
])
