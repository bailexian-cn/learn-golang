#!/bin/bash
IPADDR=`ip route get 8.8.8.8 | grep "src" | awk '{print $7}'`
HOSTNAME=`hostname`

function generateK8sConfig() {
    # 1 生成配置
    mkdir -p /etc/k8s/config
    cd /etc/k8s/config/
    # 1.1 生成ca证书
    cat << EOF > ca-csr.json
{
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "L": "Wuhan",
            "ST": "Hubei",
            "O": "k8s",
            "OU": "System"
        }
    ]
}
EOF
    cfssl gencert --initca ca-csr.json | cfssljson --bare ca
    # ca-key.pem : CA的私有key
    # ca.pem : CA证书
    # ca.csr : CA的证书请求文件
    cat << EOF > ca-config.json
{
    "signing": {
        "default": {
            "expiry": "175200h"
        },
        "profiles": {
            "kubernetes": {
                "expiry": "175200h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth",
                    "client auth"
                ]
            },
            "etcd": {
                "expiry": "175200h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth",
                            "client auth"
                ]
            }
        }
    }
}
EOF
# 1.2 生成etcd证书
    cat << EOF > etcd-csr.json
{
  "CN": "etcd-server",
  "hosts": [
    "0.0.0.0",
    "127.0.0.1",
    "${IPADDR}"
  ],
  "key": {
    "algo": "rsa",
    "size": 4096
  },
  "names": [
    {
      "C": "CN",
      "L": "Wuhan",
      "O": "k8s",
      "OU": "System",
      "ST": "Hubei"
    }
  ]
}
EOF
    cat << EOF > etcd-client-csr.json
{
  "CN": "etcd-client",
  "hosts": [
    ""
  ],
  "key": {
    "algo": "rsa",
    "size": 4096
  },
  "names": [
    {
      "C": "CN",
      "L": "Wuhan",
      "ST": "Hubei",
      "O": "k8s",
      "OU": "System"
    }
  ]
}
EOF
    cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \\
    -config=ca-config.json \\
    -profile=etcd etcd-csr.json | cfssljson -bare etcd
    cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \\
    -config=ca-config.json \\
    -profile=etcd etcd-client-csr.json | cfssljson -bare etcd-client
# 1.3 生成apiserver证书
    cat << EOF > kube-apiserver-csr.json
{
  "CN": "etcd-server",
  "hosts": [
    "0.0.0.0",
    "127.0.0.1",
    "${IPADDR}"
  ],
  "key": {
    "algo": "rsa",
    "size": 4096
  },
  "names": [
    {
      "C": "CN",
      "L": "Wuhan",
      "O": "k8s",
      "OU": "System",
      "ST": "Hubei"
    }
  ]
}
EOF
    cat << EOF > kube-apiserver-client-csr.json
{
  "CN": "kube-apiserver-client",
  "hosts": [
    ""
  ],
  "key": {
    "algo": "rsa",
    "size": 4096
  },
  "names": [
    {
      "C": "CN",
      "L": "Wuhan",
      "ST": "Hubei",
      "O": "k8s",
      "OU": "System"
    }
  ]
}
EOF
    cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \\
    -config=ca-config.json \\
    -profile=kubernetes kube-apiserver-csr.json | cfssljson -bare kube-apiserver
    cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \\
    -config=ca-config.json \\
    -profile=kubernetes kube-apiserver-client-csr.json | cfssljson -bare kube-apiserver-client
    # 1.4 生成kubelet证书
    cat << EOF > kubelet-csr.json
{
  "CN": "kubelet",
  "hosts": [
    "0.0.0.0",
    "127.0.0.1",
    "${IPADDR}"
  ],
  "key": {
    "algo": "rsa",
    "size": 4096
  },
  "names": [
    {
      "C": "CN",
      "L": "Wuhan",
      "O": "k8s",
      "OU": "System",
      "ST": "Hubei"
    }
  ]
}
EOF
    cat << EOF > kubelet-client-csr.json
{
  "CN": "kubelet-client",
  "hosts": [
    ""
  ],
  "key": {
    "algo": "rsa",
    "size": 4096
  },
  "names": [
    {
      "C": "CN",
      "L": "Wuhan",
      "ST": "Hubei",
      "O": "k8s",
      "OU": "System"
    }
  ]
}
EOF
    cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \\
    -config=ca-config.json \\
    -profile=kubernetes kubelet-csr.json | cfssljson -bare kubelet
    cfssl gencert -ca=ca.pem -ca-key=ca-key.pem \\
    -config=ca-config.json \\
    -profile=kubernetes kubelet-client-csr.json | cfssljson -bare kubelet-client
