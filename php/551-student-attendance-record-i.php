<?php
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function checkRecord($s) {

        $arr = str_split($s);
        $countA = 0;
        $countL = 0;
        $continuous = false;
        foreach ($arr as $value) {
            if ($value == 'A') {
                $countA++;
                $countL = 0;
                continue;
            }
            if ($value == 'L') {
                $countL ++;
                if ($countL > 2) {
                    $continuous = true;
                }
                continue;
            }
            $countL = 0;
        }
        return $countA <= 1 && !$continuous;
    }
}

$a = new Solution();
$a->checkRecord(
    "LALL");