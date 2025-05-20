pipeline {
    agent any
    options {
        skipDefaultCheckout()  // ปิด auto checkout ของ Jenkins
    }
    stages {
        stage('Build') {
            steps {
                // ตอนนี้ Jenkins ยังไม่ได้ checkout repo อัตโนมัติ
                // คุณต้องสั่ง checkout เองถ้าต้องการไฟล์จาก repo
                git(
                    url: 'https://github.com/jarntae/Project-serverless.git',
                    branch: 'main',
                    credentialsId: '9f140c33-5219-411d-88a1-bf6c77e17bcc'
                )
                sh 'docker-compose down || true'
                sh 'docker-compose up --build -d'
            }
        }
    }
}
