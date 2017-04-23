( function() {
	'use strict';

	angular
		.module('appShortenUrl')
		.controller('ShortenCtrl', Controller);

	function Controller($scope, ShortenUrlService) {

		$scope.saveUrl = function(url) {
			ShortenUrlService.postUrlShorten(url).then(function success(response) {
				$scope.success = true;
				$scope.url = response.data;
			}, function fail(response) {
				$scope.error = response.data.Description;
			});
		};
		
		$scope.url = {};
	}
})();