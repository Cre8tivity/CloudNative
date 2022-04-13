##LAB-10 

## Instruction 

## 1.Docker
$microk8s start 

$docker build -t mdang4/kubec-app .

- replace mdang4 with your docker username

$docker push mdang4/kubec-app:latest

- replace mdang4 with your docker username

Verified container is running

## 2.Kubernetes

$kubectl create -f deployment.yaml

$kubectl get pods

$kubectl get svc

$kubectl expose deployment my-kubec-app --type=NodePort --name=kubec-app-svc --target-port=8000

$export NODE_PORT=$(kubectl get services/kubec-app-svc -o go-template='{{(index .spec.ports 0).nodePort}}'); echo NODE_PORT=$NODE_PORT

$export NODE_IP=$(kubectl describe nodes | grep InternalIP | awk '{print $2}');echo

$minikube ip

-Ex: 192.168.64.7 <= This is my minikube IP address

$kubectl get svc

-EX: 8080:32526/TCP <= Look at the table

$minikube-ip:nodePort

-curl "http://$NODE_IP:$NODE_PORT/list?"

