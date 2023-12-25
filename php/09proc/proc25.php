<?php
function IsSquare($K){
        $w = intval(sqrt($K));
        $nat = ($w * $w == $K);
        return $nat;
}
$hisob = 1;
$miq = 0;
while($hisob <= 10){
	$n = mt_rand(1, 100);
	echo "n".$hisob." = ".$n."<br>";
	if(IsSquare($n))
		$miq++;
	$hisob++;
}
echo "Миқдори квадратҳо: ".$miq;
?>