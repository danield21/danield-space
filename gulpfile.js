const gulp = require('gulp');
const sass = require('gulp-sass');
const concat = require('gulp-concat');
const babel = require('gulp-babel');
const browserify = require('browserify');
const babelify = require('babelify');
const source = require('vinyl-source-stream');
const buffer = require('vinyl-buffer');
const uglify = require('gulp-uglify');
const sourcemaps = require('gulp-sourcemaps');
const vulcanize = require('gulp-vulcanize');

gulp.task('sass', function() {
	return gulp.src('web/sass/*.scss')
		.pipe(sass({
			'sourcemap=none': true
		}))
		.pipe(concat('app.css'))
		.pipe(gulp.dest('app/dist/css'))
});

gulp.task('templates', function() {
    return gulp.src('web/templates/src/nm-elements.html')
        .pipe(vulcanize())
        .pipe(gulp.dest('app/dist/components'));
});

gulp.task('js', function () {
    // app.js is your main JS file with all your module inclusions
    return browserify({entries: './web/js/app.js', debug: true})
        .transform("babelify", { presets: ["es2015"] })
        .bundle()
        .pipe(source('app.js'))
        .pipe(buffer())
        .pipe(sourcemaps.init())
        .pipe(uglify())
        .pipe(sourcemaps.write('./maps'))
        .pipe(gulp.dest('app/dist/js'));
});

gulp.task('copyJs', function () {
    return gulp.src([
        "./web/templates/bower_components/webcomponentsjs/webcomponents-lite.min.js"
    ]).pipe(gulp.dest("app/dist/js"))
});

gulp.task('copyView', function () {
    return gulp.src([
        "./web/view/**"
    ]).pipe(gulp.dest("app/view"))
});

gulp.task('default', ["templates", "js", "sass", "copyJs", "copyView"])