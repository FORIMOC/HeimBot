<?php
$conn = mysqli_connect("localhost", "root", "xxx", "Hoshino");
if (!$conn){
    echo "Database connected" . mysqli_connect_error();
}
?>