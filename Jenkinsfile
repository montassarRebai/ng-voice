pipeline { 
     agent any
     tools {
        go 'go1.19'
    }
     environment { 
         IMAGE_VERSION = "${env.BUILD_NUMBER}" 
         IMAGE_NAME = "aze012/go-app:${IMAGE_VERSION}" 
         GO119MODULE = 'on'
     } 
     
     stages { 
      stage('Compile') {
            steps {
                sh 'go build -o ./go-app .' 
            }
        }
         stage("Build docker image"){ 
             steps{ 
                 sh "docker build --tag=${IMAGE_NAME} ." 
             } 
         } 
         stage("Push to DockerHub") {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                sh "docker push ${IMAGE_NAME}"
                }
            } 
         }
         stage("Deploy Helm chart"){ 
             steps{
               script {
              def remote = [:]
              remote.name = 'master'
              remote.host = '192.168.1.24'
              remote.user = 'root'
              remote.password = 'monta'
              remote.allowAnyHosts = true
    
              sshCommand remote:remote, command: 'helm upgrade --install go-k8s . --set app.namespace=ng-voice --set image.tag=${IMAGE_VERSION}'
             
           
          
        }
            }
         } 
     } 
     post { 
         always { 
                     cleanWs() 
                 } 
           } 
 }
