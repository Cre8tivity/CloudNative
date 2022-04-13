#LAB-10 

## Instruction 


$minikube start 

$docker build -t mdang4/kubec-app .

$docker push mdang4/kubec-app:latest

$kubectl create -f deployment.yaml

$kubectl get pods

$kubectl get svc

$kubectl expose deployment my-kubec-app --type=NodePort --name=kubec-app-svc --target-port=8080

$export NODE_PORT=$(kubectl get services/kubec-app-svc -o go-template='{{(index .spec.ports 0).nodePort}}'); echo NODE_PORT=$NODE_PORT

$export NODE_IP=$(kubectl describe nodes | grep InternalIP | awk '{print $2}');echo

$minikube ip
192.168.64.7

$kubectl get svc
EX: 8080:32526/TCP

$minikube-ip:nodePort
curl "http://192.168.64.7:32526/list?"

