<?php
function Power1($A, $B){
        $nat = exp($B * log($A));
        return $nat;
}
function Power2($A, $N){
        $dar = 1;
        for($co = 1; $co <= $N; $co++)
                $dar *= $A;
        return $dar;
}
function Power3($A, $B){
        $but = intval($B);
        if($but == $B)
                $nat = Power2($A, $B);
        else
                $nat = Power1($A, $B);
        return $nat;
}
$A = mt_rand(-20, 20);
$B = mt_rand(-20, 20);
$C = mt_rand(-20, 20);
$P = mt_rand(-10, 10);
echo "A = ".$A;
echo "<br>B = ".$B;
echo "<br>C = ".$C;
echo "<br>P = ".$P;
echo "<br>A^P = ".Power3($A, $P);
echo "<br>B^P = ".Power3($B, $P);
echo "<br>C^P = ".Power3($C, $P);
?>