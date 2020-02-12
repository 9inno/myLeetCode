<?php
class Solution {

    /**
     * @param String $J
     * @param String $S
     * @return Integer
     */
    function numJewelsInStones($J, $S) {
        $arrJ = array_count_values(str_split($J));
        $arrS = array_count_values(str_split($S));
        $result = 0;
        foreach ($arrJ as $key => $value) {
            if (isset($arrS[$key])) {
                $result += $arrS[$key];
            }
        }

        return $result;
    }
}

$a = new Solution();
echo $a->numJewelsInStones('aA', 'aAAbbbb');