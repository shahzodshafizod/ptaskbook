<?php
const PI = 3.14;
function RadToDeg($R){
        $grad = ($R * 180) / PI;
        return $grad;
}
$hisob = 1;
while($hisob <= 5){
	$rad = mt_rand(0, 6) + 0.28;
	echo "Радиан: ".$rad."<br>";
	$grad = RadToDeg($rad);
	$grad = number_format($grad, 2);
	echo "Градус: ".$grad."<br><br>";
	$hisob++;
}
?>