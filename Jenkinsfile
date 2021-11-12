pipeline{
    agent any
    stages{
        stage("Checkout"){
            steps{
                checkout scm
            }
        }
        stage("Test"){
            steps{
                dir('src/crawling'){
                    sh('make test')
                }
            }
            steps{
                dir('src/menu'){
                    sh('make test')
                }
            }
            steps{
                dir('src/skill-server'){
                    sh('make test')
                }
            }
        }
    }
}