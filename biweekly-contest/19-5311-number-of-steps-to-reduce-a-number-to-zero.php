<?php
class Solution {

    /**
     * @param Integer $num
     * @return Integer
     */
    function numberOfSteps ($num) {
        $result = 0;
        while ($num != 0) {
            if ($num % 2 == 1) {
                $num--;
            } else {
                $num /= 2;
            }
            $result ++;
        }
        return $result;
    }
}

$a = new Solution();
echo $a->numberOfSteps(8);