# Sử dụng Docker Config

Bài này tập trung vào việc tạo và quản lý configurations trong Docker Swarm, giúp bạn quản lý cấu hình cho các services một cách an toàn và hiệu quả.

## 1. Tạo Docker Config

Docker Config cho phép bạn lưu trữ các files cấu hình, mà có thể được sử dụng bởi các services trong Swarm. Ví dụ, để tạo một config cho nginx, bạn có thể tạo một file cấu hình nginx đơn giản, sau đó tải nó lên như một Docker Config.

Tạo một file có tên `default.conf`:

```bash
nano default.conf
```

Nội dung bên trong:

```bash
This is nginx config.
```

Tải file này lên Docker config:

```bash
docker config create nginx_default_conf default.conf
```

Trong đó:
* `nginx_default_conf` là tên config.
* `default.conf` là tên file.

Ta sẽ sử dụng sẵn file `docker-compose.yml` có trong bài này để deploy service nginx lên:

```bash
docker stack deploy -c docker-compose.yml nginx
```