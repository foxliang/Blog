# 冒泡排序算法
```
    public function test() {
        $arr = array(43, 54, 62, 21, 66, 32, 78, 36, 76, 39);
        var_dump($arr);
        echo '<br/>';
        $arr = $this->bubbleSort($arr);
        var_dump($arr);
    }

    public function bubbleSort($arr) {
        $len = count($arr);
        //该层循环控制 需要冒泡的轮数
        for ($i = 1; $i < $len; $i++) {
            //该层循环用来控制每轮 冒出一个数 需要比较的次数
            for ($k = 0; $k < $len - $i; $k++) {
                if ($arr[$k] > $arr[$k + 1]) {
                    $tmp = $arr[$k + 1]; // 声明一个临时变量
                    $arr[$k + 1] = $arr[$k];
                    $arr[$k] = $tmp;
                }
            }
        }
        return $arr;
    }
```
# 快速排序
```
 public function quick_sort($arr) {
        //先判断是否需要继续进行
        $length = count($arr);
        if ($length <= 1) {
            return $arr;
        }
        $base_num = $arr[0]; //选择一个标尺 选择第一个元素
        //初始化两个数组
        $left_array = array(); //小于标尺的
        $right_array = array(); //大于标尺的
        for ($i = 1; $i < $length; $i++) {      //遍历 除了标尺外的所有元素，按照大小关系放入两个数组内
            if ($base_num > $arr[$i]) {
                //放入左边数组
                $left_array[] = $arr[$i];
            } else {
                //放入右边
                $right_array[] = $arr[$i];
            }
        }
        //再分别对 左边 和 右边的数组进行相同的排序处理方式
        //递归调用这个函数,并记录结果
        $left_array = $this->quick_sort($left_array);
        $right_array = $this->quick_sort($right_array);
        //合并左边 标尺 右边
        return array_merge($left_array, array($base_num), $right_array);
    }

    public function test() {
        $arr = array(4, 3, 1, 2, 8, 9);
        var_dump($arr);
        echo '<br/>';
        $arr = $this->quick_sort($arr);
        var_dump($arr);
    }
```
# 二分查找
```
    public function bin_search($arr, $low, $high, $k) {
        if ($low <= $high) {
            $mid = intval(($low + $high) / 2);
            if ($arr[$mid] == $k) {
                return $mid;
            } else if ($k < $arr[$mid]) {
                return $this->bin_search($arr, $low, $mid - 1, $k);
            } else {
                return $this->bin_search($arr, $mid + 1, $high, $k);
            }
        }
        return -1;
    }

    public function test() {
        $arr = array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        var_dump($arr);
        echo '<br/>';
        $arr = $this->bin_search($arr, 0, 8, 4);
        var_dump($arr);
    }
```
# 顺序查找
```
   public function seq_search($arr, $n, $k) {
        $array[$n] = $k;
        for ($i = 0; $i < $n; $i++) {
            if ($arr[$i] == $k) {
                break;
            }
        }
        if ($i < $n) {
            return $i;
        } else {
            return -1;
        }
    }

    public function test_suanfa() {
        $arr = array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        var_dump($arr);
        echo '<br/>';
        $arr = $this->seq_search($arr, 4, 4);
        var_dump($arr);
    }
```

# 线性表的删除
```
    public function delete_array_element($array, $i) {
        $len = count($array);
        for ($j = $i; $j < $len; $j ++) {
            if (isset($array[$j + 1])) {
                $array[$j] = $array[$j + 1];
            }
        }
        array_pop($array);
        return $array;
    }

    public function test() {
        $arr = array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);
        var_dump($arr);
        echo '<br/>';
        $arr = $this->delete_array_element($arr, 4);
        var_dump($arr);
    }
```
# 字符串翻转
```
    public function strrev($str) {
        $rev_str = '';
        if ($str == '') {
            return 0;
        }
        for ($i = (strlen($str) - 1); $i >= 0; $i --) {
            $rev_str .= $str[$i];
        }
        return $rev_str;
    }

    public function test() {
        $arr = 'hellow wolrd';
        var_dump($arr);
        echo '<br/>';
        $arr = $this->strrev($arr);
        var_dump($arr);
    }
```
