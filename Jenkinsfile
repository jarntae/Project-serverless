pipeline {
    agent any
    stage('Build') {
        steps {
            checkout scm  // checkout โค้ดที่ Jenkins ดึงมาให้ pipeline
            sh 'docker-compose down || true'
            sh 'docker-compose up --build -d'
    }
}