# 1.5 生成sa证书
    cat << EOF > sa-csr.json
{
  "CN":"sa",
  "hosts": [
    ""
  ],
  "key": {
    "algo": "rsa",
    "size": 4096
  },
  "names": [
    {
      "C": "CN",
      "L": "Wuhan",
      "ST": "Hubei",
      "O": "k8s",
      "OU": "System"
    }
  ]
}
EOF

    cfssl gencert -initca sa-csr.json  | cfssljson -bare sa -
    openssl x509 -in sa.pem -pubkey -noout > sa.pub
    # 1.6 生成kubeconfig
    kubectl config set-cluster k8s-demo \\
    --certificate-authority=/etc/k8s/config/ca.pem \\
    --embed-certs=true \\
    --server=https://${IPADDR}:6443 \\
    --kubeconfig=/etc/kubernetes/config/admin.kubeconfig

    kubectl config set-credentials system:admin \\
    --client-certificate=/etc/k8s/config/kube-apiserver-client.pem \\
    --embed-certs=true \\
    --client-key=/etc/k8s/config/kube-apiserver-client-key.pem \\
    --kubeconfig=/etc/kubernetes/config/admin.kubeconfig

    kubectl config set-context default \\
    --cluster=k8s-demo \\
    --user=system:admin \\
    --kubeconfig=/etc/kubernetes/config/admin.kubeconfig

    kubectl config use-context default \\
    --kubeconfig=/etc/kubernetes/config/admin.kubeconfig

    mkdir -p /root/.kube
    mv /etc/kubernetes/config/admin.kubeconfig /root/.kube/config -f
    export KUBECONFIG=/root/.kube/config
}


# 2 启动服务

function startEtcd() {
    # 2.1 启动etcd
    cat << EOF > /usr/lib/systemd/system/etcd.service
# /usr/lib/systemd/system/etcd.service
[Unit]
Description=Etcd
Wants=network-online.target
Before=kubelet.service
After=network-online.target

[Service]
Type=notify
EnvironmentFile=-/etc/sysconfig/crio
ExecStart=/usr/bin/etcd --name etcd \\
    --data-dir /var/lib/etcd \\
    --listen-client-urls https://0.0.0.0:2379 \\
    --advertise-client-urls https://0.0.0.0:2379 \\
    --trusted-ca-file /etc/k8s/config/ca.pem \\
    --cert-file /etc/k8s/config/etcd.pem \\
    --key-file /etc/k8s/config/etcd-key.pem \\
    --listen-peer-urls "" \\
    --client-cert-auth
ExecReload=/bin/kill -s HUP $MAINPID
TasksMax=infinity
LimitNOFILE=1048576
LimitNPROC=1048576
LimitCORE=infinity
OOMScoreAdjust=-999
TimeoutStartSec=0
Restart=on-abnormal
EOF
    systemctl daemon-reload
    systemctl enable etcd
    systemctl restart etcd
    # 测试etcd
    etcdctl \\
    --cert=/etc/k8s/config/etcd-client.pem \\
    --key=/etc/k8s/config/etcd-client-key.pem \\
    --cacert=/etc/k8s/config/ca.pem \\
    --endpoints=https://${IPADDR}:2379 \\
    member list
}

