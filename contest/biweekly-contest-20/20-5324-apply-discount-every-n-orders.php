<?php
class Cashier {
    public $prices = [];
    public $count = 0;
    public $n = null;
    public $discount = null;
    /**
     * @param Integer $n
     * @param Integer $discount
     * @param Integer[] $products
     * @param Integer[] $prices
     */
    function __construct($n, $discount, $products, $prices) {
        $this->n = $n;
        $this->discount = $discount;
        $this->prices = array_combine($products, $prices);

    }

    /**
     * @param Integer[] $product
     * @param Integer[] $amount
     * @return Float
     */
    function getBill($product, $amount) {
        $total = 0;
        foreach ($product as $key => $value) {
            $total += $this->prices[$value] * $amount[$key];
        }
        $this->count ++;
        if ($this->count == $this->n) {
            $this->count = 0;
            return  ($total - ($total * $this->discount) /100);
        }
        return $total;


    }
}

/**
 * Your Cashier object will be instantiated and called as such:
 * $obj = Cashier($n, $discount, $products, $prices);
 * $ret_1 = $obj->getBill($product, $amount);
 */