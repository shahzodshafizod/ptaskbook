<?php
function RectPS($x1, $y1, $x2, $y2, &$P, &$S){
	$a = abs($x1 - $x2);
	$b = abs($y1 - $y2);
	$P = 2 * ($a + $b);
	$S = $a * $b;
}
$hisob = 1;
while($hisob <= 3){
	$x1 = mt_rand(1, 50);
	$x2 = mt_rand(1, 50);
	$y1 = mt_rand(1, 50);
	$y2 = mt_rand(1, 50);
	RectPS($x1, $y1, $x2, $y2, $p, $s);
	echo "Координатаҳои тарафи якӯми росткунҷа:";
	echo "<br>x1 = ".$x1."<br>y1 = ".$y1;
	echo "<br>Координатаҳои тарафи дуюми росткунҷа:";
	echo "<br>x2 = ".$x2."<br>y2 = ".$y2;
	echo "<br>Периметри росткунҷа: ".$p;
	echo "<br>Масоҳати росткунҷа: ".$s."<br><br>";
	$hisob++;
}
?>