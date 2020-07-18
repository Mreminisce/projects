let mix = require('laravel-mix');


mix
  .setStaticPath('static/')
  .ts('resources/js/app.ts', 'static/js')
  .sass('resources/sass/app.scss', 'static/css').version();
