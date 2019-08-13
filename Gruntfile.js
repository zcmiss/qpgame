//nodejs在redhat下面的安装
//curl -sL https://rpm.nodesource.com/setup_11.x | bash -
//https://github.com/nodesource/distributions
module.exports = function(grunt) {
    // Project configuration.
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        apidoc: {
            frontEnd: {
                src: ["./app/frontEndControllers","./app/silverMerchant"],
                dest: "./assets/qpApi1.0.0",
                options:{
                    includeFilters:[".go$"],
                    log:true,
                }
            },
            admin:{
                src: ["./admin/controllers"],
                dest: "./assets/adminQpApi1.0.0",
                options:{
                    includeFilters:[".go$"],
                    log:true,
                }
            }
        },
        watch:{
            apidoc:{
                files: "./app/**/*.go",
                tasks: ['apidoc:frontEnd']
            },
            adminApidoc:{
                files: "./admin/controllers/**/*.go",
                tasks: ['apidoc:admin']
            },
        }
    });
    grunt.loadNpmTasks('grunt-apidoc');
    grunt.loadNpmTasks('grunt-contrib-watch');
};