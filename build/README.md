# Instructions

```
oc apply -f build 

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