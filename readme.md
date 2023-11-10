k3d cluster delete mycluster
k3d cluster create mycluster --api-port 6550 -p "8081:80@loadbalancer"

docker build ./service -t docker.io/library/service:v1
k3d image import docker.io/library/service:v1 -c mycluster

kubectl apply -f service/k8_resources/deployment.yml
kubectl apply -f service/k8_resources/service.yml
kubectl apply -f service/k8_resources/ingress.yml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl apply -f dashboard/sa.yml
kubectl apply -f dashboard/crb.yml
kubectl apply -f dashboard/secret.yml
kubectl -n kubernetes-dashboard create token admin-user
kubectl proxy
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/workloads?namespace=default

localhost:8081/ping