<?php
class ProductOfNumbers {
    public $numbers;
    public $result;
    /**
     */
    function __construct() {
        $this->numbers = [];
        $this->result = [];
    }

    /**
     * @param Integer $num
     * @return NULL
     */
    function add($num) {
        $this->numbers[] = $num;
        $this->result = [];
    }

    /**
     * @param Integer $k
     * @return Integer
     */
    function getProduct($k) {
        if (isset($this->result[$k])) {
            return $this->result[$k];
        }
        $count = count($this->numbers) - 1;
        $result = null;

        for ($i = $count; $i > $count - $k; $i --) {
            if ($result === null) {
                $result = $this->numbers[$i];
            }else {
                $result *= $this->numbers[$i];
            }
        }
        if (!isset($this->result[$k])) {
            $this->result[$k] = $result;
        }
        return $result;
    }
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * $obj = ProductOfNumbers();
 * $obj->add($num);
 * $ret_2 = $obj->getProduct($k);
 */