<?php
if ($status_record==1&&$data['user_id']!=80000000){
    // 除去匿名用户
    if ($data['post_type']!="message"){
        die();
    }

// 过滤单双引号
    $username=mysqli_escape_string($conn,$data['sender']['nickname']);
    $content=mysqli_escape_string($conn,$data['message']);

// 更新用户录入
    $sql=mysqli_query($conn,"SELECT * FROM Hoshino_users WHERE user_id='{$data['user_id']}' AND group_id='{$data['group_id']}'");
    if (mysqli_num_rows($sql)>0){
        $row=mysqli_fetch_assoc($sql);
        $new_freq=$row['freq']+1;
        mysqli_query($conn,"UPDATE Hoshino_users SET freq={$new_freq}, username='{$username}' WHERE user_id='{$data['user_id']}' AND group_id='{$data['group_id']}'");
    }else{
        mysqli_query($conn,"INSERT INTO Hoshino_users (user_id, username, group_id, freq) VALUES ('{$data['user_id']}', '{$username}', '{$data['group_id']}', 1)");
    }

// 更新消息录入
// 此条消息录入messages表
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
    mysqli_query($conn,"INSERT INTO Hoshino_messages (user_id, group_id, time, content, pre_5users_id) VALUES ('{$data['user_id']}', '{$data['group_id']}', {$time}, '{$content}', '{$pre_5users_id}')");

// 此条消息录入premessages表(mark:相对于此条消息的次序,1为最接近的)
    mysqli_query($conn,"DELETE FROM Hoshino_premessages WHERE mark=5 AND group_id='{$data['group_id']}'"); // 删除 5
    for ($i=4;$i>0;$i--){
        $new_mark=$i+1;
        mysqli_query($conn,"UPDATE Hoshino_premessages SET mark={$new_mark} WHERE mark={$i} AND group_id='{$data['group_id']}'"); // 1~4 -> 2~5
    }
    mysqli_query($conn,"INSERT INTO Hoshino_premessages (user_id, username, group_id, content, mark) VALUES ('{$data['user_id']}', '{$username}', '{$data['group_id']}', '{$content}', 1)"); // 插入 1

// 更新graph表
    if ($pre_5users_id!='initiating...'){
        $array=explode('_',$pre_5users_id);
        for ($i=0;$i<5;$i++){
            if ($data['user_id']==$array[$i]){
                continue;
            }
            $sql=mysqli_query($conn,"SELECT * FROM Hoshino_graph WHERE group_id='{$data['group_id']}' AND user1_id='{$data['user_id']}' AND user2_id='{$array[$i]}'");
            if (mysqli_num_rows($sql)>0){
                $row=mysqli_fetch_assoc($sql);
                $new_value=$row['value']+1;
                $sql=mysqli_query($conn,"SELECT * FROM Hoshino_users WHERE group_id='{$data['group_id']}' AND user_id='{$array[$i]}'");
                $row=mysqli_fetch_assoc($sql);
                $user2name=$row['username'];
                mysqli_query($conn,"UPDATE Hoshino_graph SET value={$new_value}, user2name='{$user2name}' WHERE group_id='{$data['group_id']}' AND user1_id='{$data['user_id']}' AND user2_id='{$array[$i]}'");
            }else{
                $sql=mysqli_query($conn,"SELECT * FROM Hoshino_users WHERE group_id='{$data['group_id']}' AND user_id='{$array[$i]}'");
                $row=mysqli_fetch_assoc($sql);
                $user2name=$row['username'];
                mysqli_query($conn,"INSERT INTO Hoshino_graph (group_id, user1_id, user1name, user2_id, user2name, value) VALUES ('{$data['group_id']}', '{$data['user_id']}', '{$data['sender']['nickname']}', '{$array[$i]}', '{$user2name}', 1)");
            }
        }
    }
}