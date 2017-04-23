( function() {
	'use strict';

	angular
		.module('appShortenUrl')
		.factory('ShortenUrlService', Service);

	function Service($http, ConfigApp) {

		function _getUrlRanking(filters) {
			var urlPath = ConfigApp.baseUrl + 'api/v1/url/ranking';

			if(filters) {
				urlPath += '?';
				angular.forEach(filters, function(value, key) {
					if(value && key) {
			        	urlPath += '&';
			        	urlPath += key + '=' + value;
			        }
			    });
			}

	        return $http.get(urlPath);
		}

		function _getUrlByAlias(alias) {
			return $http.get(ConfigApp.baseUrl + 'api/v1/url/details/' + alias);
		}

		function _postUrlShorten(url) {
			return $http.post(ConfigApp.baseUrl + 'api/v1/url/shorten', url);	
		}

		return {
			getUrlRanking : _getUrlRanking,
			getUrlByAlias : _getUrlByAlias,
			postUrlShorten : _postUrlShorten,
		};
	}
})();