pipeline {
  agent any
  tools {
    go '1.21.1'
  }
  stages {
    stage('checkout branch') {
      steps {
        git(url: 'https://github.com/edunota/edunota-backend', branch: 'main')
      }
    }

    stage('test') {
      steps {
        sh 'make test'
      }
    }

  }
}
