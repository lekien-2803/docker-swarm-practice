Khi làm việc với Docker Swarm, chúng ta nên sử dụng Portainer để dễ dàng quản lý Docker environment. Bên cạnh đó chúng ta cũng cần sử dụng một reverse-proxy cũng như là load balancer, vậy nên ta sẽ sử dụng Traefik.

Trong bài tập này, ta sẽ deploy hai service này lên cùng trong một stack. Trong đây đã có sẵn một file là `reverse-proxy`, file này chứa cấu hình của hai service trên.

## 1. Cài đặt DNS

Ở đây chúng ta sẽ cấu hình cho traefik sẽ được mở ở cổng 8888 với subdomain là `traefik.lc` và portainer được mở ở cổng 9000 với subdomain là `portainer.lc`.

Bây giờ các bạn cần sửa file `/etc/hosts` trên máy thật của các bạn để có thể mở được các subdomain này. 
### 1.1 Với hệ điều hành linux/macOS thì cần sửa dưới quyền root với câu lệnh:

```bash
nano /etc/hosts
```

### 1.2 Với hệ điều hành windows

Các bạn cần mở file hosts trong đường dẫn `C:\Windows\System32\drivers\etc`:

![Alt text](./images/hosts-windows.png)

Bởi vì file này chỉ được chỉnh sửa dưới quyền admin, vậy nên các bạn có thể mở CMD hoặc Terminal dưới quyền admin lên để sửa:

![Alt text](./images/terminal-admin.png)

Sau đó dùng lệnh:
```bash
nano hosts
```
Vào chỉnh sửa file hosts, ta sẽ thêm hai subdomain ở trên với cú pháp:

```bash
<địa chỉ ip> <tên miền>
```

![Alt text](./images/edit-hosts.png)

## 2. Deploy reverse-proxy stack

Bây giờ chúng ta sẽ vào trong máy ảo `manager01`, sau đó cd vào thư mục chứa file `reverse-proxy.yml` rồi gõ lệnh:

```bash
docker stack deploy -c reverse-proxy.yml reverse-proxy
```

Bây giờ các bạn sử dụng browser ở máy thật, nhập địa chỉ `portainer.lc:9000` thì sẽ có kết quả:

![Alt text](./images/portainer.png)

Nhập địa chỉ `traefik.lc:8888` thì có kết quả:

![Alt text](./images/traefik.png)

Như vậy là truy cập thành công.
Sau khi tạo tài khoản thì ta đăng nhập vào portainer. 

Đây là giao diện của portainer:

![Alt text](./images/portainer-dashboard.png)

Ta có thể dễ dàng xem các services, containers, images, networks, v.v... mà không cần phải gõ lệnh nhiều lần:

![Alt text](./images/portainer-menu.png)
