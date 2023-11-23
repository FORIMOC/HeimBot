import mysql.connector

# 连接到数据库
cnx = mysql.connector.connect(
    host='127.0.0.1',
    user='root',
    password='ycy253255',
    database='HoshinoV2'
)

# 创建游标
cursor = cnx.cursor()

# 执行查询语句
query = "SELECT content FROM group_messages ORDER BY created_at DESC LIMIT 20000;"
cursor.execute(query)

# 获取查询结果
results = cursor.fetchall()

# 关闭游标和数据库连接
cursor.close()
cnx.close()
i = 0
# 将结果导出到文本文件
with open('normal_msg.txt', 'w') as file:
    for result in results:
        if i % 10 == 0:
            content = result[0]
            file.write(content + '\n')
        i += 1

print("导出完成！")
