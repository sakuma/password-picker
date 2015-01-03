module.exports = (grunt)->
  pkg = grunt.file.readJSON("package.json")

  for taskName of pkg.devDependencies
    grunt.loadNpmTasks taskName  if taskName.substring(0, 6) is "grunt-"

  grunt.initConfig
    bower_concat:
      all:
        dest: 'tmp/lib.js'
        cssDest: 'tmp/lib.css'
        exclude: []
        bowerOptions:
          relative: false

    uglify:
      lib:
        src: "<%= bower_concat.all.dest %>"
        dest: "public/assets/js/lib.min.js"
      app:
        src: "assets/js/*"
        dest: "public/assets/js/app.min.js"

    cssmin:
      main:
        src: [
          "<%= bower_concat.all.cssDest %>"
          "assets/css/style.css"
        ]
        dest: "public/assets/css/application.min.css"

    copy:
      main:
        files: [
          {
            cwd: 'bower_components/materialize/dist/font/material-design-icons'
            src: ['*']
            dest: 'public/assets/font/material-design-icons'
            expand: true
          },
          {
            cwd: 'bower_components/materialize/dist/font/roboto'
            src: ['*']
            dest: 'public/assets/font/roboto'
            expand: true
          }
        ]

    clean: [
      "<%= bower_concat.all.dest %>"
      "<%= bower_concat.all.cssDest %>"
    ]

  grunt.registerTask('default', ['bower_concat', 'uglify', 'cssmin', 'copy', 'clean'])
