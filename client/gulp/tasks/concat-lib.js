var gulp = require('gulp');
var concat = require('gulp-concat');
var config = require('../config').deploy;

/* Unifica todas libs externas */
gulp.task('concat-lib', function () {
	return gulp.src(config.concatLibJsItems)
		.pipe(concat(config.concatLibJsName))
		.pipe(gulp.dest(config.distDirectory + 'js'));
});