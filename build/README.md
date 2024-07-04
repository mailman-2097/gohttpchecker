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

Details of github [webhooks](https://docs.github.com/en/webhooks/using-webhooks/creating-webhooks)

Information on creating triggers [openshift hook](https://docs.openshift.com/container-platform/4.13/cicd/builds/triggering-builds-build-hooks.html)