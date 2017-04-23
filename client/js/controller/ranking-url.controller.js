( function() {
	'use strict';

	angular
		.module('appShortenUrl')
		.controller('RankingCtrl', Controller);

	function Controller($scope, ShortenUrlService) {

		$scope.loadUrlsRanking = function() {
			$scope.filters.page = $scope.currentPage;
			$scope.filters.per_page = itemPerPage;

			ShortenUrlService.getUrlRanking($scope.filters).then(function success(response) {
				$scope.lastPage = response.data.meta.pagination.total_pages;
				$scope.urls = response.data.urls;
			}, function fail(response) {
				$scope.error = response.data.Description;
			});
		};
		
		var itemPerPage = 10;
		$scope.lastPage = 0;
		$scope.currentPage = 1;
		$scope.filters = {};
		$scope.urls = [];
		$scope.loadUrlsRanking();
	}
})();