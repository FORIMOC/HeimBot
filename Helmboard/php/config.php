<?php
$conn = mysqli_connect("localhost", "root", "ycy253255", "Hoshino");
if (!$conn){
    echo "Database connected" . mysqli_connect_error();
}
?>