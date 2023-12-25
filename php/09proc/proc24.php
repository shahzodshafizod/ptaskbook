<?php
function Even($K){
        $nat = (($K % 2) == 0);
        return $nat;
}
$hisob = 1;
$miq = 0;
while($hisob <= 10){
	$n = mt_rand(-100, 100);
	echo "n".$hisob." = ".$n."<br>";
	if(Even($n))
		$miq++;
	$hisob++;
}
echo "Миқдори ададҳои ҷуфт: ".$miq;
?>