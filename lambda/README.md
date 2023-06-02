# SSM2ENV Lambda Layer

We can create an ssm2env lambda layer which will run the ssm2env cli to export all environment variables into the lambda prior to the actual lambda being run.  This will use the `SSM2ENV_PATH` environment variable to determine where to pull variables from.  All `SSM2ENV_` environment variables can be set to configure the cli *except* export which is necessary for the layer to work.

**You must set the `AWS_LAMBDA_EXEC_WRAPPER` environment variable in your lambda to `/opt/ssm2env-wrapper.sh`**

To make use of this layer, simply create a lambda layer for the desired architecture or both and upload the provided release zips as the source.  In the future, it would be nice to see this as a public layer for consumption, but for now, it will be dependent on the consumer of this layer to create the layer themselves.

```bash
gh release download -R shopsmart/ssm2env -p 'ssm2env-lambda-layer_*.zip' -D lambda

aws lambda publish-layer-version --layer-name "ssm2env-amd64" \
  -compatible-architectures x86_64 \
  --zip-file lambda/ssm2env-lambda-layer_*_linux_amd64.zip

aws lambda publish-layer-version --layer-name "ssm2env-arm64" \
  -compatible-architectures arm64 \
  --zip-file lambda/ssm2env-lambda-layer_*_linux_arm64.zip
```
