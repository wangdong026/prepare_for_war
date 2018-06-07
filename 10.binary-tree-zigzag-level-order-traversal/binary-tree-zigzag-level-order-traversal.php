<?php
// ret [[3],[20,9],[15,7]]
function zigzagLevelOrder($arr)
{
    if (empty($arr)) {
        return [];
    }
    return recursive([], [$arr]);
}

function recursive($ret, $arr)
{
    $res = getLevelData($arr);
    if (!empty($res['ret'])) {
        if (count($ret) % 2 == 1) {
            // 翻转
            $temp = [];
            for ($i = count($res['ret']) - 1; $i >= 0; $i--) {
                $temp[] = $res['ret'][$i];
            }
            $res['ret'] = $temp;
        }
        $ret[] = $res['ret'];
    }
    if (!empty($res['nextLevelTree'])) {
        $ret = recursive($ret, $res['nextLevelTree']);
    }
    return $ret;
}

function getLevelData($arr)
{
    $ret           = [];
    $nextLevelTree = [];

    if (!empty($arr)) {
        foreach ($arr as $item) {
            $ret[] = $item['value'];
            if (!empty($item['left'])) {
                $nextLevelTree[] = $item['left'];
            }
            if (!empty($item['right'])) {
                $nextLevelTree[] = $item['right'];
            }
        }
    }
    return array(
        'ret'           => $ret,
        'nextLevelTree' => $nextLevelTree,
    );
}

$arr = [
    'value' => 3,
    'left'  => [
        'value' => 9,
        'left'  => [],
        'right' => [
            'value' => 2,
            'left'  => [
                'value' => 6,
                'left'  => [],
                'right' => [],
            ],
            'right' => [],
        ],
    ],
    'right' => [
        'value' => 20,
        'left'  => [
            'value' => 15,
            'left'  => [],
            'right' => [],
        ],
        'right' => [
            'value' => 7,
            'left'  => [],
            'right' => [],
        ],
    ],
];

var_dump(zigzagLevelOrder($arr));