# wx-play
### 配置文件

```bash
  cp ./backend/example.config.toml ./backend/config.toml 
  cp ./frontend/example.project.private.config.json ./frontend/project.private.config.json
```

### 域名配置
```bash
# const baseUrl = 'https://music.demo.net' //修改为你后端域名
vim frontend/utils/api.js
```


### 数据
```bash
# 记得修改source字段和cover字段的URL
source music.mysql
```

![image](images/1.png)
![image](images/2.png)
![image](images/3.png)