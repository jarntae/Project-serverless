pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'docker-compose down || true'  // ป้องกัน error ถ้า container ยังไม่เคยรัน
                sh 'docker-compose up --build -d'
            }
        }
    }
}
