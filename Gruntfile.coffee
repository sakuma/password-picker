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
        ]
        dest: "assets/lib.js"

    uglify:
      main:
        src: "<%= concat.dist.dest %>",
        dest: "assets/lib.min.js"

    cssmin:
      main:
        src: [
          "assets/css/normalize.css"
          "assets/css/skeleton.css"
        ]
        dest: "assets/lib.min.css"

  grunt.registerTask('default', ['concat', 'uglify', 'cssmin'])
