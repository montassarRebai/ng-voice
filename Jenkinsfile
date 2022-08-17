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
                 
                 
                 withCredentials([
                  string(credentialsId: 'addr', variable: 'ADDR'),
                  string(credentialsId: 'pswd', variable: 'PSWD'),
                  string(credentialsId: 'username', variable: 'USERNAME'),
    
    ]){
        echo "My secret addr is '${ADDR}'"
        echo "My secret pswd is '${PSWD}'"
       script {
              def remote = [:]
              remote.name = 'master'
              remote.host = '${ADDR}'
              remote.user = '${USERNAME}'
              remote.password = '${PSWD}'
              remote.allowAnyHosts = true
    
             sshCommand remote:remote, command: "pwd"
             
             sshCommand remote:remote, command: "yq eval '.image.tag = ${IMAGE_VERSION}' -i /home/monta/Desktop/go-k8s-helm/go-k8s/values.yaml"
             sshCommand remote:remote, command: "cat /home/monta/Desktop/go-k8s-helm/go-k8s/values.yaml"
             sshCommand remote:remote, command: "cd /home/monta/Desktop/go-k8s-helm/go-k8s/ ; sh git.sh"
             
          
           }
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
