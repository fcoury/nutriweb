module.exports = (grunt) ->


  grunt.initConfig
    pkg: grunt.file.readJSON('package.json')

    coffee:
      app:
        expand: true
        cwd: 'app'
        src: '**/*.coffee'
        dest: 'www/js/app'
        ext: '.js'

    concat:
      jquery:
        files:
          'www/js/vendor/jquery.js': ['app/bower_components/jquery/dist/jquery.js']

      angular:
        files:
          'www/js/vendor/angular.js': ['app/bower_components/angular/angular.js']

    watch:
      app:
        files: '**/*.coffee'
        tasks: ['coffee']

      bower:
        files: 'bower_components/**/*'
        tasks: ['concat']


  grunt.loadNpmTasks 'grunt-contrib-coffee'
  grunt.loadNpmTasks 'grunt-contrib-concat'
  grunt.loadNpmTasks 'grunt-contrib-watch'

  grunt.registerTask 'copy', ['concat']
  grunt.registerTask 'default', ['concat', 'coffee']
