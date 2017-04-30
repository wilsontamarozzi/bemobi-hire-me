var gulp = require('gulp');
var runSequence = require('run-sequence');

gulp.task('build', function (callback) {
	return runSequence('clean', 'jshint', [
		'concat-lib',
		'compress-js',
		'compress-css',
	], 'copy-other-files', 'pack', callback);
});
