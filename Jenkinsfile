pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Run unit test') {
            steps {
                sh 'docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 go test -v'
            }
        }

        stage('Build') {
            steps {
                sh 'docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.8 go build -v'
            }
        }
    }
}
