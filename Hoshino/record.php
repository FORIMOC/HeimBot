<?php
// 除去匿名用户
if ($data['user_id']==80000000){
    die();
}
// 更新用户录入
$sql=mysqli_query($conn,"SELECT * FROM Hoshino_users WHERE user_id='{$data['user_id']}' AND group_id='{$data['group_id']}'");
if (mysqli_num_rows($sql)>0){
    $row=mysqli_fetch_assoc($sql);
    $new_freq=$row['freq']+1;
    mysqli_query($conn,"UPDATE Hoshino_users SET freq={$new_freq}, username='{$data['sender']['nickname']}' WHERE user_id='{$data['user_id']}' AND group_id='{$data['group_id']}'");
}else{
    mysqli_query($conn,"INSERT INTO Hoshino_users (user_id, username, group_id, freq) VALUES ('{$data['user_id']}', '{$data['sender']['nickname']}', '{$data['group_id']}', 1)");
}

// 更新消息录入
// 此条消息录入messages库
$sql=mysqli_query($conn,"SELECT * FROM Hoshino_premessages WHERE group_id='{$data['group_id']}'");
if (mysqli_num_rows($sql)>=5){
    for ($i=0;$i<5;$i++){
        $row[$i]=mysqli_fetch_assoc($sql);
    }
    $pre_5users_id='';
    for ($i=1;$i<=5;$i++){
        for ($j=0;$j<5;$j++){
            if ($row[$j]['mark']==$i){
                $pre_5users_id.=$row[$j]['user_id'];
                if ($i!=5){
                    $pre_5users_id.='_';
                }
            }
        }
    }
}else{
    $pre_5users_id='initiating...';
}
$time=time();
mysqli_query($conn,"INSERT INTO Hoshino_messages (user_id, group_id, time, content, pre_5users_id) VALUES ('{$data['user_id']}', '{$data['group_id']}', {$time}, '{$data['message']}', '{$pre_5users_id}')");

// 此条消息录入premessages库(mark:相对于此条消息的次序,1为最接近的)
mysqli_query($conn,"DELETE FROM Hoshino_premessages WHERE mark=5 AND group_id='{$data['group_id']}'"); // 删除 5
for ($i=4;$i>0;$i--){
    $new_mark=$i+1;
    mysqli_query($conn,"UPDATE Hoshino_premessages SET mark={$new_mark} WHERE mark={$i} AND group_id='{$data['group_id']}'"); // 1~4 -> 2~5
}
mysqli_query($conn,"INSERT INTO Hoshino_premessages (user_id, group_id, mark) VALUES ('{$data['user_id']}', '{$data['group_id']}', 1)"); // 插入 1