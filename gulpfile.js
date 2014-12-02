var gulp = require('gulp');
var sass = require('gulp-sass');
var csso = require('gulp-csso');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');
var plumber = require('gulp-plumber');
var coffee = require('gulp-coffee');
var gutil = require('gulp-util');
var clean = require('gulp-clean');
var coffeelint = require('gulp-coffeelint');
var webserver = require('gulp-webserver');
var opn = require('opn');
var del = require('del');

var basePaths = {
  frontend: './sqliteweb-web',
  server: './sqliteweb-server'
};

var paths = {
  src: {
    js: [
      basePaths.frontend + '/bower_components/jquery/dist/jquery.js',
      basePaths.frontend + '/bower_components/bootstrap-sass-official/assets/javascripts/bootstrap.js',
      basePaths.frontend + '/bower_components/lodash/dist/lodash.js',
      basePaths.frontend + '/js/*.js',
      '!sqliteweb-web/js/app.min.js'
    ],
    coffee: [
      basePaths.frontend + '/js/app.coffee'
    ],
    css: [
      basePaths.frontend + '/css/app.scss'
    ],
    fonts: [
      basePaths.frontend + '/bower_components/font-awesome/fonts/*'
    ],
    dist: [
      basePaths.frontend + '/js/*.min.js',
      basePaths.frontend + '/css/*.css',
      basePaths.frontend + '/fonts/*',
      basePaths.frontend + '/*.html',
      basePaths.frontend + '/vendor/ace/**/*',
      '!sqliteweb-web/bower_components'
    ]
  },
  dest: {
    css: basePaths.frontend + '/css',
    js: basePaths.frontend + '/js',
    fonts: basePaths.frontend + '/fonts',
    dist: basePaths.server + '/static'
  },
  clean: {
    dev: [
      basePaths.frontend + '/js/*.js',
      basePaths.frontend + '/css/*.css',
      basePaths.frontend + '/fonts/*'
    ],
    dist: [
      basePaths.server + '/static/js/*.js',
      basePaths.server + '/static/css/*.css',
      basePaths.server + '/static/*.html',
      basePaths.server + '/static/ace/**/*.js'
    ]
  }
};

var server = {
  host: 'localhost',
  port: '8001'
};

// Clean local dev files
gulp.task('clean', function(cb) {
  del(paths.clean.dev);
});

// Clean dist files
gulp.task('cleanDist', function () {
  del(paths.clean.dist);
});

gulp.task('sass', function() {
  gulp.src(paths.src.css)
    .pipe(plumber())
    .pipe(sass())
    .pipe(csso())
    .pipe(gulp.dest(paths.dest.css));
});

// Lint Coffeescript files
gulp.task('lint', function () {
    gulp.src(paths.src.coffee)
        .pipe(coffeelint())
        .pipe(coffeelint.reporter())
});

// Compile Coffeescript files
gulp.task('coffee', function() {
  gulp.src(paths.src.coffee)
    .pipe(coffee({bare: true}).on('error', gutil.log))
    .pipe(gulp.dest(paths.dest.js));
});

gulp.task('compress', function() {
  gulp.src(paths.src.js)
    .pipe(concat('app.min.js'))
    .pipe(uglify())
    .pipe(gulp.dest(paths.dest.js));
});

gulp.task('watch', function() {
  gulp.watch('sqliteweb-web/css/*.scss', ['sass']);
  gulp.watch(['sqliteweb-web/js/*.coffee'], ['lint', 'coffee', 'compress']);
});

gulp.task('watchDist', function() {
  gulp.watch('sqliteweb-web/css/*.scss', ['sass']);
  gulp.watch(['sqliteweb-web/js/*.coffee'], ['lint', 'coffee', 'compress']);
  gulp.watch(['sqliteweb-web/css/app.css',
              'sqliteweb-web/js/app.min.js',
              'sqliteweb-web/index.html'], ['cleanDist', 'copyDistFiles']);
});

gulp.task('webserver', function() {
  gulp.src('sqliteweb-web')
    .pipe(webserver({
      host: server.host,
      port: server.port,
      livereload: true,
      directoryListing: false
    }));
});

gulp.task('openbrowser', function() {
  opn( 'http://' + server.host + ':' + server.port );
});

// Copy all fonts to the public folder
gulp.task('fonts', function() {
  gulp.src(paths.src.fonts)
    .pipe(gulp.dest(paths.dest.fonts));
});

// Copy all dist files to the server folder.
gulp.task('copyDistFiles', function () {
  gulp.src(paths.src.dist, {base: "./sqliteweb-web"})
    .pipe(gulp.dest(paths.dest.dist));
});

gulp.task('default', ['clean', 'sass', 'lint', 'coffee', 'compress', 'fonts', 'webserver', 'watch', 'openbrowser']);
gulp.task('dev', ['sass', 'lint', 'coffee', 'compress', 'fonts', 'cleanDist', 'copyDistFiles', 'watchDist']);
gulp.task('dist', ['sass', 'lint', 'coffee', 'compress', 'fonts', 'cleanDist', 'copyDistFiles']);
