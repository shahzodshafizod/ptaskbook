<?php
const PI = 3.14;
$grad = mt_rand(1, 359);
$rad = PI * $grad / 180;
$rad = number_format($rad, 2);
echo "Значение угла в градусах: ".$grad;
echo "<br>Значение угла в радианах: ".$rad;
?>

/*
<?php
const PI = 3.14;
$grad = mt_rand(1, 359);
$rad = PI * $grad / 180;
$rad = number_format($rad, 2);
echo "Қимати кунҷ бо градус: ".$grad;
echo "<br>Қимати кунҷ бо радиан: ".$rad;
?>
*/