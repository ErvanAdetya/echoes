pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
            when {
                changeRequest target: 'trunk'
            }
        }

        stage('Run unit test') {
            steps {
                sh 'docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 go test -v'
            }
            when {
                changeRequest target: 'trunk'
            }
        }

        stage('Build') {
            steps {
                sh 'docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 go build -v'
            }
            when {
                changeRequest target: 'trunk'
            }
        }
    }
}
