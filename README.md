# HelmBot V0.2.0

<u>2022/1/20</u>



## 此版本新增改进

- 应用echarts初步将收集的数据通过图表可视化：https://forimoc.com/forimocV2/Helmboard/


- Hoshino\_users表中增加username字段



## 上个版本的问题描述及修复

- 当用户id大于$2^{31}-1$时，user\_id会溢出导致数据存储不正确。已将user\_id和group\_id的类型替换为varchar(15)



## 此版本的数据库接口

Hoshino_status表：机器人状态控制

| 列名     | 类型        | 说明                       |
| -------- | ----------- | -------------------------- |
| id       | int(11)     | 主键                       |
| group_id | varchar(15) | 群id                       |
| status   | int(11)     | 监听状态(1为开启，0为关闭) |

Hoshino\_users表：记录每个发言用户

| 列名     | 类型         | 说明           |
| -------- | ------------ | -------------- |
| id       | int(11)      | 主键           |
| user_id  | varchar(15)  | 用户id         |
| username | varchar(255) | 用户名         |
| group_id | varchar(15)  | 所在群id       |
| freq     | int(11)      | 发送消息总次数 |

Hoshino\_messages表：记录每条消息

| 列名          | 类型          | 说明                        |
| ------------- | ------------- | --------------------------- |
| id            | int(11)       | 主键                        |
| user_id       | varchar(15)   | 发送消息的用户id            |
| group_id      | varchar(15)   | 该消息所在的群id            |
| time          | int(11)       | 时间戳                      |
| content       | varchar(1000) | 内容                        |
| pre_5users_id | varchar(80)   | 该消息的前5条消息的发送者id |

Hoshino_premessages表：缓存每条消息的前5条消息

| 列名     | 类型        | 说明                                         |
| -------- | ----------- | -------------------------------------------- |
| id       | int(11)     | 主键                                         |
| user_id  | varchar(15) | 发送消息的用户id                             |
| group_id | varchar(15) | 该消息所在的群id                             |
| mark     | int(11)     | 先后顺序标记(1~5，1为距离该条消息最近的消息) |

