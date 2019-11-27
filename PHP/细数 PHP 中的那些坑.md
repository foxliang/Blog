PHP作为“宇宙中最好的编程语言”，被黑只会加速它的改进，所以偶尔黑一下也无妨嘛~正所谓世界上只有两种语言，一种是被人黑的，另一种是没人用的。

### 字符串 == 比较类型强转隐患
```
// php 5
var_dump(md5('240610708') == md5('QNKCDZO'));//bool(true)
var_dump(md5('aabg7XSs') == md5('aabC9RqS'));//bool(true)
var_dump(sha1('aaroZmOk') == sha1('aaK1STfY'));//bool(true)
var_dump(sha1('aaO8zKZF') == sha1('aa3OFF9m'));//bool(true)
var_dump('0010e2' == '1e3');//10×10^2 = 1×10^3  bool(true)
var_dump('0x1234Ab' == '1193131');//bool(true)
var_dump('0xABCdef' == ' 0xABCdef');//bool(true)
var_dump("603E-4234" == "272E-3063");//bool(true)
var_dump('0e1' == '0e2'); //bool(true)
// php 7 含十六进制字符串不再被认为是数字 http://php.net/manual/zh/migration70.incompatible.php
var_dump('0x1234Ab' == '1193131');//bool(false)
var_dump('0xABCdef' == ' 0xABCdef');//bool(false)
var_dump("0x123" == "291");//bool(false)
var_dump(is_numeric("0x123"));//bool(false)
>>> md5('240610708')
=> "0e462097431906509019562988736854"
>>> md5('QNKCDZO')
=> "0e830400451993494058024219903391"
// php 是弱语言，会自动判断数据类型,0eXXXXXXXXXX 转成 0 了 
//来自文档：如果比较一个数字和字符串或者比较涉及到数字内容的字符串，则字符串会被转换为数值并且比较按照数值来进行。此规则也适用于 switch 语句。当用 === 或 !== 进行比较时则不进行类型转换，因为此时类型和数值都要比对。
>>> md5('QNKCDZO')==0
=> true
>>> md5('240610708')==0
=> true

// 使用 === 判断 官方都建议直接用password_hash加密
var_dump(md5('240610708') === md5('QNKCDZO'));//bool(false)
//http://bayescafe.com/php/yuebaomei-ctf.html
var_dump("42"=="0x2A");//bool(true)
var_dump("1" == "01"); // 1 == 1 -> true
var_dump("10" == "1e1"); // 10 == 10 -> true
var_dump(100 == "1e2"); // 100 == 100 -> true
var_dump("\x34\x32\x2E"=="42");//bool(true)
var_dump("\001abc");//abc
var_dump('\001abc');//\001abc
$a = "1234567";
var_dump($a['test']);//1
var_dump(in_array(false, array('xxx')));//false
empty('0');//false
"133" == "0133";
133 == "0133";
133 == 0133;    //因为0133是一个八进制数，转成十进制是91
"0133" != 91;   //字符串中的数字始终是十进制的，这个也可以理解
"0x10" == 16;   //但是!，在十六进制中上面的说法又不成立了
"1e3" == 1000;  //科学计数表示也一样
#'string' == true，而且'string' == 0，但是，true != 0
null == 0;
null < -1;
$flag = "THIS IS FLAG";

if  ("POST" == $_SERVER['REQUEST_METHOD'])
{
    $password = $_POST['password'];//420.00000e-1
    if (0 >= preg_match('/^[[:graph:]]{12,}$/', $password))
    {
        echo 'Wrong Format';
        exit;
    }

    while (TRUE)
    {
        $reg = '/([[:punct:]]+|[[:digit:]]+|[[:upper:]]+|[[:lower:]]+)/';
        if (6 > preg_match_all($reg, $password, $arr))
        break;

        $c = 0;
        $ps = array('punct', 'digit', 'upper', 'lower');
        foreach ($ps as $pt)
        {
            if (preg_match("/[[:$pt:]]+/", $password))
            $c += 1;
        }

        if ($c < 3) break;
        if ("42" == $password) echo $flag;
        else echo 'Wrong password';
        exit;
    }
}

//https://segmentfault.com/q/1010000012046306
$red_money = 143.66;
$receive_money = 14.55;
$residue_money = $red_money > $receive_money ? $red_money - $receive_money : 0;
$receive_money = $residue_money * 100;
var_dump($receive_money);//12911
var_dump((int)$receive_money);/12910
var_dump(intval(12910.9));//int(12910)
var_dump($receive_money*10000);//12910000
var_dump((int)($receive_money*10000));//12910999
var_dump(decbin($receive_money));
var_dump(decbin(12911));
ps: php 7 优化和不兼容

$goo = [
    "bar" => [
        "baz" => 100,
        "cug" => 900
    ]
];

$foo = "goo";

$$foo["bar"]["baz"];  // 实际为：($$foo)['bar']['baz']; PHP 5 中为：${$foo['bar']['baz']};https://zhuanlan.zhihu.com/p/29478077 https://github.com/justcodingnobb/fuck-php-interview https://github.com/todayqq/caseInterviewQuestions
$fn = function (?int $in) {
    return $in ?? "NULL";
};

$fn(null);
$fn(5);
$fn();  // TypeError: Too few arguments to function {closure}()
clipboard.png
```

