# Multi-service Stack

## Mục tiêu
Deploy một ứng dụng web đơn giản có front-end, back-end và database.

## Thực hành
Bài này cũng tương tự như bài 11, nhưng chúng ta sẽ có một số thay đổi:


- Ta sẽ sử dụng ứng dụng web trong thư mục `simple-app`. Tuy nhiên ta sẽ thêm phần giao diện front-end tại đường dẫn `domain.lc:8088/list-people`.

- Áp dụng kiến thức về secret trong bài 09, ta sẽ lưu mật khẩu của database trong file secret chứ không để lộ ra trong tag `enviroment`. Cùng với đó là sử dụng config trong bài 10 để cấu hình web hiển thị ở đường dẫn `domain.lc:8088/list-people`.

## 1. Khởi tạo image go-app

Build image:

```bash
docker build -t go-app:latest .
```

## 2. Tạo Docker Secret

Tạo một file tạm thời tên là `postgres_password.txt` trên máy ảo `manager01`:

```bash
vim postgres_password.txt
```

Sau đó ghi mật khẩu để đăng nhập vào của các bạn vào file này, ví dụ như của tôi sẽ là `"abc123-"`:

```bash
abc123-
```

Tạo Docker Secret từ file `postgres_password.txt`:

```bash
docker secret create postgres_password postgres_password.txt
```

* Ở đây, `postgres_password` là tên của secret.

Sau khi tạo secret, xóa file tạm để đảm bảo rằng mật khẩu không còn tồn tại trên hệ thống của bạn:

```bash
rm postgres_password.txt
```

## 3. Tạo Docker Config

Các bạn tạo một file có tên là `default.conf`:

```bash
vim default.conf
```
Nội dung bên trong sẽ là:

```nginx
server {
    listen       8090;
    server_name  localhost;

    location /list-people {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    # Cấu hình khác...
}
```

Tải file này lên Docker config:

```bash
docker config create nginx_conf default.conf
```

***Lưu ý: Nếu các bạn vẫn còn giữ Registry từ bài 11 thì không cần làm lại bước 4 và bước 5.***

## 4. Thiết lập Private Docker Registry

Chạy Private Docker Registry trên node `manager01`:

```bash
docker stack deploy -c registry.yml registry
```

## 5. Cấu hình các Node Swarm để trust Registry

Ta truy cập vào hai máy node còn lại là `manager02` và `worker01`. 

Sau đó chỉnh sửa hoặc tạo file `/etc/docker/daemon.json` trên mỗi node với lệnh:

```bash
vim /etc/docker/daemon.json
```

Với nội dung:

```bash
{
  "insecure-registries" : ["192.168.56.101:5000"]
}
```

Ta restart Docker Daemon trên hai node vừa sửa:

```bash
sudo systemctl restart docker
```

## 6. Tag và push image vào Registry

Tag image `go-app` vào Registry:

```bash
docker tag go-app localhost:5000/go-app
```

Push image vào Registry:

```bash
docker push localhost:5000/go-app
```

## 7. Deploy multi-service stack

Giờ chúng ta sẽ deploy stack:

```bash
docker stack deploy -c docker-compose.yml web-app
```
