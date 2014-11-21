"use strict"

angular.module("quarter", []).filter("quarter", function() {
	var quarter = function(input) {
		switch (input) {
			case "P":  return "@"
			case "F":  return "Final"
			case "FO": return "OT"
			case "H":  return "Half"
		}

		// Only process numeric values.
		if (isNaN(input) || input === null) return input;

		return "Q" + input
	}

	return quarter
})