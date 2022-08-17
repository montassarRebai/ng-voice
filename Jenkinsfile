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
         HELM_GIT_REPO_URL = "github.com/montassarRebai/ng-voice-helm.git"
         GIT_REPO_EMAIL = 'montassar.rebai@esprit.tn'
         GIT_REPO_BRANCH = "main"
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
         //stage("Push to DockerHub") {
            //steps {
               // withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
               // sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
               // sh "docker push ${IMAGE_NAME}"
               // }
           // } 
        // }
         stage('Updating Helm values file'){
            steps {
                script {
              def remote = [:]
              remote.name = 'master'
              remote.host = '192.168.1.24'
              remote.user = 'root'
              remote.password = 'monta'
              remote.allowAnyHosts = true
    
             sshCommand remote:remote, command: "pwd"
             
              sshCommand remote:remote, command: "yq eval '.image.tag = ${IMAGE_VERSION}' -i /home/monta/Desktop/go-k8s-helm/go-k8s/values.yaml"
          
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
