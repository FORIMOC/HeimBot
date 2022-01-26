<?php
$sql=mysqli_query($conn,"SELECT * FROM Hoshino_status WHERE group_id = '{$data['group_id']}'");
if(mysqli_num_rows($sql)>0){
    $row=mysqli_fetch_assoc($sql);
    $status_record=$row['status_record'];
    $status_warning=$row['status_warning'];
}else{
    mysqli_query($conn,"INSERT INTO Hoshino_status (group_id, status_record, status_warning) VALUES ({$data['group_id']}, 1, 1)");
}