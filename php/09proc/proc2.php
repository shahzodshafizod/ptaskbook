<?php
function PowerA234($A, &$B, &$C, &$D){
        $B = $A * $A;
        $C = $B * $A;
        $D = $B * $B;
}
$hisob = 1;
while($hisob <= 5){
	$a = mt_rand(-40, 40);
	PowerA234($a, $b, $c, $d);
	echo "Адади додашуда: ".$a;
	echo "<br>Дараҷаи дуюми адад: ".$b;
	echo "<br>Дараҷаи сеюми адад: ".$c;
	echo "<br>Дараҷаи чорӯми адад: ".$d."<br><br>";
	$hisob++;
}
?>

/*
<?php
    function PowerA234($A, &$B, &$C, &$D) {
        $B = $A * $A;
        $C = $A * $B;
        $D = $A * $C;
    }
    
    $a = 0;
    $b = 0;
    $c = 0;
    $d = 0;
    for ($i = 0; $i < 5; $i++) {
        $a = (float)readline("a = ");
        PowerA234($a, $b, $c, $d);
        echo sprintf("b = %.2f\tc = %.2f\td = %.2f\n", $b, $c, $d);
    }
?>
*/