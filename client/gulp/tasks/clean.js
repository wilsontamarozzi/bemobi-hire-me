var gulp = require('gulp');
var clean = require('gulp-clean');
var es = require('event-stream');
var config = require('../config').deploy;

gulp.task('clean', function () {
    var zipFile = gulp.src(config.packageName, {read: false})
    	.pipe(clean());

    var distDir = gulp.src(config.distDirectory)
		.pipe(clean());

	return es.merge(zipFile, distDir);
});
