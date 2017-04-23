( function() {
	'use strict';

	angular
		.module('appShortenUrl')
		.controller('RetrieveCtrl', Controller);

	function Controller($window, $scope, $routeParams, ShortenUrlService) {
		
		$scope.alias = $routeParams.alias;

		ShortenUrlService.getUrlByAlias($routeParams.alias).then(function success(response) {
			$window.location.href = response.data.url.address;
		}, function error(response) {
			switch(response.status) {
				case 404:
					$scope.error = 'Desculpe não encontramos ;(';
				break;
				default:
					$scope.error = 'Não foi possível carregar os dados!';
				break;
			}
		});
	}
})();