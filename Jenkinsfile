pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh 'echo \'build step\''
        git(url: 'https://github.com/edunota/edunota-backend', branch: 'main')
      }
    }

  }
}