image:
	docker build -t kubein . 

cluster: deletecluster
	kind create cluster --config=kind-config.yaml

deletecluster:
	kind delete cluster

ingress: 
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

