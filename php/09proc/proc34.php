<?php
function Fact($N){
        $hisob = 1;
        $fact = 1;
        while($hisob <= $N){
                $fact *= $hisob;
                $hisob++;
        }
        return $fact;
}
$hisob = 1;
while($hisob <= 5){
	$N = mt_rand(1, 14);
	$nat = Fact($N);
	echo $N."! = ".$nat."<br>";
	$hisob++;
}
?>