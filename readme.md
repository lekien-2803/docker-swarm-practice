# Thực hành với docker-swarm
*Yêu cầu: Đã có 3 máy ảo ubuntu, chi tiết có thể tham khảo tại đây.*
## I. Cơ bản
### 1. Cài đặt Docker 

Vì trong Vagrantfile đã có sẵn lệnh khởi chạy cài đặt docker rồi nên ta không cần cài thủ công nữa.

### 2. Tạo một Swarm Cluster

Khởi tạo docker swarm với lệnh:
```
docker swarm init --advertise-addr <ip của máy làm manager>
```

Sau khi lệnh được thực thi, sẽ có một câu lệnh được trả về có dạng:
```
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

### 3. Deploy một Service đơn giản

Trong bài này, chúng ta sẽ sử dụng nginx, một web server phổ biến, làm service để deploy.

Trên máy Manager, hãy mở terminal và chạy lệnh sau để tạo một service mới với tên my-web:

```
docker service create --name my-web --publish 8081:80 nginx
```

Ở đây, `--publish 8081:80`` nghĩa là chúng ta đang map cổng 8081 trên máy host (Swarm manager) tới cổng 80 trên container nginx.

