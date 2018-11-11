package main

/*
Channels verbinden die gleichzeitig ablaufenden Go-Routinen. Sie können Werte von einer Go-Routine in Channels senden und diese Werte in einer Anderen empfangen.
Desweiteren werden sie zur Synchronisation der Go-Routinen benutzt. Bei einem Lesezugriff auf einen Channel der noch keinen Wert besitzt, wartet die lesende Routine auf einen Wert.
Ebenso wartet eine Go-Routine die einen Wert in den Channel schreibt auf einen Lesezugriff einer Go-Routine.
*/
