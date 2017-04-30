var gulp = require('gulp');
var config = require('../config').deploy;
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');

gulp.task('compress-js', function () {
	return gulp.src(config.compressJsItems)
		.pipe(concat(config.compressJsName))
		.pipe(uglify({mangle: false}))
		.pipe(gulp.dest(config.distDirectory + 'js'));
});