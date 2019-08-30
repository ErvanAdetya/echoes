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

        stage('Build') {
            steps {
                sh 'docker build .'
            }
            when {
                changeRequest target: 'trunk'
            }
        }
    }
}
