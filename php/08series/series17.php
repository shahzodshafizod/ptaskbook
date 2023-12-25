<?php
$hisob = 2;
$mant = true;
$N = mt_rand(1, 30);
echo "Миқдори ададҳои маҷмӯъ = ".$N."<br>";
$B = mt_rand(1, 100) / 4;
$B = number_format($B, 2);
echo "B = ".$B."<br><br>";
$A1 = mt_rand(1, 100);
if(($A1 > $B) && $mant){
	echo "B = ".$B."<br>";
	echo "A1 = ".$A1."<br>";
	$mant = false;
}
while($hisob <= $N){
	$A2 = mt_rand(1, 100);
	if($A1 < $A2){
		if(($A2 > $B) && $mant){
			echo "B = ".$B."<br>";
			echo "<br>A".$hisob." = ".$A1."<br>";
			$mant = false;
		}
		else
			echo "A".$hisob." = ".$A1."<br>";
		$hisob++;
		$A1 = $A2;
	}
	else
		break;
}
echo "A".$hisob." = ".$A2;
?>