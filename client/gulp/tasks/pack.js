var gulp = require('gulp');
var zip = require('gulp-zip');
var config = require('../config').deploy;

gulp.task('pack', function (cb) {
	return gulp.src(config.packItems)
        .pipe(zip(config.packageName))
        .pipe(gulp.dest('./'));
});