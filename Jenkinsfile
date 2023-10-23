pipeline {
  agent any
  stages {
    stage('checkout branch') {
      steps {
        git(url: 'https://github.com/edunota/edunota-backend', branch: 'main')
      }
    }

  }
}