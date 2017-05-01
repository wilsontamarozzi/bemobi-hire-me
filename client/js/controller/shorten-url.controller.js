( function() {
	'use strict';

	angular
		.module('appShortenUrl')
		.controller('ShortenCtrl', Controller);

	function Controller($scope, ShortenUrlService) {

		$scope.saveUrl = function() {
			ShortenUrlService.postUrlShorten($scope.url).then(function success(response) {
				$scope.success = true;
				$scope.url = response.data;
			}, function fail(response) {
				$scope.success = false;
				switch(response.status) {
				case 409:
					$scope.error = 'Esse endereço opcional já existe ;(';
				break;
				default:
					$scope.error = 'Houve um erro ao salvar, mas já estámos trabalhando nisso.';
				break;
			}
			});
		};
		
		$scope.url = {};
	}
})();