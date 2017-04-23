( function() {

	angular
		.module('appShortenUrl')
		.value('ConfigApp', {
			baseUrl: 'http://localhost:8080/'
		});
})();