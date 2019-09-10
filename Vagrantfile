# -*- mode: ruby -*-
# vi: set ft=ruby :

# After loading this
# Install a pod network
# $ kubectl apply -f https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')

# Allow pods to run on the master node
# $ kubectl taint nodes --all node-role.kubernetes.io/master-

$script = <<-SCRIPT

# Install kubernetes
apt-get update && apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet kubeadm kubectl

# kubelet requires swap off 
swapoff -a

# This adds the line twice, but the second time doesn't matter
sed '/ExecStart=/a Environment="KUBELET_EXTRA_ARGS=--cgroup-driver=cgroupfs"' /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

# Get the IP address that VirtualBox has given this VM
IPADDR=`ifconfig eth1 | grep Mask | awk '{print $2}'| cut -f2 -d:`
echo This VM has IP address $IPADDR 

# Set up Kubernetes
kubeadm init --apiserver-cert-extra-sans=$IPADDR  --node-name kube

# Set up admin creds for the vagrant user 
echo Copying credentials to /home/vagrant...
sudo --user=vagrant mkdir -p /home/vagrant/.kube
cp -i /etc/kubernetes/admin.conf /home/vagrant/.kube/config
chown $(id -u vagrant):$(id -g vagrant) /home/vagrant/.kube/config

# User will need to complete the setup:
echo ""
echo As well as adding pod networking you will probably want to allow pods to run on this master node:
echo ""
echo kubectl taint nodes --all node-role.kubernetes.io/master-
echo ""
echo And install the networking plugin
echo "kubectl apply -f \"https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d \'\n\')\""
echo ""
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox" do |v|
    v.memory = 16384
    v.cpus = 2
    v.name = "serverless"
  end

  config.vm.box = "bento/ubuntu-19.04"
  config.vm.network "private_network", ip: "172.28.128.3"
  config.vm.synced_folder ".", "/serverless"
  config.vm.hostname = "serverless"
  config.vm.define "serverless"
  config.vm.provision "docker"
  config.vm.provision "shell", inline: $script
end
