# Sử dụng Docker Stack
Ở bài này chúng ta sẽ nhóm các service vào cùng một stack. Việc sử dụng stack giúp quản lý và triển khai ứng dụng một cách dễ dàng và có tổ chức. Một stack trong Docker Swarm là một tập hợp các services, networks, và các cấu hình khác được định nghĩa bằng Docker Compose file.

Để thực hiện bài tập này, chúng ta sẽ sử dụng Docker Compose trong Swarm mode để triển khai một stack bao gồm một web service (ví dụ: nginx) và một database (ví dụ: MySQL).

## 1. Chuẩn bị Docker Compose file

File `docker-compose.yml` ở trong folder này.

## 2. Deploy Stack sử dụng Docker Compose

Tại máy ảo ubuntu `manager01` ta gõ lệnh sau:

```bash
docker stack deploy -c docker-compose.yml mystack
```

Ở đây, `mystack` là tên bạn đặt cho stack của chúng ta.

