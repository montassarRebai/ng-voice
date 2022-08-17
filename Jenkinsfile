pipeline { 
     agent any
     tools {
        go 'go1.19'
    }
     environment { 
         IMAGE_VERSION = "${env.BUILD_NUMBER}" 
         IMAGE_NAME = "aze012/go-app:${IMAGE_VERSION}" 
         GO119MODULE = 'on'
         GIT_CREDS = credentials('github-cred')
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
         stage('Updating Helm values file'){
            steps {
                sshagent(credentials:['ssh-jenkins']){
               sh 'ssh  -o StrictHostKeyChecking=no  root@192.168.1.24 uptime "whoami"'
               sh 'yq eval '.image.tag = ${IMAGE_VERSION}' -i /home/monta/Desktop/go-k8s-helm/go-k8s/values.yaml'
               sh 'cat /home/monta/Desktop/go-k8s-helm/go-k8s/values.yaml'
               sh 'cd /home/monta/Desktop/go-k8s-helm/go-k8s/ ; sh git.sh'
          }
       
                
            }
        }
     } 
     //post { 
        // always { 
                     //cleanWs() 
                 //} 
          // } 
 }
