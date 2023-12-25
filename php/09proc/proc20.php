<?php
function TriangleP($a, $h){
        $b = sqrt(pow($a / 2, 2) + $h * $h);
        $nat = $a + 2 * $b;
		$nat = number_format($nat, 2);
        return $nat;
}
$hisob = 1;
while($hisob <= 3){
	$a = mt_rand(1, 100) / 4;
	$h = mt_rand(1, 100) / 4;
	$natija = TriangleP($a, $h);
	$a = number_format($a, 2);
	$h = number_format($h, 2);
	echo "Асос: ".$a;
	echo "<br>Баландӣ: ".$h;
	echo "<br>Периметри секунҷа: ".$natija."<br><br>";
	$hisob++;
}
?>