(function (window) {
	'use strict';

	var app = angular
		.module('appShortenUrl', ['ngRoute', 'ngMessages', 'ngStorage'])
		.run(run);

	function run($rootScope, $window) {
		// Seta um variavel global no projeto com o dominio para usar 
		// de referencia por todos lugares e facilitar a manutenção
		$rootScope.rootScopeHostname = $window.location.hostname + "/";
	}
})();