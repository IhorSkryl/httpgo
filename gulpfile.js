'use strict';

const port = '8080';

const gulp = require('gulp');
const rename = require('gulp-rename');
const notify = require("gulp-notify");
const notifier = require('node-notifier');

/* babel */
const babel = require('gulp-babel');

const jsmin = require('gulp-jsmin');


/* native build in single file */
const concat = require('gulp-concat');

gulp.task('native-js', () => {
  let isBabel = babel({
    presets: [require('babel-preset-es2015')]
  });

  isBabel.on('error', function(e) {
    console.log(e);
    isBabel.end();
    notifier.notify(`error JS: ${ e.message }`);
  });

  gulp.src(`./nativeJS/**/*.js`)
    .pipe(isBabel)
    .pipe(concat('native.min.js'))
    // .pipe(jsmin())
    .pipe(rename({dirname: ''}))
    .pipe(gulp.dest(`./js/`));
    // .pipe(open({uri: recacheURL}))


  require('child_process').exec(`curl localhost:${ port }/recache`, function (err, stdout, stderr) {
    console.log(stdout);
    console.log(stderr);
    notifier.notify(`recache: ${ stdout }`)
  });

});


/* watch */
gulp.task('watch', () => {
  gulp.watch([`./nativeJS/**/*`], ['native-js']);
});


/* default */
gulp.task('default', ['native-js', 'watch']);