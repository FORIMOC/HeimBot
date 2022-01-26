<?php
if ($status_warning==1){
    // 脏话 filter
    $bad=['xxx'];
    preg_match_all('/'.implode("|",$bad).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到脏话关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=your_group_id&message='.$result;
        fopen($url,'r');
    }

    // 键政 filter
    $jz=['xxx'];
    preg_match_all('/'.implode("|",$jz).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到键政关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=your_group_id&message='.$result;
        fopen($url,'r');
    }

    // 色情 filter
    $sex=['xxx'];
    preg_match_all('/'.implode("|",$sex).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到色情关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=your_group_id&message='.$result;
        fopen($url,'r');
    }

    // 学术 filter
    $academic=['xxx'];
    preg_match_all('/'.implode("|",$academic).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到学术关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=your_group_id&message='.$result;
        fopen($url,'r');
    }

    // 恶俗 filter
    $esu=['xxx'];
    preg_match_all('/'.implode("|",$esu).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到恶俗关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=your_group_id&message='.$result;
        fopen($url,'r');
    }
}