<?php
const PI = 3.14;
function CircleS($R){
        $S = PI * $R * $R;
        return $S;
}
$hisob = 1;
while($hisob <= 3){
	$r = mt_rand(1, 100) / 4;
	$natija = CircleS($r);
	$r = number_format($r, 2);
	$natija = number_format($natija, 2);
	echo "Радиуси доира: ".$r;
	echo "<br>Масоҳати доира: ".$natija."<br><br>";
	$hisob++;
}
?>