### PDO bindParam 要求第二个参数是一个引用变量
```
$dbh = new PDO('mysql:host=localhost;dbname=test', "test");
 
$query = <<<QUERY
  INSERT INTO `user` (`username`, `password`) VALUES (:username, :password);
QUERY;
$statement = $dbh->prepare($query);
 
$bind_params = array(':username' => "laruence", ':password' => "weibo");
foreach( $bind_params as $key => $value ){
    $statement->bindParam($key, $value);
}
$statement->execute();
//期望执行 sql
INSERT INTO `user` (`username`, `password`) VALUES ("laruence", "weibo");
// 实际执行 sql
INSERT INTO `user` (`username`, `password`) VALUES ("weibo", "weibo");

//第一次循环
$value = $bind_params[":username"];
$statement->bindParam(":username", &$value); //此时, :username是对$value变量的引用
 
//第二次循环
$value = $bind_params[":password"]; //oops! $value被覆盖成了:password的值
$statement->bindParam(":password", &$value);
// 解决
foreach( $bind_params as $key => &$value ) { //注意这里
    $statement->bindParam($key, $value);
}

return $statement->execute($params);
```
### PHP 引用
```
参考鸟哥一条微博
clipboard.png

    $arr = range(1,3);
    foreach($arr as &$v){
        
    }
    foreach($arr as $v){
        
    }
    print_r($arr);//[1,2,2]
    
 // 解决一
   $arr = range(1,3);
    foreach($arr as &$v){
        
    }
    unset($v);
    foreach($arr as $v){
        
    }
    print_r($arr);//[1,2,3]
    // 解决二
    $arr = range(1,3);
    foreach($arr as &$v){
        
    }
    foreach($arr as $v2){
        
    }
    print_r($arr);//[1,2,3]
    // 解决三
    $arr = range(1,3);
    foreach($arr as &$v){
        
    }
    foreach($arr as &$v){
        
    }
    print_r($arr);//[1,2,3]
array_merge vs +
//
$arr1 = array(1 => "one", "2" => "two", 3 => "three");
$arr2 = array(2 => "new two", 3 => "new three");
print_r($arr1 + $arr2);
Array
(
    [1] => one
    [2] => two
    [3] => three
)
print_r(array_merge($arr1, $arr2));
Array
(
    [0] => one
    [1] => two
    [2] => three
    [3] => new two
    [4] => new three
)
```
### 浮点数精度问题
```
var_dump(15702>=(157.02*100));//bool(false)
var_dump(11111>=(111.11*100));//bool(true)
var_dump(bcsub(15702,(157.02*100)) >= 0);//bool(true)
if(abs(15702-(157.02*100)) < 0.001) {
    echo "相等";
} else {
    echo "不相等";
}

$f = 0.58;
var_dump(intval($f * 100)); //57 0.58 * 100 = 57.999999999...
in_array switch/case 松散比较
 $arr = ['a', 'pro' => 'php', 8, true];
 var_dump(in_array(2, $arr)); // bool(true)
 var_dump(in_array('b', $arr)); // bool(true)
 var_dump(in_array(0, $arr)); // tbool(true)
 var_dump(in_array(null, $arr)); // bool(false)
var_dump(in_array(2, $arr, true)); // bool(false)
var_dump(in_array(0, $arr, true)); // bool(false)
 $name = 0;
 switch ($name) {
          case "a":
               //...
               break;
          case "b":
               //...
               break;
     }
    switch (strval($name)) {
          case "a":
               //...
               break;
          case "b":
               //...
               break;
     } 
     //http://php.net/manual/zh/types.comparisons.php#types.comparisions-loose
 function test($var) 
{ 
switch (true) 
{ 
case 'apple' === $var: 
echo 'apple', PHP_EOL; 
break; 
case 0 === $var: 
echo '0', PHP_EOL; 
break; 
default: 
echo 'default', PHP_EOL; 
} 
}    
  $arr = array('0', 0, 'apple');

foreach ($arr as $value)
{
    test($value);
}   
strpos
function getReferer($link)
{
    $refMap = [
        'baidu' => '百度',
        'sougou' => '搜狗',
        '360' => '360',
        'google' => '谷歌'
    ];
    foreach ($refMap as $key => $value) {
        if (strpos($link, $key) !== false) {
            return $value;
        }
    }
    return '其他';
}
// https://secure.php.net/manual/zh/function.strpos.php 如果 needle 不是一个字符串，那么它将被转换为整型并被视为字符的顺序值。
echo getReferer('https://www.google.com/search?workd=google');//360 
// 解决
function getReferer($link)
{
    $refMap = [
        'baidu' => '百度',
        'sougou' => '搜狗',
        '360' => '360',
        'google' => '谷歌'
    ];
    foreach ($refMap as $key => $value) {
        if (mb_strpos($link, $key) !== false) {
        //if (strpos($link, strval($key)) !== false) {
            return $value;
        }
    }
    return '其他';
}
```
### PHP 不同版本 curl 文件上传
```
//PHP的cURL支持通过给CURL_POSTFIELDS传递关联数组（而不是字符串）来生成multipart/form-data的POST请求 https://blog.zsxsoft.com/post/5
if (class_exists('\CURLFile')) {
    $field = array('fieldname' => new \CURLFile(realpath($filepath)));
} else {
    $field = array('fieldname' => '@' . realpath($filepath));
}
foreach 顺序
$arr=[];
$arr[2] = 2;
$arr[1]  = 1;
$arr[0]  = 0;
foreach ($arr as $key => $val) {
echo $val;// 2 1 0 
}
while (list($key, $v) = each($arr)) {
   //获取不到  foreach会自动reset，each之前, 先reset数组的内部指针
}
for($i=0,$l=count($arr); $i<$l; $i++) {
    echo $arr[$i];// 0 1 2
}
json_decode
>>> json_decode('php')
=> null
// 对非 json 字符串并非返回 null 
>>> json_decode('0x123')
=> 291
echo json_encode(["name" => "php", "age" => "22"]) . "\n";// {"name":"php","age":"22"}
echo json_encode([]) . "\n";//[] 返回这个会让 APP 崩溃
echo json_encode((object)[]) . "\n";//{}
 >>> $a = 0.1 + 0.7
  => 0.8
  >>> printf('%.20f', $a)
  => 0.79999999999999993339
  >>> json_encode($a)
  => "0.7999999999999999"
  >>> \YaJson::encode($a)//https://github.com/seekerliu/laravel-another-json
  => "0.8"
ini_set('serialize_precision', 14);
 $a = 0.1 + 0.7;
 echo json_encode($a);//0.8
echo json_decode(0.7999999999999999);//0.8
strtotime('-x month')
date_default_timezone_set('Asia/Shanghai');
$t = strtotime('2017-08-31');
echo date('Ym',strtotime('- 1 month',$t));//201707
echo date('Ym',strtotime('- 2 month',$t));//201707 ?
// 
$first_day_of_month = date('Y-m',strtotime('2017-08-31')) . '-01 00:00:01';
$t = strtotime($first_day_of_month);
echo date('Ym',strtotime('- 1 month',$t));//201707
echo date('Ym',strtotime('- 2 month',$t));//201706
echo date("Ym", strtotime("-2 month", strtotime("first day of 2017-08-31")));//201706
BOM
//json 解析成 null 写代码时指定 utf-8 without bom 
function remove_utf8_bom($text)
{
    $bom = pack('H*','EFBBBF');
    $text = preg_replace("/^$bom/", '', $text);
    return $text;
}
// ps:PHP导出Excel 可能会乱码，需要写入 BOM头
$content = pack('H*','EFBBBF');
fwrite($fp, $content);
```
### PHP解析大整数
```
$shopId = 17978812896666957068;
var_dump($shopId);//float(1.7978812896667E+19)
$shopId= number_format(17978812896666957068);
var_dump($shopId);//17978812896666957824
$shopId= strval(17978812896666957068);
var_dump($shopId);

$shopId = 17978812896666957068 . '';
var_dump($shopId);//float(1.7978812896667E+19)
$shopId = '17978812896666957068';
var_dump($shopId);

// 输出
// string(20) "17978812896666957068"
超过PHP最大表示范围的纯整数，在MySQL中可以使用bigint/varchar保存，MySQL在查询出来的时候会将其使用string类型保存的。 对于赋值，在PHP里，如果遇到有大整数需要赋值的话，不要尝试用整型类型去赋值$var = '17978812896666957068';
```

