<?php
$conn = mysqli_connect("localhost", "root", "ycy253255", "Hoshino");
$crontab_conn = mysqli_connect("localhost", "root", "ycy253255", "crontab");
if (!$conn){
    echo "Database connected" . mysqli_connect_error();
}
if (!$crontab_conn){
    echo "Database connected" . mysqli_connect_error();
}
?>