# Thực hành với docker-swarm
## I. Easy
### 1. Triển khai các ứng dụng lên docker-swarm với 3 máy ubuntu.
- Yêu cầu: triển khai một số image đơn giản có sẵn image trên docker hub như portainer, traefik, registry

Đầu tiên cần có 3 máy ảo ubuntu trên oracle virtual box. Ta sẽ sử dụng vagrant để tạo 3 máy này với cấu hình 2gb ram, 1 cpu, bên cạnh đó cài sẵn docker và docker-compose lên 3 máy này.

Vagrantfile để tạo 3 máy ảo ubuntu nói trên ở đây.

cd vào folder `3-ubuntu` nơi chứa Vagrantfile, mở terminal lên và dùng lệnh để khởi tạo 3 máy ảo:
```
vagrant up
```

Lần sau khi đã tạo được máy ảo rồi thì vẫn dùng lệnh này để khởi động máy ảo, bởi lệnh này vừa để khởi động, vừa có thể cập nhật mới.

Nếu muốn tắt máy ảo đi thì dùng lệnh:
```
vagrant halt
```

Muốn xóa hẳn máy ảo thì dùng lệnh:
```
vagrant destroy
```

Sau khi khởi tạo được 3 máy ảo, vẫn tại terminal đó, ta dùng lệnh để truy cập vào máy ảo:
```
vagrant ssh <tên máy ảo>
```

Ví dụ ở đây ta có ba máy tên là: `manager01`, `manager02`, `worker01`. Ta sẽ truy cập vào `manager01`, nhớ là chuyển sang root để có quyền cao nhất:
```
sudo -i
```

Ta cd vào `/home/vagrant`

