# Instructions

```
oc apply -f ims.yaml
oc apply -f bc.yaml

oc describe bc/ghc

oc apply -f - <<EOF
kind: Secret
apiVersion: v1
metadata:
  name: github-webhook-secret
  namespace: ghc
  creationTimestamp: null
data:
  WebHookSecretKey: <>
EOF


```