function startApiserver() {
    cat << EOF > /usr/lib/systemd/system/kube-apiserver.service
# /usr/lib/systemd/system/kube-apiserver.service
[Unit]
Description=kube-apiserver
Wants=network-online.target
Before=kubelet.service
After=network-online.target

[Service]
Type=notify
EnvironmentFile=-/etc/sysconfig/crio
ExecStart=kube-apiserver \\
          --advertise-address=${IPADDR} \\
          --bind-address=0.0.0.0 \\
          --service-cluster-ip-range=172.16.0.0/16 \\
          --service-node-port-range=30000-32767 \\
          --etcd-servers=https://${IPADDR}:2379 \\
          --etcd-cafile=/etc/k8s/config/ca.pem \\
          --etcd-certfile=/etc/k8s/config/etcd-client.pem \\
          --etcd-keyfile=/etc/k8s/config/etcd-client-key.pem \\
          --client-ca-file=/etc/k8s/config/ca.pem \\
          --tls-cert-file=/etc/k8s/config/kube-apiserver.pem \\
          --tls-private-key-file=/etc/k8s/config/kube-apiserver-key.pem \\
          --kubelet-certificate-authority=/etc/k8s/config/ca.pem \\
          --kubelet-client-certificate=/etc/k8s/config/kubelet-client.pem \\
          --kubelet-client-key=/etc/k8s/config/kubelet-client-key.pem \\
          --service-account-key-file=/etc/k8s/config/sa.pub \\
          --service-account-signing-key-file=/etc/k8s/config/sa-key.pem \\
          --service-account-issuer=api
ExecReload=/bin/kill -s HUP $MAINPID
TasksMax=infinity
LimitNOFILE=1048576
LimitNPROC=1048576
LimitCORE=infinity
OOMScoreAdjust=-999
TimeoutStartSec=0
Restart=on-abnormal
EOF
    systemctl daemon-reload
    systemctl enable kube-apiserver
    systemctl restart kube-apiserver
}

function startControllerManager() {
    cat << EOF > /usr/lib/systemd/system/kube-controller-manager.service
# /usr/lib/systemd/system/kube-controller-manager.service
[Unit]
Description=kube-controller-manager
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
ExecStart=kube-controller-manager \\
          --bind-address=0.0.0.0 \\
          --cluster-cidr=10.0.0.0/8 \\
          --master=https://${IPADDR}:6443 \\
          --service-cluster-ip-range=172.16.0.0/16 \\
          --kubeconfig=/root/.kube/config \\
          --service-account-private-key-file=/etc/k8s/config/sa-key.pem \\
          --cluster-signing-cert-file=/etc/k8s/config/ca.pem \\
          --cluster-signing-key-file=/etc/k8s/config/ca-key.pem
ExecReload=/bin/kill -s HUP $MAINPID
TasksMax=infinity
LimitNOFILE=1048576
LimitNPROC=1048576
LimitCORE=infinity
OOMScoreAdjust=-999
TimeoutStartSec=0
Restart=on-abnormal
EOF
    systemctl daemon-reload
    systemctl enable kube-controller-manager
    systemctl restart kube-controller-manager
}

function startScheduler() {
    cat << EOF > /usr/lib/systemd/system/kube-scheduler.service
# /usr/lib/systemd/system/kube-scheduler.service
[Unit]
Description=kube-scheduler
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
ExecStart=kube-scheduler \\
          --master=https://${IPADDR}:6443 \\
          --bind-address=0.0.0.0 \\
          --kubeconfig=/root/.kube/config \\
          --authorization-kubeconfig=/root/.kube/config
ExecReload=/bin/kill -s HUP $MAINPID
TasksMax=infinity
LimitNOFILE=1048576
LimitNPROC=1048576
LimitCORE=infinity
OOMScoreAdjust=-999
TimeoutStartSec=0
Restart=on-abnormal
EOF
    systemctl daemon-reload
    systemctl enable kube-scheduler.service
    systemctl restart kube-scheduler.service
}

function startCrio() {
    # 2.5 启动crio
    cat << EOF > /usr/lib/systemd/system/crio.service
# /usr/lib/systemd/system/crio.service
[Unit]
Description=Container Runtime Interface for OCI (CRI-O)
Documentation=https://github.com/cri-o/cri-o
Wants=network-online.target
Before=kubelet.service
After=network-online.target

[Service]
Type=notify
EnvironmentFile=-/etc/sysconfig/crio
Environment=GOTRACEBACK=crash
ExecStart=/usr/bin/crio \\
          --metrics-socket /run/crio/crio.sock\\
          $CRIO_CONFIG_OPTIONS \\
          $CRIO_RUNTIME_OPTIONS \\
          $CRIO_STORAGE_OPTIONS \\
          $CRIO_NETWORK_OPTIONS \\
          $CRIO_METRICS_OPTIONS
ExecReload=/bin/kill -s HUP $MAINPID
TasksMax=infinity
LimitNOFILE=1048576
LimitNPROC=1048576
LimitCORE=infinity
OOMScoreAdjust=-999
TimeoutStartSec=0
Restart=on-abnormal

[Install]
WantedBy=multi-user.target
Alias=cri-o.service

# /etc/systemd/system/crio.service.d/10-mco-default-madv.conf
[Service]
Environment="GODEBUG=x509ignoreCN=0,madvdontneed=1"

# /etc/systemd/system/crio.service.d/10-mco-profile-unix-socket.conf
[Service]
Environment="ENABLE_PROFILE_UNIX_SOCKET=true"

# /etc/systemd/system/crio.service.d/20-nodenet.conf
[Service]
Environment="CONTAINER_STREAM_ADDRESS=${IPADDR}"
EOF
    systemctl daemon-reload
    systemctl enable crio
    systemctl restart crio
}

