<?php
function Sign($X){
        $nat = $X / abs($X);
        return $nat;
}
$A = mt_rand(-100, 100);
$B = mt_rand(-100, 100);
echo "A = ".$A;
echo "<br>B = ".$B;
$natija = Sign($A) + Sign($B);
echo "<br>Sign(A) + Sign(B) = ".$natija;
?>