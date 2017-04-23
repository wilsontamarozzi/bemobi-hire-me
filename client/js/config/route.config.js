( function() {

	angular
		.module('appShortenUrl')
		.config(Config);

	function Config($routeProvider, $locationProvider) {
		$routeProvider.when('/', {
			templateUrl: 'view/url-new.view.html',
			controller: 'ShortenCtrl'
		});

		$routeProvider.when('/ranking', {
			templateUrl: 'view/url-ranking.view.html',
			controller: 'RankingCtrl'
		});

		$routeProvider.when('/:alias', {
			templateUrl: 'view/url-redirect.view.html',
			controller: 'RetrieveCtrl'
		});

		$routeProvider.otherwise({redirectTo: '/'});
		//$locationProvider.html5Mode(true);
		//$locationProvider.hashPrefix('!');
	}
})();