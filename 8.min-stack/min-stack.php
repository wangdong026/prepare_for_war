<?php
class stack
{

    private $mStack = [];

    public function push($val)
    {
        $min = $val;
        if (count($this->mStack) > 0 && $this->getMin() < $val) {
            $min = $this->getMin();
        }
        $this->mStack[] = array(
            'value' => $val,
            'min'   => $min,
        );
    }

    public function pop()
    {
        if (count($this->mStack) == 0) {
            return;
        }
        unset($this->mStack[count($this->mStack) - 1]);
    }

    public function getMin()
    {
        return $this->mStack[count($this->mStack) - 1]['min'];
    }

    public function getTop()
    {
        return $this->mStack[count($this->mStack) - 1]['value'];
    }
}

$mStack = new stack();
$mStack->push(-2);
$mStack->push(0);
$mStack->push(-3);
echo $mStack->getMin() . PHP_EOL;
$mStack->pop();
echo $mStack->getTop() . PHP_EOL;
echo $mStack->getMin() . PHP_EOL;