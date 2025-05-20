pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                cleanWs() 
                git(
                    url: 'https://github.com/jarntae/Project-serverless.git',
                    branch: 'main',
                    credentialsId: 'github-pat'
                )
            }
        }

        stage('Build') {
            steps {
                sh 'docker-compose down || true'  // ป้องกัน error ถ้า container ยังไม่เคยรัน
                sh 'docker-compose up --build -d'
            }
        }
    }
}
