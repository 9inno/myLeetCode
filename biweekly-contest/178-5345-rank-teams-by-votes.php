<?php
class Solution {

    /**
     * @param String[] $votes
     * @return String
     */
    function rankTeams($votes) {
        $tmp = [];
        foreach ($votes as $vote) {
            $arr = str_split($vote);
            for ($i = 0; $i < count($arr) ; $i ++) {
                isset($tmp[$arr[$i]][$i]) ?   $tmp[$arr[$i]][$i]++ : $tmp[$arr[$i]][$i] = 1;
            }
        }
        $result = '';
        foreach ($tmp as $value) {
            arsort($value);
            $result .= key($value);
        }
        return $result;
    }
}

$a = new Solution();
echo $a->rankTeams(["WXYZ","XYZW"]);