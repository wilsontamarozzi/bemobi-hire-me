module.exports = {
	deploy: {
		distDirectory : 'dist/',
		testHintItems : ['js/**/*.js'],
	    packageName: 'zipped_dist.zip',
	    packItems: ['dist/**'],
	    concatLibJsName : 'lib.min.js',
	    concatLibJsItems : [
			'node_modules/angular/angular.min.js',
			'node_modules/angular-route/angular-route.min.js',
			'node_modules/angular-messages/angular-messages.min.js',
			'node_modules/ngstorage/ngStorage.min.js',
		],
		compressJsName : 'all.min.js',
		compressJsItems : [
			'js/app.js',
			'js/**/*.js'
		],
		compressCssName : 'styles.min.css',
		compressCssItems : [
			'node_modules/bootstrap/dist/css/bootstrap.min.css',
			'css/layout.css',
		]
    }
};