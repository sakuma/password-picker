module.exports = (grunt)->

  grunt.loadNpmTasks("grunt-contrib-concat")
  grunt.loadNpmTasks("grunt-contrib-uglify")
  grunt.loadNpmTasks("grunt-contrib-cssmin")

  grunt.initConfig
    pkg: grunt.file.readJSON("package.json")
    concat:
      dist:
        src: [
          "bower_components/vue/dist/vue.min.js"
          "bower_components/superagent/superagent.js"
          "bower_components/jquery/dist/jquery.js"
          "bower_components/materialize/dist/js/materialize.js"
        ]
        dest: "assets/lib.js"

    uglify:
      main:
        src: "<%= concat.dist.dest %>",
        dest: "assets/lib.min.js"

    cssmin:
      main:
        src: [
          "bower_components/materialize/dist/css/materialize.css"
          "assets/css/style.css"
        ]
        dest: "assets/application.min.css"

  grunt.registerTask('default', ['concat', 'uglify', 'cssmin'])
