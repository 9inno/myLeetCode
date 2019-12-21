<?php
class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Float
     */
    function findMedianSortedArrays($nums1, $nums2) {
        $n = count($nums1);
        $m = count($nums2);
        if ($n > $m) {
            return $this->findMedianSortedArrays($nums2, $nums1);
        }

        $iMin =0;
        $iMax = $n;

        while ($iMin <= $iMax) {
            $i = intval(($iMin+ $iMax) / 2);
            $j = intval(($n + $m + 1) / 2 - $i);
            if ($j !== 0 && $i !== $n && $nums2[$j-1] > $nums1[$i]) {
                $iMin = $i + 1;
            } elseif ($i !== 0 && $j !== $m && $nums1[$i-1] > $nums2[$j] ) {
                $iMax = $i -1;
            } else {
                if ($i === 0) {
                    $maxLeft = $nums2[$j-1];
                } elseif ($j === 0) {
                    $maxLeft = $nums1[$i-1];
                } else {
                    $maxLeft = max($nums1[$i-1], $nums2[$j-1]);
                }

                if (($n + $m) % 2 === 1) {
                    return $maxLeft;
                }

                if ($i === $n) {
                    $minRight = $nums2[$j];
                } elseif ($j === $m) {
                    $minRight = $nums1[$i];
                } else {
                    $minRight = min($nums1[$i], $nums2[$j]);
                }

                return ($maxLeft + $minRight) / 2;
            }

        }
        return 0;
    }
}



//test
$a = new Solution();
print_r($a->findMedianSortedArrays([1,3], [2]));


/**
 *  解题思路：
 *   中位数是按顺序排列的一组数据中居于中间位置的数
 *   将 长度为n的数组 num1 和 长度为m的num2 同时分成两个部分 切分位置为i,j
 *   条件一 ：保证 n + m 为偶数时  两边数量相等  即 i + j = n - i + m -j   则 j = (n + m ) / 2 -i
 *            m + n 为奇数时  左边比右边数量多1   即 i + J - 1 = n - i + m -j 则 j = (n + m + 1) /2 -i
 *            因 i,j 为位置下标 为整数 所以需要对 ( n + m + 1)/2 向下取整  那么  intval(n + m ) / 2 = intval((n + m + 1) /2)
 *            则 可以合并奇偶情况  j = intval((n+m+1)/2) -i
 *            另 保证  0 <= i <= n  情况下 要保证  0 <= j <= m  必须 要 n <= m
 *            当i =0 时  j = intval((n+m+1)/2) 即  (n + m + 1)/2 <= m  n<=m-1
 *            当i =n 时 (n + m + 1) /2 - n >= 0  n <= m-1
 *   条件二 ： 要满足 max($num1[$i-1], $num2[$j-1]) <= min($num1[$i], $num2[$j])
 *             因为 num1 num2 为有序数组  满足 $num1[$i-1] < $num1[$i] && $num2[$j-1] < $num2[$j]
 *             所以只要判断 $num1[$i -1] < $num2[$j] 和 $num2[$j-1] < $num1[$i]
 *             $num1[$i -1] < $num2[$j]  的情况 需要增加i 减少j 确保满足条件1 且保证 不越界 即 i != n  j!=0
 *             $num2[$j-1] < $num1[$i]  的情况 需要减少i 增加j 确保满足条件1 且保证 不越界 即 j != m  i!=0
 *             边界情况  当 i == 0|| j ==0 时   左半部分最大值 为 num2[$j-1] || num1[$i-1]
 *             当 i == n || j == m 时 右边最最小值 为 num2[$j] || num1[$i]
 *
 *    初始化 i的值 用二分法
 *
 *    时间复杂度O(log(min(n,m)))
 *
 */