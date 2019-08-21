
Velero gives you tools to back up and restore your Kubernetes cluster resources and persistent volumes. It also takes snapshots of your cluster’s Persistent Volumes using your cloud provider’s block storage snapshot features, and can then restore your cluster’s objects and Persistent Volumes to a previous state.

```
# git clone https://github.com/heptio/velero

# aws s3api create-bucket --bucket velero-bkp --region eu-central-1 --create-bucket-configuration LocationConstraint=eu-central-1

# aws iam create-user --user-name velero
```

Attach policies to give velero the necessary permissions

```
# BUCKET=velero-bkp
# cat > velero-policy.json <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeVolumes",
                "ec2:DescribeSnapshots",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:CreateSnapshot",
                "ec2:DeleteSnapshot"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:DeleteObject",
                "s3:PutObject",
                "s3:AbortMultipartUpload",
                "s3:ListMultipartUploadParts"
            ],
            "Resource": [
                "arn:aws:s3:::${BUCKET}/*"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:ListBucket"
            ],
            "Resource": [
                "arn:aws:s3:::${BUCKET}"
            ]
        }
    ]
}
EOF

# aws iam put-user-policy \
  --user-name velero \
  --policy-name velero \
  --policy-document file://velero-policy.json
```

Create an access key for the user
```
# aws iam create-access-key --user-name velero
```

Create a Velero-specific credentials file (credentials-velero) in your local directory
```
[default]
aws_access_key_id=<AWS_ACCESS_KEY_ID>
aws_secret_access_key=<AWS_SECRET_ACCESS_KEY>
```

Apply some basic prerequisites (e.g. CustomResourceDefinitions, namespaces, and RBAC): 
```
# kubectl apply -f examples/common/00-prereqs.yaml
```

Create Secret
```
# kubectl create secret generic cloud-credentials \
    --namespace velero \
    --from-file cloud=credentials-velero
```

In examples/aws/05-backupstoragelocation.yaml: Replace <YOUR_BUCKET> and <YOUR_REGION>
```
# kubectl apply -f examples/aws/05-backupstoragelocation.yaml
```

In examples/aws/06-volumesnapshotlocation.yaml: Replace <YOUR_REGION>
```
# kubectl apply -f examples/aws/06-volumesnapshotlocation.yaml
```

Deployment
```
# kubectl apply -f examples/aws/10-deployment.yaml
```
Check to see that Velero deployment have been successfully created:
```
# kubectl get deployments -l component=velero --namespace=velero
```
Let's install the Velero client:
```
# brew install velero
```

Deploy sample nginx for testing `nginx-example-with-pv.yaml`
```
---
apiVersion: v1
kind: Namespace
metadata:
  name: nginx-example
  labels:
    app: nginx

---
apiVersion: v1
kind: Service
metadata:
  name: nginx
  namespace: nginx-example
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  selector:
    app: nginx
  type: LoadBalancer

---
apiVersion: apps/v1beta2
kind: StatefulSet
metadata:
  name: web
  namespace: nginx-example
spec:
  serviceName: "nginx"
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.15-alpine
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi        
```

```
# kubectl create -f nginx-example-with-pv.yaml
# kubectl -n nginx-example exec web-0 -- sh -c 'echo $(hostname) > /usr/share/nginx/html/index.html'
# kubectl get svc -n nginx-example -l app=nginx -o wide
# curl <lb_ip>
# kubectl get pvc --namespace nginx-example
# nginx_pv_name=$(kubectl get pv -o jsonpath='{.items[?(@.spec.claimRef.name=="www-web-0")].metadata.name}')
# kubectl label pv $nginx_pv_name app=nginx
```

Let’s create a backup
```
# velero backup create nginx-backup --selector app=nginx --exclude-namespaces velero,default,kube-system,ingress-nginx,kube-public,monitoring --snapshots-volume
```

With velero backup get you will also get a list of all available backups:
```
# velero backup get
# velero backup describe nginx-backup --details
```

Now, let's simulate a disaster: 
```
# kubectl delete namespace nginx-example
```
Check that the nginx service/deployment/pv are gone. You might need to wait for a few minutes for the namespace to be fully cleaned up.

Restore Our Resources
```
# velero restore create --from-backup nginx-backup
# velero restore get
NAME                   BACKUP    STATUS      WARNINGS   ERRORS    CREATED                         SELECTOR
nginx-20190413211418   nginx     Completed   0          0         2019-04-13 21:14:16 +0530 IST   <none>
```

NOTE: The restore can take a few moments to finish. During this time, the STATUS column reads InProgress.

After a successful restore, the STATUS column is Completed, and WARNINGS and ERRORS are 0. All objects in the nginx-example namespace should be just as they were before you deleted them.

If there are errors or warnings, you can look at them in detail:
```
velero restore describe nginx-20190413211418
```

Delete backup
```
velero delete backup nginx-backup
```


Delete Velero
```
kubectl delete namespace/velero clusterrolebinding/velero
kubectl delete crds -l component=velero
```
