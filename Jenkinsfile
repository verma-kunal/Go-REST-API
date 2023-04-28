pipeline {
    agent any
    tools{
        go 'go1.20'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = `${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}`
        DOCKERHUB = "vkunal"
    }
    stages {
        stage('Checkout') {
            checkout scm
        }
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u github.com/gin-gonic/gin'
	            sh 'go get -u github.com/stretchr/testify'
            }
        }
        stage('Unit-tests') {
            steps {
                echo "Running unit tests"
                sh 'go test ./src/server_test.go'
            }
        }
        stage('Build') { 
            steps {
                echo "Compiling and building the app"
                sh 'go build -o go-rest-api ./src/main.go'
            }
        }
        stage('SonarQube Code analysis') {
            def scannerHome = tool 'SonarScanner 4.0';
            steps {
                withSonarQubeEnv('sonarqube-server') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
            // send alerts/notification:
            post {
                always {
                    script{
                        def reportTasks = scanForIssues tool: sonarqubeServer(
                        serverId: 'sonarqube-server',
                        installationName: 'SonarQube'
                        ), type: 'BUG,VULNERABILITY'

                        if (reportTasks.size() > 0) {
                            emailext subject: "SonarQube Scan Result for ${env.JOB_NAME} - ${env.BUILD_NUMBER}",
                            body: "${Jenkins.getInstance().getRootUrl()}job/${env.JOB_NAME}/${env.BUILD_NUMBER}/console\n\n${reportTasks.size()} issues detected.",
                            recipientProviders: [[$class: 'DevelopersRecipientProvider']]
                        }
                    }
                }
            }
        }
        stage('Docker Build') {
            steps {
                sh 'docker build -t go-app .'
            }
        }
        stage('Push Docker image') {
            steps {
                withCredentials([string(credentialsId: 'dockerhub', variable: 'DOCKERHUB')]) {
                    sh 'docker login -u $DOCKERHUB'
                    sh 'docker push $DOCKERHUB/go-app'
                }
            }
        }
    }
}
