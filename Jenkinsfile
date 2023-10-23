pipeline {
  agent any
  tools { go '1.21.0' }
  stages {
    stage('checkout branch') {
      steps {
        git(url: 'https://github.com/edunota/edunota-backend', branch: 'main')
      }
    }

    stage('test') {
      steps {
        sh 'go version'
      }
    }

  }
}