### curl获取网页内容出现乱码
```
curl_setopt($ch,CURLOPT_ENCODING,'gzip')//）如果抓取的网页进行了gzip压缩 加入gzip解析
$data = file_get_contents("compress.zlib://".$url); //​​​​  Header 里 Accept-Encoding:gzip 是告诉对方服务器使用 Gzip 进行传输。 ​​​​
trim 中文乱码

echo rawurlencode('河北省');//%E6%B2%B3%E5%8C%97%E7%9C%81
echo rawurlencode('广东省');//%E5%B9%BF%E4%B8%9C%E7%9C%81
echo rtrim('河北省', '省');//河北
echo rtrim('广东省', '省');//广��
//省的十六制作表示是e7 9c 81，而东的十六进制表示是e4 b8 9c，都出现了9c，哦~正因为是rtrim，所以rtrim('广东省', '省')的时候把“东”的十六进制表示的最后一位也被trim掉了。同理rtrim('河北省', '省')不存在此问题

echo trim_cn('广东省','省');//广东
function trim_cn($str, $trim, $charset = 'UTF-8') {
    $len = mb_strlen($str, $charset);
    if (!$len)
        return;
    
    $t1 = $t2 = false;$o=$l=0;
    for($i=0;$i<$len;$i++)
    {
        $str1 = mb_substr($str, $i, 1, $charset);
        $str2 = mb_substr($str, $len-$i-1, 1, $charset);
        if($str1 == $trim && !$t1) $o++; else $t1 = true;
        if($str2 == $trim && !$t2) $l++; else $t2 = true;
    }
    return mb_substr($str, $o, ($len-$l-$o), $charset);;
}
```
### __callStatic
```
//在对象中调用一个不可访问方法时，__call() 会被调用
//在静态上下文中调用一个不可访问方法时，__callStatic() 会被调用 目标方法非 public 时__callStatic 才会起作用。 
class A{
    public static function __callStatic($name, $arguments)
    {
        echo $name.'静态方法不存在!';
    }

    public function test()
    {
        echo 'test 方法';
    }
}

A::test();//test 方法
```
### mb_substr字符编码
```
$str = '北京市朝阳区';
var_dump(mb_substr($str,0,3));//预期输出是：string(9) "北京市"，但是输出确是：string(3) "北"
//mb_substr这个函数在操作的时候如果没有传字符编码，则按照默认的内部编码操作字符串。PHP5.6之前的默认编码都是ISO-8859-1，5.6之后的才是UTF8，UTF8模式下，一个中文字符占3个字节，而ISO-8859-1则是按照一个字节进行处理，所以自然取出来的是异常的字符串。
var_dump(mb_substr($str,0,3,'UTF-8'));//string(9) "北京市"
url参数中的+替换为空格
$name=str_replace('%20','+',$_GET['name']);
```
### 安全 base64
```
 function urlsafeB64Decode($input)
    {
        $remainder = strlen($input) % 4;
        if ($remainder) {
            $padlen = 4 - $remainder;
            $input .= str_repeat('=', $padlen);
        }
        return base64_decode(strtr($input, '-_', '+/'));
    }
 function urlsafeB64Encode($input)
    {
        return str_replace('=', '', strtr(base64_encode($input), '+/', '-_'));
    }
array Undefined offset
```
### php -a
```
php > $a=[];
php > echo $a[1];
PHP Notice:  Undefined offset: 1 in php shell code on line 1
php > $a=null;
php > echo $a[1];
php > $a=4;
php > echo $a[1];
```
### 数组比较
```
array("foo", "bar") != array("bar", "foo");  //这个时候，array就是数组
array("foo" => 1, "bar" => 2) == array("bar" => 2, "foo" => 1);  //这个时候，array又变成了无序hash表
 

$first  = array("foo" => 123, "bar" => 456);
$second = array("foo" => 456, "bar" => 123);
var_dump(array_diff($first, $second));
==> array()
```
### 数组序列化
```
$arrA = array('a', 'b', 'c');
echo json_encode($arrA) . "\n";

$arrB = array('a' => 1, 'b' => 2, 'c' => 3);
echo json_encode($arrB) . "\n";
["a","b","c"]
{"a":1,"b":2,"c":3}
$arrA = array(1 => 'a', 2 => 'b', 3 => 'c');
echo json_encode($arrA) . "\n";

$arrA = array(0 => 'a', 2 => 'b', 3 => 'c');
echo json_encode($arrA) . "\n";
{"1":"a","2":"b","3":"c"}
{"0":"a","2":"b","3":"c"}
//仅当数组的key是从0开始的整数，并且key连续不间断，对其序列化的结果才是JS中的数组，也就是通常认为的PHP索引数组。
json 解析错误
删除控制符
function strip_control_characters($str){
    return preg_replace('/[\x00-\x1F\x7F-\x9F]/u', '', $str);
} 
删除BOM解决：sed -i '1 s/^\xef\xbb\xbf//' file
```
### 32位系统2038年问题
```
1、日期字符串转换为时间戳
$obj = new DateTime("2050-12-31 23:59:59");
echo $obj->format("U"); // 2556115199

// 2、时间戳转换为日期字符串
$obj = new DateTime("@2556115199"); // 这里时间戳前要写一个@符号
$timezone = timezone_open('Asia/HONG_KONG'); // 设置时区
$obj->setTimezone($timezone); 
echo $obj->format("Y-m-d H:i:s"); // 2050-12-31 23:59:59

// 而且DateTime还可以有其他玩法
$obj = new DateTime("2050-12-31 23:59:59");
echo $obj->format("Y/m/d H:i:s"); // 换种方式输入时间字符串2050/12/31 23:59:59

var_dump(strtotime("2050-12-31 23:59:59"));//false

```

[十个 PHP 开发者最容易犯的错误](https://segmentfault.com/a/1190000014126990 )

