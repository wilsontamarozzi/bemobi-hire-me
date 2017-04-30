var gulp = require('gulp');
var config = require('../config').deploy;
var es = require('event-stream');
var rename = require('gulp-rename');

/* Copia os fonts para Dist */
gulp.task('copy-other-files', function () {
	var index = gulp.src('index-prod.html')
		.pipe(rename('index.html'))
		.pipe(gulp.dest(config.distDirectory));

	var images = gulp.src('img/**/*.*')
		.pipe(gulp.dest(config.distDirectory + 'img'));

	var fonts = gulp.src('node_modules/bootstrap/fonts/*.*')
		.pipe(gulp.dest(config.distDirectory + 'fonts'));

	var views = gulp.src('view/**/*.html')
		.pipe(gulp.dest(config.distDirectory + 'view'));

	/*var startServer = gulp.src('start-server.bat')
		.pipe(gulp.dest(config.distDirectory));*/

	return es.merge(index, images, fonts, views);
});