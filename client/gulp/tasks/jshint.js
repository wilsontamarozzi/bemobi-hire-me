var gulp = require('gulp');
var jshint = require('gulp-jshint');
var config = require('../config').deploy;

/* Verifica erros de syntax no javascript */
gulp.task('jshint', function () {
	return gulp.src(config.testHintItems)
		.pipe(jshint())
		.pipe(jshint.reporter('default'));
});