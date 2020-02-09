<?php
class Solution {

    /**
     * @param Integer[] $arr
     * @param Integer $k
     * @param Integer $threshold
     * @return Integer
     */
    function numOfSubarrays($arr, $k, $threshold) {
        $len = count($arr) - $k + 1;
        $result = 0;
        $i = 0;
        $j = $k - 1;
        $sum = 0;
        for ($z = 0; $z < $k; $z ++) {
            $sum+= $arr[$z];
        }
        while ($i < $len) {
            if ($sum / $k >= $threshold) {
                $result ++;
            }
            $sum -= $arr[$i];
            $sum += $arr[$j+1];
            $i ++;
            $j ++;
        }
        return $result;
    }
}

$a = new Solution();
echo $a->numOfSubarrays([2,2,2,2,5,5,5,8],3 ,4);