<?php
class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return Integer
     */
    function minSteps($s, $t) {
        $arrS = array_count_values(str_split( $s));
        $arrT = array_count_values(str_split( $t));
        $result = 0;
        foreach ($arrS as $key => $value) {
            if (isset($arrT[$key])){
                $arrT[$key] = abs($arrS[$key] - $arrT[$key]);
            } else {
                $arrT[$key] = $arrS[$key];
            }
        }
        foreach ($arrT as $v) {
            $result += $v;
        }

        return intval($result / 2);
    }
}


$a = new Solution();
echo $a->minSteps('friend', 'family');