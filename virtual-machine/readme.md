# Hướng dẫn nhanh về vagrant
Đầu tiên cần có 3 máy ảo ubuntu trên oracle virtual box. Ta sẽ sử dụng vagrant để tạo 3 máy này với cấu hình 2gb ram, 1 cpu, bên cạnh đó cài sẵn docker và docker-compose lên 3 máy này.

Với hệ điều hành Windows, các ta nên dùng PowerShell hoặc CMD chạy dưới quyền admin. Còn đối với Linux/MacOS thì nên chạy dưới quyền root.

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

Các lệnh cơ bản của vagrant có thể tham khảo thêm ở [đây](https://gist.github.com/wpscholar/a49594e2e2b918f4d0c4).