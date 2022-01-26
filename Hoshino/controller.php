<?php
if ($data['user_id']==2097517935){
    if ($data['message']=="开启监听"){
        $sql=mysqli_query($conn,"UPDATE Hoshino_status SET status=1");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino.jpg]"}');
        }
    }
    if ($data['message']=="关闭监听"){
        $sql=mysqli_query($conn,"UPDATE Hoshino_status SET status=0");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino_sleep.jpg]"}');
        }
    }
}