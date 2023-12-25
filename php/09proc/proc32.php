<?php
const PI = 3.14;
function DegToRad($D){
        $rad = PI * $D / 180;
        return $rad;
}
$hisob = 1;
while($hisob <= 5){
	$grad = mt_rand(1, 359);
	echo "Градус: ".$grad."<br>";
	$rad = DegToRad($grad);
	$rad = number_format($rad, 2);
	echo "Радиан: ".$rad."<br><br>";
	$hisob++;
}
?>