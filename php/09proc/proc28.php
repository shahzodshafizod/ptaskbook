<?php
function IsPrime($K){
        $i = 2;
        $mant = true;
        while($i <= sqrt($K)){
                if(($K % $i) == 0)
                        $mant = false;
                $i++;
        }
        return $mant;
}
$hisob = 1;
$miq = 0;
while($hisob <= 10){
	$n = mt_rand(2, 1000000);
	echo "n".$hisob." = ".$n."<br>";
	if(IsPrime($n))
		$miq++;
	$hisob++;
}
echo "Миқдори ададҳои оддӣ: ".$miq;
?>