#  Sử dụng Docker Registry

Từ đầu đến giờ chúng ta mới chỉ deploy những service đơn giản, sử dụng những images có sẵn trên Docker Hub nên các service này luôn chạy ổn định, kể cả khi chúng ta có đột ngột tắt một máy ảo đi chăng nữa, Swarm cũng sẽ tự điều chỉnh service đồng đều cho hai máy còn lại.

Tuy nhiên, trong thực tế web app của chúng ta phải tự làm, do đó image cho web app đó chúng ta phải tự build từ `Dockerfile`.

Ta sẽ cùng làm bài tập này để hiểu lý do vì sao cần sử dụng Docker Registry.

## 1. Sử dụng local image

Trong folder bài này có chứa một folder `simple-app` có chứa một golang web app, một file `init-data.sql` để khởi tạo bảng và dữ liệu, `docker-compose.yml` và `Dockerfile`.

Giờ ta sẽ cd vào repo này, nơi có chứa cả `docker-compose.yml` và `Dockerfile`.

Trước hết là cần phải build ra được image cho web app từ `Dockerfile`:


```bash
docker build -t go-app:latest .
```

* `go-app` là tên image và nó phải trùng với tag image của service `app` ở trong file docker-compose.yml

Sau khi có image, ta sẽ deploy lên bằng lệnh:

```bash
docker stack deploy -c docker-compose.yml my-stack
```


