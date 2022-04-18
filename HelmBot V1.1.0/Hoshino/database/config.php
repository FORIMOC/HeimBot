<?php
$conn = mysqli_connect("localhost", "root", "xxx", "Hoshino");
$crontab_conn = mysqli_connect("localhost", "root", "xxx", "crontab");
if (!$conn){
    echo "Database connected" . mysqli_connect_error();
}
if (!$crontab_conn){
    echo "Database connected" . mysqli_connect_error();
}
?>