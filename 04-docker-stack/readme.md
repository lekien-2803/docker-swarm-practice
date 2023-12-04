#  Sử dụng Docker Stack

Ở bài này chúng ta sẽ deploy một ứng dụng web đơn giản có kết nối với database.

Đầu tiên các bạn hãy clone repo này về:

```bash
git clone https://github.com/lekien-2803/postgres-init-people
```

Trong repo này có chứa web app golang và file `init-data.sql` để khởi tạo bảng và dữ liệu. 

Giờ ta sẽ cd vào repo này, nơi có chứa cả `docker-compose.yml` và `Dockerfile`.

Trước hết là cần phải build ra được image cho web app từ `Dockerfile`:

```bash
docker build -t my-golang-app:lastest .
# my-golang-app là tên image và nó phải trùng với tag image của service app ở trong file docker-compose.yml
```

Sau khi có image, ta sẽ deploy lên bằng lệnh:

```bash
docker stack deploy -c docker-compose.yml my-stack
```
