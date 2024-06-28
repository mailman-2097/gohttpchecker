# Instructions

```
oc apply -f ims.yaml
oc apply -f bc.yaml

oc describe bc/ghc

https://api.d0.ocp.dev.education.nsw.gov.au:6443/apis/build.openshift.io/v1/namespaces/ghc/buildconfigs/ghc/webhooks/<secret>/github

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