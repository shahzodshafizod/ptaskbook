<?php
function TrianglePS($a, &$P, &$S){
        $P = 3 * $a;
        $S = $a * $a * sqrt(3) / 4;
}
$hisob = 1;
while($hisob <= 3){
	$a = mt_rand(1, 100) / 4;
	TrianglePS($a, $p, $s);
	$a = number_format($a, 2);
	$p = number_format($p, 2);
	$s = number_format($s, 2);
	echo "Тарафи секунҷаи баробартараф: ".$a;
	echo "<br>Периметри секунҷа: ".$p;
	echo "<br>Масоҳати секунҷа: ".$s."<br><br>";
	$hisob++;
}
?>