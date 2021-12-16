# 5.03
```shell
λ ~/Github/devops-with-k8s/project/ master kubectl apply -k github.com/fluxcd/flagger/kustomize/linkerd
customresourcedefinition.apiextensions.k8s.io/alertproviders.flagger.app created
customresourcedefinition.apiextensions.k8s.io/canaries.flagger.app created
customresourcedefinition.apiextensions.k8s.io/metrictemplates.flagger.app created
serviceaccount/flagger created
clusterrole.rbac.authorization.k8s.io/flagger created
clusterrolebinding.rbac.authorization.k8s.io/flagger created
deployment.apps/flagger created
λ ~/Github/devops-with-k8s/project/ master kubectl -n linkerd rollout status deploy/flagger
deployment "flagger" successfully rolled out
λ ~/Github/devops-with-k8s/project/ master kubectl create ns test && \
  kubectl apply -f https://run.linkerd.io/flagger.yml
namespace/test created
deployment.apps/load created
configmap/frontend created
deployment.apps/frontend created
service/frontend created
deployment.apps/podinfo created
service/podinfo created
λ ~/Github/devops-with-k8s/project/ master kubectl -n test rollout status deploy podinfo
deployment "podinfo" successfully rolled out
λ ~/Github/devops-with-k8s/project/ master kubectl -n test port-forward svc/frontend 8080
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
Handling connection for 8080
^C%

λ ~/Github/devops-with-k8s/project/ master cat <<EOF | kubectl apply -f -
apiVersion: flagger.app/v1beta1
kind: Canary
metadata:
  name: podinfo
  namespace: test
spec:
  targetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: podinfo
  service:
    port: 9898
  analysis:
    interval: 10s
    threshold: 5
    stepWeight: 10
    maxWeight: 100
    metrics:
    - name: request-success-rate
      thresholdRange:
        min: 99
      interval: 1m
    - name: request-duration
      thresholdRange:
        max: 500
      interval: 1m
EOF
canary.flagger.app/podinfo created
λ ~/Github/devops-with-k8s/project/ master kubectl -n test get ev --watch
LAST SEEN   TYPE      REASON                  OBJECT                                MESSAGE
66s         Normal    ScalingReplicaSet       deployment/load                       Scaled up replica set load-7f97579865 to 1
66s         Normal    ScalingReplicaSet       deployment/frontend                   Scaled up replica set frontend-6957977dc7 to
1
66s         Normal    Injected                deployment/load                       Linkerd sidecar proxy injected
66s         Normal    Injected                deployment/frontend                   Linkerd sidecar proxy injected
66s         Normal    SuccessfulCreate        replicaset/load-7f97579865            Created pod: load-7f97579865-fq7f8
66s         Normal    SuccessfulCreate        replicaset/frontend-6957977dc7        Created pod: frontend-6957977dc7-bd8tm
66s         Normal    Scheduled               pod/frontend-6957977dc7-bd8tm         Successfully assigned test/frontend-6957977dc
7-bd8tm to k3d-k3s-default-agent-1
66s         Normal    Scheduled               pod/load-7f97579865-fq7f8             Successfully assigned test/load-7f97579865-fq
7f8 to k3d-k3s-default-agent-0
66s         Normal    ScalingReplicaSet       deployment/podinfo                    Scaled up replica set podinfo-7bfd46f477 to 1
66s         Normal    Injected                deployment/podinfo                    Linkerd sidecar proxy injected
66s         Normal    Scheduled               pod/podinfo-7bfd46f477-p9vjc          Successfully assigned test/podinfo-7bfd46f477
-p9vjc to k3d-k3s-default-agent-0
66s         Normal    SuccessfulCreate        replicaset/podinfo-7bfd46f477         Created pod: podinfo-7bfd46f477-p9vjc
66s         Normal    Pulled                  pod/frontend-6957977dc7-bd8tm         Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
66s         Normal    Pulled                  pod/load-7f97579865-fq7f8             Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
66s         Normal    Created                 pod/frontend-6957977dc7-bd8tm         Created container linkerd-init
66s         Normal    Pulled                  pod/podinfo-7bfd46f477-p9vjc          Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
65s         Normal    Created                 pod/load-7f97579865-fq7f8             Created container linkerd-init
65s         Normal    Created                 pod/podinfo-7bfd46f477-p9vjc          Created container linkerd-init
65s         Normal    Started                 pod/frontend-6957977dc7-bd8tm         Started container linkerd-init
65s         Normal    Started                 pod/load-7f97579865-fq7f8             Started container linkerd-init
65s         Normal    Started                 pod/podinfo-7bfd46f477-p9vjc          Started container linkerd-init
65s         Normal    Pulled                  pod/frontend-6957977dc7-bd8tm         Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
65s         Normal    Created                 pod/frontend-6957977dc7-bd8tm         Created container linkerd-proxy
65s         Normal    Started                 pod/frontend-6957977dc7-bd8tm         Started container linkerd-proxy
65s         Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:16:53 +0000 UTC: e4337dbbdbc68f69112985681170acaa
65s         Normal    Pulled                  pod/load-7f97579865-fq7f8             Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
65s         Normal    Pulled                  pod/podinfo-7bfd46f477-p9vjc          Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
65s         Normal    Created                 pod/podinfo-7bfd46f477-p9vjc          Created container linkerd-proxy
65s         Normal    Created                 pod/load-7f97579865-fq7f8             Created container linkerd-proxy
65s         Normal    Started                 pod/load-7f97579865-fq7f8             Started container linkerd-proxy
65s         Normal    Started                 pod/podinfo-7bfd46f477-p9vjc          Started container linkerd-proxy
64s         Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:16:54 +0000 UTC: b076e81fc7fcbcdfa56d3e6b751adc20
64s         Normal    Pulling                 pod/podinfo-7bfd46f477-p9vjc          Pulling image "quay.io/stefanprodan/podinfo:1
.7.0"
64s         Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:16:54 +0000 UTC: e57246f3669ce23faaed9a9a024235d9
64s         Normal    Pulling                 pod/frontend-6957977dc7-bd8tm         Pulling image "nginx:alpine"
63s         Normal    Pulling                 pod/load-7f97579865-fq7f8             Pulling image "buoyantio/slow_cooker:1.2.0"
62s         Normal    Pulled                  pod/podinfo-7bfd46f477-p9vjc          Successfully pulled image "quay.io/stefanprod
an/podinfo:1.7.0" in 2.413720837s
62s         Normal    Created                 pod/podinfo-7bfd46f477-p9vjc          Created container podinfod
62s         Normal    Started                 pod/podinfo-7bfd46f477-p9vjc          Started container podinfod
60s         Normal    Pulled                  pod/load-7f97579865-fq7f8             Successfully pulled image "buoyantio/slow_coo
ker:1.2.0" in 3.076456638s
60s         Normal    Created                 pod/load-7f97579865-fq7f8             Created container slow-cooker
60s         Normal    Started                 pod/load-7f97579865-fq7f8             Started container slow-cooker
60s         Normal    Pulled                  pod/frontend-6957977dc7-bd8tm         Successfully pulled image "nginx:alpine" in 3
.99650895s
60s         Normal    Created                 pod/frontend-6957977dc7-bd8tm         Created container nginx
60s         Normal    Started                 pod/frontend-6957977dc7-bd8tm         Started container nginx
2s          Normal    Synced                  canary/podinfo                        all the metrics providers are available!
1s          Warning   Synced                  canary/podinfo                        podinfo-primary.test not ready: waiting for r
ollout to finish: observed deployment generation less than desired generation
1s          Normal    ScalingReplicaSet       deployment/podinfo-primary            Scaled up replica set podinfo-primary-f75dff6
9 to 1
1s          Normal    Injected                deployment/podinfo-primary            Linkerd sidecar proxy injected
1s          Normal    SuccessfulCreate        replicaset/podinfo-primary-f75dff69   Created pod: podinfo-primary-f75dff69-h8r65
1s          Normal    Scheduled               pod/podinfo-primary-f75dff69-h8r65    Successfully assigned test/podinfo-primary-f7
5dff69-h8r65 to k3d-k3s-default-agent-1
1s          Normal    Pulled                  pod/podinfo-primary-f75dff69-h8r65    Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
1s          Normal    Created                 pod/podinfo-primary-f75dff69-h8r65    Created container linkerd-init
1s          Normal    Started                 pod/podinfo-primary-f75dff69-h8r65    Started container linkerd-init
0s          Normal    Pulled                  pod/podinfo-primary-f75dff69-h8r65    Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
0s          Normal    Created                 pod/podinfo-primary-f75dff69-h8r65    Created container linkerd-proxy
0s          Normal    Started                 pod/podinfo-primary-f75dff69-h8r65    Started container linkerd-proxy
0s          Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:17:58 +0000 UTC: 771b5d0d80561c06e2f341cf9c723b0a
0s          Normal    Pulling                 pod/podinfo-primary-f75dff69-h8r65    Pulling image "quay.io/stefanprodan/podinfo:1
.7.0"
0s          Normal    Pulled                  pod/podinfo-primary-f75dff69-h8r65    Successfully pulled image "quay.io/stefanprod
an/podinfo:1.7.0" in 2.72230545s
0s          Normal    Created                 pod/podinfo-primary-f75dff69-h8r65    Created container podinfod
0s          Normal    Started                 pod/podinfo-primary-f75dff69-h8r65    Started container podinfod
0s          Normal    Synced                  canary/podinfo                        all the metrics providers are available!
0s          Normal    ScalingReplicaSet       deployment/podinfo                    Scaled down replica set podinfo-7bfd46f477 to
 0
0s          Normal    Killing                 pod/podinfo-7bfd46f477-p9vjc          Stopping container linkerd-proxy
1s          Normal    Killing                 pod/podinfo-7bfd46f477-p9vjc          Stopping container podinfod
1s          Normal    SuccessfulDelete        replicaset/podinfo-7bfd46f477         Deleted pod: podinfo-7bfd46f477-p9vjc
0s          Normal    Synced                  canary/podinfo                        Initialization done! podinfo.test
^C%

λ ~/Github/devops-with-k8s/project/ master kubectl -n test get svc
NAME              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
frontend          ClusterIP   10.43.0.158     <none>        8080/TCP   2m21s
podinfo-canary    ClusterIP   10.43.210.196   <none>        9898/TCP   77s
podinfo-primary   ClusterIP   10.43.217.125   <none>        9898/TCP   77s
podinfo           ClusterIP   10.43.14.24     <none>        9898/TCP   2m21s
λ ~/Github/devops-with-k8s/project/ master kubectl -n test set image deployment/podinfo \
  podinfod=quay.io/stefanprodan/podinfo:1.7.1
deployment.apps/podinfo image updated
λ ~/Github/devops-with-k8s/project/ master kubectl -n test get ev --watch
LAST SEEN   TYPE      REASON                  OBJECT                                MESSAGE
3m40s       Normal    ScalingReplicaSet       deployment/load                       Scaled up replica set load-7f97579865 to 1
3m40s       Normal    ScalingReplicaSet       deployment/frontend                   Scaled up replica set frontend-6957977dc7 to
1
3m40s       Normal    Injected                deployment/load                       Linkerd sidecar proxy injected
3m40s       Normal    Injected                deployment/frontend                   Linkerd sidecar proxy injected
3m40s       Normal    SuccessfulCreate        replicaset/load-7f97579865            Created pod: load-7f97579865-fq7f8
3m40s       Normal    SuccessfulCreate        replicaset/frontend-6957977dc7        Created pod: frontend-6957977dc7-bd8tm
3m40s       Normal    Scheduled               pod/frontend-6957977dc7-bd8tm         Successfully assigned test/frontend-6957977dc
7-bd8tm to k3d-k3s-default-agent-1
3m40s       Normal    Scheduled               pod/load-7f97579865-fq7f8             Successfully assigned test/load-7f97579865-fq
7f8 to k3d-k3s-default-agent-0
3m40s       Normal    ScalingReplicaSet       deployment/podinfo                    Scaled up replica set podinfo-7bfd46f477 to 1
3m40s       Normal    Scheduled               pod/podinfo-7bfd46f477-p9vjc          Successfully assigned test/podinfo-7bfd46f477
-p9vjc to k3d-k3s-default-agent-0
3m40s       Normal    SuccessfulCreate        replicaset/podinfo-7bfd46f477         Created pod: podinfo-7bfd46f477-p9vjc
3m40s       Normal    Pulled                  pod/frontend-6957977dc7-bd8tm         Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
3m40s       Normal    Pulled                  pod/load-7f97579865-fq7f8             Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
3m40s       Normal    Created                 pod/frontend-6957977dc7-bd8tm         Created container linkerd-init
3m40s       Normal    Pulled                  pod/podinfo-7bfd46f477-p9vjc          Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
3m39s       Normal    Created                 pod/load-7f97579865-fq7f8             Created container linkerd-init
3m39s       Normal    Created                 pod/podinfo-7bfd46f477-p9vjc          Created container linkerd-init
3m39s       Normal    Started                 pod/frontend-6957977dc7-bd8tm         Started container linkerd-init
3m39s       Normal    Started                 pod/load-7f97579865-fq7f8             Started container linkerd-init
3m39s       Normal    Started                 pod/podinfo-7bfd46f477-p9vjc          Started container linkerd-init
3m39s       Normal    Pulled                  pod/frontend-6957977dc7-bd8tm         Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
3m39s       Normal    Created                 pod/frontend-6957977dc7-bd8tm         Created container linkerd-proxy
3m39s       Normal    Started                 pod/frontend-6957977dc7-bd8tm         Started container linkerd-proxy
3m39s       Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:16:53 +0000 UTC: e4337dbbdbc68f69112985681170acaa
3m39s       Normal    Pulled                  pod/load-7f97579865-fq7f8             Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
3m39s       Normal    Pulled                  pod/podinfo-7bfd46f477-p9vjc          Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
3m39s       Normal    Created                 pod/podinfo-7bfd46f477-p9vjc          Created container linkerd-proxy
3m39s       Normal    Created                 pod/load-7f97579865-fq7f8             Created container linkerd-proxy
3m39s       Normal    Started                 pod/load-7f97579865-fq7f8             Started container linkerd-proxy
3m39s       Normal    Started                 pod/podinfo-7bfd46f477-p9vjc          Started container linkerd-proxy
3m38s       Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:16:54 +0000 UTC: b076e81fc7fcbcdfa56d3e6b751adc20
3m38s       Normal    Pulling                 pod/podinfo-7bfd46f477-p9vjc          Pulling image "quay.io/stefanprodan/podinfo:1
.7.0"
3m38s       Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:16:54 +0000 UTC: e57246f3669ce23faaed9a9a024235d9
3m38s       Normal    Pulling                 pod/frontend-6957977dc7-bd8tm         Pulling image "nginx:alpine"
3m37s       Normal    Pulling                 pod/load-7f97579865-fq7f8             Pulling image "buoyantio/slow_cooker:1.2.0"
3m36s       Normal    Pulled                  pod/podinfo-7bfd46f477-p9vjc          Successfully pulled image "quay.io/stefanprod
an/podinfo:1.7.0" in 2.413720837s
3m36s       Normal    Created                 pod/podinfo-7bfd46f477-p9vjc          Created container podinfod
3m36s       Normal    Started                 pod/podinfo-7bfd46f477-p9vjc          Started container podinfod
3m34s       Normal    Pulled                  pod/load-7f97579865-fq7f8             Successfully pulled image "buoyantio/slow_coo
ker:1.2.0" in 3.076456638s
3m34s       Normal    Created                 pod/load-7f97579865-fq7f8             Created container slow-cooker
3m34s       Normal    Started                 pod/load-7f97579865-fq7f8             Started container slow-cooker
3m34s       Normal    Pulled                  pod/frontend-6957977dc7-bd8tm         Successfully pulled image "nginx:alpine" in 3
.99650895s
3m34s       Normal    Created                 pod/frontend-6957977dc7-bd8tm         Created container nginx
3m34s       Normal    Started                 pod/frontend-6957977dc7-bd8tm         Started container nginx
2m35s       Warning   Synced                  canary/podinfo                        podinfo-primary.test not ready: waiting for r
ollout to finish: observed deployment generation less than desired generation
2m35s       Normal    ScalingReplicaSet       deployment/podinfo-primary            Scaled up replica set podinfo-primary-f75dff6
9 to 1
2m35s       Normal    Injected                deployment/podinfo-primary            Linkerd sidecar proxy injected
2m35s       Normal    SuccessfulCreate        replicaset/podinfo-primary-f75dff69   Created pod: podinfo-primary-f75dff69-h8r65
2m35s       Normal    Scheduled               pod/podinfo-primary-f75dff69-h8r65    Successfully assigned test/podinfo-primary-f7
5dff69-h8r65 to k3d-k3s-default-agent-1
2m35s       Normal    Pulled                  pod/podinfo-primary-f75dff69-h8r65    Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
2m35s       Normal    Created                 pod/podinfo-primary-f75dff69-h8r65    Created container linkerd-init
2m35s       Normal    Started                 pod/podinfo-primary-f75dff69-h8r65    Started container linkerd-init
2m34s       Normal    Pulled                  pod/podinfo-primary-f75dff69-h8r65    Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
2m34s       Normal    Created                 pod/podinfo-primary-f75dff69-h8r65    Created container linkerd-proxy
2m34s       Normal    Started                 pod/podinfo-primary-f75dff69-h8r65    Started container linkerd-proxy
2m34s       Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:17:58 +0000 UTC: 771b5d0d80561c06e2f341cf9c723b0a
2m33s       Normal    Pulling                 pod/podinfo-primary-f75dff69-h8r65    Pulling image "quay.io/stefanprodan/podinfo:1
.7.0"
2m30s       Normal    Pulled                  pod/podinfo-primary-f75dff69-h8r65    Successfully pulled image "quay.io/stefanprod
an/podinfo:1.7.0" in 2.72230545s
2m30s       Normal    Created                 pod/podinfo-primary-f75dff69-h8r65    Created container podinfod
2m30s       Normal    Started                 pod/podinfo-primary-f75dff69-h8r65    Started container podinfod
2m26s       Normal    Synced                  canary/podinfo                        all the metrics providers are available!
2m26s       Normal    ScalingReplicaSet       deployment/podinfo                    Scaled down replica set podinfo-7bfd46f477 to
 0
2m26s       Normal    Killing                 pod/podinfo-7bfd46f477-p9vjc          Stopping container linkerd-proxy
2m26s       Normal    Killing                 pod/podinfo-7bfd46f477-p9vjc          Stopping container podinfod
2m26s       Normal    SuccessfulDelete        replicaset/podinfo-7bfd46f477         Deleted pod: podinfo-7bfd46f477-p9vjc
2m25s       Normal    Synced                  canary/podinfo                        Initialization done! podinfo.test
6s          Normal    Synced                  canary/podinfo                        New revision detected! Scaling up podinfo.tes
t
6s          Normal    ScalingReplicaSet       deployment/podinfo                    Scaled up replica set podinfo-69c49997fd to 1
5s          Normal    Injected                deployment/podinfo                    Linkerd sidecar proxy injected
5s          Normal    SuccessfulCreate        replicaset/podinfo-69c49997fd         Created pod: podinfo-69c49997fd-vbtwv
5s          Normal    Scheduled               pod/podinfo-69c49997fd-vbtwv          Successfully assigned test/podinfo-69c49997fd
-vbtwv to k3d-k3s-default-agent-0
5s          Normal    Pulled                  pod/podinfo-69c49997fd-vbtwv          Container image "cr.l5d.io/linkerd/proxy-init
:v1.4.0" already present on machine
5s          Normal    Created                 pod/podinfo-69c49997fd-vbtwv          Created container linkerd-init
5s          Normal    Started                 pod/podinfo-69c49997fd-vbtwv          Started container linkerd-init
4s          Normal    Pulled                  pod/podinfo-69c49997fd-vbtwv          Container image "cr.l5d.io/linkerd/proxy:stab
le-2.11.1" already present on machine
4s          Normal    Created                 pod/podinfo-69c49997fd-vbtwv          Created container linkerd-proxy
4s          Normal    Started                 pod/podinfo-69c49997fd-vbtwv          Started container linkerd-proxy
4s          Normal    IssuedLeafCertificate   serviceaccount/default                issued certificate for default.test.serviceac
count.identity.linkerd.cluster.local until 2021-12-17 12:20:28 +0000 UTC: 5651924450dd7db5145c612796f57347
3s          Normal    Pulling                 pod/podinfo-69c49997fd-vbtwv          Pulling image "quay.io/stefanprodan/podinfo:1
.7.1"
0s          Normal    Pulled                  pod/podinfo-69c49997fd-vbtwv          Successfully pulled image "quay.io/stefanprod
an/podinfo:1.7.1" in 2.735889751s
0s          Normal    Created                 pod/podinfo-69c49997fd-vbtwv          Created container podinfod
0s          Normal    Started                 pod/podinfo-69c49997fd-vbtwv          Started container podinfod
0s          Normal    Synced                  canary/podinfo                        Starting canary analysis for podinfo.test
0s          Normal    Synced                  canary/podinfo                        Advance podinfo.test canary weight 10
0s          Normal    Synced                  canary/podinfo                        Advance podinfo.test canary weight 20
0s          Normal    Synced                  canary/podinfo                        Advance podinfo.test canary weight 30
0s          Normal    Synced                  canary/podinfo                        Advance podinfo.test canary weight 40
0s          Normal    Synced                  canary/podinfo                        Advance podinfo.test canary weight 50
0s          Normal    Synced                  canary/podinfo                        (combined from similar events): Advance podin
fo.test canary weight 60
0s          Normal    Synced                  canary/podinfo                        (combined from similar events): Advance podin
fo.test canary weight 70
0s          Normal    Synced                  canary/podinfo                        (combined from similar events): Advance podin
fo.test canary weight 80
0s          Normal    Synced                  canary/podinfo                        (combined from similar events): Advance podin
fo.test canary weight 90
0s          Normal    Synced                  canary/podinfo                        (combined from similar events): Advance podin
fo.test canary weight 100
0s          Normal    Synced                  canary/podinfo                        (combined from similar events): Copying podin
fo.test template spec to podinfo-primary.test
0s          Normal    ScalingReplicaSet       deployment/podinfo-primary            Scaled up replica set podinfo-primary-74c9548
5fd to 1
0s          Normal    Injected                deployment/podinfo-primary            Linkerd sidecar proxy injected
0s          Normal    SuccessfulCreate        replicaset/podinfo-primary-74c95485fd   Created pod: podinfo-primary-74c95485fd-vx4
8b
0s          Normal    Scheduled               pod/podinfo-primary-74c95485fd-vx48b    Successfully assigned test/podinfo-primary-
74c95485fd-vx48b to k3d-k3s-default-server-0
0s          Normal    Pulled                  pod/podinfo-primary-74c95485fd-vx48b    Container image "cr.l5d.io/linkerd/proxy-in
it:v1.4.0" already present on machine
0s          Normal    Created                 pod/podinfo-primary-74c95485fd-vx48b    Created container linkerd-init
0s          Normal    Started                 pod/podinfo-primary-74c95485fd-vx48b    Started container linkerd-init
0s          Normal    Pulled                  pod/podinfo-primary-74c95485fd-vx48b    Container image "cr.l5d.io/linkerd/proxy:st
able-2.11.1" already present on machine
0s          Normal    Created                 pod/podinfo-primary-74c95485fd-vx48b    Created container linkerd-proxy
0s          Normal    Started                 pod/podinfo-primary-74c95485fd-vx48b    Started container linkerd-proxy
0s          Normal    IssuedLeafCertificate   serviceaccount/default                  issued certificate for default.test.service
account.identity.linkerd.cluster.local until 2021-12-17 12:22:18 +0000 UTC: 86dca5e02ad7864c9faaf716da5eb5d3
0s          Normal    Pulling                 pod/podinfo-primary-74c95485fd-vx48b    Pulling image "quay.io/stefanprodan/podinfo
:1.7.1"
0s          Normal    Pulled                  pod/podinfo-primary-74c95485fd-vx48b    Successfully pulled image "quay.io/stefanpr
odan/podinfo:1.7.1" in 3.162521696s
0s          Normal    Created                 pod/podinfo-primary-74c95485fd-vx48b    Created container podinfod
0s          Normal    Started                 pod/podinfo-primary-74c95485fd-vx48b    Started container podinfod
0s          Normal    ScalingReplicaSet       deployment/podinfo-primary              Scaled down replica set podinfo-primary-f75
dff69 to 0
0s          Normal    SuccessfulDelete        replicaset/podinfo-primary-f75dff69     Deleted pod: podinfo-primary-f75dff69-h8r65
0s          Normal    Killing                 pod/podinfo-primary-f75dff69-h8r65      Stopping container linkerd-proxy
0s          Normal    Killing                 pod/podinfo-primary-f75dff69-h8r65      Stopping container podinfod
0s          Normal    Synced                  canary/podinfo                          (combined from similar events): Routing all
 traffic to primary
0s          Normal    ScalingReplicaSet       deployment/podinfo                      Scaled down replica set podinfo-69c49997fd
to 0
0s          Normal    Synced                  canary/podinfo                          (combined from similar events): Promotion c
ompleted! Scaling down podinfo.test
0s          Normal    SuccessfulDelete        replicaset/podinfo-69c49997fd           Deleted pod: podinfo-69c49997fd-vbtwv
0s          Normal    Killing                 pod/podinfo-69c49997fd-vbtwv            Stopping container linkerd-proxy
1s          Normal    Killing                 pod/podinfo-69c49997fd-vbtwv            Stopping container podinfod
0s          Warning   Unhealthy               pod/podinfo-69c49997fd-vbtwv            Readiness probe failed: Get "http://10.42.2
.55:4191/ready": dial tcp 10.42.2.55:4191: connect: connection refused
0s          Warning   Unhealthy               pod/podinfo-69c49997fd-vbtwv            Liveness probe failed: Get "http://10.42.2.
55:4191/live": dial tcp 10.42.2.55:4191: connect: connection refused
^C%

λ ~/Github/devops-with-k8s/project/ master watch kubectl -n test get canary
λ ~/Github/devops-with-k8s/project/ master kubectl -n test get trafficsplit podinfo -o yaml
apiVersion: split.smi-spec.io/v1alpha2
kind: TrafficSplit
metadata:
  creationTimestamp: "2021-12-16T12:17:46Z"
  generation: 12
  name: podinfo
  namespace: test
  ownerReferences:
  - apiVersion: flagger.app/v1beta1
    blockOwnerDeletion: true
    controller: true
    kind: Canary
    name: podinfo
    uid: d2a33bc7-8ecb-41d9-aaa0-6f672b129650
  resourceVersion: "84307"
  uid: 1ff7f5d5-281c-4fdb-b533-973e7aafef39
spec:
  backends:
  - service: podinfo-canary
    weight: "0"
  - service: podinfo-primary
    weight: "100"
  service: podinfo
λ ~/Github/devops-with-k8s/project/ master watch linkerd viz -n test stat deploy --from deploy/load
λ ~/Github/devops-with-k8s/project/ master kubectl -n test port-forward svc/frontend 8080
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
Handling connection for 8080
^C%
λ ~/Github/devops-with-k8s/project/ master kubectl delete -k github.com/fluxcd/flagger/kustomize/linkerd && \
  kubectl delete ns test
customresourcedefinition.apiextensions.k8s.io "alertproviders.flagger.app" deleted
customresourcedefinition.apiextensions.k8s.io "canaries.flagger.app" deleted
customresourcedefinition.apiextensions.k8s.io "metrictemplates.flagger.app" deleted
serviceaccount "flagger" deleted
clusterrole.rbac.authorization.k8s.io "flagger" deleted
clusterrolebinding.rbac.authorization.k8s.io "flagger" deleted
deployment.apps "flagger" deleted
namespace "test" deleted
λ ~/Github/devops-with-k8s/project/ master
```