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

var basePaths = {};
var paths = {
  src: {
    js: [],
    css: [],
    html: []
  },
  dst: {

  },
  clean: {

  }
};

var sourcePaths = {
  distFiles: [
    './sqliteweb-web/js/*.min.js',
    './sqliteweb-web/css/*.css',
    './sqliteweb-web/*.html',
    './sqliteweb-web/vendor/ace/**/*',
    '!sqliteweb-web/bower_components'
  ],
  filesToClean: [
    './sqliteweb-web/js/*.js',
    './sqliteweb-web/css/*.css'
  ],
  distFilesToClean: [
    './sqliteweb-server/static/js/*.js',
    './sqliteweb-server/static/css/*.css',
    './sqliteweb-server/static/*.html',
    './sqliteweb-server/static/ace/**/*.js'
  ]
};

var server = {
  host: 'localhost',
  port: '8001'
};

gulp.task('clean', function(cb) {
    gulp.src(sourcePaths.filesToClean, {read: false})
        .pipe(clean());
});

gulp.task('sass', function() {
  gulp.src('sqliteweb-web/css/app.scss')
    .pipe(plumber())
    .pipe(sass())
    .pipe(csso())
    .pipe(gulp.dest('sqliteweb-web/css'));
});

// Lint Coffeescript files
gulp.task('lint', function () {
    gulp.src('./sqliteweb-web/js/app.coffee')
        .pipe(coffeelint())
        .pipe(coffeelint.reporter())
});

// Compile Coffeescript files
gulp.task('coffee', function() {
  gulp.src('./sqliteweb-web/js/app.coffee')
    .pipe(coffee({bare: true}).on('error', gutil.log))
    .pipe(gulp.dest('./sqliteweb-web/js'));
});

gulp.task('compress', function() {
  gulp.src([
    'sqliteweb-web/bower_components/jquery/dist/jquery.js',
    'sqliteweb-web/bower_components/bootstrap-sass-official/assets/javascripts/bootstrap.js',
    'sqliteweb-web/bower_components/lodash/dist/lodash.js',
    'sqliteweb-web/js/*.js',
    '!sqliteweb-web/js/app.min.js'
  ])
    .pipe(concat('app.min.js'))
    .pipe(uglify())
    .pipe(gulp.dest('sqliteweb-web/js'));
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

// Clean dist files
gulp.task('cleanDist', function () {
  del(sourcePaths.distFilesToClean);
});

// Copy all dist files to the server folder.
gulp.task('copyDistFiles', function () {
  gulp.src(sourcePaths.distFiles, {base: "./sqliteweb-web"})
    .pipe(gulp.dest('./sqliteweb-server/static'));
});

gulp.task('default', ['clean', 'sass', 'lint', 'coffee', 'compress', 'webserver', 'watch', 'openbrowser']);
gulp.task('dev', ['sass', 'lint', 'coffee', 'compress', 'cleanDist', 'copyDistFiles', 'watchDist']);
//gulp.task('dist', ['sass', 'lint', 'coffee', 'compress', 'cleanDist', 'copyDistFiles']);
