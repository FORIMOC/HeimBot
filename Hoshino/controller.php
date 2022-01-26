<?php
if ($data['user_id']==your_qq){
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

    if ($data['message']=="查询群聊"){
        header("content-type:application/json");
        $url='http://127.0.0.1:5701/get_group_list';
        $data=json_decode(file_get_contents($url));
        $result='Hoshino所监听的群聊:';
        for ($i=0;$i<sizeof($data->data);$i++){
            $result.='\n'.$data->data[$i]->group_name.'('.$data->data[$i]->member_count.'/'.$data->data[$i]->max_member_count.')\n'.$data->data[$i]->group_id;
        }
        exit('{"reply": "'.$result.'"}');
    }

    if (preg_match('/^查询群聊重合度 \d+ \d+$/',$data['message'])){
        preg_match_all('/^查询群聊重合度 (\d+) (\d+)$/',$data['message'],$match);
        $group_id1=$match[1][0];
        $group_id2=$match[2][0];
        header("content-type:application/json");
        $url='http://127.0.0.1:5701/get_group_member_list?group_id='.$group_id1;
        $data1=json_decode(file_get_contents($url));
        $url='http://127.0.0.1:5701/get_group_member_list?group_id='.$group_id2;
        $data2=json_decode(file_get_contents($url));
        if (!isset($data1->data)||!isset($data2->data)){
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
        $result='重合率:\n'.$group_id1.' | '.$count.'/'.sizeof($data1->data).'\n'.$group_id2.' | '.$count.'/'.sizeof($data2->data);
        exit('{"reply": "'.$result.'"}');
    }
}