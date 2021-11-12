pipeline{
    agent any
    tools{
        go 'go-tool'
    }
    environment{
        GO111MODULE = 'on'
    }
    stages{
        stage("Checkout"){
            steps{
                checkout scm
            }
        }
        stage("Test"){
            parallel {
                stage('crawling test'){
                    steps{
                        dir('src/crawling'){
                            sh('make test')
                        }
                    }
                }
                stage('menu test'){
                    steps{
                        dir('src/menu'){
                            sh('make test')
                        }
                    }
                }
                stage('skill server test'){
                    steps{
                        dir('src/skill-server'){
                            sh('make test')
                        }
                    }
                }
            }
        }
    }
}