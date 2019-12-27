<?php
class Solution {
    public  $count = 0;
    /**
     * @param Integer $m
     * @param Integer $n
     * @return int
     */
    function uniquePaths($m, $n) {

        if ($m === 1 && $n === 1) {
            $this->count++;
            return $this->count;
        }

        if ($n > 1) {

            $this->uniquePaths($m , $n -1);
        }
        if($m > 1) {
            $this->uniquePaths($m-1, $n);
        }
        return $this->count;
    }


    function uniquePaths_dp($m, $n) {
        $arr = [];
        for ($i = 1; $i <= $m; $i++) {
            for ($j = 1; $j <= $n; $j++) {
                if ($i == 1 || $j == 1) {
                    $arr[$i][$j] = 1;
                } else {
                    $arr[$i][$j] = $arr[$i-1][$j] + $arr[$i][$j-1];
                }
            }
        }
        return $arr[$m][$n];
    }
}

$a = new Solution();
echo $a->uniquePaths_dp(10,10);

/**
 *   解题思路：
 *     m * n 个格子  构造成 m * n 的二维数组
 *     第 (m,n) 个格子的走法 等于 (m-1, n) + (m, n-1) 的和
 *    所以将二维数组填满 取 $arr[$m][$n]
 *    时间复杂度 O( m * n)
 *    递归超时啦 哈哈哈哈。
 */