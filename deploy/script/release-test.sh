need_start_server_shell=(
  #rpc
  user-rpc-test.sh
)
for i in ${need_start_server_shell[*]} ; do
  chmod +x $i
  ./$i
done

docker ps

docker exec -it etcd3 etcdctl get --prefix ""