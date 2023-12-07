# Tạo một Swarm Cluster

Khởi tạo docker swarm với lệnh:
```bash
docker swarm init --advertise-addr 192.168.56.101
```

Trong đó `192.168.56.101` là ip của máy `manager01`.

Sau khi lệnh được thực thi, sẽ có một câu lệnh được trả về có dạng:
```bash 
Swarm initialized: current node (rh9krtfftxhjvb25t7k6y8gxs) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-0t08e9l15ta4refv1y7jms78er7iguz8wsdrdjunkitylh5wrf-63zl0jxp7retpn7ov6sequk3e 192.168.56.101:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

*Lưu ý: đoạn token này ở máy của bạn có thể khác ví dụ.*

Lấy token dành cho worker:
```
docker swarm join-token worker
```

Lấy token dành cho manager:
```
docker swarm join-token manager
```

Mở thêm tab terminal để truy cập vào máy `manager02` và `worker01`. Tại máy `worker01` ta sẽ sử dụng đoạn lệnh
```
docker swarm join --token SWMTKN-1-0t08e9l15ta4refv1y7jms78er7iguz8wsdrdjunkitylh5wrf-63zl0jxp7retpn7ov6sequk3e 192.168.56.101:2377
```
để biến máy `worker01` trở thành worker node trong docker swarm, làm tương tự với máy `manager02` nhưng sử dụng token dành cho manager.

Sử dụng lệnh để kiểm tra thông tin các node:
```
docker node ls

ID                            HOSTNAME    STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
29fdrzibeubab1jc6nf25lotd *   manager01   Ready     Active         Leader           24.0.2
n0ddmq3i2knka9k0scnpl90q8     manager02   Ready     Active         Reachable        24.0.2
uejoqcuzhomza7jaztjvyzzfj     worker01    Ready     Active                          24.0.2
```