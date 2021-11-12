pipeline{
    agent any
    stages{
        stage("Checkout"){
            steps{
                checkout scm
            }
        }
        stage("Jenkins test"){
            steps{
                dir("terraform"){
                    sh('pwd')
                }
            }
        }
    }
}