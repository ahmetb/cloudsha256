# cloudsha256

Sample Go application to be deployed on Cloud Run to test upload speeds.

Public image at `gcr.io/ahmetb-public/cloudsha256:latest`.

Try out:

```
head -c 10000000 < /dev/urandom > myfile
curl -X POST --data-binary @myfile https://cloudsha256-dpyb4duzqq-uc.a.run.app/
```
