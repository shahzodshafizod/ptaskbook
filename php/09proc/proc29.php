<?php
function DigitCount($K){
        $miq = 0;
        while($K > 0){
                $miq++;
                $K = intval($K / 10);
        }
        return $miq;
}
$hisob = 1;
while($hisob <= 5){
	$N = mt_rand(1, 1000000);
	echo "N".$hisob." = ".$N."<br>";
	$raq = DigitCount($N);
	echo "Миқдори рақамҳои адад: ".$raq."<br><br>";
	$hisob++;
}
?>