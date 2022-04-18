<?php
$white_list = [];
$target_list = [];
$sql=mysqli_query($conn,"SELECT * FROM Hoshino_whitelist");
for ($i=0;$i<mysqli_num_rows($sql);$i++){
    $row[$i]=mysqli_fetch_assoc($sql);
    array_push($white_list, $row[$i]['user_id']);
}
$sql=mysqli_query($conn,"SELECT * FROM Hoshino_targetlist");
for ($i=0;$i<mysqli_num_rows($sql);$i++){
    $row[$i]=mysqli_fetch_assoc($sql);
    array_push($target_list, $row[$i]['group_id']);
}

$sql=mysqli_query($conn,"SELECT * FROM Hoshino_status WHERE group_id = '{$data['group_id']}'");
if(mysqli_num_rows($sql)>0){
    $row=mysqli_fetch_assoc($sql);
    $status_record=$row['status_record'];
    $status_warning=$row['status_warning'];
    $status_history=$row['status_history'];
    $authorized=(in_array($data['user_id'], $white_list)&&in_array($data['group_id'], $target_list))??false;
}else{
    mysqli_query($conn,"INSERT INTO Hoshino_status (group_id, status_record, status_warning, status_history) VALUES ({$data['group_id']}, 0, 1, 0)");
}
// 授权用户控制层
if ($authorized){

    if (preg_match('/^Hoshino$/i',$data['message'])){
        exit('{"reply": "HelmBot:QQ群监听机器人 V1.0.2[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino_icon.jpg]开启监听/关闭监听 [群号]\n开启报警/关闭报警\n开启推送/关闭推送\n查询群聊\n查询群聊重合度 [群号1] [群号2]"}');
    }

    if (preg_match('/^开启监听 \d+$/',$data['message'])){
        preg_match_all('/^开启监听 (\d+)$/',$data['message'],$match);
        $group_id=$match[1][0];
        $sql=mysqli_query($conn,"UPDATE Hoshino_status SET status_record=1 WHERE group_id='{$group_id}'");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino.jpg]"}');
        }else{
            exit('{"reply": "不存在此群"}');
        }
    }
    if (preg_match('/^关闭监听 \d+$/',$data['message'])){
        preg_match_all('/^关闭监听 (\d+)$/',$data['message'],$match);
        $group_id=$match[1][0];
        $sql=mysqli_query($conn,"UPDATE Hoshino_status SET status_record=0 WHERE group_id='{$group_id}'");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino_sleep.jpg]"}');
        }else{
            exit('{"reply": "不存在此群"}');
        }
    }
    if ($data['message']=="开启报警"){
        $sql=mysqli_query($conn,"UPDATE Hoshino_status SET status_warning=1");
        foreach ($target_list as $key => $value){
            mysqli_query($conn,"UPDATE Hoshino_status SET status_warning=0 WHERE group_id='{$value}'");
        }
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino.jpg]"}');
        }
    }
    if ($data['message']=="关闭报警"){
        $sql=mysqli_query($conn,"UPDATE Hoshino_status SET status_warning=0");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino_sleep.jpg]"}');
        }
    }
    if ($data['message']=="开启推送"){
        $sql=mysqli_query($crontab_conn,"UPDATE crontab_status SET status=1");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino.jpg]"}');
        }
    }
    if ($data['message']=="关闭推送"){
        $sql=mysqli_query($crontab_conn,"UPDATE crontab_status SET status=0");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino_sleep.jpg]"}');
        }
    }
    if ($data['message']=="开启历史"){
        $sql=mysqli_query($conn, "UPDATE Hoshino_status SET status_history=1 WHERE group_id='{$data['group_id']}'");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino.jpg]"}');
        }
    }
    if ($data['message']=="关闭历史"){
        $sql=mysqli_query($conn, "UPDATE Hoshino_status SET status_history=0 WHERE group_id='{$data['group_id']}'");
        if ($sql){
            exit('{"reply": "[CQ:image,file=https://forimoc.com/QQ_robots/Hoshino/images/Hoshino_sleep.jpg]"}');
        }
    }

