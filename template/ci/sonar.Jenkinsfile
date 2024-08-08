@Library('zing-jenkins') _

MAKE = 'make -f ci/Makefile'

node('docker') {
    currentBuild.displayName = "#${env.BUILD_NUMBER} @${env.NODE_NAME})"

    configFileProvider([
        configFile(fileId: 'global', variable: 'GLOBAL'),
    ]) {
        global = load env.GLOBAL
    }

    stage('Checkout') {
        checkout scm
    }

    try {
        stage('Tests') {
            ansiColor('xterm') {
                sh("${MAKE} unit-test-containerized")
            }
        }

        stage('Vulnerabilities') {
            ansiColor('xterm') {
                sh("${MAKE} vuln-sarif")
            }
        }

        stage('Update SonarQube Main Branch') {
            SHA = sh(script: 'git rev-parse HEAD | cut -c 1-8', returnStdout: true).trim()
            branch = sh(script: 'git rev-parse --abbrev-ref HEAD', returnStdout: true).trim()

            String scannerHome = tool 'SonarScanner'
            String cmdline = [
                "${scannerHome}/bin/sonar-scanner",
                '-Dproject.settings=./ci/sonar-project.properties',
                "-Dsonar.projectVersion=${SHA}",
                "-Dsonar.branch.name=${branch}",
            ].join(' ')

            withSonarQubeEnv('sonarqube.zenoss.io') {
                sh(cmdline)
            }
        }
    } finally {
        stage('Clean test environment') {
            sh("${MAKE} down")
        }
    }
}
