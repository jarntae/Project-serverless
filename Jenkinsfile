pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                checkout scm  // checkout โค้ดจาก repo
                sh 'docker-compose down || true'
                sh 'docker-compose up --build -d'
            }
        }
    }
}


