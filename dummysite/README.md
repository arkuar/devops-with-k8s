# 5.01

Apply all the configurations in the manifests folder with `kubectl apply -f manifests`

Applying dummysite.yaml at the same time probably fails, so run `kubectl apply -f manifest/dummysite.yaml` and you should see a new deployment being created by the controller

The created deployment has port `3000` open so if you wish to visit it with your browser, port forward with `kubectl port-forward <dummysite-dep-name> 3000:3000`
