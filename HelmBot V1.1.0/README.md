# HelmBot V0.2.0

<u>2022/4/18</u>



## 此版本新增改进

- 新增ht功能：可以以保存到本地图片的方式记录历史事件

  用法如下：

```shell
ht -l(查询所有历史人物)
ht [tag] [key] [截图](记录历史)
ht -ti [tag](查询某位历史人物)
ht -ki [tag] [key](查询某人物的某历史)
```

- 重构项目结构

```shell
Hoshino
  controller # 机器人控制层
  database   # 数据库连接管理
  ht_record  # ht图片保存目录
  images     # 图片
  src        # 主要功能目录
  test       # 测试脚本目录
  util       # 通用功能脚本目录
index.php    # 机器人入口
```

