<?php
function InvertDigits(&$K){
        $adad = 0;
        while($K > 0){
                $adad = $adad * 10 + $K % 10;
                $K = intval($K / 10);
        }
        $K = $adad;
}
$hisob = 1;
while($hisob <= 5){
	$N = mt_rand(1, 1000000);
	echo "Адади додашуда: ".$N;
	InvertDigits($N);
	echo "<br>Адад бо чаппагӣ: ".$N."<br><br>";
	$hisob++;
}
?>