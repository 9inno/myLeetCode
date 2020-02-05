<?php
class Solution {

    /**
     * @param String $text
     * @return Integer
     */
    function maxNumberOfBalloons($text) {
        $list = [
            'b' => 0,
            'a' => 0,
            'l' => 0,
            'o' => 0,
            'n' => 0,
        ];

        $len = strlen($text);
        for ($i = 0; $i < $len; $i++) {
            if (isset($list[$text[$i]])) {
                $list[$text[$i]] ++;
            }
        }

        $list['l'] = intval($list['l'] / 2);
        $list['o'] = intval($list['o'] / 2);

        sort($list);
        return $list[0];
    }
}

$a = new Solution();
echo $a->maxNumberOfBalloons('loonbalxballpoon');