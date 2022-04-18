<?php
if ($status_history==1){
    if ($data['message']=="ht") {
        exit('{"reply": "ht -l(查询所有历史人物)\nht [tag] [key] [截图](记录历史)\nht -ti [tag](查询某位历史人物)\nht -ki [tag] [key](查询某人物的某历史)"}');
    }

    if (substr($data['message'], 0, 3) == "ht " && substr($data['message'], 0, 4) != "ht -") {
        if (preg_match_all('/^ht (\S+?) (\S+?) (.+)/', $data['message'], $match)) {
            $tag = addslashes($match[1][0]);
            $keyword = addslashes($match[2][0]);
            $value = addslashes($match[3][0]);
            if (filter($tag) || filter($keyword) || filter($value)) {
                exit('{"reply": "fxxk off,hacker[CQ:face,id=13]"}');
            }
            preg_match_all('/\[CQ:image,file=.*?,url=(.*?)]/', $value, $match);
            if (!isset($match[1][0])){
                exit('{"reply": "不是截图[CQ:face,id=179]"}');
            }
            $content=file_get_contents($match[1][0]);
            $filename=$tag.'_'.$keyword;
            $url='/var/www/html/QQ_robots/Hoshino/ht_record/'.$filename;
            $sql = mysqli_query($conn, "SELECT * FROM Hoshino_history WHERE tag='{$tag}' AND keyword='{$keyword}'");
            if (mysqli_num_rows($sql) > 0) {
                if (in_array($data['user_id'], $white_list)) {
                    file_put_contents($url,$content);
                    $sql2 = mysqli_query($conn, "UPDATE Hoshino_history SET time = {$data['time']}, definer = '{$data['sender']['nickname']}', value = '{$filename}' WHERE tag = '{$tag}' AND keyword = '{$keyword}'");
                    if ($sql2) {
                        exit('{"reply": "Hoshino重写了这段历史[CQ:face,id=284]"}');
                    }
                } else {
                    exit('{"reply": "你无权修改Hoshino的记忆[CQ:face,id=179]"}');
                }
            } else {
                file_put_contents($url,$content);
                mysqli_query($conn, "INSERT INTO Hoshino_history (time, definer, tag, keyword, value) VALUES ({$data['time']}, '{$data['sender']['nickname']}', '{$tag}', '{$keyword}', '{$filename}')");
                exit('{"reply": "Hoshino记住了这段历史"}');
            }
        } else {
            exit('{"reply": "ht [tag] [key] [截图][CQ:face,id=179]"}');
        }
    }

    if (isset($data['group_id'])) {
        $tag = $data['message'];
        $sql4 = mysqli_query($conn, "SELECT * FROM Hoshino_history WHERE tag = '{$tag}'");
        if (mysqli_num_rows($sql4) > 0) {
            for ($i = 0; $i < mysqli_num_rows($sql4); $i++) {
                $row[$i] = mysqli_fetch_assoc($sql4);
            }
            $num = rand(0, mysqli_num_rows($sql4) - 1);
            $url='https://forimoc.com/QQ_robots/Hoshino/ht_record/'.$row[$num]['value'];
            $result = $tag . '之' . $row[$num]['keyword'].'[CQ:image,file='.$url.']';
            exit('{"reply": "' . $result . '"}');
        }
    }

    if (substr($data['message'], 0, 7) == "ht -ki ") {
        $word = substr($data['message'], 7, strlen($data['message']) - 7);
        preg_match_all('/^(\S+?) (.+)/', $word, $match);
        $tag = $match[1][0];
        $keyword = $match[2][0];
        if (filter($tag) || filter($keyword)) {
            exit('{"reply": "fxxk off,hacker[CQ:face,id=13]"}');
        }
        $keyword = mysqli_real_escape_string($conn, $keyword);
        $tag = mysqli_escape_string($conn, $tag);
        $sql5 = mysqli_query($conn, "SELECT * FROM Hoshino_history WHERE tag = '{$tag}' AND keyword='{$keyword}'");
        if (mysqli_num_rows($sql5) > 0) {
            $row = mysqli_fetch_assoc($sql5);
            $time = date('Y-m-d H:i:s', $row['time'] + 8 * 3600);
            $url='https://forimoc.com/QQ_robots/Hoshino/ht_record/'.$row['value'];
            exit('{"reply": "时间: ' . $time . '\n键: [' . $row['tag'] . ']-[' . $row['keyword'] . '] -> 值: [[CQ:image,file='.$url.']]\n定义者: ' . $row['definer'] . '"}');
        } else {
            exit('{"reply": "' . $tag . '之' . $keyword . '是什么[CQ:face,id=32]"}');
        }
    }

    if (substr($data['message'], 0, 7) == "ht -ti ") {
        $tag = substr($data['message'], 7, strlen($data['message']) - 7);
        if (filter($tag)) {
            exit('{"reply": "fxxk off,hacker[CQ:face,id=13]"}');
        }
        $tag = mysqli_real_escape_string($conn, $tag);
        $sql5 = mysqli_query($conn, "SELECT * FROM Hoshino_history WHERE tag='{$tag}'");
        if (mysqli_num_rows($sql5) > 0) {
            $result = '[' . $tag . ']:';
            for ($i = 0; $i < mysqli_num_rows($sql5); $i++) {
                $row[$i] = mysqli_fetch_assoc($sql5);
                $keyword = $row[$i]['keyword'];
                $result .= '\n' . $keyword;
            }
            exit('{"reply": "' . $result . '"}');
        } else {
            exit('{"reply": "' . $tag . '是什么[CQ:face,id=32]"}');
        }
    }

    if ($data['message'] == "ht -l") {
        $sql = mysqli_query($conn, 'SELECT * FROM Hoshino_history');
        $tag_array = array();
        for ($i = 0; $i < mysqli_num_rows($sql); $i++) {
            $row[$i] = mysqli_fetch_assoc($sql);
            if (!in_array($row[$i]['tag'], $tag_array)) {
                array_push($tag_array, $row[$i]['tag']);
            }
        }
        $result = "所有历史人物:";
        for ($i = 0; $i < sizeof($tag_array); $i++) {
            $result .= '\n' . $tag_array[$i];
        }
        exit('{"reply": "' . $result . '"}');
    }
}