k3d cluster delete mycluster
k3d cluster create mycluster --api-port 6550 -p "8081:80@loadbalancer"

kubectl proxy
open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/workloads?namespace=default

localhost:8081/ping