"use strict"

angular.module("Picks.User.Login", [
	"ngCookies",
	"ngRoute"
])

.config(["$routeProvider", function($routeProvider) {
	$routeProvider.when("/user/login", {
		controller:  "Picks.User.LoginController",
		templateUrl: "/app/user/login/login.partial.html"
	})
}])

// login.service.js
angular.module("Picks.User.Login").service("LoginService", ["jsonrpc", function(jsonrpc) {
	this.auth = function(username, pin) {
		var data = {
			Username: username,
			PIN: pin
		}
		return jsonrpc("User.Auth", data)
	}

	this.usernames = function() {
		return jsonrpc("User.Usernames")
	}

	this.picksClosed = function() {
		return jsonrpc("Picks.Closed")
	}
}])

// login.controller.js
angular.module("Picks.User.Login").controller("Picks.User.LoginController", [
	"$cookieStore",
	"$location",
	"$log",
	"$rootScope",
	"$scope",
	"LoginService",
	function($cookieStore, $location, $log, $rootScope, $scope, LoginService) {
		var defaultOption = "Select Your Name..."

		$scope.Form = {}
		$scope.PicksClosed = true
		$scope.Usernames = []

		LoginService.picksClosed()
			.success(function(data, status) {
				$scope.PicksClosed = data.Closed
			})
			.error(function(data, status) {
				$log.warn("error status: %s", status)
				$log.warn("error data: ", data)
			})

		LoginService.usernames()
			.success(function(data, status) {
				$scope.Usernames = data.result.Usernames
				$scope.Usernames.unshift(defaultOption)
				$scope.Form.Username = defaultOption
			})
			.error(function(data, status) {
				$log.warn("error status: %s", status)
				$log.warn("error data: ", data)
			})

		$scope.doLogin = function() {
			$scope.Alert = {}

			if (!angular.isDefined($scope.Form.Username) || $scope.Form.Username == defaultOption) {
				$scope.Alert.Danger = "Please select your name"
				return
			}

			if (!angular.isDefined($scope.Form.PIN)) {
				$scope.Alert.Danger = "Please provide your PIN"
				return
			}

			LoginService.auth($scope.Form.Username, $scope.Form.PIN)
				.success(function(data, status) {
					if (data.error != null) {
						$scope.Alert.Danger = data.error
						return
					}

					var user = {
						Id:       data.result.Id,
						IsAdmin:  data.result.IsAdmin,
						Username: data.result.Username,
						Theme:    data.result.Theme,
						Beer:     data.result.Beer
					}
					$rootScope.User = user
					$rootScope.Theme = user.Theme
					$cookieStore.put("User", user)

					if ($scope.PicksClosed) {
						$location.path("/picks/viewall")
					} else {
						$location.path("/picks/submit")
					}
				}).error(function(data, status) {
					$log.warn("error status: %s", status)
					$log.warn("error data: ", data)
				})
		}
	}
])
