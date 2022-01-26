# HelmBot V1.0.1

<u>2022/1/27</u>



## 修正错误

- 数据表`Hoshino_graph`中`user1name`和`user2name`长度错误地设置的过短：已改为`varchar(255)`



## 修正的数据库

- `Hoshino_graph`表：存储Helmboard可视化图表的数据来源

| 列名      | 类型         | 说明                     |
| --------- | ------------ | ------------------------ |
| id        | int(11)      | 主键                     |
| group_id  | varchar(15)  | 群id                     |
| user1_id  | varchar(15)  | 用户1id                  |
| user1name | varchar(255) | 用户2名称                |
| user2_id  | varchar(15)  | 用户2id                  |
| user2name | varchar(255) | 用户2名称                |
| value     | int(11)      | 用户1对用户2的社交倾向值 |


