<?php

$servername = "localhost";  // 替换为你的数据库主机名
$username = "root";     // 替换为你的数据库用户名
$password = "ycy253255";     // 替换为你的数据库密码
$dbname = "Heimbot";       // 替换为你的数据库名称

// 创建连接
$conn = new mysqli($servername, $username, $password, $dbname);

// 检查连接是否成功
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
}


$tables = array();
$result = $conn->query("SHOW TABLES");

if ($result->num_rows > 0) {
    while ($row = $result->fetch_row()) {
        $tables[] = $row[0];
    }
} else {
    echo "没有找到任何表.";
}


foreach ($tables as $table) {
    $conn->query("TRUNCATE TABLE `$table`");
}

echo "所有表的数据已清空.";

// 关闭连接
$conn->close();

