在使用php编程的场景中，总有一些情况下需要将数组转为json字符串，就需要用到php自带的json_encode函数。

但是当数组中含有中文字符串时，转出来的结果却是以下结果（unicode字符串）:

<?php

    public function test_json()
    {
        $tmpArr = array(
            'id' => 1,
            'name' => 'fox',
            'desc' => '这里是描述'
        );
        $tmpJson = json_encode($tmpArr);
        echo $tmpJson;
    }

输出：{"id":1,"name":"fox","desc":"\u8fd9\u91cc\u662f\u63cf\u8ff0"}

要想中文不被转为unicode字符串，只需要给json_encode函数中传入一个参数JSON_UNESCAPED_UNICODE即可，如下：

<?php

        $tmpArr = array(
            'id' => 1,
            'name' => 'fox',
            'desc' => '这里是描述'
        );
        $tmpJson = json_encode($tmpArr, JSON_UNESCAPED_UNICODE);
        echo $tmpJson;
?>

输出：{"id":1,"name":"fox","desc":"这里是描述"}

但是PHP版本<5.4.0中并不支持以上参数JSON_UNESCAPED_UNICODE。
