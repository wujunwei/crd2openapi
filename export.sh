#!/bin/bash

NS=kube-system

resourceList=(
deploy
services
endpoints
ingress
secrets
pvc
cm
)

printList(){
  for aa in ${resourceList[@]};
  do
    aList=$(kubectl  -n $NS get $aa |grep -v NAME  |awk '{print $1}')
    if [ ! "${aList[*]}"x == "x" ];then
      [ -d ./$aa ] || mkdir ./$aa
      for i in $aList;
      do
        echo $aa $i
        kubectl -n $NS get $aa $i -o yaml > $aa/$i.yaml
      done
    fi
  done
}

# create namespaces yaml
kubectl  get namespaces $NS -o yaml > namespaces.yaml

# create pv yaml
pvList=$(kubectl get pv |grep "$NS/" |awk '{print $1}')
if [ ! "${pvList[*]}"x == "x" ];then
  for i in ${pvList[@]}
  do
    echo pv $i
    kubectl get pv $i -o yaml > $i.pv.yaml
  done
fi

printList