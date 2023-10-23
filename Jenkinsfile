pipeline {
  agent any
  tools { go '1.21' }
  environment {
      GO114MODULE = 'on'
      CGO_ENABLED = 0 
      GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
  }
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