if ($data['user_id']==2097517935){
    if (preg_match('/^qq白名单$/i',$data['message'])){
        $result='Hoshino的qq白名单如下:';
        foreach ($white_list as $key => $value){
            $result.='\n'.$value;
        }
        exit('{"reply": "'.$result.'"}');
    }
    if (preg_match('/^qq群白名单$/i',$data['message'])){
        $result='Hoshino的qq群白名单如下:';
        foreach ($target_list as $key => $value){
            $result.='\n'.$value;
        }
        exit('{"reply": "'.$result.'"}');
    }
    if (preg_match_all('/^添加qq白名单 (\d+)$/i',$data['message'],$match)){
        $user_id=$match[1][0];
        $sql=mysqli_query($conn,"SELECT * FROM Hoshino_whitelist WHERE user_id='{$user_id}'");
        if (mysqli_num_rows($sql)>0){
            exit('{"reply": "已在白名单中"}');
        }else{
            $sql=mysqli_query($conn,"INSERT INTO Hoshino_whitelist (user_id) VALUES ('{$user_id}')");
            if ($sql){
                exit('{"reply": "添加成功[CQ:face,id=282]"}');
            }
        }
    }
    if (preg_match_all('/^删除qq白名单 (\d+)$/i',$data['message'],$match)){
        $user_id=$match[1][0];
        $sql=mysqli_query($conn,"SELECT * FROM Hoshino_whitelist WHERE user_id='{$user_id}'");
        if (mysqli_num_rows($sql)==0){
            exit('{"reply": "不在白名单中"}');
        }else{
            $sql=mysqli_query($conn,"DELETE FROM Hoshino_whitelist WHERE user_id='{$user_id}'");
            if ($sql){
                exit('{"reply": "删除成功[CQ:face,id=282]"}');
            }
        }
    }
    if (preg_match_all('/^添加qq群白名单 (\d+)$/i',$data['message'],$match)){
        $group_id=$match[1][0];
        $sql=mysqli_query($conn,"SELECT * FROM Hoshino_targetlist WHERE group_id='{$group_id}'");
        if (mysqli_num_rows($sql)>0){
            exit('{"reply": "已在白名单中"}');
        }else{
            $sql=mysqli_query($conn,"INSERT INTO Hoshino_targetlist (group_id) VALUES ('{$group_id}')");
            if ($sql){
                exit('{"reply": "添加成功[CQ:face,id=282]"}');
            }
        }
    }
    if (preg_match_all('/^删除qq群白名单 (\d+)$/i',$data['message'],$match)){
        $group_id=$match[1][0];
        $sql=mysqli_query($conn,"SELECT * FROM Hoshino_targetlist WHERE group_id='{$group_id}'");
        if (mysqli_num_rows($sql)==0){
            exit('{"reply": "不在白名单中"}');
        }else{
            $sql=mysqli_query($conn,"DELETE FROM Hoshino_targetlist WHERE group_id='{$group_id}'");
            if ($sql){
                exit('{"reply": "删除成功[CQ:face,id=282]"}');
            }
        }
    }

}


    if ($data['message']=="查询群聊"){
        header("content-type:application/json");
        $url='http://127.0.0.1:5701/get_group_list';
        $data=json_decode(file_get_contents($url));
        $result='Hoshino所监听的群聊:';
        for ($i=0;$i<sizeof($data->data);$i++){
            if (!in_array($data->data[$i]->group_id, $target_list)){
                $result.='\n'.$data->data[$i]->group_name.'('.$data->data[$i]->member_count.'/'.$data->data[$i]->max_member_count.')\n'.$data->data[$i]->group_id;
            }
        }
        exit('{"reply": "'.$result.'"}');
    }

    if (preg_match_all('/^查询群聊重合度 (\d+) (\d+)$/',$data['message'],$match)){
        $group_id1=$match[1][0];
        $group_id2=$match[2][0];
        header("content-type:application/json");
        $url='http://127.0.0.1:5701/get_group_member_list?group_id='.$group_id1;
        $data1=json_decode(file_get_contents($url));
        $url='http://127.0.0.1:5701/get_group_member_list?group_id='.$group_id2;
        $data2=json_decode(file_get_contents($url));
        if (in_array($group_id1, $target_list)||in_array($group_id2, $target_list)||!isset($data1->data)||!isset($data2->data)){
            exit('{"reply": "不存在此群聊"}');
        }
        $count=0;
        for ($i=0;$i<sizeof($data1->data);$i++){
            for ($j=0;$j<sizeof($data2->data);$j++){
                if ($data1->data[$i]->user_id==$data2->data[$j]->user_id){
                    $count++;
                }
            }
        }
        $result='重合率:\n'.$group_id1.' | '.$count.'/'.sizeof($data1->data).'('.round($count*100/sizeof($data1->data),2).'%)\n'.$group_id2.' | '.$count.'/'.sizeof($data2->data).'('.round($count*100/sizeof($data2->data),2).'%)';
        exit('{"reply": "'.$result.'"}');
    }
}