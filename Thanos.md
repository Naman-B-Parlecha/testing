1. Fix MinIO mc Command Errors
The quickstart.sh script uses outdated mc commands (mc config host add and mc config host rm). Newer versions of mc use mc alias set and mc alias remove. Update the script to fix this.

Edit `scripts/quickstart.sh`:
Find the MinIO setup section (around lines 45–50):

``` bash
${MC_EXECUTABLE} alias set tmp http://${MINIO_ENDPOINT} ${MINIO_ACCESS_KEY} ${MINIO_SECRET_KEY}
${MC_EXECUTABLE} mb tmp/${MINIO_BUCKET}
${MC_EXECUTABLE} alias remove tmp
```

2. Install MinIO
``` bash
go install github.com/minio/mc@latest
export MC_EXECUTABLE=/home/naman/go/bin/mc
```

3. make bucket.yml
```bash
cat ~/Desktop/opensource/thanos/data/bucket.yml
```
add to it
   
   ```bash
   type: S3
    config:
      bucket: thanos
      endpoint: 127.0.0.1:9000
      insecure: true
      signature_version2: true
      access_key: THANOS
      secret_key: ITSTHANOSTIME
   ```
4. ``` bash
   export PROMETHEUS_EXECUTABLE=/home/naman/go/bin/prometheus-v0.54.1
   export MINIO_ENABLED=1
   export MINIO_EXECUTABLE=/home/naman/go/bin/minio-v0.0.0-20241014163537-3da7c9cce3de
   export MC_EXECUTABLE=/home/naman/go/bin/mc
   cd ~/Desktop/opensource/thanos
   make quickstart
   ```
