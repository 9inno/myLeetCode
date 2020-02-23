<?php
class Solution {

    /**
     * @param Integer[] $arr
     * @return Integer[]
     */
    function sortByBits($arr) {
        $tmp = [];
        $result = [];
        foreach ($arr as $value) {
            $tmpValue = $value;
            $i = 0;
            while ($tmpValue != 0) {
                $i++;
                $tmpValue &= ($tmpValue - 1);
            }
            $tmp[$i][] = $value;
        }
        ksort($tmp);
        foreach ($tmp as $item) {
            $tmpItem = $item;
            sort($tmpItem);
            $result = array_merge($result, $tmpItem);
        }
        return $result;
    }
}

$a = new Solution();
print_r($a->sortByBits([10,100,1000,10000]));