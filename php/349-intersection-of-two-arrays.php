<?php
class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[]
     */
    function intersection($nums1, $nums2) {
        $map = [];
        $result = [];
        foreach ($nums1 as $item) {
            $map[$item] = true;
        }
        foreach ($map as $key => $value) {
            if (in_array($key, $nums2)) {
                $result[] = $key;
            }
        }
        return $result;
    }
}