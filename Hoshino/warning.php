<?php
if ($status_warning==1){
    // 脏话 filter
    $bad=["傻逼","妈的","操你妈","弱智","脑残","智障","狗东西","狗逼","你妈死了","你妈逼"];
    preg_match_all('/'.implode("|",$bad).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到脏话关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=781815666&message='.$result;
        fopen($url,'r');
    }

    // 键政 filter
    $jz=['中共','土共','社工','入关','国党','兔友','神友','六四','纳粹','视奸'];
    preg_match_all('/'.implode("|",$jz).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到键政关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=781815666&message='.$result;
        fopen($url,'r');
    }

    // 色情 filter
    $sex=['奶子','阴道','阴蒂','阴茎','贞操锁','乳房'];
    preg_match_all('/'.implode("|",$sex).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到色情关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=781815666&message='.$result;
        fopen($url,'r');
    }

    // 学术 filter
    $academic=['python','javascript','C\+\+','java','ruby','go','php','rust','django','css','html','flask','nodejs','angular','vue','react','cms','wordpress',
        'ctf','acm','数模','前端','后端','前后端','力扣','面试','cpc','leetcode','ieee',
        'sql','注入','ssti','ssrf','xss','rce','撞库',
        '保研','考研','实验室','论文','大作业','绩点'];
    preg_match_all('/'.implode("|",$academic).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到学术关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=781815666&message='.$result;
        fopen($url,'r');
    }

    // 恶俗 filter
    $esu=['杨程瑜','唐嘉雄','夏赫绅','骆泽文','徐志','康哲鸣','方陶松','白睿宁','温腾飞','江润汉'];
    preg_match_all('/'.implode("|",$esu).'/i',$data['message'],$match);
    if (isset($match[0][0])) {
        $result='警告: 群'.$data['group_id'].'检测到恶俗关键词['.implode("|",$match[0]).']';
        echo $result;
        $result=urlencode($result);
        $url='http://127.0.0.1:5701/send_group_msg?group_id=781815666&message='.$result;
        fopen($url,'r');
    }
}