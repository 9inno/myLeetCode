<?php
class Solution {

    /**
     * @param Integer $hour
     * @param Integer $minutes
     * @return Float
     */
    function angleClock($hour, $minutes) {

        $angle = abs($hour * 30 + $minutes * 0.5 - $minutes * 6);
        if ($angle > 180) {
            $angle = 360 - $angle;
        }
        return $angle;
    }
}

$a = new Solution();
echo $a->angleClock(12, 30);