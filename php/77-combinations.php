<?php
class Solution {

    /**
     * @param Integer $n
     * @param Integer $k
     * @return Integer[][]
     */
    function combine($n, $k) {
        $result = [];
        $this->recursion($result, $n, $k, 1, []);
        return $result;
    }

    function recursion(&$result, $n, $k, $p, $tmp) {
        if ($p > $n + 1) {
            return;
        }

        if (count($tmp) === $k) {
            $result[] = $tmp;
            return;
        }

        $this->recursion($result, $n, $k, $p + 1, array_merge($tmp, [$p]));
        $this->recursion($result, $n, $k, $p + 1, $tmp);
    }
}

$a = new Solution();
print_r($a->combine(4,2));


/**
 *
 *  解题思路：
 *    每个元素 有 选和不选两种情况
 *    选够 k 个 则结束
 *    TODO 性能不高 优化
 */