function startKubelet() {
    # 2.6 启动kubelet
    cat << EOF > /etc/kubernetes/kubelet.conf
apiVersion: kubelet.config.k8s.io/v1beta1
authentication:
  anonymous:
    enabled: false
  webhook:
    cacheTTL: 2m0s
    enabled: true
  x509:
    clientCAFile: /etc/k8s/config/ca.pem
authorization:
  mode: Webhook
  webhook:
    cacheAuthorizedTTL: 5m0s
    cacheUnauthorizedTTL: 30s
cgroupDriver: systemd
clusterDNS:
- 169.254.20.10
clusterDomain: cluster.local
containerLogMaxFiles: 3
containerLogMaxSize: 100Mi
contentType: application/vnd.kubernetes.protobuf
cpuCFSQuota: true
cpuCFSQuotaPeriod: 100ms
cpuManagerPolicy: none
cpuManagerReconcilePeriod: 10s
enableControllerAttachDetach: true
enableDebuggingHandlers: true
enforceNodeAllocatable:
- pods
eventBurst: 10
eventRecordQPS: 5
evictionPressureTransitionPeriod: 0s
failSwapOn: true
fileCheckFrequency: 20s
hairpinMode: promiscuous-bridge
healthzBindAddress: 127.0.0.1
healthzPort: 10248
httpCheckFrequency: 20s
imageGCHighThresholdPercent: 85
imageGCLowThresholdPercent: 80
imageMinimumGCAge: 2m0s
iptablesDropBit: 15
iptablesMasqueradeBit: 14
kind: KubeletConfiguration
kubeAPIBurst: 100
kubeAPIQPS: 50
kubeReserved:
  cpu: 500m
  ephemeral-storage: 10Gi
  memory: 500Mi
  pid: "1000"
makeIPTablesUtilChains: true
maxOpenFiles: 1000000
maxPods: 250
nodeLeaseDurationSeconds: 40
nodeStatusReportFrequency: 10s
nodeStatusUpdateFrequency: 10s
oomScoreAdj: -999
podCIDR: 10.150.0.0/16
podPidsLimit: -1
port: 10250
registryBurst: 10
registryPullQPS: 5
rotateCertificates: true
runtimeRequestTimeout: 30s
serializeImagePulls: true
staticPodPath: /etc/kubernetes/manifests
streamingConnectionIdleTimeout: 4h0m0s
syncFrequency: 1m0s
systemReserved:
  cpu: 500m
  ephemeral-storage: 10Gi
  memory: 500Mi
  pid: "1000"
volumeStatsAggPeriod: 1m0s
EOF

    # 关闭swap
    sed -i 's/.*swap.*/#&/' /etc/fstab
    swapoff -a

    mkdir -p /etc/kubernetes/manifests

    cat << EOF > /usr/lib/systemd/system/kubelet.service
# /usr/lib/systemd/system/kubelet.service
[Unit]
Description=kubelet
Wants=network-online.target
After=network-online.target

[Service]
Type=notify
ExecStart=kubelet \\
        --bootstrap-kubeconfig="" \\
        --kubeconfig=/root/.kube/config  \\
        --config=/etc/kubernetes/kubelet.conf \\
        --node-ip=${IPADDR} \\
        --container-runtime=remote \\
        --container-runtime-endpoint=/run/crio/crio.sock \\
        --runtime-cgroups=/system.slice/crio.service \\
        --v=2
ExecReload=/bin/kill -s HUP $MAINPID
TasksMax=infinity
LimitNOFILE=1048576
LimitNPROC=1048576
LimitCORE=infinity
OOMScoreAdjust=-999
TimeoutStartSec=0
Restart=on-abnormal
EOF
    systemctl daemon-reload
    systemctl enable kubelet.service
    systemctl restart kubelet.service

    kubectl label nodes ${HOSTNAME} node-role.kubernetes.io/master=true --overwrite

    kubectl get nodes
}

generateK8sConfig
startEtcd
startApiserver
startControllerManager
startScheduler
startCrio
startKubelet
