"use strict"

angular.module("jsonrpc", [])

.provider("jsonrpc", function() {
	var path = "/rpc"

	this.setPath = function(rpcPath) {
		path = rpcPath
	}

	this.$get = [ "$http", function($http) {
		function doRequest(method, data) {
			var id = method + " " + (new Date).toJSON()
			var payload = {
				jsonrpc: "2.0",
				id:      id,
				method:  method,
				params:  []
			}
			if (angular.isDefined(data)) {
				payload.params.push(data)
			}
			return $http.post(path, payload)
		}
		return doRequest
	}]
})
