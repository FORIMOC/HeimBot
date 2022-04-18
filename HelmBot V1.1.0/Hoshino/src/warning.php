<?php
if ($status_warning==1&&!in_array($data['user_id'], $white_list)){
    $tag=['脏话','键政','色情','学术','恶俗'];
    //脏话
    $keyword[0]=['傻逼','妈的','操你妈','弱智','脑残','智障','狗东西','狗逼','你妈死了','你妈逼'];
    //键政
    $keyword[1]=['中共','土共','社工','入关','国党','兔友','神友','六四','纳粹','视奸','开倒车'];
    //色情
    $keyword[2]=['奶子','阴道','阴蒂','阴茎','贞操锁','乳房'];
    //学术
    $keyword[3]=['python','javascript','C\+\+','java','ruby','golang','database','django','css','flask','nodejs','angular','vue','react','cms','wordpress',
        'ctf','acm','数模','前端','后端','前后端','力扣','面试','cpc','leetcode','ieee','比赛','竞赛',
        'sql','注入','ssti','ssrf','xss','撞库',
        '保研','考研','实验室','论文','大作业','绩点','上岸','直博','直硕','读研','学硕','专硕','调剂','复试','跨考','双非',
        '非全','研究生','推免','夏令营','必修','选修','专业','综排','科研','作弊','网课'];
    //恶俗
    $keyword[4]=['杨程瑜','唐嘉雄','夏赫绅','庞博','骆泽文','徐志','康哲鸣','方陶松','白睿宁','温腾飞','江润汉'];
    $flag=0;
    $group_name=$data['group_id'];
    $content=json_decode(file_get_contents('http://127.0.0.1:5701/get_group_list'));
    $group_info=$content->data;
    foreach ($group_info as $key => $value){
        if ($data['group_id']==$value->group_id){
            $group_name=$value->group_name;
            break;
        }
    }
    $group_name=addslashes($group_name);
    $result=urlencode('**['.$group_name.']-('.$data['group_id'].')检测到关键词**');
    $message=$data['message'];
    for ($i=0;$i<sizeof($keyword);$i++){
        preg_match_all('/'.implode("|",$keyword[$i]).'/i',$message,$match);
        if (isset($match[0][0])){
            $flag=1;
            $match_deduplicated=[];
            for ($j=0,$k=0;$j<sizeof($match[0]);$j++){
                if (!in_array($match[0][$j],$match_deduplicated)){
                    $match_deduplicated[$k]=$match[0][$j];
                    $k++;
                }
            }
            $result.='%0a'.urlencode($tag[$i].': '.implode("|",$match_deduplicated));
        }
    }
    if ($flag) {
        $result.='%0a'.urlencode($data['sender']['nickname'].'('.$data['user_id'].') > '.$data['message']);
        $broadcast_list = ['781815666'];
        send('Hoshino', $broadcast_list, $result);
    }else{
        die();
    }
}