var gulp = require('gulp');
var config = require('../config').deploy;
var concat = require('gulp-concat');
var cssmin = require('gulp-cssmin');
var es = require('event-stream');

gulp.task('compress-css', function () {
	return gulp.src(config.compressCssItems)
		.pipe(concat(config.compressCssName))
		.pipe(cssmin())
		.pipe(gulp.dest(config.distDirectory + 'css'));
});