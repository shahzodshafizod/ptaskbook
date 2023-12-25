<?php
const PI = 3.14;
$rad = mt_rand(0, 6) + 0.28;
$grad = 180 * $rad / PI;
$grad = intval($grad);
echo "Значение угла в радианах: ".$rad;
echo "<br>Значение угла в градусах: ".$grad;
?>

/*
<?php
const PI = 3.14;
$rad = mt_rand(0, 6) + 0.28;
$grad = 180 * $rad / PI;
$grad = intval($grad);
echo "Қимати кунҷ бо радиан: ".$rad;
echo "<br>Қимати кунҷ бо градус: ".$grad;
?>
*/